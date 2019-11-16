<script>
  export let alternate = false;
</script>

<div
  class="p-2 w-full flex {alternate ? 'bg-gray-200 hover:bg-blue-100' : 'bg-white hover:bg-blue-100'}">
  <a href="#" class="w-2/12">Title Here</a>
  <div class="w-8/12 truncate">This is the very long summary text</div>
  <div class="w-1/12">1999</div>
  <div class="w-1/12 text-center">7.6</div>
</div>
