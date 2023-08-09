<script>
    import TextField from "@smui/textfield";
    import { AreOsuAuthCredentialsSet } from "wailsjs/go/app/App"

    let clientIdTextField;
    let clientSecretTextField;

    const areCredentialsSet = AreOsuAuthCredentialsSet();

    if (areCredentialsSet) {
        // Display asterisks in the text fields
        clientIdTextField.placeholder = "*****";
        clientSecretTextField.placeholder = "*****";
    }

</script>

<TextField bind:this={clientIdTextField} label="Client ID" />
<TextField bind:this={clientSecretTextField} label="Client Secret" />