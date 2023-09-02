<script>
  import Sidebar from "./components/Sidebar.svelte";
  import WarningCard from "./components/WarningCard.svelte";
  import { clientId, clientSecret, credentialsSet, clientValid, warningMessage } from "./store";
  import { GetCredentials, AreOsuAuthCredentialsSet, CreateNewClient, IsClientValid } from "../wailsjs/go/app/App"
  import { onMount, onDestroy } from "svelte";
  import { Router, Route } from 'svelte-routing';
  import Interface from "./routes/Interface.svelte";
  import Settings from "./routes/Settings.svelte"
  import Chart from "./routes/Chart.svelte";
  import Contact from "./routes/Contact.svelte";
  
  let errorMessage = '';
  warningMessage.subscribe(value => {
    errorMessage = value;
  });

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
            warningMessage.set("An error occurred while creating a new client. Please verify your credentials are correct.");
        }
    }
  });
</script>

<main class="main">
  <Sidebar />
  <div class="page-content">
    <Router>
      <Route path="/contact" component={Contact} />
      <Route path="/chart" component={Chart} />
      <Route path="/settings" component={Settings} />
      <Route path="/" ><Interface /></Route>
    </Router>
  </div>

</main>

<WarningCard message={errorMessage} />

<style>
  .main {
    display: flex;
  }
  .page-content {
    margin-left: 4rem;
    height: 100vh;
    flex-grow: 1;
  }
</style>