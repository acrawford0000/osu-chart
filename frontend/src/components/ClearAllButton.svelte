<script>
  import { clearPlayers } from '../store';
  import AppIcon from './AppIcon.svelte';
</script>
<button class="button-quiet button-danger" on:click={clearPlayers}><AppIcon name="close" size={14}/>Clear All</button>
