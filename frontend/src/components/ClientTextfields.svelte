<script>
    import TextField from "@smui/textfield";
    import SaveButton, { Label } from "@smui/button"
    import Icon from "@smui/textfield/icon";
    import { SaveOsuAuthCredentials } from "../../wailsjs/go/app/App"
    import { credentialsSet, clientValid, clientId, clientSecret } from '../store';
    import WarningCard from "./WarningCard.svelte";

    let errorMessage = '';
    let showSecret = false;

    async function onSave() {
        if ($clientValid) {
            return;
        }
        if ($clientId === "" || $clientSecret === "") {
            errorMessage = "Client ID and Client Secret must not be empty";
            return;
        }
        try {
            await SaveOsuAuthCredentials($clientId, $clientSecret);
            // Credentials were successfully saved and a new client was created
            credentialsSet.set(true);
            clientValid.set(true);
        } catch (error) {
            // An error occurred while saving the credentials or creating a new client
            console.error(error);
            errorMessage = "An error occurred while saving your API credentials: " + error.message;
        }
    }

    function toggleShowSecret() {
        showSecret = !showSecret;
    }

</script>

<form on:submit|preventDefault={onSave}>
    <div style="display: flex; flex-direction: column;">
        <WarningCard bind:message={errorMessage} />
        <TextField 
            label="Client ID"
            bind:value={$clientId}
        />
            <TextField 
                label="Client Secret"
                type={showSecret ? 'text' : 'password'}
                bind:value={$clientSecret} 
            >
            
            <Icon class="material-icons" role=button slot="trailingIcon" on:click={toggleShowSecret}>{showSecret ? 'visibility_off' : 'visibility'}</Icon>
            
            </TextField>

        <SaveButton
            type="submit"
            touch variant="unelevated"
        >
        <Label>Save</Label>
        </SaveButton>
    </div>
    
</form>