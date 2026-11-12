<script>
  import { onDestroy } from 'svelte';
  import { BrowserOpenURL } from '../../wailsjs/runtime';
  import AppIcon from './AppIcon.svelte';
  import Picture from '../assets/images/ac_pfp.png';
  let copied = false;
  let copyError = '';
  let timer;
  function openLink(event) { event.preventDefault(); BrowserOpenURL(event.currentTarget.href); }
  async function copyDiscord() {
    try { await navigator.clipboard.writeText('acrawford69'); copied = true; copyError = ''; clearTimeout(timer); timer = setTimeout(() => copied = false, 2200); }
    catch { copyError = 'Could not copy. My Discord username is acrawford69.'; }
  }
  onDestroy(() => clearTimeout(timer));
</script>
<article class="creator-card">
  <span class="creator-tag">THE PERSON BEHIND THE CHARTS</span>
  <img class="creator-avatar" src={Picture} alt="ACrawford"/>
  <h1>Andrew Crawford</h1><span class="handle">@ACrawford</span>
  <p>Built for the love of osu!<br/>A little perspective on how far we’ve come.</p>
  <div class="creator-links"><a href="https://www.youtube.com/@ACrawford00" on:click={openLink} class="button button-secondary"><AppIcon name="video" size={17}/>YouTube</a><a href="https://github.com/acrawford0000" on:click={openLink} class="button button-secondary"><AppIcon name="github" size={17}/>GitHub</a></div>
  <button class="discord-button" on:click={copyDiscord}><AppIcon name={copied ? 'check' : 'copy'} size={14}/>{copied ? 'Copied to clipboard' : 'acrawford69 on Discord'}</button>
  <span class="copy-status" role="status">{copyError || (copied ? 'Discord username copied.' : '')}</span>
</article>
<style>
  .creator-card { width: 390px; max-width: calc(100vw - var(--sidebar-width) - 44px); background: #191b25f5; border: 1px solid #b7a9ff2b; box-shadow: 0 24px 80px #05050b70; padding: 34px 30px 24px; border-radius: 22px; text-align: center; }
  .creator-tag { display: block; font-size: 8px; font-weight: 650; letter-spacing: 1.8px; color: var(--accent); margin-bottom: 27px; }
  .creator-avatar { width: 100px; height: 100px; border-radius: 25px; border: 1px solid var(--border); margin-bottom: 22px; }
  h1 { font-size: 25px; letter-spacing: -.6px; margin-bottom: 7px; }
  .handle { font-size: 11px; color: var(--accent); }
  p { font-size: 12px; line-height: 1.8; margin: 24px 0; }
  .creator-links { display: flex; gap: 10px; }
  .creator-links a { flex: 1; font-size: 11px; }
  .discord-button { margin: 19px auto 0; background: transparent; color: var(--muted); font-size: 11px; padding: 5px; font-weight: 400; }
  .copy-status { min-height: 12px; display: block; color: var(--muted); font-size: 9px; margin-top: 5px; }
  @media(max-width:600px) { .creator-card { padding: 28px 20px 18px; } h1 { font-size: 22px; } }
</style>
