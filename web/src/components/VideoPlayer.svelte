<script>
  import Hls from "hls.js";
  import { onMount } from "svelte";
  import { sidebarVisible } from "../store/state";
  import "./VideoPlayer.css";

  let video;

  onMount(() => {
    if (Hls.isSupported()) {
      var hls = new Hls();
      hls.loadSource("http://localhost:8080/media/index.m3u8");
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, function() {
        console.log("manifest parsed");
        video.play();
      });
    } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
      video.src = "http://localhost:8080/media/index.m3u8";
      video.addEventListener("loadedmetadata", function() {
        console.log("metadata laoded");
        video.play();
      });
    }
  });
</script>

<style>
  .sidebar {
    width: calc(100% - 300px);
  }
</style>

<!-- VideoPlayer is the core video component that plays HLS video streams -->
<div
  class="flex items-center justify-center sm:p-4 w-screen {$sidebarVisible && 'sidebar'}">
  <video
    id="video"
    class="w-full"
    bind:this={video}
    controlsList="nodownload noremoteplayback" controls/>
</div>
<!-- <track src="/captions_file.vtt" label="English" kind="captions" srclang="en-us" default > -->
