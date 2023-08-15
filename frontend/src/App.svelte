<script>
  import { Router, Route, navigate } from "svelte-routing";
  import Interface from "./routes/Interface.svelte";
  import Settings from "./routes/Settings.svelte"
  import Chart from "./routes/Chart.svelte";
  import Tab, { Icon, Label } from '@smui/tab';
  import TabBar from '@smui/tab-bar';

  
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
