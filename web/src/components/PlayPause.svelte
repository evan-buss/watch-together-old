<script>
  import { user } from "../store/state";
  export let video;

  let icon = "la la-play";

  let isStreamer = $user.type === "streamer";

  video.addEventListener("play", () => {
    console.log("playing");
    icon = "la la-pause";
  });

  video.addEventListener("pause", () => {
    console.log("paused");
    icon = "la la-play";
  });

  function toggle() {
    // Doesn't do anything if not the streamer
    if (!isStreamer) return;
    if (video.paused) {
      video.play();
    } else {
      video.pause();
    }
  }
</script>

<div
  on:click={toggle}
  class="flex px-2 mr-4 rounded-full bg-white h-8 w-8 items-center
  justify-center {isStreamer ? 'cursor-pointer' : 'cursor-not-allowed'}">
  <i class={icon} />
</div>
