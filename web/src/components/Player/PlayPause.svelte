<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { user } from "../../store/state";

  export let isPaused = true;

  const dispatcher = createEventDispatcher();

  let isStreamer = $user.type === "streamer";
  $: icon = isPaused ? "la la-play" : "la la-pause";

  // video.addEventListener("play", () => {
  //   console.log("playing");
  //   icon = "la la-pause";
  // });

  // video.addEventListener("pause", () => {
  //   console.log("paused");
  //   icon = "la la-play";
  // });

  // onMount(() => {
  //   if (!player) {
  //     return;
  //   }
  //   player.listenTo(player, "pause", () => {
  //     socket.send("pause", {});
  //     icon = "la la-play";
  //   });

  //   player.listenTo(player, "play", () => {
  //     socket.send("play", {});
  //     icon = "la la-pause";
  //   });
  // });

  function toggle() {
    // Doesn't do anything if not the streamer
    if (isStreamer) dispatcher("toggle");
  }
</script>

<div
  on:click={toggle}
  class="flex px-2 mr-4 rounded-full bg-white h-8 w-8 items-center
  justify-center {isStreamer ? 'cursor-pointer' : 'cursor-not-allowed'}">
  <i class={icon} />
</div>
