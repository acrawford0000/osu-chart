<script>
    import TextField from "@smui/textfield";
    import Icon from '@smui/textfield/icon'
    import AddButton, { Label } from "@smui/button"
    import { addPlayer } from "../store";
    import { writable } from "svelte/store";


export const username = writable('');
let textfield;

function handleIconClick() {
        username.set('');
        textfield.focus();
    }

    function handleAddPlayer() {
        addPlayer($username);
        username.set('');
    }

    function handleKeyDown(event) {
        if (event.key === 'Enter') {
            handleAddPlayer();
        }
    }

</script>

<TextField
    bind:value={$username}
    bind:this={textfield}
    on:keydown={handleKeyDown}
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