<script>
  import StepRunMessageArgs from "./StepRunMessageArgs.svelte";

  export let step;
</script>

<StepRunMessageArgs msgArgs={step.msgArgs} hasSubject />
