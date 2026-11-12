<script>
  import { username, addPlayer, authState } from '../store';
  import AppIcon from './AppIcon.svelte';
  async function submit() {
    const submitted = $username;
    if (await addPlayer(submitted) && $username === submitted) username.set('');
  }
</script>
<form class="AddPlayer" on:submit|preventDefault={submit}>
  <label class="username-field"><span>Find a player</span><input bind:value={$username} aria-label="Enter a username" placeholder="Enter an osu! username" autocomplete="off" spellcheck="false"/></label>
  <button class="button-primary" type="submit" disabled={$authState.status === 'loading' || $authState.status === 'saving'}><AppIcon name="plus" size={17}/>Add Player</button>
</form>
<style>
  .AddPlayer { display: flex; align-items: flex-end; gap: 12px; }
  .username-field { flex: 1; }
  button { height: 44px; white-space: nowrap; }
  @media(max-width:600px) { .AddPlayer { align-items: stretch; flex-direction: column; } }
</style>
