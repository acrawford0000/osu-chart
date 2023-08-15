<script>
    import Radio from '@smui/radio'
    import FormField from '@smui/form-field'
    import { SetGameMode } from "../../wailsjs/go/app/App"
    import { selectedMode } from '../store';

    let modes = [
        'standard', 'taiko', 'mania', 'catch'
    ]

    function handleSelect(event) {
        // Update the store with the new selected game mode
        selectedMode.set(event.target.value);
        SetGameMode(event.target.value);
    }
    
</script>

{#each modes as mode}
    <FormField>
        <Radio
            bind:group={$selectedMode}
            value={mode}
            on:change={handleSelect}
        />
        <span slot="label">
            {mode}
        </span>
    </FormField>
{/each}
