<script>
  import Callout from "../components/base/Callout.svelte";
  import Link from "../components/base/Link.svelte";
  import RunsList from "../components/runs/RunsList.svelte";
  import Layout from "../layouts/Layout.svelte";
</script>

<Layout title="Overview" let:project>
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
    <div class="mt-5">
      <Callout
        color="yellow"
        title="No Runs currently On Going"
        icon={`<path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" xclip-rule="evenodd" />`}>
        <p>
          Head over to the
          <Link className="text-yellow-800 underline" to="/runs">
            Runs page
          </Link>
          to create new Run.
        </p>
      </Callout>
    </div>
  </div>

  <div class="mt-8">
    <h3 class="text-lg leading-6 font-medium text-gray-900">Last 3 Runs</h3>
    <div class="mt-5 bg-white shadow overflow-hidden sm:rounded-md">
      <ul>
        <RunsList {project} limit={3} />
      </ul>
    </div>
  </div>
</Layout>
