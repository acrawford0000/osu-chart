<script>
  import { link as routerLink, useLocation } from 'svelte-routing';
  import { selectedMode, authState } from '../store';
  import AppIcon from './AppIcon.svelte';
  import { GAME_MODES } from '../lib/metrics';
  import modeOsu from '../assets/images/mode-osu.png';
  import modeTaiko from '../assets/images/mode-taiko.png';
  import modeMania from '../assets/images/mode-mania.png';
  import modeFruits from '../assets/images/mode-fruits.png';
  const images = [modeOsu, modeTaiko, modeMania, modeFruits];
  const links = [
    { icon: 'players', text: 'Players', path: '/' },
    { icon: 'chart', text: 'Chart', path: '/chart' },
    { icon: 'settings', text: 'Settings', path: '/settings' },
    { icon: 'contact', text: 'Contact', path: '/contact' }
  ];
  const location = useLocation();
</script>

<aside class="sidebar">
  <a href="/" use:routerLink class="brand" aria-label="osu!Chart home">
    <span class="brand-mark">o!</span><span class="brand-name">osu!<strong>Chart</strong><small>YOUR PROGRESS, TOGETHER</small></span>
  </a>
  <div class="nav-label">WORKSPACE</div>
  <nav aria-label="Main navigation">
    {#each links as item}
      <a href={item.path} use:routerLink class="sidebar-item" aria-label={item.text} title={item.text} aria-current={$location.pathname === item.path ? 'page' : undefined}>
        <AppIcon name={item.icon}/><span>{item.text}</span>
        {#if $location.pathname === item.path}<span class="nav-active-dot"></span>{/if}
      </a>
    {/each}
  </nav>
  <div class="sidebar-bottom">
    <div class="nav-label">DEFAULT GAME MODE</div>
    <div class="mode-grid" role="group" aria-label="Default game mode">
      {#each GAME_MODES as mode, index}
        <button class="mode-button" class:selected={$selectedMode === mode.key} aria-pressed={$selectedMode === mode.key} title={mode.label} on:click={() => selectedMode.set(mode.key)}>
          <img src={images[index]} alt=""/><span>{mode.label}</span>
        </button>
      {/each}
    </div>
    <div class="connection"><span class="status-dot" class:ready={$authState.status === 'ready'}></span><span>{$authState.status === 'ready' ? 'Connected to osu!' : $authState.status === 'loading' ? 'Connecting…' : 'API setup needed'}</span></div>
  </div>
</aside>

<style>
  .sidebar { position: fixed; inset: 0 auto 0 0; width: var(--sidebar-width); background: var(--sidebar); border-right: 1px solid var(--border); display: flex; flex-direction: column; padding: 28px 16px 18px; z-index: 50; overflow-y: auto; }
  .brand { display: flex; gap: 12px; align-items: center; color: var(--text); margin: 0 8px 42px; white-space: nowrap; }
  .brand-mark { display: grid; place-items: center; width: 38px; height: 38px; flex-shrink: 0; border-radius: 11px; background: var(--accent-strong); font-size: 23px; font-weight: 750; letter-spacing: -2px; color: white; }
  .brand-name { font-size: 18px; letter-spacing: -.5px; }
  .brand-name strong { font-weight: 650; }
  .brand-name small { display: block; font-size: 6.5px; letter-spacing: 1.5px; color: var(--muted); margin-top: 4px; }
  .nav-label { font-size: 9px; font-weight: 650; letter-spacing: 1.6px; color: #6e758d; margin: 0 12px 12px; }
  nav { display: grid; gap: 5px; }
  .sidebar-item { display: flex; align-items: center; position: relative; gap: 13px; height: 46px; border-radius: 9px; padding: 0 13px; color: var(--muted); font-size: 13px; font-weight: 550; }
  .sidebar-item:hover { background: #ffffff05; color: var(--text); }
  .sidebar-item[aria-current='page'] { background: var(--accent-wash); color: #c3b6ff; }
  .sidebar-item[aria-current='page']::before { content: ''; position: absolute; left: -16px; top: 11px; bottom: 11px; width: 3px; border-radius: 0 3px 3px 0; background: var(--accent); }
  .nav-active-dot { width: 5px; height: 5px; border-radius: 50%; background: var(--accent); margin-left: auto; }
  .sidebar-bottom { margin-top: auto; padding-top: 48px; }
  .mode-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 6px; }
  .mode-button { padding: 8px; min-width: 0; gap: 7px; background: transparent; border: 1px solid var(--border); color: var(--muted); font-size: 10px; font-weight: 500; }
  .mode-button img { width: 18px; height: 18px; }
  .mode-button.selected { color: #ddd6ff; border-color: #aa9aff65; background: var(--accent-wash); }
  .connection { display: flex; align-items: center; gap: 8px; border-top: 1px solid var(--border); margin-top: 22px; padding: 19px 10px 0; color: var(--muted); font-size: 10px; }
  @media(max-width:1099px) { .sidebar { padding: 24px 12px 18px; } .brand { margin: 0 auto 36px; } .brand-name,.nav-label,.sidebar-item > span,.mode-button > span,.connection > span:last-child { display: none; } .sidebar-item { justify-content: center; padding: 0; } .sidebar-item[aria-current='page']::before { left: -12px; } .mode-grid { grid-template-columns: 1fr; } .mode-button { padding: 9px; } .connection { justify-content: center; padding: 16px 0 0; } }
  @media(max-width:600px) { .sidebar { padding-left: 8px; padding-right: 8px; } .brand-mark { width: 34px; height: 34px; } .sidebar-item[aria-current='page']::before { left: -8px; } }
</style>
