<script>
  import { authState, saveCredentials, normalizeError, clientId, clientSecret } from '../store';
  import AppIcon from './AppIcon.svelte';
  let errorMessage = '';
  let success = false;
  let showSecret = false;
  $: busy = $authState.status === 'saving' || $authState.status === 'loading';
  async function onSave() {
    errorMessage = ''; success = false;
    try { success = await saveCredentials(); }
    catch (error) { errorMessage = normalizeError(error).message; }
  }
</script>
<form on:submit|preventDefault={onSave} class="field-stack">
  <label>Client ID<input aria-label="Client ID" bind:value={$clientId} disabled={busy} placeholder="Your application’s Client ID" inputmode="numeric" autocomplete="off" on:input={() => success = false}/></label>
  <label>Client Secret<div class="secret-field"><input aria-label="Client Secret" type={showSecret ? 'text' : 'password'} value={$clientSecret} disabled={busy} placeholder="Your application’s Client Secret" autocomplete="off" spellcheck="false" on:input={event => { clientSecret.set(event.currentTarget.value); success = false; }}/><button class="icon-button" type="button" aria-label={showSecret ? 'Hide Client Secret' : 'Show Client Secret'} aria-pressed={showSecret} on:click={() => showSecret = !showSecret}><AppIcon name="eye" size={17}/></button></div></label>
  {#if errorMessage}<div class="form-message error" role="alert">{errorMessage}</div>{/if}
  {#if success}<div class="form-message success" role="status"><AppIcon name="check" size={15}/>Credentials verified and saved.</div>{/if}
  <div class="save-row"><button class="button-primary" type="submit" disabled={busy}>{#if $authState.status === 'saving'}<span class="spinner"></span>Verifying…{:else}<AppIcon name="check" size={16}/>Save credentials{/if}</button></div>
</form>
<style>
  .secret-field { position: relative; margin-top: 8px; }
  .secret-field input { margin-top: 0; padding-right: 46px; }
  .secret-field button { position: absolute; right: 4px; top: 4px; }
  .form-message { font-size: 12px; line-height: 1.6; padding: 10px 12px; border-radius: 8px; display: flex; gap: 8px; align-items: center; }
  .error { color: var(--danger); background: #ff95a50b; border: 1px solid #ff95a52a; }
  .success { color: var(--success); background: #71d4b10b; border: 1px solid #71d4b12a; }
  .save-row { display: flex; padding-top: 3px; }
</style>
