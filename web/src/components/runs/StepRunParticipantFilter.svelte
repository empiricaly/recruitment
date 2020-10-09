<script>
  import CodeMirror from "../editors/CodeMirror.svelte";

  export let step;
</script>

<div class="border">
  <CodeMirror bind:value={step.filterArgs.js} mode="javascript" />
</div>
