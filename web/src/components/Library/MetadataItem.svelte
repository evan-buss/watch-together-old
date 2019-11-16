<script>
  export let alternate = false;
  export let movie;
  console.log(movie);
</script>

<div
  class="p-2 w-full flex {alternate ? 'bg-gray-200 hover:bg-blue-100' : 'bg-white hover:bg-blue-100'}">
  <a href={movie.url} target="_blank" rel="noopener" class="w-2/12 truncate">
    {movie.title}
  </a>
  <div class="w-8/12 truncate">{movie.summary}</div>
  <div class="w-1/12">{movie.year}</div>
  <div class="w-1/12 text-center">{movie.rating}</div>
</div>
