<script>
  import { Router, Route, navigate } from "svelte-routing";
  import Interface from "./routes/Interface.svelte";
  import Settings from "./routes/Settings.svelte"
  import Chart from "./routes/Chart.svelte";
  import Tab, { Icon, Label } from '@smui/tab';
  import TabBar from '@smui/tab-bar';
  
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

  export let url = "/";
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

<!-- <div class="blurredBg"/>

<main>
  <TextField/>

  <IconButton/>

  <List/>

  <ClearAllButton/>

  <GameMode/>

  <ClientTextfields/>

</main>

<style>

  :blurredBg {
    background-image: url({bg});
    background-size: cover;
    background-position: center;
    filter: blur(5px);
  }

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
 -->