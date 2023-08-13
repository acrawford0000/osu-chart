<script>
    import { onMount } from 'svelte'
    import TextField from "@smui/textfield";
    import SaveButton, { Label } from "@smui/button"
    import { AreOsuAuthCredentialsSet, SetOsuAuthCredentials } from "../../wailsjs/go/app/App"

    let clientId = "";
    let clientSecret = "";
    let clientIdPlaceholder = "";
    let clientSecretPlaceholder = "";

    onMount(async () => {
        const areCredentialsSet = await AreOsuAuthCredentialsSet();
        if (areCredentialsSet) {
            // Display asterisks in the text fields
            clientIdPlaceholder = "*****";
            clientSecretPlaceholder = "*****";
        }
    });

    function onSave() {
        if (clientId === "" || clientSecret === "") {
            alert("Client ID and Client Secret must not be empty");
            return;
        }
        SetOsuAuthCredentials(clientId, clientSecret);
    }

    function clearCredentials() {
        SetOsuAuthCredentials("", "")
    }

</script>

<form on:submit|preventDefault={onSave}>
    <div style="display: flex; flex-direction: column;">
        <TextField 
            label="Client ID"
            placeholder={clientIdPlaceholder}

            bind:value={clientId}
        />
        <TextField 
            label="Client Secret"
            placeholder={clientSecretPlaceholder}

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