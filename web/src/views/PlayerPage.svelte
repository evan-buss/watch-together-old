<script>
  import { onMount } from "svelte";
  import NavBar from "../components/NavBar.svelte";
  import Sidebar from "../components/Sidebar/Sidebar.svelte";
  import ClapprPlayer from "../components/Player/ClapprPlayer.svelte";
</script>

<div class="h-full max-h-full flex flex-col justify-between md:flex-row">
  <ClapprPlayer />
  <Sidebar />
</div>
