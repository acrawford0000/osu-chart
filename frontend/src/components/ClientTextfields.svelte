<script>
    import TextField from "@smui/textfield";
    import SaveButton, { Label } from "@smui/button"
    import { SaveOsuAuthCredentials } from "../../wailsjs/go/app/App"
    import { credentialsSet } from '../store';

    let clientId = "";
    let clientSecret = "";

    async function onSave() {
        if (clientId === "" || clientSecret === "") {
            alert("Client ID and Client Secret must not be empty");
            return;
        }
        try {
            await SaveOsuAuthCredentials(clientId, clientSecret);
            // Credentials were successfully saved and a new client was created
            credentialsSet.set(true);
        } catch (error) {
            // An error occurred while saving the credentials or creating a new client
            console.error(error);
            alert("An error occurred while saving your API credentials: " + error.message);
        }
    }


</script>

<form on:submit|preventDefault={onSave}>
    <div style="display: flex; flex-direction: column;">
        <TextField 
            label="Client ID"
            bind:value={clientId}
        />
        <TextField 
            label="Client Secret"
            bind:value={clientSecret} 

        />
        <SaveButton
            type="submit"
            touch variant="unelevated"
        >
        <Label>Save</Label>
        </SaveButton>


    </div>
    
</form>