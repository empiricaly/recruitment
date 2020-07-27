<script context="module">
  import { pathToRegexp } from "path-to-regexp";

  const regexes = {
    overview: pathToRegexp("/projects/:project_id/overview"),
    runs: pathToRegexp("/projects/:project_id/runs/:run_id?"),
    templates: pathToRegexp("/projects/:project_id/templates/:template_id?"),
    participants: pathToRegexp(
      "/projects/:project_id/participants/:participant_id?"
    )
  };
</script>

<script>
  import Logo from "../components/base/Logo.svelte";
  import Link from "../components/base/Link.svelte";
  import Button from "../components/base/Button.svelte";
  import Header from "./Header.svelte";
  import { fade, fly } from "svelte/transition";
  import { quadInOut } from "svelte/easing";
  import { path } from "../lib/routing.js";

  export let title = null;
  export let overtitle = null;
  export let action = null;

  let projectName = "Speed Dating";
  let projectID = "jk3n21kj";

  let sidebarOpen = false;

  const menuitems = [
    {
      title: "Overview",
      path: `/projects/${projectID}/overview`,
      regex: regexes.overview,
      icon: `<path d="M19 21H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9a1 1 0 0 1-1 1zM6 19h12V9.157l-6-5.454-6 5.454V19z"/>`
    },
    {
      title: "Runs",
      path: `/projects/${projectID}/runs`,
      regex: regexes.runs,
      icon: `<path d="M16.394 12L10 7.737v8.526L16.394 12zm2.982.416L8.777 19.482A.5.5 0 0 1 8 19.066V4.934a.5.5 0 0 1 .777-.416l10.599 7.066a.5.5 0 0 1 0 .832z"/>`
    },
    {
      title: "Templates",
      path: `/projects/${projectID}/templates`,
      regex: regexes.templates,
      icon: `<path d="M15.728 9.686l-1.414-1.414L5 17.586V19h1.414l9.314-9.314zm1.414-1.414l1.414-1.414-1.414-1.414-1.414 1.414 1.414 1.414zM7.242 21H3v-4.243L16.435 3.322a1 1 0 0 1 1.414 0l2.829 2.829a1 1 0 0 1 0 1.414L7.243 21z" />`
    },
    {
      title: "Participants",
      path: `/projects/${projectID}/participants`,
      regex: regexes.participants,
      icon: `<path d="M4 22a8 8 0 1 1 16 0h-2a6 6 0 1 0-12 0H4zm8-9c-3.315 0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6zm0-2c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4z" />`
    }
  ];

  function toggleSidebar() {
    sidebarOpen = !sidebarOpen;
  }
</script>

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
            <nav class="mt-5 px-2">
              {#each menuitems as item}
                <Link
                  to={item.path}
                  className={'group flex items-center px-2 py-2 leading-6 text-base font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + (item.regex.test($path) ? 'text-white bg-gray-900 focus:bg-gray-700' : 'mt-1 text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
                  <svg
                    class="mr-3 h-6 w-6 text-gray-300 group-hover:text-gray-300
                    group-focus:text-gray-300 transition ease-in-out
                    duration-150"
                    viewBox="0 0 24 24"
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
                  <p class="text-base leading-6 font-medium text-white">Me</p>
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
          {projectName}
        </Link>
        <nav class="mt-3 flex-1 px-2 bg-gray-800">
          {#each menuitems as item}
            <Link
              to={item.path}
              className={'group flex items-center px-2 py-2 text-sm leading-5 font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + (item.regex.test($path) ? 'text-white bg-gray-900 focus:bg-gray-700' : 'text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
              <svg
                class="mr-3 h-6 w-6 text-gray-300 group-hover:text-gray-300
                group-focus:text-gray-300 transition ease-in-out duration-150"
                viewBox="0 0 24 24"
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
                    d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904
                    0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0
                    018 0z" />
                </svg>
              </span>
            </div>
            <div class="ml-3">
              <p class="text-sm leading-5 font-medium text-white">Me</p>
              <p
                class="text-xs leading-4 font-medium text-gray-300
                group-hover:text-gray-200 transition ease-in-out duration-150">
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
        class="-ml-0.5 -mt-0.5 h-12 w-12 inline-flex items-center justify-center
        rounded-md text-gray-500 hover:text-gray-900 focus:outline-none
        focus:bg-gray-200 transition ease-in-out duration-150"
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
          <Header {title} {overtitle} {action} on:click />
          <!-- <div
            class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex justify-between
            items-baseline">
            <h1 class="text-2xl font-semibold text-gray-900 leading-snug">
              {#if overtitle}
                <div
                  class="uppercase tracking-wide font-medium text-gray-500 mr-1
                  text-sm">
                  {overtitle}
                </div>
              {/if}
              {title}
            </h1>
            {#if action}
              <Button on:click text={action} />
            {/if}
            <slot name="actions" />
          </div> -->
        {/if}
        <slot name="header" />
        <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
          <div class="py-4">
            <slot />
          </div>
        </div>
      </div>
    </main>
  </div>
</div>
