<script>
  import { onMount } from "svelte";
  import { messages } from "../store/chat.js";
  import Sidebar from "../components/Sidebar.svelte";
  import VideoPlayer from "../components/VideoPlayer.svelte";

  let socket = null;
  let sidebarVisible = true;

  onMount(() => {
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onmessage = message => {
      console.log("socket message received");
      messages.update(state => [
        ...state,
        { sender: "evan", message: message.data, sent: true }
      ]);
    };

    return () => {
      socket.close();
    };
  });

  function handleToggle() {
    console.log("TOGGLE ME");
    sidebarVisible = !sidebarVisible;
  }
</script>

<div class="h-screen flex">
  <VideoPlayer full={!sidebarVisible} />
  <Sidebar
    visible={sidebarVisible}
    on:sendMessage={event => socket.send(event.detail.value)}
    on:toggleSidebar={handleToggle} />

</div>
