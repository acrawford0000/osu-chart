<script>
    import ClearAllButton, { Label } from "@smui/button"
    import { clearPlayers } from "../store";
</script>

<ClearAllButton class="clear-all-button" on:click={clearPlayers}>
    <Label>Clear All</Label>
</ClearAllButton>
