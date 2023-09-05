<script>
    import LineChart from "../components/PlayerChart.svelte"
    import StatSelector from "../components/StatSelectorV2.svelte";
    import { isLoading } from "../store";
    import CircularProgress from "@smui/circular-progress"
    import LoadingScreen from "../components/LoadingScreen.svelte";

</script>

<main class="main">
    {#if $isLoading}
        <div class="loading-container">
            <LoadingScreen/>
            <div class="progress-container">
                <CircularProgress style="height: 64px; width: 64px;" indeterminate />
            </div>
        </div>
    {:else}
        <LineChart/>
        <StatSelector/>
    {/if}
</main>

<style>
    .main {
        align-self: center;
        padding: 20px;
        margin: 10px 0;
    }
    .loading-container {
        display: flex;
        justify-content: center;
        align-items: center;
        position: relative;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
    .progress-container {
        position: absolute;
    }
</style>