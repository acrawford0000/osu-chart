<script>
  import PlayerList from '../components/PlayerList.svelte';
  import ClearAllButton from '../components/ClearAllButton.svelte';
  import AddPlayer from '../components/AddPlayer.svelte';
  import AppIcon from '../components/AppIcon.svelte';
  import { players, authState, selectedMode } from '../store';
  import { GAME_MODES } from '../lib/metrics';
</script>
<main class="page">
  <header class="page-heading">
    <div><div class="eyebrow">YOUR WORKSPACE</div><h1>Better together.</h1><p>Bring your players together. Discover the story behind the stats.</p></div>
    <span class="status-pill"><span class="status-dot" class:ready={$authState.status === 'ready'}></span>{$authState.status === 'ready' ? 'API connected' : 'Set up your API in Settings'}</span>
  </header>
  <section class="panel add-panel" aria-label="Add players">
    <div class="add-intro"><span class="section-icon"><AppIcon name="players" size={22}/></span><div><h2>Build your comparison</h2><p>Add a player using their osu! username.</p></div></div>
    <AddPlayer/>
    <div class="add-footnote">Adding players in <strong>{GAME_MODES.find(mode => mode.key === $selectedMode)?.label}</strong>. Change the default in the sidebar, or choose a mode for each player.</div>
  </section>
  <div class="section-heading"><h2>Your players <span class="count">{$players.length}</span></h2>{#if $players.length}<ClearAllButton/>{/if}</div>
  {#if $players.length}
    <PlayerList/>
  {:else}
    <div class="empty-state"><span class="empty-icon"><AppIcon name="players" size={26}/></span><h2>Your comparison starts here</h2><p>Add yourself, a friend, or a player you follow. Their saved history will be ready to explore in Chart.</p></div>
  {/if}
</main>
<style>
  .add-panel { padding: 26px; margin-bottom: 34px; }
  .add-intro { display: flex; align-items: center; gap: 14px; margin-bottom: 24px; }
  .add-intro h2 { margin-bottom: 3px; }
  .add-intro p { font-size: 12px; margin: 0; }
  .section-icon { width: 44px; height: 44px; background: var(--accent-wash); color: var(--accent); border-radius: 12px; display: grid; place-items: center; flex-shrink: 0; }
  .add-footnote { margin-top: 15px; color: var(--muted); font-size: 11px; line-height: 1.6; }
  .add-footnote strong { color: #c7c0e7; font-weight: 500; }
  .section-heading { display: flex; align-items: center; justify-content: space-between; margin-bottom: 18px; }
  .section-heading h2 { margin: 0; display: flex; align-items: center; gap: 9px; }
  .count { display: inline-grid; place-items: center; min-width: 24px; height: 23px; border-radius: 6px; border: 1px solid var(--border); color: var(--muted); font-size: 11px; font-weight: 500; }
  @media(max-width:600px) { .add-panel { padding: 20px; } }
</style>
