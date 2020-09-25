<script context="module">
  import { pathToRegexp } from "path-to-regexp";
  import { query } from "svelte-apollo";
  import { quadInOut } from "svelte/easing";
  import { fade, fly } from "svelte/transition";
  import Link from "../components/base/Link.svelte";
  import Loader from "../components/base/Loader.svelte";
  import Logo from "../components/base/Logo.svelte";
  import { client } from "../lib/apollo";
  import { user } from "../lib/auth.js";
  import { GET_PROJECT } from "../lib/queries";
  import { params, path } from "../lib/routing.js";
  import Header from "./Header.svelte";

  const regexes = {
    overview: pathToRegexp("/projects/:project_id/overview"),
    runs: pathToRegexp("/projects/:project_id/runs/:run_id?"),
    templates: pathToRegexp("/projects/:project_id/templates/:template_id?"),
    participants: pathToRegexp(
      "/projects/:project_id/participants/:participant_id?"
    ),
  };
</script>

<script>
  export let title = null;
  export let titleUpdatable = false;
  export let overtitle = null;
  export let action = null;
  export let facts = [];
  export let actions = [];

  let sidebarOpen = false;

  $: projectID = $params.projectID;

  $: project =
    projectID &&
    query(client, { query: GET_PROJECT, variables: { projectID } });

  $: menuitems = projectID
    ? [
        {
          title: "Overview",
          path: `/projects/${projectID}/overview`,
          regex: regexes.overview,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M573.48 219.91L310.6 8a35.85 35.85 0 0 0-45.19 0L2.53 219.91a6.71 6.71 0 0 0-1 9.5l14.2 17.5a6.82 6.82 0 0 0 9.6 1L64 216.72V496a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16V216.82l38.8 31.29a6.83 6.83 0 0 0 9.6-1l14.19-17.5a7.14 7.14 0 0 0-1.11-9.7zM240 480V320h96v160zm240 0H368V304a16 16 0 0 0-16-16H224a16 16 0 0 0-16 16v176H96V190.92l187.71-151.4a6.63 6.63 0 0 1 8.4 0L480 191z"/></svg>`,
        },
        {
          title: "Runs",
          path: `/projects/${projectID}/runs`,
          regex: regexes.runs,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M145.35 207a8 8 0 0 0-11.35 0l-71 71-39-39a8 8 0 0 0-11.31 0L1.35 250.34a8 8 0 0 0 0 11.32l56 56a8 8 0 0 0 11.31 0l88-88a8 8 0 0 0 0-11.32zM62.93 384c-17.67 0-32.4 14.33-32.4 32s14.73 32 32.4 32a32 32 0 0 0 0-64zm82.42-337A8 8 0 0 0 134 47l-71 71-39-39a8 8 0 0 0-11.31 0L1.35 90.34a8 8 0 0 0 0 11.32l56 56a8 8 0 0 0 11.31 0l88-88a8 8 0 0 0 0-11.32zM503 400H199a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h304a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8zm0-320H199a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h304a8 8 0 0 0 8-8V88a8 8 0 0 0-8-8zm0 160H199a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h304a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8z"/></svg>`,
        },
        {
          title: "Templates",
          path: `/projects/${projectID}/templates`,
          regex: regexes.templates,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M417.8 315.5l20-20c3.8-3.8 10.2-1.1 10.2 4.2V464c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V112c0-26.5 21.5-48 48-48h292.3c5.3 0 8 6.5 4.2 10.2l-20 20c-1.1 1.1-2.7 1.8-4.2 1.8H48c-8.8 0-16 7.2-16 16v352c0 8.8 7.2 16 16 16h352c8.8 0 16-7.2 16-16V319.7c0-1.6.6-3.1 1.8-4.2zm145.9-191.2L251.2 436.8l-99.9 11.1c-13.4 1.5-24.7-9.8-23.2-23.2l11.1-99.9L451.7 12.3c16.4-16.4 43-16.4 59.4 0l52.6 52.6c16.4 16.4 16.4 43 0 59.4zm-93.6 48.4L403.4 106 169.8 339.5l-8.3 75.1 75.1-8.3 233.5-233.6zm71-85.2l-52.6-52.6c-3.8-3.8-10.2-4-14.1 0L426 83.3l66.7 66.7 48.4-48.4c3.9-3.8 3.9-10.2 0-14.1z"/></svg>`,
        },
        {
          title: "Participants",
          path: `/projects/${projectID}/participants`,
          regex: regexes.participants,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"><path d="M544 224c44.2 0 80-35.8 80-80s-35.8-80-80-80-80 35.8-80 80 35.8 80 80 80zm0-128c26.5 0 48 21.5 48 48s-21.5 48-48 48-48-21.5-48-48 21.5-48 48-48zM320 256c61.9 0 112-50.1 112-112S381.9 32 320 32 208 82.1 208 144s50.1 112 112 112zm0-192c44.1 0 80 35.9 80 80s-35.9 80-80 80-80-35.9-80-80 35.9-80 80-80zm244 192h-40c-15.2 0-29.3 4.8-41.1 12.9 9.4 6.4 17.9 13.9 25.4 22.4 4.9-2.1 10.2-3.3 15.7-3.3h40c24.2 0 44 21.5 44 48 0 8.8 7.2 16 16 16s16-7.2 16-16c0-44.1-34.1-80-76-80zM96 224c44.2 0 80-35.8 80-80s-35.8-80-80-80-80 35.8-80 80 35.8 80 80 80zm0-128c26.5 0 48 21.5 48 48s-21.5 48-48 48-48-21.5-48-48 21.5-48 48-48zm304.1 180c-33.4 0-41.7 12-80.1 12-38.4 0-46.7-12-80.1-12-36.3 0-71.6 16.2-92.3 46.9-12.4 18.4-19.6 40.5-19.6 64.3V432c0 26.5 21.5 48 48 48h288c26.5 0 48-21.5 48-48v-44.8c0-23.8-7.2-45.9-19.6-64.3-20.7-30.7-56-46.9-92.3-46.9zM480 432c0 8.8-7.2 16-16 16H176c-8.8 0-16-7.2-16-16v-44.8c0-16.6 4.9-32.7 14.1-46.4 13.8-20.5 38.4-32.8 65.7-32.8 27.4 0 37.2 12 80.2 12s52.8-12 80.1-12c27.3 0 51.9 12.3 65.7 32.8 9.2 13.7 14.1 29.8 14.1 46.4V432zM157.1 268.9c-11.9-8.1-26-12.9-41.1-12.9H76c-41.9 0-76 35.9-76 80 0 8.8 7.2 16 16 16s16-7.2 16-16c0-26.5 19.8-48 44-48h40c5.5 0 10.8 1.2 15.7 3.3 7.5-8.5 16.1-16 25.4-22.4z"/></svg>`,
        },
      ]
    : [];

  function toggleSidebar() {
    sidebarOpen = !sidebarOpen;
  }
