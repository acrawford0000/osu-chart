import { writable } from "svelte/store";
import { get } from "svelte/store";
import { GetUser, IsClientValid, GetStatsUpdates, SavePlayerData } from "../wailsjs/go/app/App";

// Set up various stores
export const selectedMode = writable('standard');
export const username = writable('');
export const credentialsSet = writable(false);
export const clientValid = writable(false);
export const selectedStat = writable('count_rank_ss');
export const clientId = writable('');
export const clientSecret = writable('');
export const warningMessage = writable('');
export const isLoading = writable(false);
export const loadingCount = writable(0);

// Set up the player store and all functions that go with it
export const players = writable([]);

export async function addPlayer(username) {
  // Increment loadingCount and set isLoading to true
  loadingCount.update(n => n + 1);
  isLoading.set(true);

  // Check if there are any players without data, if so, update them. This is an extra precaution.
  await updatePlayerData();
  
  // Call the GetUser function from your backend to get the user object for the given username
  try {
  const user = await GetUser(username);

  // Add the user object to the store
  players.update(currentPlayers => [...currentPlayers, { ...user, selectedMode: get(selectedMode)}]);

  // Fetch stats for added player
  await fetchPlayerStats();
  } catch (error) {
    warningMessage.set("An error occurred while getting user data: " + error.message);

    // Decrement loadingCount and if 0, set isLoading to false
    loadingCount.update(n => n - 1);
    loadingCount.subscribe(value => {
      if (value === 0) {
        isLoading.set(false);
      }
    });
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
  const unsubscribe = players.subscribe(currentPlayers => {
    allPlayers = currentPlayers;
  });
  unsubscribe();
  

  // Loop through all players and fetch their stats
  for (const player of allPlayers) {
    // Only fetch stats for players that don't already have stats
    if (!player.stats) {
      try {
        // Call the GetStatsUpdates function from the backend to get the player's stats using their user ID
        const stats = await GetStatsUpdates(player.id);

        // Update the player object in the store with the returned stats
        players.update(currentPlayers => currentPlayers.map(p => {
          if (p.username === player.username) {
            return { ...p, stats };
          }
          return p;
        }));
      } catch (error) {
        console.error(`Error fetching stats for player ${player.username}:`, error);
      }
    }
  }
  // Decrement loadingCount and if 0, set isLoading to false
  loadingCount.update(n => n - 1);
  loadingCount.subscribe(value => {
    if (value === 0) {
      isLoading.set(false);
    }
  });
}

export function removePlayer(username) {
  players.update(currentPlayers => currentPlayers.filter(player => player.username !== username));
}

export function clearPlayers() {
  players.set([]);
}

export function isClientValid() {
  IsClientValid().then(set => {
    clientValid.set(set);
  });
}

export async function savePlayerData(playerData: string) {
  try {
    // Call the backend function with the player data
    await SavePlayerData(playerData);
    console.log('Successfully saved player data');
  } catch (error) {
    console.error('Error saving player data:', error);
    throw error;
  }
}