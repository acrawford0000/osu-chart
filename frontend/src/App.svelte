<script>
  import Sidebar from "./components/Sidebar.svelte";
  import PageContent from "./components/PageContent.svelte";
  import WarningCard from "./components/WarningCard.svelte";
  import { clientId, clientSecret, credentialsSet, clientValid } from "./store";
  import { GetCredentials, AreOsuAuthCredentialsSet, CreateNewClient, IsClientValid } from "../wailsjs/go/app/App"
  import { onMount } from "svelte";
  import { EventsOn, EventsEmit } from '../wailsjs/runtime'
 
  onMount(async () => {
    


    if (await AreOsuAuthCredentialsSet()) {
        try {
            const data = await GetCredentials();
            clientId.set(data.client_id);
            clientSecret.set(data.client_secret);
            credentialsSet.set(true);
        } catch (err) {
            console.error('Failed to get credentials:', err);
        }
    }
    if (await AreOsuAuthCredentialsSet() && !(await IsClientValid())) {
        try {
            await CreateNewClient();
        } catch (err) {
            console.error('Failed to create new client:', err);
            errorMessage = "An error occurred while creating a new client. Please verify your credentials are correct."
        }
    }
});

  let errorMessage = '';
</script>

<main>
  <Sidebar />
  <PageContent />

</main>

<WarningCard message={errorMessage}/>

<style>
  .main {
    display: flex;
  }
</style>