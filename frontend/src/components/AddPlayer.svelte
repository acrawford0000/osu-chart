<script>
    import TextField from "@smui/textfield";
    import Icon from '@smui/textfield/icon'
    import AddButton, { Label } from "@smui/button"
    import WarningCard from "./WarningCard.svelte";
    import { username, addPlayer, credentialsSet, isClientValid, clientValid } from "../store";
    import { onMount } from "svelte";

    let textfield;
    let errorMessage = '';

    function handleIconClick() {
        username.set('');
        textfield.focus();
    }

    onMount(async () => {
        await isClientValid;
    });

    async function handleAddPlayer() {
        await isClientValid();
        if (!$credentialsSet) {
            alert("Please set your API credentials in the settings before adding a player.");
            return;
        }
        if (!$clientValid) {
            alert("Client is invalid. Please check your credentials and try again.");
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
    style="width: 80%"
    >
    <Label>Add Player</Label>
</AddButton>

<WarningCard message={errorMessage}/>