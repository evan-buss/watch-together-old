<script>
  import MetadataItem from "./MetadataItem.svelte";
</script>

<!-- Backdrop Overlay -->
<div class="fixed z-10 top-0 left-0 w-screen h-screen bg-gray-200 opacity-75" />
<!-- Metadata Search Menu -->
<div
  class="fixed z-20 top-0 left-0 w-screen h-screen flex flex-col items-center
  justify-center">
  <!-- Center Modal -->
  <div class="w-full max-w-3xl flex flex-col bg-white shadow-lg rounded-lg p-4">
    <!-- Search Header -->
    <div class="flex flex-row">
      <div class="w-full md:w-4/6 pr-3 mb-6 md:mb-0">
        <label
          class="block uppercase tracking-wide text-gray-700 text-xs font-bold
          mb-2 ml-1"
          for="grid-first-name">
          Title
        </label>
        <input
          class="appearance-none block w-full bg-gray-200 text-gray-700 border
          border-gray-200 rounded py-2 px-4 mb-3 leading-tight
          focus:outline-none focus:bg-white"
          id="grid-first-name"
          type="text"
          placeholder="The Matrix" />
      </div>
      <div class="w-full md:w-2/6">
        <label
          class="block uppercase tracking-wide text-gray-700 text-xs font-bold
          mb-2 ml-1"
          for="grid-last-name">
          Year
        </label>
        <input
          class="appearance-none block w-full bg-gray-200 text-gray-700 border
          border-gray-200 rounded py-2 px-4 leading-tight focus:outline-none
          focus:bg-white focus:border-gray-500"
          id="grid-last-name"
          type="text"
          placeholder="1999" />
      </div>
    </div>
    <!-- Table Header -->
    <div class="w-full flex mb-1">
      <div class="w-2/12 text-gray-700">Title</div>
      <div class="w-8/12 text-gray-700">Summary</div>
      <div class="w-1/12 text-gray-700">Year</div>
      <div class="w-1/12 text-gray-700">Rating</div>
    </div>
    <!-- Table Body -->
    <div class="w-full rounded border-2 border-gray-300 flex flex-col">
      <MetadataItem />
      <MetadataItem alternate={true} />
      <MetadataItem />
      <MetadataItem alternate={true} />
    </div>
  </div>
</div>
