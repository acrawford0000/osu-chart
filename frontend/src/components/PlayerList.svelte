<script>
  import { players, removePlayer } from "../store";
  let allPlayers = [];

  players.subscribe((value) => {
    allPlayers = value;
  });
</script>

<div class="player-list">
  {#each allPlayers as player}
    <div
      class="player-card"
      style="background-image:url({player.cover_url}); background-size:cover; background-position:center;"
    >
      <div class="player-card-overlay"></div>
      <div class="player-card-content">
        <img class="player-avatar" src={player.avatar_url} alt="Avatar" />
        <div class="player-info">
          <div class="player-username">{player.username}</div>
          <div class="player-id">ID: {player.id}</div>
        </div>
        <span
          class="material-icons player-remove-button"
          tabindex="0"
          role="button"
          on:click={() => removePlayer(player.username)}
          on:keydown={(e) => {
            if (e.key === "Enter" || e.key === " ") {
              removePlayer(player.username);
            }
          }}>close</span
        >
      </div>
    </div>
  {/each}
</div>

<style>
  .player-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
  .player-card {
    width: 300px;
    height: 150px;
    margin: 10px;
    border-radius: 10px;
    position: relative;
    overflow: hidden;
  }
  .player-card:hover {
    box-shadow: 0 0 11px rgba(82, 80, 80, 0.7);
  }
  .player-card-overlay {
    width: 100%;
    height: 100%;
    position: absolute;
    z-index: 0;
    background-color: rgba(0, 0, 0, 0.75);
  }
  .player-card-content {
    width: 100%;
    height: 100%;
    position: absolute;
    z-index: 1;
    display: flex;
    align-items: center;
  }
  .player-avatar {
    width: 50px;
    height: 50px;
    margin-left: 10px;
    border-radius: 50%;
  }
  .player-info {
    margin-left: 10px;
  }
  .player-username {
    font-size: 18px;
    font-weight: bold;
  }
  .player-id {
    font-size: 12px;
  }
  .player-remove-button {
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
  }
</style>
