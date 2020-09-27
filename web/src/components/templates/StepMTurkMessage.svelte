<script>
  import StepMessageArgs from "./StepMessageArgs.svelte";

  export let step;

  console.log("stepstepstepstep", step);
</script>

<StepMessageArgs bind:msgArgs={step.msgArgs} />
