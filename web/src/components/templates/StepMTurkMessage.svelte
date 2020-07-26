<script>
  import StepMessageArgs from "./StepMessageArgs.svelte";
  export let step;

  $: msgArgs = step.args[0];
</script>

<StepMessageArgs bind:msgArgs />
