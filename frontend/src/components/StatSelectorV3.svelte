<script>
    import { model } from "../../wailsjs/go/models"
    import { selectedStat } from '../store';
    import { onMount } from 'svelte';

    let stats = [];

    onMount(() => {
        // Create an instance of the UserStats class
        const userStats = new model.UserStats
        // Get the properties of the UserStats class
        stats = Object.keys(userStats).filter(key => typeof userStats[key] !== 'function');
    });

    function handleSelect(event) {
        // Update the store with the new selected game mode
        selectedStat.set(event.target.value);
        console.log("selectedStats:", $selectedStat);
    }
</script>

<div class="select">
    <label for="stats-select">Select Stat:</label>
    <select id="stats-select" bind:value={$selectedStat} on:change={handleSelect}>
        {#each stats as stat}
            <option value={stat}>{stat}</option>
        {/each}
    </select>
</div>

<style>
    .select {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 10px;
    }
    
    /* Dark mode styling */
    body {
        background-color: #121212;
        color: #fff;
    }
    
    select {
        background-color: #424242;
        color: #fff;
        border: none;
        padding: 8px 16px;
        border-radius: 4px;
    }
    
    option {
        background-color: #424242;
        color: #fff;
    }
</style>
