<script>
  import { players, removePlayer, fetchPlayerStats, selectedMode } from "../store";
  import { SetGameMode } from "../../wailsjs/go/app/App"
  import { get } from "svelte/store";

  let allPlayers = [];

  players.subscribe((value) => {
    allPlayers = value;
  });

  let modes = ["standard", "taiko", "mania", "catch"];

  async function handleModeSelect(player, event) {
  // Get the current selected mode from the store
  let currentSelectedMode = get(selectedMode);

  // Set the game mode in the backend to the selected mode
  await SetGameMode(event.target.value);

  // Clear the player's stats and update their selected mode
  players.update((currentPlayers) =>
    currentPlayers.map((p) => {
      if (p.username === player.username) {
        return { ...p, stats: null, selectedMode: event.target.value };
      }
      return p;
    })
  );

  // Fetch stats for the selected player
  await fetchPlayerStats();

  // Set the game mode in the backend back to its original value
  await SetGameMode(currentSelectedMode);
}

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
          <div class="player-country">        
            <img class="player-flag" src= {`https://s.ppy.sh/images/flags/${player.country.code.toLowerCase()}.gif`}
            alt={player.country.name}/>
            {player.country.name}
          </div>
          <!-- <div class="player-id">ID: {player.id}</div> -->
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
        <select on:change={(event) => handleModeSelect(player, event)}
          class="mode-select"
          value={player.selectedMode}
          >
          {#each modes as mode}
            <option value={mode}>{mode}</option>
          {/each}
        </select>
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
    position: absolute;
    left: 50%;
    top: 50%;
    width: 60%;
    transform: translate(-50%, -50%);
    align-content: center;
    justify-content: center;
  }
  .player-username {
    font-size: 18px;
    font-weight: bold;
  }
  .player-country {
    font-size: 12px;
  }
  .player-flag {
    width: 16px;
    height: 11px;
    margin-right: 2%;
  }
/*   .player-id {
    font-size: 12px;
  } */
  .player-remove-button {
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
  }
  select {
    position: absolute;
    bottom: 5%;
    left: 2%;
    border: none;
    background-color: white;
  }
  select option {
    border: none;
    background-color: transparent;
  }
</style>
