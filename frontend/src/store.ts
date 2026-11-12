import { writable, get } from 'svelte/store';
import * as api from '../wailsjs/go/app/App';
import { createPlayerState } from './lib/playerState';
import { createAuthState } from './lib/authState';
export { normalizeError } from './lib/errors';

export const selectedMode = writable('standard');
export const selectedStat = writable('count_rank_ss');
export const username = writable('');
export const warningMessage = writable('');

const authentication = createAuthState(api);
export const { authState, clientId, clientSecret, credentialsSet, clientValid, saveCredentials } = authentication;
export async function initializeCredentials() {
  await authentication.initialize();
  const state = get(authState);
  if (state.error) warningMessage.set(state.error.message);
}

const comparison = createPlayerState(api, message => warningMessage.set(message));
export const { players, loadingCount, isLoading, changePlayerMode, retryPlayer, removePlayer, clearPlayers, fetchPlayerStats } = comparison;
export async function addPlayer(name: string) {
  const state = get(authState);
  if (state.status !== 'ready') {
    warningMessage.set(state.status === 'loading' || state.status === 'saving'
      ? 'Please wait while your API credentials are being verified.'
      : 'Set and verify your API credentials in Settings before adding players.');
    return false;
  }
  return comparison.addPlayer(name, get(selectedMode));
}
