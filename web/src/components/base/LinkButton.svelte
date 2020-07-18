<script>
  import Link from "./Link.svelte";

  export let to;
  export let full = false;
  export let secondary = false;
  export let replace = false;
  export let className = "";
  const activeClassName = "";

  $: css = secondary
    ? "text-indigo-700 bg-indigo-100 hover:bg-indigo-50 focus:border-indigo-300 focus:shadow-outline-indigo active:bg-indigo-200"
    : "text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700";
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
