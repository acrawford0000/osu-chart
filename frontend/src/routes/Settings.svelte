<script>
    import { writable } from 'svelte/store';
    import ClientTextfields from "../components/ClientTextfields.svelte";
    import GameMode from "../components/GameMode.svelte";
    import StatSelector from '../components/StatSelector.svelte';

    const clientId = writable('');
    const clientSecret = writable('');
</script>

<main>
    <div style="display: flex; justify-content: center;">
        <div style="width: 50%;">
            <div class="section">
                <h2 class="section-title">Client Credentials</h2>
                <ClientTextfields/>
            </div>
            <div class="section">
                <h2 class="section-title">Game Mode</h2>
                <GameMode/>
            </div>
            <div class="section">
                <h2 class="section-title">Stats</h2>
                <StatSelector/>
            </div>
        </div>
    </div>
</main>

<style>
    .section {
        background-color: #272727;
        padding: 20px;
        margin: 10px 0;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    }
    
    .section-title {
        font-size: 24px;
        margin-bottom: 10px;
    }
</style>