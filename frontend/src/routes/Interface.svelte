<script>
    import PlayerList from '../components/PlayerList.svelte';
    import ClearAllButton from '../components/ClearAllButton.svelte';
    import TextField, { username } from '../components/AddPlayer.svelte';
    import { players } from '../store';
    

    let showButton = false;

    $: showButton = $players.length > 0;
</script>

<main>
    
    <TextField/>
  
    <PlayerList/>

    {#if showButton}
        <ClearAllButton/>
    {/if}


</main>

