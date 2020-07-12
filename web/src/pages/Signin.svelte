<script>
  import Label from "../components/form/Label.svelte";
  import Input from "../components/form/Input.svelte";
  import Button from "../components/form/Button.svelte";

  let username = "";
  let password = "";

  function submit(event) {
    event.preventDefault();
    console.log("submit", username, password);
  }
</script>

<div class="min-h-screen bg-white flex">
  <div
    class="flex-1 flex flex-col justify-center py-12 px-4 sm:px-6 lg:flex-none
    lg:px-20 xl:px-24">
    <div class="mx-auto w-full max-w-sm">
      <div>
        <!-- <img
          class="h-12 w-auto"
          src="/img/logos/workflow-mark-on-white.svg"
          alt="Workflow" /> -->
        <h2 class="mt-6 text-3xl leading-9 font-extrabold text-gray-900 w-64">
          Sign in
        </h2>
      </div>

      <div class="mt-8">

        <div class="mt-6">
          <form action="#" method="POST" on:submit={submit}>
            <div>
              <Label forID="username" text="Username" />
              <Input
                bind:value={username}
                id="username"
                autocomplete="username"
                focus
                required
                placeholder="Enter your username" />
            </div>

            <div class="mt-6">
              <Label forID="password" text="Password" />
              <Input
                bind:value={password}
                id="password"
                autocomplete="current-password"
                type="password"
                required
                placeholder="Enter your password" />
            </div>

            <div class="mt-8">
              <Button type="submit" text="Sign in" full />
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
  <div class="hidden lg:block relative w-0 flex-1 bg-gray-100">
    <!-- <img
      class="absolute inset-0 h-full w-full object-cover"
      src="https://source.unsplash.com/7eOm_tJMSPE/2503x1633"
      alt="" /> -->
  </div>
</div>
