<script>
  import ProjectLine from "../components/projects/ProjectLine.svelte";
  import Label from "../components/form/Label.svelte";
  import Input from "../components/form/Input.svelte";
  import Button from "../components/form/Button.svelte";
  import Callout from "../components/base/Callout.svelte";

  let newProject = false;
  let name = "";

  const projects = [
    {
      id: "njeqqlqw",
      name: "Speed Dating"
    },
    {
      id: "k4l32lld",
      name: "Networks and Reputation and Networks and Reputation"
    }
  ];

  function handleNewProject(event) {
    event.preventDefault();
  }
</script>

<main class="flex justify-center items-center h-full w-full">

  {#if projects.length === 0}
    <div class="w-64 px-4 py-4 rounded-md">
      <Callout color="yellow">
        You have no projects yet.
        <br />
        Create one now!
      </Callout>
    </div>
  {/if}
  <div class="w-64 px-4 py-4 rounded-md max-h-full overflow-auto">
    {#if newProject || projects.length === 0}
      <h1 class="font-semibold">New Project</h1>
      <form class="mt-3" on:submit={handleNewProject}>
        <!-- <Label forID="name" text="New Project" /> -->
        <Input
          focus
          bind:value={name}
          id="name"
          required
          placeholder="Enter Project Name" />
        <div class="mt-4">
          <Button type="submit" text="Create Project" full />
        </div>
        {#if projects.length !== 0}
          <div class="mt-2">
            <Button
              on:click={() => (newProject = false)}
              text="Cancel"
              full
              secondary />
          </div>
        {/if}
      </form>
    {:else}
      <h1 class="font-semibold pl-2">Projects</h1>
      <ul class="mt-3">
        {#each projects as project}
          <ProjectLine id={project.id} name={project.name} />
        {/each}
      </ul>

      <div class="mt-4">
        <Button
          on:click={() => (newProject = true)}
          text="Create a new Project"
          full />
      </div>
    {/if}

  </div>
</main>
