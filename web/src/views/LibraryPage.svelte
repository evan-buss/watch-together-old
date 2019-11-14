<script>
  import { onMount } from "svelte";
  import VideoCard from "../components/VideoCard.svelte";
  import NavBar from "../components/NavBar.svelte";

  let movies = [];

  onMount(async () => {
    const response = await fetch(
      "http://localhost:8080/movie/year?search=2019"
    );
    movies = await response.json();
  });
</script>

<div class="flex flex-row flex-wrap justify-center bg-gray-300 py-4">
  {#each movies as movie}
  <VideoCard {movie} />
  {/each}
</div>
