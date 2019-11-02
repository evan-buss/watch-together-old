<script>
  import { onMount } from "svelte";
  import Sidebar from "./components/Sidebar.svelte";
  import VideoPlayer from "./components/VideoPlayer.svelte";

  onMount(async () => {
    // Connect to web-socket here
  });
</script>

<div class="h-screen flex">
  <VideoPlayer/>
  <Sidebar />
</div>
