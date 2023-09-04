<script>
    import LineChart from "../components/PlayerChart.svelte"
    import StatSelector from "../components/StatSelector.svelte";

</script>

<main class="main">
    <LineChart/>
    <StatSelector/>
</main>

<style>
    .main {
        align-self: center;
        padding: 20px;
        margin: 10px 0;
    }
</style>