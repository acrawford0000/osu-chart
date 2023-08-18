<script>
    export let message = '';
    let showMessage = true;
    const hideMessage = () => {
        showMessage = false;
    }
</script>

<style>
    .warning-card {
        width: 90%;
        padding: 20px;
        margin: 10px auto;
        background-color: #f44336;
        color: white;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
        text-align: center;
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
