<script>
    import Radio from '@smui/radio'
    import FormField from '@smui/form-field'
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

<div class="radio-grid">
    {#each stats as stat}
        <FormField>
            <Radio
                bind:group={$selectedStat}
                value={stat}
                on:change={handleSelect}
            />
            <span slot="label">
                {stat}
            </span>
        </FormField>
    {/each}
</div>

<style>
    .radio-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 10px;
    }
</style>