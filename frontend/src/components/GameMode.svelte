<script>
    import Radio from '@smui/radio'
    import FormField from '@smui/form-field'
    import { SetGameMode } from "../../wailsjs/go/app/App"

    let selectedMode = 'standard';

    let modes = [
        'standard', 'taiko', 'mania', 'catch'
    ]

    function handleSelect(event) {
        selectedMode = event.target.value;
        SetGameMode(selectedMode);
    }
    
</script>

{#each modes as mode}
    <FormField>
        <Radio
            bind:group={selectedMode}
            value={mode}
        />
        <span slot="label">
            {mode}
        </span>
    </FormField>
{/each}
