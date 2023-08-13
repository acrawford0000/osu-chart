<script>
    import { writable } from 'svelte/store';
    import ClientTextfields from "../components/ClientTextfields.svelte";
    import GameMode from "../components/GameMode.svelte";

    const clientId = writable('');
    const clientSecret = writable('');
</script>

<main>
    <div style="display: flex; justify-content: center;">
        <div style="width: 50%;">
            <ClientTextfields/>
        </div>
    </div>
    
    <GameMode/>
</main>