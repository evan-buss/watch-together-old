<script>
  import Hls from "hls.js";
  import { onMount } from "svelte";

  import { user } from "../store/state";
  import Volume from "./Volume.svelte";
  import VideoProgress from "./VideoProgress.svelte";
  import PlayPause from "./PlayPause.svelte";
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
      });
    } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
      video.src = "http://localhost:8080/media/index.m3u8";
      video.addEventListener("loadedmetadata", function() {
        console.log("metadata loaded");
      });
    }
  });
</script>

<!-- TODO: Certain video aspec ratios cut off shit. Depends on the source -->
<div
  class="video-container md:px-4 md:pt-4 flex flex-col flex-grow items-center
  justify-center">
  <div class="w-full lg:w-11/12 xl:w-10/12">
    <video id="video" />
    {#if video !== null}
      <div class="block flex flex-row items-center px-4 md:px-0">
        <PlayPause {video} />
        <VideoProgress {video} />
        <Volume {video} />
      </div>
    {/if}
  </div>
</div>

<!-- <track src="/captions_file.vtt" label="English" kind="captions" srclang="en-us" default > -->
