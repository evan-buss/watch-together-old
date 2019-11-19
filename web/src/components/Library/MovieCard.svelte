<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { notifications } from "../Notification/notification";

  const dispatch = createEventDispatcher();
  export let metadata;
  export let filterText = "";

  let movie;
  let summary;
  let isHover = false;

  let fileName = metadata.location
    .split("\\")
    .pop()
    .split("/")
    .pop();

  $: matchesFilter =
    (movie &&
      movie.title.toLowerCase().indexOf(filterText.toLowerCase()) !== -1) ||
    metadata.location.toLowerCase().indexOf(filterText.toLowerCase()) !== -1;

  // Fetch metadata from the server on load
  onMount(async () => {
    if (metadata.metadata != -1) {
      const response = await fetch(
        `http://localhost:8080/id/${metadata.metadata}`
      );
      movie = await response.json();
      summary =
        movie.summary.length > 120
          ? movie.summary.substring(0, 120) + "..."
          : movie.summary;
    }
  });

  function playMovie() {
    if (movie) {
      notifications.addPersistant({
        type: "movie",
        title: movie.title,
        message: `${movie.title} now playing!`,
        poster: movie.poster
      });
    } else {
      notifications.addPersistant({
        type: "movie",
        title: fileName,
        message: `${fileName} now playing!`
      });
    }
    dispatch("play");
  }
</script>

{#if matchesFilter}
  {#if movie}
    <div
      on:mouseleave={() => (isHover = false)}
      class="w-full h-56 max-w-sm rounded-lg shadow-xl my-2 sm:mx-2 bg-gray-100
      flex flex-row relative">
      <!-- Poster -->
      {#if isHover}
        <div
          on:click={playMovie}
          class="absolute top-0 left-0 h-full w-full rounded-lg opacity-75
          bg-black flex items-center justify-center cursor-pointer">
          <i
            class="opacity-100 la la-play-circle text-6xl text-white
            hover:text-green-500" />
        </div>
      {/if}
      <img
        on:mouseenter={() => (isHover = true)}
        src={movie.poster}
        alt="{movie.title} poster"
        class="rounded-lg rounded-r-none object-cover object-left-top" />
      <!-- Content Container -->
      <div class="p-2 h-full flex flex-col justify-between">
        <!-- Title and Summary Container -->
        <div>
          <a
            href={movie.url}
            target="_blank"
            rel="noopener"
            class="font-bold text-xl text-gray-800 mb-2 break-words">
            {movie.title}
          </a>
          <div class="text-base text-gray-600">{summary}</div>
        </div>

        <!-- Details -->
        <div class="flex items-center justify-between items-baseline px-2">
          <div class="text-base text-gray-900 font-bold">{movie.year}</div>
          <i
            on:click={() => dispatch('open')}
            class="la la-cog p-1 hover:text-green-600 hover:shadow rounded-full
            cursor-pointer" />
          <span>
            {#if movie.rating !== undefined}
              <i class="la la-star text-yellow-500 text-xl" />
            {/if}
            <span class="text-base text-gray-600">
              {movie.rating === undefined ? 'Not available' : movie.rating}
            </span>
          </span>
        </div>
      </div>
    </div>
  {:else}
    <div
      on:mouseleave={() => (isHover = false)}
      class="w-full h-56 max-w-sm rounded-lg shadow-xl my-2 sm:mx-2 bg-gray-100
      flex flex-row relative">
      {#if isHover}
        <div
          on:click={playMovie}
          class="absolute top-0 left-0 h-full w-full rounded-lg opacity-75
          bg-black flex items-center justify-center cursor-pointer">
          <i
            class="opacity-100 la la-play-circle text-6xl text-white
            hover:text-green-500" />
        </div>
      {/if}
      <!-- Image -->
      <div
        on:mouseenter={() => (isHover = true)}
        class="w-5/12 bg-gray-900 rounded-lg rounded-r-none text-6xl text-white
        font-mono flex justify-center items-center select-none">
        ?
      </div>
      <!-- Movie Details -->
      <div class="w-7/12 h-full flex flex-col justify-between ">
        <!-- Title -->
        <div class="p-2 font-bold text-xl text-gray-800 mb-2 break-words">
          {fileName}
        </div>
        <!-- Summary -->
        <div class="p-2 text-base flex-grow text-gray-600">
          Unable to automatically determine metadata. Please enter manually.
        </div>
        <!-- Bottom Bar -->
        <div class="flex items-center justify-between items-baseline">
          <div
            on:click={() => dispatch('open')}
            class="p-2 w-full text-white bg-red-500 hover:bg-red-400
            cursor-pointer text-center rounded-br-lg select-none">
            Edit Metadata
          </div>
        </div>
      </div>
    </div>
  {/if}
{/if}
