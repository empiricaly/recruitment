<script context="module">
  import Loader from "./components/base/Loader.svelte";
  import Alert from "./components/overlays/Alert.svelte";
  import Notification from "./components/overlays/Notification.svelte";
  import { loggingIn } from "./lib/auth.js";
  import { route } from "./lib/routing.js";
</script>

{#if $loggingIn}
  <main>
    <Loader />
  </main>
{:else}
  <main class="antialiased bg-gray-100" use:route />
{/if}
<Alert />
<Notification />

<style>
  main {
    width: 100vw;
    height: 100vh;
  }
</style>
