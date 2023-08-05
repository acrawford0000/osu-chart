import { writable } from "svelte/store";
import { get } from "svelte/store";



export const players = writable([]);

export async function addPlayer(username) {
  // Make the API request to get the user object for the given username
  const response = await fetch(`https://osu.ppy.sh/api/v2/users/${username}`);
  const user = await response.json();

  // Add the user object to the store
  players.update(currentPlayers => [...currentPlayers, user]);
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
