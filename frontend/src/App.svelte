<script>
  import { Router, Route, navigate } from "svelte-routing";
  import Interface from "./routes/Interface.svelte";
  import Settings from "./routes/Settings.svelte"
  import Chart from "./routes/Chart.svelte";
  import Tab, { Icon, Label } from '@smui/tab';
  import TabBar from '@smui/tab-bar';
  import WarningCard from "./components/WarningCard.svelte";
  import { clientId, clientSecret, credentialsSet, clientValid } from "./store";
  import { GetCredentials, AreOsuAuthCredentialsSet, CreateNewClient, IsClientValid } from "../wailsjs/go/app/App"
  import { onMount } from "svelte";
 
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
  export let url = "/";
  let tabs = [
    {
      icon: 'list',
      label: 'Players',
      path: '/'
    },
    {
      icon: 'settings',
      label: 'Settings',
      path: '/settings'
    },
    {
      icon: 'stacked_line_chart',
      label: 'Chart',
      path: '/chart'
    },
  ];
  
  let active = tabs[0];
  
  function handleTabClick(tab) {
    active = tab;
    navigate(tab.path);
  }

</script>


<Router {url}>
  <nav>
    <TabBar {tabs} let:tab bind:active>
      <Tab {tab} on:click={() => handleTabClick(tab)}>
        <Icon class="material-icons">{tab.icon}</Icon>
        <Label>{tab.label}</Label>
      </Tab>
    </TabBar>
  </nav>
  <div>
    <Route path="/chart" component={Chart} />
    <Route path="/settings" component={Settings} />
    <Route path="/" ><Interface /></Route>
  </div>
</Router>

<WarningCard message={errorMessage}/>