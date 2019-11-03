<script>
  import Message from "./Message.svelte";

  let messages = [{ name: "Evan", message: "lmao", sent: true }];
  let value = "";
  let visible = true;

  let sendMessage = event => {
    if (event.which === 13) {
      event.preventDefault();
      messages = [...messages, { name: "Evan", message: value, sent: true }];
      value = "";
    }
  };
</script>

<style>

</style>

<!-- https://svelte.dev/examples#update -> Desired scrolling behavior -->
<!-- Toggle close and open tab -->
<div
  on:click={() => (visible = !visible)}
  class=" left-0 mt-1 -mx-8 border rounded rounded-r-none p-2 hover:bg-gray-300
  pointer h-10">
  X
</div>
<!-- TODO: Make visible visible from other classes. Need to adjust video sizing if hidden... -->
{#if visible}
  <div
    class="fixed flex flex-col items-center justify-between text-light-grey
    right-0 w-1/6 h-screen border-l border-light-grey">
    <div class="border-b w-full">
      <h1 class="text-3xl text-center">Live Chat</h1>
    </div>
    <div class="h-full w-full overflow-x-hidden overflow-y-scroll">
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
{/if}
