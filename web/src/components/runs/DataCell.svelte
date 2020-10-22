<script>
  export let participant;
  export let key;

  const d = participant.data.find((d) => d.key === key);
</script>

{#if d}{d.val}{/if}
