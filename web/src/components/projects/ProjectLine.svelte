<script>
  import Link from "../base/Link.svelte";

  export let name;
  export let id;
</script>

<li class="mt-2">
  <Link
    to="/projects/{id}/overview"
    className="flex px-2 py-2 bg-gray-50 text-gray-700 rounded
    hover:bg-gray-200 hover:text-gray-900 border-mint-200 border">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 448 512"
      width="12"
      fill="currentColor"
      class="ml-1 mr-2 text-mint-600"><path
        d="M218.101 38.101L198.302 57.9c-4.686 4.686-4.686 12.284 0 16.971L353.432 230H12c-6.627 0-12 5.373-12 12v28c0 6.627 5.373 12 12 12h341.432l-155.13 155.13c-4.686 4.686-4.686 12.284 0 16.971l19.799 19.799c4.686 4.686 12.284 4.686 16.971 0l209.414-209.414c4.686-4.686 4.686-12.284 0-16.971L235.071 38.101c-4.686-4.687-12.284-4.687-16.97 0z" /></svg>
    {name}
  </Link>
</li>
