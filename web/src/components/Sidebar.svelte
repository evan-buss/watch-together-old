<script>
  import Message from "./Message.svelte";
  import { getSocket } from "../store/socket";
  import { sidebarVisible, user } from "../store/state";
  import {
    beforeUpdate,
    afterUpdate,
    onMount,
    createEventDispatcher
  } from "svelte";
  const dispatch = createEventDispatcher();
  let socket = getSocket($user.ip);

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
    console.log(socket);
    
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

<style>
  .message-area:hover {
    overflow-y: scroll;
  }
</style>

<div
  class="{!$sidebarVisible && 'hidden'} h-full w-full md:max-w-xs flex flex-col
  items-center justify-between text-light-grey md:border-l border-gray-400">
  <div class="message-area h-full w-full overflow-hidden" bind:this={div}>
    <ul class="flex flex-col p-2">
      {#each messages as message}
        <Message details={message} />
      {/each}
    </ul>
  </div>
  <div class="p-2 border-t border-gray-400 w-full">
    <textarea
      bind:value={message}
      on:keypress={sendMessage}
      class="rounded resize-none rounded-lg focus:shadow-md focus:outline-none w-full p-2
      text-gray-600"
      rows="3" />
  </div>
</div>
