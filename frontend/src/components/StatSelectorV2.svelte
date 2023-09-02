<script>
    import Select, { Option } from '@smui/select';
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
        selectedStat.set(event.detail);
        console.log("selectedStats:", $selectedStat);
    }
</script>

<div class="select">
    <Select 
    bind:value={$selectedStat} 
    label="Select Stat" 
    on:change={handleSelect}>
        {#each stats as stat}
            <Option value={stat}>{stat}</Option>
        {/each}
    </Select>
</div>

<style>
    .select{
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 10px;
    }
</style>
