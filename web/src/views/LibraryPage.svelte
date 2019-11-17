<script>
  import { onMount } from "svelte";
  import MovieCard from "../components/Library/MovieCard.svelte";
  import MetadataModal from "../components/Library/MetadataModal.svelte";
  import NavBar from "../components/NavBar.svelte";

  let showMetadata = false;
  let movies = [];
  let filterText = "";
  let netError = false;
  let metadataItem;
  // FIXME: Hack because the way flexbox works...
  // $: filteredMovies = movies.filter(
  //   item => item.title.toLowerCase().indexOf(filterText.toLowerCase()) !== -1
  // );
  // $: flexAlign =
  //   filteredMovies.length < 4 ? "lg:justify-start" : "lg:justify-between";

  onMount(async () => {
    await getLibrary();
  });

  async function getLibrary() {
    movies = [];
    const response = await fetch("http://localhost:5228/library");
    try {
      movies = await response.json();
    } catch (error) {
      netError = true;
    }
  }

  function updateMetadata(event) {
    showMetadata = false;
    fetch("http://localhost:5228/library", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ id: metadataItem.id, metadata: event.detail })
    })
      .then(data => {
        if (data.status === 200) {
          return data.json();
        } else {
          throw Error("cannot complete");
        }
      })
      .then(async () => {
        // FIXME: I want to ideally just update the changed card in place but I cannot figure out how to make it work.
        // For now I am just reloading the entire library on movie update
        await getLibrary();
      })
      .catch(error => {
        console.log(error);
      });
  }
</script>

<div class="bg-gray-300">
  <!-- {#if showMetadata} -->
  {#if true}
    <MetadataModal
      on:close={() => (showMetadata = false)}
      on:update={updateMetadata} />
  {/if}

  <div class="lg:w-10/12 mx-auto">
    <!-- Top Panel Controls -->
    <div
      class="full flex flex-col sm:flex-row justify-between items-center px-4">
      <h1 class="text-3xl font-bold font-sans py-4">Movie Library</h1>
      <input
        type="text"
        bind:value={filterText}
        class="appearance-none outline-none rounded p-2 w-full max-w-sm sm:w-64
        focus:shadow-outline"
        placeholder="Search" />
    </div>
    <!-- Movie Cards List -->
    {#if netError}
      <h1 class="text-center text-2xl font-bold font-sans text-gray-500 mt-12">
        Please run
        <code class="text-gray-800 bg-white p-2">watch-together init</code>
        to generate your library
      </h1>
    {:else if movies.length === 0}
      <h1 class="text-center text-2xl font-bold font-sans text-gray-500 mt-12">
        Unable to locate any movies.
      </h1>
    {:else}
      <!-- <div class="flex flex-wrap justify-center items-center {flexAlign}"> -->
      <div class="flex flex-wrap justify-center items-center">
        {#each movies as metadata (metadata.id)}
          <MovieCard
            {metadata}
            {filterText}
            on:open={() => {
              showMetadata = true;
              metadataItem = metadata;
            }} />
        {/each}
      </div>
    {/if}
  </div>
</div>
