<script>
  import { onMount } from "svelte";
  import { getPath, listen, push, replace as rplc } from "../../lib/routing.js";

  export let to;
  export let replace = false;
  export let className = "";
  export let activeClassName = "router-link-active";

  let href = to;
  let computedClassName = className;

  const handleNavigate = (e) => {
    e.preventDefault();
    const fnc = replace ? rplc : push;
    fnc(to);
  };

  function computeClassName() {
    if (getPath() === to) {
      computedClassName = `${className} ${activeClassName}`;
    } else {
      computedClassName = className;
    }
  }

  onMount(() => {
    computeClassName();
    listen(computeClassName);
    href = to;
  });
</script>

<a class={computedClassName} {href} on:click={handleNavigate}>
  <slot />
</a>
