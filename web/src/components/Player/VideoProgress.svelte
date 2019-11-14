<script>
  import { tweened } from "svelte/motion";
  import { cubicOut } from "svelte/easing";
  export let video = null;
  let percentage = 0;

  const progress = tweened(0, {
    duration: 200,
    easing: cubicOut
  });

  // video.addEventListener("play", () => {});
  // video.addEventListener("pause", () => {});
  // video.ontimeupdate = function() {
  //   // percentage = Math.round((video.currentTime / video.duration) * 100);
  //   percentage = video.currentTime / video.duration;
  //   progress.set(percentage);
  // };
</script>

<style>
  /* progress::-moz-progress-bar {
    background-color: red !important;
  }

  progress {
    color: red;
  }

  progress::-webkit-progress-value {
    background-color: red !important;
  } */
</style>

<progress class="w-full mr-4" value={$progress} />
