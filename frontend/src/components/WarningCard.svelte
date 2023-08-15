<script>
    export let message = '';
</script>

<style>
    .warning-card {
        height: 75%;
        width: 95%;
        padding: 20px;
        margin: 10px auto;
        background-color: #f44336;
        color: white;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
        text-align: center;
        z-index: 9999;
    }
</style>

{#if message}
    <div class="warning-card">
        {message}
    </div>
{/if}
