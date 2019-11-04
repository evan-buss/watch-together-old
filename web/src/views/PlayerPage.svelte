<script>
  import { onMount } from "svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";

  let sidebarVisible = true;
  console.log("player page loaded");
</script>

<div class="h-screen flex pt-16">
  <VideoPlayer full={!sidebarVisible} />
  <Sidebar
    visible={sidebarVisible}
    on:toggleSidebar={() => (sidebarVisible = !sidebarVisible)} />

</div>
