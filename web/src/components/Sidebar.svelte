<script>
  import Message from "./Message.svelte";
</script>

<style>

</style>

<div
  class="fixed flex flex-col items-center justify-between text-light-grey
  right-0 w-1/6 h-screen border-l border-light-grey">
  <div class="border-b w-full">
    <h1 class="text-3xl text-center">Live Chat</h1>
  </div>
  <div class="h-full w-full overflow-x-hidden overflow-y-scroll">
    <ul class="flex flex-col">
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message details={{ name: 'Evan', message: 'lmao', sent: true }} />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message details={{ name: 'Evan', message: 'lmao', sent: true }} />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
      <Message />
    </ul>
  </div>
  <div class="p-2 border-t w-full">
    <textarea
      class="rounded resize-none border-2 border-blue-400 w-full p-1"
      rows="3" />
  </div>
</div>
