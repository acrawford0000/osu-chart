<script>
  import Sidebar from "./components/Sidebar.svelte";
  import WarningCard from "./components/WarningCard.svelte";
  import { initializeCredentials, warningMessage } from './store';
  import { onMount } from "svelte";
  import { Router, Route } from 'svelte-routing';
  import Interface from "./routes/Interface.svelte";
  import Settings from "./routes/Settings.svelte"
  import Chart from "./routes/Chart.svelte";
  import Contact from "./routes/Contact.svelte";
  
  onMount(() => { void initializeCredentials(); });
</script>

<Router>
  <main class="main">
    <Sidebar />
    <div class="page-content">
      <Route path="/contact" component={Contact} />
      <Route path="/chart" component={Chart} />
      <Route path="/settings" component={Settings} />
      <Route path="/" ><Interface /></Route>
    </div>
  </main>
</Router>

<WarningCard message={$warningMessage} />

<style>
  .main {
    display: flex;
  }
  .page-content {
    margin-left: var(--sidebar-width);
    min-height: 100vh;
    flex-grow: 1;
    min-width: 0;
  }
</style>
