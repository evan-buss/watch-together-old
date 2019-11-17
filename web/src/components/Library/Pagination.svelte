<script>
  import { createEventDispatcher } from "svelte";

  export let total;
  export let perPage;
  export let currentPage;

  const dispatch = createEventDispatcher();

  function prev() {
    if (currentPage > 0) {
      currentPage--;
      dispatch("page", currentPage);
    }
  }

  function next() {
    if (currentPage < total / perPage - 1) {
      currentPage++;
      dispatch("page", currentPage);
    }
  }

  $: max =
    currentPage >= total / perPage - 1 ? total : (currentPage + 1) * perPage;
</script>

<div class="w-full flex justify-end items-center m-1">
  <!-- Results Info -->
  <div class="text-gray-500 mr-2">
    Showing {currentPage * perPage}-{max} of {total}
  </div>
  <!-- Left Arrow -->
  <i
    on:click={prev}
    class="la la-angle-double-left p-2 m-1 shadow rounded bg-gray-100
    cursor-pointer" />
  <!-- Right Arrow -->
  <i
    on:click={next}
    class="la la-angle-double-right p-2 m-1 shadow rounded bg-gray-100
    cursor-pointer" />
</div>
