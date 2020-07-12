<script>
  import Layout from "../layouts/Layout.svelte";
  import Link from "../components/base/Link.svelte";
</script>

<Layout title="Overview">
  <div>
    <!-- <h3 class="text-lg leading-6 font-medium text-gray-900">Last 30 days</h3> -->
    <div class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-3">
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <dl>
            <dt class="text-sm leading-5 font-medium text-gray-500 truncate">
              Participants
            </dt>
            <dd class="mt-1 text-3xl leading-9 font-semibold text-gray-900">
              2,892
            </dd>
          </dl>
        </div>
      </div>
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <dl>
            <dt class="text-sm leading-5 font-medium text-gray-500 truncate">
              Runs
            </dt>
            <dd class="mt-1 text-3xl leading-9 font-semibold text-gray-900">
              32
            </dd>
          </dl>
        </div>
      </div>
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <dl>
            <dt class="text-sm leading-5 font-medium text-gray-500 truncate">
              Templates
            </dt>
            <dd class="mt-1 text-3xl leading-9 font-semibold text-gray-900">
              6
            </dd>
          </dl>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-8">
    <h3 class="text-lg leading-6 font-medium text-gray-900">On Going Runs</h3>
    <div class="mt-5 rounded-md bg-yellow-50 p-4">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg
            class="h-5 w-5 text-yellow-400"
            viewBox="0 0 20 20"
            fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75
              1.334-.213 2.98-1.742 2.98H4.42c-1.53
              0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012
              0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <div class="ml-3">
          <h3 class="text-sm leading-5 font-medium text-yellow-800">
            No Runs currently On Going
          </h3>
          <div class="mt-2 text-sm leading-5 text-yellow-700">
            <p>
              Head over to the
              <Link className="text-yellow-800 underline" to="/runs">
                Runs page
              </Link>
              to create new Run.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-8">
    <h3 class="text-lg leading-6 font-medium text-gray-900">Last 3 Runs</h3>
    <div class="mt-5 bg-white shadow overflow-hidden sm:rounded-md">
      <ul>
        <li>
          <a
            href="#"
            class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
            transition duration-150 ease-in-out">
            <div class="px-4 py-4 sm:px-6">
              <div class="flex items-center justify-between">
                <div
                  class="text-sm leading-5 font-medium text-indigo-600 truncate">
                  Back End Developer
                </div>
                <div class="ml-2 flex-shrink-0 flex">
                  <span
                    class="px-2 inline-flex text-xs leading-5 font-semibold
                    rounded-full bg-green-100 text-green-800">
                    Full-time
                  </span>
                </div>
              </div>
              <div class="mt-2 sm:flex sm:justify-between">
                <div class="sm:flex">
                  <div
                    class="mr-6 flex items-center text-sm leading-5
                    text-gray-500">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3
                        0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0
                        00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015
                        5v1H1v-1a5 5 0 015-5z" />
                    </svg>
                    Engineering
                  </div>
                  <div
                    class="mt-2 flex items-center text-sm leading-5
                    text-gray-500 sm:mt-0">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        fill-rule="evenodd"
                        d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0
                        010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                        clip-rule="evenodd" />
                    </svg>
                    Remote
                  </div>
                </div>
                <div
                  class="mt-2 flex items-center text-sm leading-5 text-gray-500
                  sm:mt-0">
                  <svg
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    viewBox="0 0 20 20"
                    fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2
                      2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0
                      00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
                      clip-rule="evenodd" />
                  </svg>
                  <span>
                    Closing on
                    <time datetime="2020-01-07">January 7, 2020</time>
                  </span>
                </div>
              </div>
            </div>
          </a>
        </li>
        <li class="border-t border-gray-200">
          <a
            href="#"
            class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
            transition duration-150 ease-in-out">
            <div class="px-4 py-4 sm:px-6">
              <div class="flex items-center justify-between">
                <div
                  class="text-sm leading-5 font-medium text-indigo-600 truncate">
                  Back End Developer
                </div>
                <div class="ml-2 flex-shrink-0 flex">
                  <span
                    class="px-2 inline-flex text-xs leading-5 font-semibold
                    rounded-full bg-green-100 text-green-800">
                    Full-time
                  </span>
                </div>
              </div>
              <div class="mt-2 sm:flex sm:justify-between">
                <div class="sm:flex">
                  <div
                    class="mr-6 flex items-center text-sm leading-5
                    text-gray-500">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3
                        0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0
                        00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015
                        5v1H1v-1a5 5 0 015-5z" />
                    </svg>
                    Engineering
                  </div>
                  <div
                    class="mt-2 flex items-center text-sm leading-5
                    text-gray-500 sm:mt-0">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        fill-rule="evenodd"
                        d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0
                        010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                        clip-rule="evenodd" />
                    </svg>
                    Remote
                  </div>
                </div>
                <div
                  class="mt-2 flex items-center text-sm leading-5 text-gray-500
                  sm:mt-0">
                  <svg
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    viewBox="0 0 20 20"
                    fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2
                      2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0
                      00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
                      clip-rule="evenodd" />
                  </svg>
                  <span>
                    Closing on
                    <time datetime="2020-01-07">January 7, 2020</time>
                  </span>
                </div>
              </div>
            </div>
          </a>
        </li>
        <li class="border-t border-gray-200">
          <a
            href="#"
            class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
            transition duration-150 ease-in-out">
            <div class="px-4 py-4 sm:px-6">
              <div class="flex items-center justify-between">
                <div
                  class="text-sm leading-5 font-medium text-indigo-600 truncate">
                  Back End Developer
                </div>
                <div class="ml-2 flex-shrink-0 flex">
                  <span
                    class="px-2 inline-flex text-xs leading-5 font-semibold
                    rounded-full bg-green-100 text-green-800">
                    Full-time
                  </span>
                </div>
              </div>
              <div class="mt-2 sm:flex sm:justify-between">
                <div class="sm:flex">
                  <div
                    class="mr-6 flex items-center text-sm leading-5
                    text-gray-500">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3
                        0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0
                        00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015
                        5v1H1v-1a5 5 0 015-5z" />
                    </svg>
                    Engineering
                  </div>
                  <div
                    class="mt-2 flex items-center text-sm leading-5
                    text-gray-500 sm:mt-0">
                    <svg
                      class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                      viewBox="0 0 20 20"
                      fill="currentColor">
                      <path
                        fill-rule="evenodd"
                        d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0
                        010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                        clip-rule="evenodd" />
                    </svg>
                    Remote
                  </div>
                </div>
                <div
                  class="mt-2 flex items-center text-sm leading-5 text-gray-500
                  sm:mt-0">
                  <svg
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    viewBox="0 0 20 20"
                    fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2
                      2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0
                      00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
                      clip-rule="evenodd" />
                  </svg>
                  <span>
                    Closing on
                    <time datetime="2020-01-07">January 7, 2020</time>
                  </span>
                </div>
              </div>
            </div>
          </a>
        </li>
      </ul>
    </div>
  </div>
</Layout>
