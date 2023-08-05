<script>
    import ClearAllButton, { Label } from "@smui/button"
    import { clearPlayers } from "../store";
</script>

<ClearAllButton on:click={clearPlayers}>
    <Label>Clear All</Label>
</ClearAllButton>