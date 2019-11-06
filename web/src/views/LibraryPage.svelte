<script>
  import { onMount } from "svelte";
  import VideoCard from "../components/VideoCard.svelte";
  import NavBar from "../components/NavBar.svelte";
  let photos = [];
  onMount(async () => {
    const response = await fetch(
      "https://jsonplaceholder.typicode.com/photos?albumId=5"
    );
    photos = await response.json();
  });
</script>

<div class="flex flex-row flex-wrap justify-center bg-gray-900 py-4">
  {#each photos as photo}
    <VideoCard src={photo.url} title={photo.title} />
  {/each}
</div>
