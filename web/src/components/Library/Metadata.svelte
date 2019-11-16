<script>
  import { createEventDispatcher } from "svelte";
  import MetadataItem from "./MetadataItem.svelte";

  const dispatch = createEventDispatcher();

  let searchTitle = "";
  let searchYear = "";
  let items = [];
  let noResults = false;

  async function search() {
    items = [];
    let data = {};
    if (searchTitle != "") {
      data.title = searchTitle;
    }
    if (searchYear != "") {
      data.year = searchYear;
    }
    let query = new URLSearchParams(data).toString();
    let resp = await fetch(`http://localhost:8080/?${query}`);
    resp.json().then(data => {
      if (data) {
        noResults = false;
        items = data;
      } else {
        noResults = true;
      }
    });
  }
</script>

<!-- Metadata Search Menu -->
<div
  class="z-10 fixed top-0 left-0 w-screen h-screen flex flex-col items-center
  justify-center">
  <!-- Backdrop Overlay -->
  <div
    on:click={() => dispatch('close')}
    class="fixed top-0 left-0 w-full h-full bg-gray-800 opacity-75
    cursor-pointer" />
  <!-- Center Modal -->
  <div
    class="z-20 w-full max-w-4xl flex flex-col bg-white shadow-lg rounded-lg p-4">
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
          bind:value={searchTitle}
          class="appearance-none block w-full bg-gray-200 text-gray-700 border
          border-gray-200 rounded py-2 px-4 mb-3 leading-tight
          focus:outline-none focus:bg-white"
          id="grid-first-name"
          type="text"
          placeholder="The Matrix" />
      </div>
      <div class="w-full md:w-1/6 pr-3">
        <label
          class="block uppercase tracking-wide text-gray-700 text-xs font-bold
          mb-2 ml-1"
          for="grid-last-name">
          Year
        </label>
        <input
          bind:value={searchYear}
          class="appearance-none block w-full bg-gray-200 text-gray-700 border
          border-gray-200 rounded py-2 px-4 leading-tight focus:outline-none
          focus:bg-white focus:border-gray-500"
          id="grid-last-name"
          type="text"
          placeholder="1999" />
      </div>
      <div class="w-full md:w-1/6 self-end">
        <button
          on:click={search}
          style="height: 38px;"
          class="appearance-none w-full bg-blue-400 hover:bg-blue-300 text-white
          py-2 px-4 mb-3 leading-tight rounded">
          Search
        </button>
      </div>
    </div>

    {#if noResults}
      <div class="text-center p-6 text-red-500">
        Could not find any movies matching that criteria.
      </div>
    {:else if items.length === 0}
      <div class="text-center p-6 text-gray-500">
        Search an item to get started
      </div>
    {:else}
      <!-- Table Header -->
      <div class="w-full flex mb-1">
        <div class="w-2/12 text-gray-700">Title</div>
        <div class="w-7/12 text-gray-700">Summary</div>
        <div class="w-1/12 text-gray-700">Year</div>
        <div class="w-1/12 text-gray-700">Rating</div>
      </div>
      <!-- Table Body -->
      <div class="w-full rounded border-2 border-gray-300 flex flex-col">
        {#each items as movie, i}
          <MetadataItem alternate={i % 2 === 0} {movie} on:update />
        {/each}
      </div>
    {/if}

  </div>
</div>
