<script>
    import Select, { Option } from '@smui/select';
    import { model } from "../../wailsjs/go/models";
    import { selectedStat } from '../store';

    let stats = [];
    // Create an instance of the UserStats class
    const userStats = new model.UserStats;
    // Get the properties of the UserStats class
    stats = Object.keys(userStats).filter(key => typeof userStats[key] !== 'function');

    let selectedValue = $selectedStat;
    $: selectedStat.set(selectedValue);
</script>
<div style="background-color: #1e1f22; padding-top: 10px; padding-bottom: 10px; border-bottom-left-radius: 10px; border-bottom-right-radius: 10px;">
    <Select style="background-color: #2c2f33" bind:value={selectedValue}
    >
        {#each stats as stat}
            <Option value={stat}>{stat}</Option>
        {/each}
    </Select>
</div>
