<script>
    import TextField from "@smui/textfield";
    import Icon from '@smui/textfield/icon';
    import AddButton, { Label } from "@smui/button";
    import { username, addPlayer, clientValid, credentialsSet, players, savePlayerData, warningMessage } from "../store";
    import { IsClientValid, AreOsuAuthCredentialsSet } from "../../wailsjs/go/app/App";
    import { get } from 'svelte/store';

    let textfield;
    function handleIconClick() {
        username.set('');
        textfield.focus();
    }

    async function handleAddPlayer() {
        if (!$credentialsSet) {
            warningMessage.set("Please set your API credentials in the settings before adding a player.")
            return;
        }
        if (!IsClientValid) {
            warningMessage.set("Client is invalid. Please check your credentials and try again.");
            return;
        }
        const currentPlayers = get(players);
        if (currentPlayers.some(player => player.username.toLowerCase() === $username.toLowerCase())) {
            warningMessage.set("A player with that username already exists.");
            
            return
        }
        addPlayer($username);
        username.set('');
        $clientValid = true;
    }

    function handleKeydown(event) {
        if (event.key === 'Enter') {
            handleAddPlayer();
            textfield.focus();
        }
    }


/*     function savePlayersToFile() {
        const data = JSON.stringify(players, null, 2);
    }

    // Function to send player data to the backend
    async function savePlayersToFile() {
    // Get all players from the store
    let allPlayers = [];
    const unsubscribe = players.subscribe(currentPlayers => {
        allPlayers = currentPlayers;
    });
    unsubscribe();

    // Convert player data to JSON
    const playerData = JSON.stringify(allPlayers);

    // Send player data to the backend
    try {
        await savePlayerData(playerData);
        console.log('Successfully sent player data to the backend');
    } catch (error) {
        console.error('Error sending player data to the backend:', error);
    }
    } */

</script>

<div class="AddPlayer">
    <TextField
        bind:value={$username}
        bind:this={textfield}
        on:keydown={handleKeydown}
        class="shape-filled"
        variant="filled"
        style="width: 100%; background-color: #1e1f22;"
        label="Enter a username"
        >
        <Icon 
        class="material-icons" role=button slot="trailingIcon" on:click={handleIconClick}>clear</Icon> 

    </TextField>
    <AddButton 
        on:click={handleAddPlayer}
        touch variant="unelevated"
        style="width: 40%; height: 56px; margin-left: 5%"
        >
        <Label>Add Player</Label>
    </AddButton>
</div>

<style>
    .AddPlayer {
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>