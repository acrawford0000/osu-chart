<script>
    import TextField from "@smui/textfield";
    import Icon from '@smui/textfield/icon'
    import AddButton, { Label } from "@smui/button"
    import GetStatsButton from "@smui/button"
    import WarningCard from "./WarningCard.svelte";
    import { username, addPlayer, fetchPlayerStats, clientValid, credentialsSet, players } from "../store";
    import { IsClientValid, AreOsuAuthCredentialsSet } from "../../wailsjs/go/app/App"
    import { get } from 'svelte/store'

    let textfield;
    let errorMessage = '';

    function handleIconClick() {
        username.set('');
        textfield.focus();
    }

    async function handleAddPlayer() {
        if (!$credentialsSet) {
            errorMessage = "Please set your API credentials in the settings before adding a player.";
            return;
        }
        if (!IsClientValid) {
            errorMessage = "Client is invalid. Please check your credentials and try again.";
            return;
        }
        const currentPlayers = get(players);
        if (currentPlayers.some(player => player.username.toLowerCase() === $username.toLowerCase())) {
            errorMessage = "A player with that username already exists.";
            
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
</script>

<TextField
    bind:value={$username}
    bind:this={textfield}
    on:keydown={handleKeydown}
    class="shape-filled"
    variant="filled"
    style="width: 90%"
    label="Enter a username"
    >
    <Icon class="material-icons" role=button slot="trailingIcon" on:click={handleIconClick}>clear</Icon> 

</TextField>

<AddButton 
    on:click={handleAddPlayer}
    touch variant="unelevated"
    style="width: 40%"
    >
    <Label>Add Player</Label>
</AddButton>

<GetStatsButton
    on:click={fetchPlayerStats}
    color="secondary"
    touch variant="unelevated"
    style="width: 40%"
    >
    <Label>Get Stats</Label>
</GetStatsButton>

<WarningCard message={errorMessage}/>