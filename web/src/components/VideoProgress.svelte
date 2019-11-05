<script>
  import { tweened } from "svelte/motion";
  import { cubicOut } from "svelte/easing";
  export let video;
  let percentage = 0;

  const progress = tweened(0, {
    duration: 200,
    easing: cubicOut
  });

  video.ontimeupdate = function() {
    // percentage = Math.round((video.currentTime / video.duration) * 100);
    percentage = video.currentTime / video.duration;
    progress.set(percentage);
  };
</script>

<progress class="w-full mr-4" value={$progress} />
