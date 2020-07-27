<script>
  import CodeMirror from "../editors/CodeMirror.svelte";

  export let step;

  $: filterArgs = step.args[0];
</script>

<div class="border">
  <CodeMirror bind:value={filterArgs.js} mode="javascript" />
</div>
