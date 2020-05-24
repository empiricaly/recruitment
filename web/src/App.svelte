<script>
  import Router from "./lib/routing.js";
  import Home from "./pages/Home.svelte";
  import Welcome from "./pages/Welcome.svelte";
  import Link from "./components/base/Link.svelte";

  function create(node) {
    const router = new Router({
      target: node,
      mode: "history",
      routes: [
        {
          path: "/",
          component: Home
        }
      ]
    });

    return {
      destroy() {
        router.destroy();
      }
    };
  }
</script>

<main>
  <Link to="/">Home</Link>
  <div use:create />
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }
</style>
