<script>
  import { onMount } from "svelte";
  import NavBar from "../components/NavBar.svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";
</script>


<div class="h-full max-h-full flex flex-col justify-between md:flex-row">
  <VideoPlayer/>
  <Sidebar />
</div>
