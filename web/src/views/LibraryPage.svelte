<script>
  import { onMount } from "svelte";
  import MovieCard from "../components/Library/MovieCard.svelte";
  import NavBar from "../components/NavBar.svelte";

  let movies = [];
  let filterText = "";
  let netError = false;
  // FIXME: Hack because the way flexbox works...
  // $: filteredMovies = movies.filter(
  //   item => item.title.toLowerCase().indexOf(filterText.toLowerCase()) !== -1
  // );
  // $: flexAlign =
  //   filteredMovies.length < 4 ? "lg:justify-start" : "lg:justify-between";

  onMount(async () => {
    const response = await fetch("http://localhost:5228/library");
    try {
      movies = await response.json();
    } catch (error) {
      netError = true;
    }
  });
</script>

<div class="bg-gray-300">
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
        {#each movies as metadata}
          <MovieCard {metadata} />
        {/each}
      </div>
    {/if}

  </div>
</div>
