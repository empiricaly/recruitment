<script>
  import Link from "../base/Link.svelte";

  export let name;
  export let id;
</script>

<li class="mt-2">
  <Link
    to="/projects/{id}/overview"
    className="block px-2 py-2 bg-gray-50 text-gray-700 rounded-md
    hover:bg-gray-200 hover:text-gray-900">
    {name}
  </Link>
</li>
