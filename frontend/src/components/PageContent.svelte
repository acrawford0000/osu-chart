<script>
    import { Router, Route } from 'svelte-routing';
    import Interface from "../routes/Interface.svelte";
    import Settings from "../routes/Settings.svelte"
    import Chart from "../routes/Chart.svelte";
    import Contact from "../routes/Contact.svelte";
</script>

<div class="page-content">
<Router>
    <Route path="/contact" component={Contact} />
    <Route path="/chart" component={Chart} />
    <Route path="/settings" component={Settings} />
    <Route path="/" ><Interface /></Route>
</Router>
</div>
  
<style>
    .page-content {
        margin-left: 4rem;
        height: 100vh;
    }
</style>