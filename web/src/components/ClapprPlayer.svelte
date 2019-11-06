<script>
  import Clappr from "clappr";
  import { onMount } from "svelte";

  let player;
  export let h;
  export let w;

  onMount(() => {
    player = new Clappr.Player({
      source: `${window.location.origin}/media/index.m3u8`,
      width: "100%",
      height: "100%",
      poster:
        "https://upload.wikimedia.org/wikipedia/commons/4/4d/Rembrandt_-_The_Anatomy_Lesson_of_Dr_Nicolaes_Tulp.jpg",
      parentId: "#player",
      resizable: true
    });
  });
</script>

<div id="player" class="w-full h-64 md:h-full md:p-4" />
