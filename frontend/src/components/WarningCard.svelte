<script>
  import { warningMessage } from '../store';
  import AppIcon from './AppIcon.svelte';
  export let message = '';
  function dismiss() { if ($warningMessage === message) warningMessage.set(''); message = ''; }
</script>
{#if message}<div class="warning-card" role="alert"><span class="warning-symbol">!</span><span>{message}</span><button class="icon-button" aria-label="Dismiss notification" on:click={dismiss}><AppIcon name="close" size={16}/></button></div>{/if}
<style>
  .warning-card { position: fixed; bottom: 22px; left: calc(var(--sidebar-width) + 24px); right: 24px; margin: 0 auto; max-width: 620px; padding: 13px 13px 13px 17px; border: 1px solid #ff95a545; background: #281f2c; color: #f7c5ce; border-radius: 12px; box-shadow: 0 12px 50px #0005; display: flex; align-items: center; gap: 12px; text-align: left; font-size: 12px; line-height: 1.6; z-index: 100; }
  .warning-symbol { display: grid; place-items: center; width: 21px; height: 21px; flex-shrink: 0; border-radius: 50%; background: #ff95a51c; color: var(--danger); font-weight: 700; }
  button { margin-left: auto; }
</style>
