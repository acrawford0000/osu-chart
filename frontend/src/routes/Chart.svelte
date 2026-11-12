<script>
  import { link as routerLink } from 'svelte-routing';
  import LineChart from '../components/PlayerChart.svelte';
  import StatSelector from '../components/StatSelectorV2.svelte';
  import AppIcon from '../components/AppIcon.svelte';
  import { isLoading, loadingCount, players, selectedStat } from '../store';
  import { getStat } from '../lib/metrics';
  $: metric = getStat($selectedStat);
  $: snapshots = $players.reduce((sum,player) => sum + (player.stats?.length || 0),0);
</script>
<main class="page">
  <header class="page-heading"><div><div class="eyebrow">THE BIGGER PICTURE</div><h1>Every snapshot tells a story.</h1><p>Compare your journey, one milestone at a time.</p></div><span class="status-pill"><AppIcon name="chart" size={13}/>{$players.length} {$players.length === 1 ? 'player' : 'players'} · {snapshots.toLocaleString()} snapshots</span></header>
  <section class="panel chart-panel" aria-label="Player comparison">
    <div class="chart-toolbar"><div><h2>{metric.label} over time</h2><span class="subtle">{metric.description}</span></div><StatSelector/></div>
    <div class="chart-content"><LineChart/>{#if !$players.length}<div class="chart-empty"><span class="empty-icon"><AppIcon name="chart" size={25}/></span><h2>A little history goes a long way.</h2><p>Add players to start comparing their saved stats.</p><a class="button button-primary" href="/" use:routerLink>Choose your players<AppIcon name="arrow" size={15}/></a></div>{/if}</div>
    <footer class="chart-footnote"><span class="hint-dot"></span>Hover to compare the nearest saved stats for each player.<span class="zoom-hint">Scroll to zoom · Drag to pan</span></footer>
  </section>
  {#if $isLoading}<div class="loading-container" role="status" aria-live="polite"><span class="spinner"></span>Loading {$loadingCount} {$loadingCount === 1 ? 'player' : 'players'}…</div>{/if}
</main>
<style>
  .chart-panel { overflow: hidden; }
  .chart-toolbar { padding: 24px 26px; display: flex; align-items: center; justify-content: space-between; gap: 20px; border-bottom: 1px solid var(--border); }
  .chart-toolbar h2 { margin-bottom: 5px; font-size: 16px; }
  .chart-toolbar .subtle { font-size: 11px; }
  .chart-content { position: relative; }
  .chart-empty { position: absolute; inset: 0; display: flex; flex-direction: column; align-items: center; justify-content: center; padding: 30px; text-align: center; background: #191b25ed; }
  .chart-empty p { font-size: 12px; }
  .chart-empty .button { margin-top: 10px; font-size: 12px; }
  .chart-footnote { border-top: 1px solid var(--border); padding: 16px 26px; font-size: 10px; line-height: 1.6; color: var(--muted); display: flex; gap: 8px; align-items: center; flex-wrap: wrap; }
  .hint-dot { width: 4px; height: 4px; background: var(--accent); border-radius: 50%; }
  .zoom-hint { margin-left: auto; color: #777f98; }
  .loading-container { display: flex; gap: 8px; align-items: center; margin-top: 14px; color: var(--muted); font-size: 12px; }
  @media(max-width:750px) { .chart-toolbar { align-items: flex-start; flex-direction: column; padding: 20px; } .chart-footnote { padding: 15px 20px; } .zoom-hint { width: 100%; margin: 0; padding-left: 12px; } }
</style>
