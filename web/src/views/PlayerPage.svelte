<script>
  import { onMount } from "svelte";
  import NavBar from "../components/NavBar.svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";
</script>

<style>
  .content {
    height: calc(100vh - 4em);
    margin-top: 4em;
  }
</style>

<div class="content flex">
  <VideoPlayer />
  <Sidebar />
</div>
