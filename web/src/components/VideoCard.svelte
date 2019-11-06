<script>
  export let src;
  export let title = "Default Title";
</script>

<div class="w-full max-w-xs md:w-64 rounded-lg shadow-2xl my-4 sm:mx-4 bg-gray-200">
  <img {src} alt="" class="w-full rounded-lg rounded-b-none" />
  <div class="px-6 py-4">
    <div class="font-bold text-xl mb-2 truncate">{title}</div>
    <div class="text-base text-gray-700">Summary {title}</div>
  </div>
</div>
