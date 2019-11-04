<script>
  import Message from "./Message.svelte";
  import { createSocket } from "../store/socket";
  import {
    beforeUpdate,
    afterUpdate,
    onMount,
    createEventDispatcher
  } from "svelte";
  const dispatch = createEventDispatcher();

  export let visible;

  let messages = [];
  let socket = createSocket();
  let value = "";

  let div;
  let autoscroll;

  beforeUpdate(() => {
    autoscroll =
      div && div.offsetHeight + div.scrollTop > div.scrollHeight - 20;
  });

  afterUpdate(() => {
    if (autoscroll) div.scrollTo(0, div.scrollHeight);
  });

  socket.bind("message", data => {
    messages = [...messages, data];
  });

  function sendMessage(event) {
    if (event.which === 13) {
      event.preventDefault();
      if (value !== "") {
        socket.send("message", { sender: "evan", message: value, sent: true });
        value = "";
      }
    }
  }
</script>

<style>
  #sidebar {
    width: 300px;
  }
</style>

<div class="flex flex-col">
  <!-- Toggle Button -->
  <div
    on:click={() => dispatch('toggleSidebar')}
    class="left-0 mt-1 border rounded rounded-r-none p-2 hover:bg-gray-300
    pointer h-10">
    X
  </div>

  <!-- Sidebar -->
  <div
    id="sidebar"
    class="fixed flex flex-col items-center justify-between text-light-grey
    right-0 h-screen border-l border-light-grey {visible ? '' : 'hidden'}">
    <div class="border-b w-full">
      <h1 class="text-3xl text-center">Live Chat</h1>
    </div>
    <div
      class="h-full w-full overflow-x-hidden overflow-y-scroll"
      bind:this={div}>
      <ul class="flex flex-col">
        {#each messages as message}
          <Message details={message} />
        {/each}
      </ul>
    </div>
    <div class="p-2 border-t w-full">
      <textarea
        bind:value
        on:keypress={sendMessage}
        class="rounded resize-none border-2 border-blue-400 w-full p-1"
        rows="3" />
    </div>
  </div>
</div>
