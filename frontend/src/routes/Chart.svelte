<script>
    import LineChart from "../components/PlayerChart.svelte"
    import StatSelector from "../components/StatSelector.svelte";
    import { isLoading } from "../store";
    import CircularProgress from "@smui/circular-progress"

</script>

<main class="main">
    {#if $isLoading}
        <div class="loading-container">
            <CircularProgress style="height: 32px; width: 32px;" indeterminate />
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
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
    .loading-container::before {
        content: "";
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-image: linear-gradient(to bottom right, transparent, #FFC0CB);
        animation: wipe 2s linear infinite;
    }
    @keyframes wipe {
        from { transform: translateX(-100%); }
        to { transform: translateX(100%); }
    }
</style>