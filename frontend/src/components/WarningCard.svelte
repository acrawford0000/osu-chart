<script>
    import { onDestroy } from "svelte";
    import { warningMessage } from "../store";
    export let message = '';
    let showMessage = true;
    const hideMessage = () => {
        showMessage = false;
        warningMessage.set('');
    }

    onDestroy(() => {
        warningMessage.set('');
    })

    $: if (message) {
        console.log('New message: ', message);
        showMessage = true;
    }
</script>

<style>
    .warning-card {
        position:fixed;
        bottom: 0;
        left: 50%;
        transform:translateX(-50%);
        width: 70%;
        max-width: 500px;
        padding: 20px;
        margin: 10px auto;
        background-color: #f44336;
        color: white;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
        text-align: center;
        z-index: 9999;
    }
    .close-button {
        float: right;
        cursor: pointer;
    }
</style>

{#if message && showMessage}
    <div class="warning-card">
        <span
            class="material-icons close-button"
            tabindex="0"
            role="button"
            on:click={hideMessage}
            on:keydown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                    hideMessage();
                }
            }}
        >close</span>
        {message}
    </div>
{/if}
