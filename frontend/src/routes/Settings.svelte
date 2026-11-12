<script>
  import ClientTextfields from '../components/ClientTextfields.svelte';
  import AppIcon from '../components/AppIcon.svelte';
  import { authState } from '../store';
  import { BrowserOpenURL } from '../../wailsjs/runtime';
  function openRegistration(event) { event.preventDefault(); BrowserOpenURL(event.currentTarget.href); }
</script>
<main class="page">
  <header class="page-heading"><div><div class="eyebrow">MAKE YOURSELF AT HOME</div><h1>Settings</h1><p>Your connection to the osu! universe.</p></div><span class="status-pill"><span class="status-dot" class:ready={$authState.status === 'ready'}></span>{$authState.status === 'ready' ? 'Connection verified' : 'Connection setup'}</span></header>
  <div class="settings-grid">
    <section class="panel panel-pad"><div class="panel-heading"><span class="setting-icon"><AppIcon name="lock" size={20}/></span><div><h2>API credentials</h2><p>Connect your osu! application to load player profiles.</p></div></div><ClientTextfields/><div class="privacy-note"><AppIcon name="lock" size={13}/><span>Your credentials stay on this computer. A failed save keeps your previous working credentials.</span></div></section>
    <section class="panel panel-pad setup-panel"><div class="eyebrow">FIRST TIME HERE?</div><h2>A quick connection</h2><p>Create an osu! OAuth application, then copy its credentials here.</p><ol class="setup-steps"><li><span>1</span><div><strong>Open your osu! account</strong><p>Sign in and open the OAuth application settings.</p></div></li><li><span>2</span><div><strong>Create an application</strong><p>Name it osu!Chart. Leave the callback URL empty.</p></div></li><li><span>3</span><div><strong>Copy, paste, connect</strong><p>Enter your Client ID and Client Secret, then save to verify the connection.</p></div></li></ol><a class="button button-secondary" href="https://osu.ppy.sh/home/account/edit#new-oauth-application" on:click={openRegistration}>Open osu! account settings<AppIcon name="external" size={15}/></a></section>
  </div>
</main>
<style>
  .settings-grid { display: grid; grid-template-columns: 1.1fr 1fr; gap: 24px; align-items: start; }
  .panel-heading { display: flex; align-items: flex-start; gap: 13px; margin-bottom: 28px; }
  .panel-heading h2 { margin-bottom: 5px; }
  .panel-heading p,.setup-panel > p { font-size: 12px; margin-bottom: 0; }
  .setting-icon { background: var(--accent-wash); color: var(--accent); padding: 11px; display: flex; border-radius: 12px; }
  .privacy-note { display: flex; gap: 7px; align-items: flex-start; font-size: 10px; line-height: 1.7; color: var(--muted); border-top: 1px solid var(--border); padding-top: 20px; margin-top: 25px; }
  .privacy-note :global(svg) { flex-shrink: 0; margin-top: 2px; }
  .setup-steps { list-style: none; padding: 0; margin: 26px 0; display: grid; gap: 23px; }
  li { display: flex; gap: 13px; }
  li > span { display: grid; place-items: center; width: 26px; height: 26px; border: 1px solid var(--border); border-radius: 8px; color: var(--accent); flex-shrink: 0; font-size: 11px; }
  li strong { display: block; font-size: 12px; font-weight: 550; padding-top: 3px; }
  li p { margin: 6px 0 0; font-size: 11px; }
  .setup-panel .button { font-size: 11px; }
  @media(max-width:1050px) { .settings-grid { grid-template-columns: 1fr; } }
</style>
