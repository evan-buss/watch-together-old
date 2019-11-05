<script>
  import { onMount } from "svelte";
  import NavBar from "../components/NavBar.svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";
</script>


<div class="flex flex-col sm:flex-row">
  <VideoPlayer />
  <Sidebar />
</div>
