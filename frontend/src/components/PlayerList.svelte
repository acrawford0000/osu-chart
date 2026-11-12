<script>
  import { players, removePlayer, changePlayerMode, retryPlayer } from '../store';
  import { GAME_MODES } from '../lib/metrics';
  import AppIcon from './AppIcon.svelte';
</script>
<div class="player-list">
  {#each $players as player (player.key)}
    <article class="player-card panel">
      <div class="player-cover" style:background-image={player.cover_url ? `url(${player.cover_url})` : 'none'}></div>
      <button class="icon-button player-remove" aria-label={`Remove ${player.username}`} title="Remove player" on:click={() => removePlayer(player.key)}><AppIcon name="close" size={16}/></button>
      <div class="player-body">
        <div class="player-identity">
          {#if player.avatar_url}<img class="player-avatar" src={player.avatar_url} alt=""/>{:else}<span class="avatar-placeholder">{player.username.slice(0,1).toUpperCase()}</span>{/if}
          <div class="player-info"><h3>{player.username}</h3><span class="player-country">{player.country?.name || 'osu! player'}</span></div>
        </div>
        <div class="player-status" aria-live="polite">
          {#if player.status === 'loading'}<span class="spinner"></span>Loading saved history…
          {:else if player.status === 'empty'}No saved history for this mode.
          {:else if player.status === 'error'}<span class="error-text" title={player.error?.message}>Could not load data.</span><button class="button-quiet" on:click={() => retryPlayer(player.key)}>Retry</button>
          {:else}<span class="status-dot ready"></span>{(player.stats?.length || 0).toLocaleString()} saved snapshots{/if}
        </div>
        <label class="mode-label">Game mode<select aria-label={`Game mode for ${player.username}`} class="mode-select" value={player.selectedMode} on:change={event => changePlayerMode(player.key,event.currentTarget.value)}>{#each GAME_MODES as mode}<option value={mode.key}>{mode.label}</option>{/each}</select></label>
      </div>
    </article>
  {/each}
</div>
<style>
  .player-list { display: grid; grid-template-columns: repeat(auto-fill,minmax(min(270px,100%),1fr)); gap: 18px; }
  .player-card { position: relative; overflow: hidden; }
  .player-cover { height: 72px; background-size: cover; background-position: center; background-color: #2a2543; opacity: .55; }
  .player-cover::after { content: ''; position: absolute; inset: 0 0 auto; height: 73px; background: linear-gradient(0deg,var(--surface),#191b2540); }
  .player-remove { position: absolute; top: 10px; right: 10px; background: #101116c9; color: #d4d4e5; }
  .player-remove:hover { color: var(--danger); }
  .player-body { position: relative; padding: 0 22px 22px; }
  .player-identity { display: flex; align-items: center; gap: 12px; margin-top: -19px; }
  .player-avatar,.avatar-placeholder { width: 54px; height: 54px; border-radius: 14px; border: 3px solid var(--surface); flex-shrink: 0; background: #302946; }
  .avatar-placeholder { display: grid; place-items: center; color: var(--accent); font-size: 21px; }
  .player-info { min-width: 0; padding-top: 12px; }
  h3 { margin: 0 0 3px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
  .player-country { color: var(--muted); font-size: 11px; }
  .player-status { min-height: 44px; display: flex; gap: 7px; align-items: center; color: var(--muted); font-size: 11px; }
  .error-text { color: var(--danger); }
  .player-status button { font-size: 11px; padding: 3px 7px; }
  .mode-label { border-top: 1px solid var(--border); padding-top: 14px; font-size: 10px; color: var(--muted); font-weight: 500; }
  .mode-select { min-height: 38px; padding: 8px 10px; font-size: 12px; margin-top: 7px; }
</style>
