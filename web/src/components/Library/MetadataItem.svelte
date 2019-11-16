<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let alternate = false;
  export let movie;
</script>

<div
  class="p-2 w-full flex {alternate ? 'bg-gray-200 hover:bg-blue-100' : 'bg-white hover:bg-blue-100'}">
  <a href={movie.url} target="_blank" rel="noopener" class="w-2/12 truncate">
    {movie.title}
  </a>
  <div class="w-7/12 truncate">{movie.summary}</div>
  <div class="w-1/12">{movie.year}</div>
  <div class="w-1/12 text-center">{movie.rating}</div>
  <button
    on:click={() => dispatch('update', movie.id)}
    class="w-1/12 p-1 bg-blue-400 hover:bg-blue-300 rounded-sm text-white">
    Update
  </button>
</div>
