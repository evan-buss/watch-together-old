<script>
  import { onMount } from "svelte";
  import NavBar from "../components/NavBar.svelte";
  import Sidebar from "../components/Sidebar.svelte";
  import ClapprPlayer from "../components/ClapprPlayer.svelte";

  let h;
  let w;
</script>

<div
  class="h-full max-h-full flex flex-col justify-between md:flex-row"
  bind:clientHeight={h}
  bind:clientWidth={w}>
  <!-- <div class="flex flex-col justify-center items-center w-full"> -->

  <!-- </div> -->

  <ClapprPlayer {h} {w} />
  <Sidebar />
</div>