</script>

{#await $project}
  <Loader />
{:then result}
  <div class="h-screen flex overflow-hidden bg-gray-100">
    <!-- Off-canvas menu for mobile -->
    {#if sidebarOpen}
      <div class="md:hidden">
        <div class="fixed inset-0 flex z-40">
          <div class="fixed inset-0" transition:fade={{ duration: 300 }}>
            <div class="absolute inset-0 bg-gray-600 opacity-75" />
          </div>
          <div
            class="relative flex-1 flex flex-col max-w-xs w-full bg-gray-800"
            transition:fly={{ x: -320, opacity: 1, duration: 300, easing: quadInOut }}>
            <div class="absolute top-0 right-0 -mr-14 p-1">
              <button
                on:click={toggleSidebar}
                class="flex items-center justify-center h-12 w-12 rounded-full
                  focus:outline-none focus:bg-gray-600"
                aria-label="Close sidebar">
                <svg
                  class="h-6 w-6 text-white"
                  stroke="currentColor"
                  fill="none"
                  viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            <div class="flex-1 h-0 pt-5 pb-4 overflow-y-auto">
              <div class="flex-shrink-0 flex items-center px-4">
                <Logo />
              </div>

              <Link
                to="/projects"
                className="mt-5 flex items-center flex-shrink-0 px-4 text-white
                py-2 bg-gray-700 font-medium hover:text-gray-200">
                {result.data.project.name}
              </Link>

              <nav class="mt-5 px-2">
                {#each menuitems as item}
                  <Link
                    to={item.path}
                    className={'group flex items-center px-2 py-2 leading-6 text-base font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + (item.regex.test($path) ? 'text-white bg-gray-900 focus:bg-gray-700' : 'mt-1 text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
                    <svg
                      class="mr-3 h-6 w-6 text-gray-300
                        group-hover:text-gray-300 group-focus:text-gray-300
                        transition ease-in-out duration-150"
                      fill="currentColor">
                      {@html item.icon}
                    </svg>
                    {item.title}
                  </Link>
                {/each}
              </nav>
            </div>
            <div class="flex-shrink-0 flex bg-gray-700 p-4">
              <Link to="/profile" className="flex-shrink-0 group block">
                <div class="flex items-center">
                  <div>
                    <span
                      class="inline-block h-10 w-10 rounded-full overflow-hidden
                        bg-gray-100">
                      <svg
                        class="h-full w-full text-gray-300"
                        fill="currentColor"
                        viewBox="0 0 24 24">
                        <path
                          d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004
                          15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0
                          11-8 0 4 4 0 018 0z" />
                      </svg>
                    </span>
                  </div>
                  <div class="ml-3">
                    <p class="text-base leading-6 font-medium text-white">
                      {$user.name}
                    </p>
                    <p
                      class="text-sm leading-5 font-medium text-gray-400
                        group-hover:text-gray-300 transition ease-in-out
                        duration-150">
                      View profile
                    </p>
                  </div>
                </div>
              </Link>
            </div>
          </div>
          <div class="flex-shrink-0 w-14">
            <!-- Force sidebar to shrink to fit close icon -->
          </div>
        </div>
      </div>
    {/if}

    <!-- Static sidebar for desktop -->
    <div class="hidden md:flex md:flex-shrink-0">
      <div class="flex flex-col w-64 bg-gray-800">
        <div class="h-0 flex-1 flex flex-col pt-5 pb-4 overflow-y-auto">
          <div class="flex items-center flex-shrink-0 px-4 text-white">
            <Logo />
          </div>
          <Link
            to="/projects"
            className="mt-5 flex items-center flex-shrink-0 px-4 text-white py-2
            bg-gray-700 font-medium hover:text-gray-200">
            {result.data.project.name}
          </Link>
          <nav class="mt-3 flex-1 px-2 bg-gray-800">
            {#each menuitems as item}
              <Link
                to={item.path}
                className={'group flex items-center px-2 py-2 text-sm leading-5 font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + (item.regex.test($path) ? 'text-white bg-gray-900 focus:bg-gray-700' : 'text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
                <svg
                  class="mr-3 h-6 w-6 text-gray-300 group-hover:text-gray-300
                    group-focus:text-gray-300 transition ease-in-out
                    duration-150"
                  fill="currentColor">
                  {@html item.icon}
                </svg>
                {item.title}
              </Link>
            {/each}
          </nav>
        </div>
        <div class="flex-shrink-0 flex bg-gray-700 p-4">
          <Link to="/profile" className="flex-shrink-0 w-full group block">
            <div class="flex items-center">
              <div>
                <span
                  class="inline-block h-9 w-9 rounded-full overflow-hidden
                    bg-gray-100">
                  <svg
                    class="h-full w-full text-gray-300"
                    fill="currentColor"
                    viewBox="0 0 24 24">
                    <path
                      d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004
                      15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0
                      11-8 0 4 4 0 018 0z" />
                  </svg>
                </span>
              </div>
              <div class="ml-3">
                <p class="text-sm leading-5 font-medium text-white">
                  {$user.name}
                </p>
                <p
                  class="text-xs leading-4 font-medium text-gray-300
                    group-hover:text-gray-200 transition ease-in-out
                    duration-150">
                  View profile
                </p>
              </div>
            </div>
          </Link>
        </div>
      </div>
    </div>
    <div class="flex flex-col w-0 flex-1 overflow-hidden">
      <div class="md:hidden pl-1 pt-1 sm:pl-3 sm:pt-3">
        <button
          on:click={toggleSidebar}
          class="-ml-0.5 -mt-0.5 h-12 w-12 inline-flex items-center
            justify-center rounded-md text-gray-500 hover:text-gray-900
            focus:outline-none focus:bg-gray-200 transition ease-in-out
            duration-150"
          aria-label="Open sidebar">
          <svg
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
      </div>
      <main
        class="flex-1 relative z-0 overflow-y-auto focus:outline-none"
        tabindex="0">
        <div class="pt-2 pb-6 md:py-6">
          {#if title}
            <Header
              project={result.data.project}
              bind:title
              {overtitle}
              {actions}
              {facts}
              {action}
              on:click
              {titleUpdatable}>
              <div slot="posttitle">
                <slot name="posttitle" />
              </div>
            </Header>
          {/if}
          <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
            <slot name="header" project={result.data.project} />
            <div class="py-4">
              <slot project={result.data.project} />
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
{:catch error}
  Error loading Project: {error}
{/await}
