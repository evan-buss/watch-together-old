<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let video;

  function down() {
    if (video.volume > 0) {
      video.volume -= 0.05;
    }
  }

  function up() {
    if (video.volume < 1) {
      video.volume += 0.05;
    }
  }

  let prevVolume;
  function mute() {
    if (video.volume > 0) {
      prevVolume = video.volume;
      video.volume = 0;
    } else {
      video.volume = prevVolume;
    }
  }
</script>

<div class="flex flex-row justify-center items-center my-4">
  <div
    on:click={down}
    class="flex rounded-full bg-white h-8 w-8 items-center justify-center
    cursor-pointer">
    <i class="la la-volume-down" />
  </div>
  <div
    on:click={mute}
    class="text-white hover:text-red-500 mx-4 cursor-pointer">
    {Math.round(video.volume * 100)}
  </div>
  <div
    on:click={up}
    class="flex rounded-full bg-white h-8 w-8 items-center justify-center
    cursor-pointer">
    <i class="la la-volume-up" />
  </div>
</div>
