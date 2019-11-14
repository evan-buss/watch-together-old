<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let volume;

  function down() {
    if (volume > 0) {
      dispatch("setVolume", {
        volume: volume -= 1
      });
    }
  }

  function up() {
    if (volume < 1) {
      dispatch("setVolume", {
        volume: volume += 1
      });
    }
  }

  // let prevVolume;
  // function mute() {
  //   if (video.volume > 0) {
  //     prevVolume = video.volume;
  //     video.volume = 0;
  //   } else {
  //     video.volume = prevVolume;
  //   }
  // }
</script>

<div class="flex flex-row justify-center items-center">
  <div
    on:click={down}
    class="flex rounded-full bg-white h-8 w-8 items-center justify-center
    cursor-pointer">
    <i class="la la-volume-down" />
  </div>
  <!-- on:click={mute} -->
  <div class="text-black hover:text-red-500 mx-4 cursor-pointer">{volume}</div>
  <div
    on:click={up}
    class="flex rounded-full bg-white h-8 w-8 items-center justify-center
    cursor-pointer">
    <i class="la la-volume-up" />
  </div>
</div>
