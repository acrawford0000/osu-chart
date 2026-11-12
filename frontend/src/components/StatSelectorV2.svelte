<script>
  import { selectedStat } from '../store';
  import { STAT_GROUPS } from '../lib/metrics';
</script>
<label class="stat-selector">Metric
  <select id="chart-metric" aria-label="Chart metric" bind:value={$selectedStat}>
    {#each STAT_GROUPS as group}<optgroup label={group.label}>{#each group.stats as stat}<option value={stat.key}>{stat.label}</option>{/each}</optgroup>{/each}
  </select>
</label>
<style>
  .stat-selector { display: flex; align-items: center; gap: 12px; color: var(--muted); font-weight: 500; font-size: 11px; }
  select { margin: 0; min-height: 38px; width: 210px; padding-top: 8px; padding-bottom: 8px; font-size: 12px; }
  @media(max-width:600px) { .stat-selector { width: 100%; } select { flex: 1; width: auto; } }
</style>
