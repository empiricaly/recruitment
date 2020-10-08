<script>
  import Link from "./Link.svelte";

  export let to;
  export let full = false;
  export let secondary = false;
  export let replace = false;
  export let className = "";
  const activeClassName = "";

  $: css = secondary
    ? "text-mint-700 bg-mint-100 hover:bg-mint-50 focus:border-mint-300 focus:shadow-outline-mint active:bg-mint-200"
    : "text-white bg-mint-600 hover:bg-mint-500 focus:outline-none focus:border-mint-700 focus:shadow-outline-mint active:bg-mint-700";
</script>

<span class:w-full={full} class="block rounded-md shadow-sm">
  <Link
    {to}
    {replace}
    className="{css}
    {className} flex w-full justify-center px-4 py-2 border border-transparent
    text-sm font-medium rounded-md transition duration-150 ease-in-out">
    <slot />
  </Link>
</span>
