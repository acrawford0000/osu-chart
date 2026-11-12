import { writable, derived, get } from 'svelte/store';
import { normalizeError } from './errors.js';

export function createAuthState(api) {
  const authState = writable({ status: 'loading', hasCredentials: false, error: null });
  const clientId = writable('');
  const clientSecret = writable('');
  const credentialsSet = derived(authState, state => state.hasCredentials);
  const clientValid = derived(authState, state => state.status === 'ready');
  let operation = 0;

  async function initialize() {
    if (get(authState).status === 'saving') return;
    const request = ++operation;
    authState.set({ status: 'loading', hasCredentials: false, error: null });
    let hasCredentials = false;
    try {
      const data = await api.GetCredentials();
      if (request !== operation) return;
      clientId.set(data.client_id || '');
      clientSecret.set(data.client_secret || '');
      hasCredentials = !!(data.client_id && data.client_secret);
      if (!hasCredentials) {
        authState.set({ status: 'missing', hasCredentials: false, error: null });
        return;
      }
      if (!(await api.IsClientValid())) await api.CreateNewClient();
      if (request !== operation) return;
      if (!(await api.IsClientValid())) throw 'AUTH_INVALID: Could not initialize the API client. Check your credentials in Settings.';
      if (request === operation) authState.set({ status: 'ready', hasCredentials: true, error: null });
    } catch (error) {
      if (request === operation) authState.set({ status: 'error', hasCredentials, error: normalizeError(error) });
    }
  }

  async function saveCredentials() {
    const previous = get(authState);
    if (previous.status === 'saving' || previous.status === 'loading') return false;
    const id = get(clientId).trim();
    const secret = get(clientSecret).trim();
    if (!/^[1-9]\d*$/.test(id) || !secret) throw 'AUTH_INVALID: Enter a positive numeric Client ID and a non-empty Client Secret.';
    const request = ++operation;
    authState.set({ ...previous, status: 'saving', error: null });
    try {
      await api.SaveOsuAuthCredentials(id, secret);
      if (request === operation) {
        clientId.set(id); clientSecret.set(secret);
        authState.set({ status: 'ready', hasCredentials: true, error: null });
      }
      return true;
    } catch (error) {
      if (request === operation) authState.set({ ...previous, error: normalizeError(error) });
      throw error;
    }
  }
  return { authState, clientId, clientSecret, credentialsSet, clientValid, initialize, saveCredentials };
}
