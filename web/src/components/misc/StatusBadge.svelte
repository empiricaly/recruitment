<script context="module">
  const statusColors = {
    created: "gray",
    running: "green",
    done: "green",
    terminated: "orange",
    failed: "red"
  };
  const statusText = {
    created: "Created",
    running: "Running",
    done: "Finished",
    terminated: "Canceled",
    failed: "Failed!"
  };
  const statusIcon = {
    created: `<path d="M11 11V7h2v4h4v2h-4v4h-2v-4H7v-2h4zm1 11C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16z"/>`,
    running: `<path d="M16.394 12L10 7.737v8.526L16.394 12zm2.982.416L8.777 19.482A.5.5 0 0 1 8 19.066V4.934a.5.5 0 0 1 .777-.416l10.599 7.066a.5.5 0 0 1 0 .832z"/>`,
    done: `<path d="M10 15.172l9.192-9.193 1.415 1.414L10 18l-6.364-6.364 1.414-1.414z"/>`,
    terminated: `<path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/>`,
    failed: `<path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm-1-5h2v2h-2v-2zm0-8h2v6h-2V7z"/>`
  };
</script>

<script>
  import Badge from "../base/Badge.svelte";
  export let status;
</script>

<Badge
  text={statusText[status]}
  color={statusColors[status]}
  icon={statusIcon[status]} />
