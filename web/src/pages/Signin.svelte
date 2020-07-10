<script>
  import Input from "../components/form/Input.svelte";
  import Button from "../components/form/Button.svelte";

  let email = "";
  let password = "";

  function submit(event) {
    event.preventDefault();
    console.log("submit", email, password);
  }
</script>

<section class="w-full h-full flex flex-col justify-center items-center">
  <form
    action="#"
    class="w-128 shadow-xl max-w-xs rounded-md border"
    on:submit={submit}>
    <h1
      class="p-8 py-4 bg-gray-800 text-gray-100 text-xl font-bold bg-color
      rounded-t-md font-mono text-center">
      Please authenticate
    </h1>

    <div class="bg-gray-700 p-8 py-10 rounded-b-md">
      <div>
        <Input
          bind:value={email}
          type="email"
          focus
          required
          placeholder="Email" />
      </div>

      <div class="mt-4">
        <Input
          bind:value={password}
          type="password"
          required
          placeholder="Password" />
      </div>

      <div class="mt-4">
        <Button type="submit" text="Sign in" />
      </div>
    </div>
  </form>
</section>
