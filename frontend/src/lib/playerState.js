import { writable, derived, get } from 'svelte/store';
import { normalizeError } from './errors.js';

// Keep requests and cache ownership together. Components invoke actions rather
// than balancing counters or changing the backend's global game mode.
export function createPlayerState(api, reportError = (_message) => {}) {
  const players = writable([]);
  const loadingCount = derived(players, value => value.filter(player => player.status === 'loading').length);
  const isLoading = derived(loadingCount, count => count > 0);
  const requests = new Map();
  const historyCache = new Map();
  const keyFor = name => name.trim().toLowerCase();
  const find = key => get(players).find(player => player.key === key);
  const update = (key, patch) => players.update(value => value.map(player => player.key === key ? { ...player, ...patch } : player));

  async function loadPlayer(key, mode, force = false) {
    const existing = find(key);
    if (!existing) return false;
    const request = Symbol(key);
    requests.set(key, request);
    const current = () => requests.get(key) === request && !!find(key);
    update(key, { selectedMode: mode, status: 'loading', error: null, stats: existing.selectedMode === mode ? existing.stats : null });
    try {
      const profile = existing.id ? existing : await api.GetUser(existing.username, mode);
      if (!current()) return false;
      if (!profile?.id) throw 'INVALID_RESPONSE: osu! returned an incomplete player profile. Please retry.';
      if (get(players).some(player => player.key !== key && player.id === profile.id)) {
        throw 'DUPLICATE_PLAYER: This player is already in your comparison.';
      }
      update(key, { ...profile, key, selectedMode: mode, stats: existing.selectedMode === mode ? existing.stats : null, status: 'loading', error: null });
      const cacheKey = `${profile.id}:${mode}`;
      const cached = historyCache.get(cacheKey);
      const age = cached ? Date.now() - cached.savedAt : Infinity;
      const stats = !force && cached && age >= 0 && age < 5 * 60 * 1000
        ? cached.stats
        : await api.GetStatsUpdates(profile.id, mode);
      if (!current()) return false;
      if (!Array.isArray(stats)) throw 'INVALID_RESPONSE: osu!track returned unreadable history. Please retry.';
      historyCache.set(cacheKey, { stats, savedAt: Date.now() });
      update(key, { stats, status: stats.length ? 'ready' : 'empty', error: null });
      return true;
    } catch (error) {
      if (current()) {
        const problem = normalizeError(error, 'Could not load this player. Please retry.');
        update(key, { status: 'error', error: problem });
        reportError(`${existing.username}: ${problem.message}`);
      }
      return false;
    } finally {
      if (requests.get(key) === request) requests.delete(key);
    }
  }

  async function addPlayer(name, mode = 'standard') {
    const username = name.trim();
    if (!username) { reportError('Enter an osu! username first.'); return false; }
    const key = keyFor(username);
    if (get(players).some(player => player.key === key || keyFor(player.username) === key)) {
      reportError('A player with that username is already in your comparison.');
      return false;
    }
    players.update(value => [...value, { key, username, selectedMode: mode, stats: null, status: 'loading', error: null }]);
    return loadPlayer(key, mode);
  }
  function changePlayerMode(key, mode) { return loadPlayer(key, mode); }
  function retryPlayer(key) {
    const player = find(key);
    return player ? loadPlayer(key, player.selectedMode, true) : Promise.resolve(false);
  }
  function removePlayer(key) {
    requests.delete(key);
    players.update(value => value.filter(player => player.key !== key));
  }
  function clearPlayers() {
    requests.clear();
    historyCache.clear();
    players.set([]);
  }
  async function fetchPlayerStats() {
    await Promise.all(get(players).filter(player => player.status === 'error' || !player.stats)
      .map(player => loadPlayer(player.key, player.selectedMode, true)));
  }
  return { players, loadingCount, isLoading, addPlayer, changePlayerMode, retryPlayer, removePlayer, clearPlayers, fetchPlayerStats };
}
