<script>
  import { onMount } from "svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";

  let sidebarVisible = true;
</script>

<div class="h-screen flex">
  <VideoPlayer full={!sidebarVisible} />
  <Sidebar
    visible={sidebarVisible}
    on:toggleSidebar={() => (sidebarVisible = !sidebarVisible)} />

</div>
