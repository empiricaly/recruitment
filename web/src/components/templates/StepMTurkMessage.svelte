<script>
  import StepMessageArgs from "./StepMessageArgs.svelte";

  export let step;
</script>

<StepMessageArgs bind:msgArgs={step.msgArgs} />
