<script>
  import { onMount } from "svelte";
  import Router from "../../lib/routing.js";

  export let to;
  export let replace = false;
  export let className = "";
  export let activeClassName = "router-link-active";

  let href = "";
  let computedClassName = className;

  const handleNavigate = e => {
    e.preventDefault();
    Router[replace ? "replace" : "push"](to);
  };

  function computeClassName() {
    if (Router.history.location.pathname === to) {
      computedClassName = `${className} ${activeClassName}`;
    } else {
      computedClassName = className;
    }
  }

  onMount(() => {
    computeClassName();
    Router.listen(computeClassName);
    href = Router.mode === "hash" ? `/#${to}` : to;
  });
</script>

<a class={computedClassName} {href} on:click={handleNavigate}>
  <slot />
</a>
