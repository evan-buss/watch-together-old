<script>
  export let movie;

  let summary =
    movie.summary.length > 140
      ? movie.summary.substring(0, 140) + "..."
      : movie.summary;

  let isHover = false;
</script>

<div
  on:mouseleave={() => (isHover = false)}
  class="w-full h-56 max-w-sm rounded-lg shadow-xl my-2 sm:mx-2 bg-gray-100 flex
  flex-row relative">
  <!-- Poster -->
  {#if isHover}
    <div
      on:click={() => console.log(movie.title)}
      class="absolute top-0 left-0 h-full w-full rounded-lg opacity-75 bg-black
      flex items-center justify-center cursor-pointer">
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
