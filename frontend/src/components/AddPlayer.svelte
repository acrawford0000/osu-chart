<script>
    import TextField from "@smui/textfield";
    import Icon from '@smui/textfield/icon'
    import AddButton, { Label } from "@smui/button"
    import GetStatsButton from "@smui/button"
    import WarningCard from "./WarningCard.svelte";
    import { username, addPlayer, fetchPlayerStats } from "../store";
    import { IsClientValid, AreOsuAuthCredentialsSet } from "../../wailsjs/go/app/App"

    let textfield;
    let errorMessage = '';

    function handleIconClick() {
        handleAddPlayer;
        textfield.focus();
    }

    async function handleAddPlayer() {
        const areCredentialsSet = await AreOsuAuthCredentialsSet();
        if (!areCredentialsSet) {
            errorMessage = "Please set your API credentials in the settings before adding a player.";
            return;
        }
        const isClientValid = await IsClientValid();
        if (!isClientValid) {
            errorMessage = "Client is invalid. Please check your credentials and try again.";
            return;
        }
        addPlayer($username);
        username.set('');
    }

    function handleKeydown(event) {
        if (event.key === 'Enter') {
            handleAddPlayer();
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