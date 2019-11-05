<script>
  import Message from "./Message.svelte";
  import { createSocket } from "../store/socket";
  import { sidebarVisible, user } from "../store/state";
  import {
    beforeUpdate,
    afterUpdate,
    onMount,
    createEventDispatcher
  } from "svelte";
  const dispatch = createEventDispatcher();
  let socket = createSocket();

  let messages = [];
  let message = "";

  // Control the chat scroll position (either stick when scrolled or keep bottom in sync)
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
      if (message !== "") {
        socket.send("message", {
          sender: $user.name,
          message: message
        });
        message = "";
      }
    }
  }
</script>

<div
  class="{!$sidebarVisible && 'hidden'} h-full w-full md:max-w-xs flex flex-col items-center
  justify-between text-light-grey border-l border-light-grey">
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
      bind:value={message}
      on:keypress={sendMessage}
      class="rounded resize-none border-2 border-blue-400 w-full p-1"
      rows="3" />
  </div>
</div>
