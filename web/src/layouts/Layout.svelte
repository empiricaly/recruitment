<script>
  import Link from "../components/base/Link.svelte";
  import Button from "../components/form/Button.svelte";
  import { fade, fly } from "svelte/transition";
  import { quadInOut } from "svelte/easing";
  import { path } from "../lib/routing.js";

  export let title = null;
  export let action = null;

  let sidebarOpen = false;

  const menuitems = [
    {
      title: "Overview",
      path: "/overview",
      icon: `<path d="M19 21H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9a1 1 0 0 1-1 1zM6 19h12V9.157l-6-5.454-6 5.454V19z"/>`
    },
    {
      title: "Runs",
      path: "/runs",
      icon: `<path d="M16.394 12L10 7.737v8.526L16.394 12zm2.982.416L8.777 19.482A.5.5 0 0 1 8 19.066V4.934a.5.5 0 0 1 .777-.416l10.599 7.066a.5.5 0 0 1 0 .832z"/>`
    },
    {
      title: "Templates",
      path: "/templates",
      icon: `<path d="M15.728 9.686l-1.414-1.414L5 17.586V19h1.414l9.314-9.314zm1.414-1.414l1.414-1.414-1.414-1.414-1.414 1.414 1.414 1.414zM7.242 21H3v-4.243L16.435 3.322a1 1 0 0 1 1.414 0l2.829 2.829a1 1 0 0 1 0 1.414L7.243 21z" />`
    },
    {
      title: "Participants",
      path: "/participants",
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
              <svg
                alt="Empirica Recruitment"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                width="200"
                viewBox="0 0 118 13">
                <path
                  d="M6.486 8.412l-1.2-.375c-.265.586-.833.955-1.653.955-1.143
                  0-1.781-.633-1.893-1.652h4.881c.018-.182.03-.317.03-.48
                  0-1.958-1.249-3.276-3.124-3.276C1.611 3.584.352 4.961.352
                  6.9c0 1.928 1.212 3.305 3.28 3.305 1.407 0 2.427-.691
                  2.854-1.793zM3.51 4.785c.984 0 1.681.586 1.763
                  1.518h-3.51c.159-.955.78-1.518 1.747-1.518zM8.959
                  10V6.309c0-1.026.516-1.5 1.289-1.5.78 0 1.272.474 1.272
                  1.5V10h1.359V6.309c0-1.026.504-1.5 1.277-1.5.78 0 1.272.474
                  1.272 1.5V10h1.36V6.215c0-1.635-.845-2.59-2.35-2.59-.944
                  0-1.553.398-1.905
                  1.031h-.04c-.37-.685-1.044-1.031-1.923-1.031-.867
                  0-1.459.34-1.804.99h-.047l-.07-.826h-1.05V10h1.36zm10.195
                  2.086v-2.83h.041c.416.539 1.014.861 1.957.861 1.6 0 2.86-1.154
                  2.86-3.222 0-2.069-1.278-3.252-2.848-3.252-.99
                  0-1.705.404-2.074
                  1.107h-.047l-.041-.96h-1.213v8.296h1.365zm1.723-3.147c-.984
                  0-1.734-.72-1.734-2.015V6.87c0-1.242.744-2.01 1.734-2.01 1.002
                  0 1.729.768 1.729 2.04 0 1.265-.715 2.038-1.729
                  2.038zm3.984-6.123h1.535V1.48h-1.535v1.336zM26.303
                  10V3.79h-1.354V10h1.354zm2.543 0V6.525c0-1.037.545-1.482
                  1.465-1.482h.65v-1.36h-.633c-.75 0-1.336.364-1.617
                  1.126h-.041l-.082-1.02h-1.102V10h1.36zm2.86-7.184h1.534V1.48h-1.535v1.336zM33.145
                  10V3.79h-1.353V10h1.354zm7.22-2.033L39.17 7.58c-.299.885-.867
                  1.365-1.81 1.365-1.208 0-1.882-.873-1.882-2.05
                  0-1.178.674-2.051 1.881-2.051.932 0 1.483.427 1.788
                  1.248l1.2-.44c-.556-1.365-1.494-2.05-2.976-2.05-1.957 0-3.27
                  1.37-3.27 3.293 0 1.916 1.313 3.293 3.24 3.293 1.512 0
                  2.561-.792
                  3.024-2.221zm1.13-2.795l1.148.398c.212-.492.575-.761
                  1.19-.761.685 0 1.172.38 1.172 1.177v.258h-1.729c-1.424
                  0-2.25.791-2.25 1.916 0 1.16.809 1.945 2.18 1.945.867 0
                  1.547-.28
                  1.928-.92h.047l.035.815h1.101V6.045c0-1.547-1.002-2.408-2.449-2.408-1.236
                  0-2.021.597-2.373 1.535zm1.963 3.79c-.762
                  0-1.113-.339-1.113-.866 0-.522.334-.862
                  1.113-.862h1.547v.317c0 .926-.656 1.412-1.547
                  1.412zM51.076.919h-.474v10.758h.474V.918zM56.385
                  10V6.549c0-1.26.697-1.87 1.758-1.87h.656v-.908h-.645c-.955
                  0-1.652.551-1.857
                  1.489h-.047l-.094-1.389h-.68V10h.909zm8.93-1.564l-.797-.252c-.276.726-.973
                  1.2-1.98 1.2-1.425
                  0-2.175-.843-2.233-2.179h5.156c.012-.123.017-.21.017-.322
                  0-1.951-1.195-3.205-3.005-3.205-1.887 0-3.088 1.347-3.088
                  3.275 0 1.91 1.154 3.24 3.175 3.24 1.372 0 2.362-.691
                  2.754-1.757zM62.46 4.48c1.213 0 2.033.756 2.098
                  2.004h-4.237c.135-1.236.908-2.004 2.139-2.004zm9.92
                  3.551l-.803-.258c-.287.967-.96 1.577-2.092 1.577-1.388
                  0-2.209-1.014-2.209-2.414 0-1.395.815-2.415 2.21-2.415 1.1 0
                  1.757.528 2.085
                  1.454l.803-.3c-.492-1.282-1.406-1.986-2.877-1.986-1.898 0-3.14
                  1.348-3.14 3.247 0 1.898 1.242 3.246 3.117 3.246 1.511 0
                  2.507-.82 2.906-2.15zM74.338 10V6.549c0-1.26.697-1.87
                  1.758-1.87h.656v-.908h-.645c-.955 0-1.652.551-1.857
                  1.489h-.047L74.11 3.87h-.68V10h.909zm7.605-6.129v3.516c0
                  1.312-.674 1.98-1.681 1.98-1.05
                  0-1.664-.656-1.664-1.916v-3.58h-.909v3.686c0 1.634.868 2.607
                  2.391 2.607 1.084 0 1.74-.55 2.016-1.353h.04L82.214
                  10h.639V3.871h-.909zm2.116-1.377h1.113V1.48h-1.113v1.014zM85.066
                  10V3.871h-.908V10h.908zm.762-5.297h1.049v3.533c0 1.213.627
                  1.764 1.746 1.764h1.084v-.873h-.984c-.633
                  0-.932-.275-.932-.914v-3.51h1.916v-.832h-1.916v-1.7h-.914v1.7h-1.049v.832zM91.5
                  10V6.484c0-1.312.639-1.98 1.588-1.98.978 0 1.553.656 1.553
                  1.916V10h.902V6.484c0-1.312.644-1.98 1.594-1.98.978 0
                  1.547.656 1.547
                  1.916V10h.908V6.314c0-1.634-.82-2.607-2.274-2.607-1.037
                  0-1.693.533-2.004
                  1.371h-.04c-.305-.879-1.002-1.371-2.005-1.371-1.013
                  0-1.652.527-1.921
                  1.354h-.041l-.076-1.19h-.64V10h.909zm14.865-1.564l-.797-.252c-.275.726-.972
                  1.2-1.98 1.2-1.424
                  0-2.174-.843-2.233-2.179h5.157c.011-.123.017-.21.017-.322
                  0-1.951-1.195-3.205-3.006-3.205-1.886 0-3.087 1.347-3.087
                  3.275 0 1.91 1.154 3.24 3.175 3.24 1.371 0 2.362-.691
                  2.754-1.757zm-2.853-3.956c1.213 0 2.033.756 2.097
                  2.004h-4.236c.135-1.236.908-2.004 2.139-2.004zm5.039
                  5.52V6.484c0-1.312.668-1.98 1.681-1.98 1.049 0 1.664.656 1.664
                  1.916V10h.909V6.314c0-1.634-.867-2.607-2.391-2.607-1.09
                  0-1.74.55-2.016
                  1.354h-.041l-.076-1.19h-.638V10h.908zm4.957-5.297h1.049v3.533c0
                  1.213.627 1.764 1.746 1.764h1.084v-.873h-.985c-.632
                  0-.931-.275-.931-.914v-3.51h1.916v-.832h-1.916v-1.7h-.914v1.7h-1.049v.832z"
                  fill="#fff" />
              </svg>
            </div>
            <nav class="mt-5 px-2">
              {#each menuitems as item}
                <Link
                  to={item.path}
                  className={'group flex items-center px-2 py-2 leading-6 text-base font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + ($path === item.path ? 'text-white bg-gray-900 focus:bg-gray-700' : 'mt-1 text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
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
          <svg
            alt="Empirica Recruitment"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            width="200"
            viewBox="0 0 118 13">
            <path
              d="M6.486 8.412l-1.2-.375c-.265.586-.833.955-1.653.955-1.143
              0-1.781-.633-1.893-1.652h4.881c.018-.182.03-.317.03-.48
              0-1.958-1.249-3.276-3.124-3.276C1.611 3.584.352 4.961.352 6.9c0
              1.928 1.212 3.305 3.28 3.305 1.407 0 2.427-.691 2.854-1.793zM3.51
              4.785c.984 0 1.681.586 1.763 1.518h-3.51c.159-.955.78-1.518
              1.747-1.518zM8.959 10V6.309c0-1.026.516-1.5 1.289-1.5.78 0
              1.272.474 1.272 1.5V10h1.359V6.309c0-1.026.504-1.5 1.277-1.5.78 0
              1.272.474 1.272 1.5V10h1.36V6.215c0-1.635-.845-2.59-2.35-2.59-.944
              0-1.553.398-1.905
              1.031h-.04c-.37-.685-1.044-1.031-1.923-1.031-.867
              0-1.459.34-1.804.99h-.047l-.07-.826h-1.05V10h1.36zm10.195
              2.086v-2.83h.041c.416.539 1.014.861 1.957.861 1.6 0 2.86-1.154
              2.86-3.222 0-2.069-1.278-3.252-2.848-3.252-.99 0-1.705.404-2.074
              1.107h-.047l-.041-.96h-1.213v8.296h1.365zm1.723-3.147c-.984
              0-1.734-.72-1.734-2.015V6.87c0-1.242.744-2.01 1.734-2.01 1.002 0
              1.729.768 1.729 2.04 0 1.265-.715 2.038-1.729
              2.038zm3.984-6.123h1.535V1.48h-1.535v1.336zM26.303
              10V3.79h-1.354V10h1.354zm2.543 0V6.525c0-1.037.545-1.482
              1.465-1.482h.65v-1.36h-.633c-.75 0-1.336.364-1.617
              1.126h-.041l-.082-1.02h-1.102V10h1.36zm2.86-7.184h1.534V1.48h-1.535v1.336zM33.145
              10V3.79h-1.353V10h1.354zm7.22-2.033L39.17 7.58c-.299.885-.867
              1.365-1.81 1.365-1.208 0-1.882-.873-1.882-2.05 0-1.178.674-2.051
              1.881-2.051.932 0 1.483.427 1.788
              1.248l1.2-.44c-.556-1.365-1.494-2.05-2.976-2.05-1.957 0-3.27
              1.37-3.27 3.293 0 1.916 1.313 3.293 3.24 3.293 1.512 0 2.561-.792
              3.024-2.221zm1.13-2.795l1.148.398c.212-.492.575-.761 1.19-.761.685
              0 1.172.38 1.172 1.177v.258h-1.729c-1.424 0-2.25.791-2.25 1.916 0
              1.16.809 1.945 2.18 1.945.867 0 1.547-.28
              1.928-.92h.047l.035.815h1.101V6.045c0-1.547-1.002-2.408-2.449-2.408-1.236
              0-2.021.597-2.373 1.535zm1.963 3.79c-.762 0-1.113-.339-1.113-.866
              0-.522.334-.862 1.113-.862h1.547v.317c0 .926-.656 1.412-1.547
              1.412zM51.076.919h-.474v10.758h.474V.918zM56.385
              10V6.549c0-1.26.697-1.87 1.758-1.87h.656v-.908h-.645c-.955
              0-1.652.551-1.857
              1.489h-.047l-.094-1.389h-.68V10h.909zm8.93-1.564l-.797-.252c-.276.726-.973
              1.2-1.98 1.2-1.425
              0-2.175-.843-2.233-2.179h5.156c.012-.123.017-.21.017-.322
              0-1.951-1.195-3.205-3.005-3.205-1.887 0-3.088 1.347-3.088 3.275 0
              1.91 1.154 3.24 3.175 3.24 1.372 0 2.362-.691 2.754-1.757zM62.46
              4.48c1.213 0 2.033.756 2.098 2.004h-4.237c.135-1.236.908-2.004
              2.139-2.004zm9.92 3.551l-.803-.258c-.287.967-.96 1.577-2.092
              1.577-1.388 0-2.209-1.014-2.209-2.414 0-1.395.815-2.415 2.21-2.415
              1.1 0 1.757.528 2.085
              1.454l.803-.3c-.492-1.282-1.406-1.986-2.877-1.986-1.898 0-3.14
              1.348-3.14 3.247 0 1.898 1.242 3.246 3.117 3.246 1.511 0 2.507-.82
              2.906-2.15zM74.338 10V6.549c0-1.26.697-1.87
              1.758-1.87h.656v-.908h-.645c-.955 0-1.652.551-1.857
              1.489h-.047L74.11 3.87h-.68V10h.909zm7.605-6.129v3.516c0
              1.312-.674 1.98-1.681 1.98-1.05
              0-1.664-.656-1.664-1.916v-3.58h-.909v3.686c0 1.634.868 2.607 2.391
              2.607 1.084 0 1.74-.55 2.016-1.353h.04L82.214
              10h.639V3.871h-.909zm2.116-1.377h1.113V1.48h-1.113v1.014zM85.066
              10V3.871h-.908V10h.908zm.762-5.297h1.049v3.533c0 1.213.627 1.764
              1.746 1.764h1.084v-.873h-.984c-.633
              0-.932-.275-.932-.914v-3.51h1.916v-.832h-1.916v-1.7h-.914v1.7h-1.049v.832zM91.5
              10V6.484c0-1.312.639-1.98 1.588-1.98.978 0 1.553.656 1.553
              1.916V10h.902V6.484c0-1.312.644-1.98 1.594-1.98.978 0 1.547.656
              1.547 1.916V10h.908V6.314c0-1.634-.82-2.607-2.274-2.607-1.037
              0-1.693.533-2.004
              1.371h-.04c-.305-.879-1.002-1.371-2.005-1.371-1.013
              0-1.652.527-1.921
              1.354h-.041l-.076-1.19h-.64V10h.909zm14.865-1.564l-.797-.252c-.275.726-.972
              1.2-1.98 1.2-1.424
              0-2.174-.843-2.233-2.179h5.157c.011-.123.017-.21.017-.322
              0-1.951-1.195-3.205-3.006-3.205-1.886 0-3.087 1.347-3.087 3.275 0
              1.91 1.154 3.24 3.175 3.24 1.371 0 2.362-.691
              2.754-1.757zm-2.853-3.956c1.213 0 2.033.756 2.097
              2.004h-4.236c.135-1.236.908-2.004 2.139-2.004zm5.039
              5.52V6.484c0-1.312.668-1.98 1.681-1.98 1.049 0 1.664.656 1.664
              1.916V10h.909V6.314c0-1.634-.867-2.607-2.391-2.607-1.09
              0-1.74.55-2.016
              1.354h-.041l-.076-1.19h-.638V10h.908zm4.957-5.297h1.049v3.533c0
              1.213.627 1.764 1.746 1.764h1.084v-.873h-.985c-.632
              0-.931-.275-.931-.914v-3.51h1.916v-.832h-1.916v-1.7h-.914v1.7h-1.049v.832z"
              fill="#fff" />
          </svg>
        </div>
        <!-- Sidebar component, swap this element with another sidebar if you like -->
        <nav class="mt-5 flex-1 px-2 bg-gray-800">
          {#each menuitems as item}
            <Link
              to={item.path}
              className={'group flex items-center px-2 py-2 text-sm leading-5 font-medium rounded-md focus:outline-none transition ease-in-out duration-150 ' + ($path === item.path ? 'text-white bg-gray-900 focus:bg-gray-700' : 'text-gray-300 hover:text-white hover:bg-gray-700 focus:text-white focus:bg-gray-700')}>
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
          <div
            class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex justify-between">
            <h1 class="text-2xl font-semibold text-gray-900">{title}</h1>
            {#if action}
              <Button on:click text={action} />
            {/if}
          </div>
        {/if}
        <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
          <div class="py-4">
            <slot />
          </div>
        </div>
      </div>
    </main>
  </div>
</div>
