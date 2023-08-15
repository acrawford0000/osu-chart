<script>
    import PlayerList from '../components/PlayerList.svelte';
    import ClearAllButton from '../components/ClearAllButton.svelte';
    import TextField from '../components/AddPlayer.svelte';
    import { players, credentialsSet } from '../store';
    import { onMount } from "svelte";
    import { IsClientValid } from "../../wailsjs/go/app/App"
    import WarningCard from "../components/WarningCard.svelte";
    

    let showButton = false;

    $: showButton = $players.length > 0;

    let showWarning = false;

onMount(async () => {
    const isClientValid = await IsClientValid();
    if (!isClientValid && $credentialsSet) {
        // Client was not successfully created on startup
        showWarning = true;
    }
  });
</script>

<main>
    <div style="width: 100%;" class="add">
        <TextField/>
    </div>
        <PlayerList/>
        {#if showWarning}
            <WarningCard message="Warning: osu! API client was not successfully created on startup. Please check your API credentials in the settings." />
        {/if}
        {#if showButton}
            <ClearAllButton/>
        {/if}
</main>

<style>
    .add {
        align-self: center;
        padding: 20px;
        margin: 10px 0;
    }
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
</style>