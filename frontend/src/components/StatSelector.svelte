<script>
    import Radio from '@smui/radio'
    import FormField from '@smui/form-field'
    import { model } from "../../wailsjs/go/models"
    import { selectedStats } from '../store';
    import { onMount } from 'svelte';

    let stats = [];

    onMount(() => {
        // Get the properties of the UserStats class
        stats = Object.keys(model.UserStats.prototype).filter(key => typeof model.UserStats.prototype[key] !== 'function');
    });

    function handleSelect(event) {
        // Update the store with the new selected game mode
        selectedStats.set(event.target.value);
    }
</script>

{#each stats as stat}
    <FormField>
        <Radio
            bind:group={$selectedStats}
            value={stat}
            on:change={handleSelect}
        />
        <span slot="label">
            {stat}
        </span>
    </FormField>
{/each}
