<script>
    import Radio from '@smui/radio'
    import FormField from '@smui/form-field'
    import { SetGameMode } from "../../wailsjs/go/app/App"

    let selectedMode = 'osu';

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


<!-- <FormField>
    <Radio
      bind:group={selectedMode}
      value="osu"
      on:change={handleSelect} 
    />
    <span slot="label">osu!standard</span>
  </FormField>

  <FormField>
    <Radio
      bind:group={selectedMode}
      value="taiko"
      on:change={handleSelect}
    />
    <span slot="label">osu!taiko</span>
  </FormField>

  <FormField>
    <Radio
      bind:group={selectedMode}
      value="mania"
      on:change={handleSelect}
    />
    <span slot="label">osu!mania</span>
  </FormField>

  <FormField>
    <Radio
      bind:group={selectedMode}
      value="fruits"
      on:change={handleSelect}
    />
    <span slot="label">osu!catch</span>
  </FormField> -->