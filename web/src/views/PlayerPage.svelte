<script>
  import { onMount } from "svelte";
  import { messages } from "../store/chat.js";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";

  let sidebarVisible = true;

  function handleToggle() {
    console.log("TOGGLE ME");
    sidebarVisible = !sidebarVisible;
  }
</script>

<div class="h-screen flex">
  <VideoPlayer full={!sidebarVisible} />
  <Sidebar
    visible={sidebarVisible}
    on:toggleSidebar={handleToggle} />

</div>
