<script>
    import { writable } from 'svelte/store';
    import ClientTextfields from "../components/ClientTextfields.svelte";
    import StatSelector from '../components/StatSelector.svelte';
    import Instructions, { Label } from "@smui/button";

    const clientId = writable('');
    const clientSecret = writable('');
    let showInstructions = false;
    function toggleInstructions() {
        showInstructions = !showInstructions;
    }
</script>

<main>
    <div style="display: flex; justify-content: space-evenly;">
            <div class="section" style="width: 20%;">
                <h2 class="section-title">Client Credentials</h2>
                <ClientTextfields/>
            </div>
            <div class="section" style="width: 50%;">
                <div>
                    <Instructions
                        on:click={toggleInstructions}
                        color="secondary"
                        touch variant="unelevated"
                        >
                        <Label>Show Instructions</Label>
                    </Instructions>
                    {#if showInstructions}
                    <div>
                        <ol>
                        <li>Make sure you are signed into the osu! site.</li>
                        <li>Click this link: <a href="https://osu.ppy.sh/home/account/edit#new-oauth-application">https://osu.ppy.sh/home/account/edit#new-oauth-application</a>.</li>
                        <li>Click on "New OAuth Application".</li>
                        <li>Enter "osu!-chart" as the name and leave the "Application Callback URLs" field blank.</li>
                        <li>Click "Register Application".</li>
                        <li>Copy and paste the "Client ID" and "Client Secret". Make sure not to share these.</li>
                        </ol>
                    </div>
                    {/if}
                </div>
            </div>
        </div>
        <div class="section">
            <h2 class="section-title">Stats</h2>
            <StatSelector/>
        </div>
</main>

<style>
    .section {
        background-color: #1e1f22;
        padding: 3rem;
        margin: 10px;
        height: 300px;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
    }
    
    .section-title {
        font-size: 24px;
        margin-bottom: 10px;
    }
    ol {
        text-align: left;
    }
</style>