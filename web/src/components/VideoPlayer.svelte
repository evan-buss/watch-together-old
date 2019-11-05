<script>
  import Hls from "hls.js";
  import { onMount } from "svelte";

  import { user } from "../store/state";
  import Volume from "./Volume.svelte";
  import VideoProgress from "./VideoProgress.svelte";
  import "./VideoPlayer.css";

  let video = null;

  let controls = $user.type === "streamer";

  onMount(() => {
    video = document.getElementById("video");
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

<!-- VideoPlayer is the core video component that plays HLS video streams -->
<div class="flex flex-col items-end justify-center sm:p-4 w-screen">
  <video id="video" class="w-full" bind:this={video} {controls} />
  <div class="flex w-full flex-row items-center justify-between">
    {#if video !== null}
      <VideoProgress {video} />
      <Volume {video} />
    {:else}Loading Video{/if}
  </div>
</div>
<!-- <track src="/captions_file.vtt" label="English" kind="captions" srclang="en-us" default > -->
