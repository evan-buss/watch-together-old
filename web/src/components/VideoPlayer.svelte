<script>
  import Hls from "hls.js";
  import { onMount } from "svelte";

  onMount(() => {
    var video = document.getElementById("video");
    if (Hls.isSupported()) {
      var hls = new Hls();
      hls.loadSource("http://localhost:8081/media/index.m3u8");
      hls.attachMedia(video);
      hls.on(Hls.Events.MANIFEST_PARSED, function() {
        console.log("manifest parsed");
        video.play();
      });
    } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
      video.src = "http://localhost:8081/media/index.m3u8";
      video.addEventListener("loadedmetadata", function() {
        console.log("metadata laoded");
        video.play();
      });
    }
  });
</script>

<div class="flex items-center justify-center w-5/6 p-4">
  <video id="video" class="inline" controls />
</div>
<!-- <track src="/captions_file.vtt" label="English" kind="captions" srclang="en-us" default > -->
