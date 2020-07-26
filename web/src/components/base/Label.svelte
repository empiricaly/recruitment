<script>
  import Question from "./Question.svelte";

  export let forID = "";
  export let text;
  export let optional = false;
  export let white = false;
  export let question = null;
</script>

<label
  for={forID}
  class="mb-1 block text-sm font-medium leading-5 {white ? 'text-gray-50' : 'text-gray-700'}
  flex justify-between">
  <div>
    {text}
    {#if question}
      <Question text={question} />
    {/if}
  </div>
  {#if optional}
    <div class="text-cool-gray-300">optional</div>
  {/if}
</label>
