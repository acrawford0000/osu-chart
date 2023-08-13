import { writable } from "svelte/store";
import { get } from "svelte/store";
import { GetUser } from "../wailsjs/go/app/App";
import { AreOsuAuthCredentialsSet } from "../wailsjs/go/app/App"



export const players = writable([]);

export async function addPlayer(username) {
  // Check if credentials are set, do not allow player to be added without credentials set
  if (AreOsuAuthCredentialsSet()) {
    await updatePlayerData();
  // Call the GetUser function from your backend to get the user object for the given username
  const user = await GetUser(username);

  // Add the user object to the store
  players.update(currentPlayers => [...currentPlayers, user]);
  } else {
    // Credentials are not set, just update players array with the entered username
    alert("Must set Client ID and Client Secret in settings");
    return;
  }
}

async function updatePlayerData() {
  // Get current value of players store
  const currentPlayers = get(players);

  // Iterate over players array
  for (let i = 0; i < currentPlayers.length; i++) {
    const player = currentPlayers[i];
    // Check if player only has a username (i.e. their data was not fetched)
    if (Object.keys(player).length === 1 && player.username) {
      // Get player data and update players array
      const updatedPlayer = await GetUser(player.username);
      players.update(currentPlayers => {
        currentPlayers[i] = updatedPlayer;
        return currentPlayers;
      });
    }
  }
}


export async function fetchPlayerStats() {
  // Get all players from the store
  let allPlayers = [];
  players.subscribe(currentPlayers => {
    allPlayers = currentPlayers;
  });

  // Loop through all players and fetch their stats
  for (const player of allPlayers) {
    // Only fetch stats for players that don't already have stats
    if (!player.stats) {
      // Make the API request to get the player's stats using their user ID
      const response = await fetch(`https://osutrack-api.ameo.dev/stats_history?user=${player.id}&mode=0`);
      const stats = await response.json();

      // Update the player object in the store with the returned stats
      players.update(currentPlayers => currentPlayers.map(p => {
        if (p.username === player.username) {
          return { ...p, stats };
        }
        return p;
      }));
    }
  }
}

export function removePlayer(username) {
  players.update(currentPlayers => currentPlayers.filter(player => player.username !== username));
}

export function clearPlayers() {
  players.set([]);
}
