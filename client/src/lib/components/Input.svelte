<script lang="ts">
  let {
    value = $bindable(),
    isPending,
    placeholder
  }: { value: string; isPending: boolean; placeholder: string } = $props();
</script>

<div class="relative">
  <div class="absolute inset-0 rounded-lg bg-gradient-to-r from-purple-600 to-blue-500 p-0.5">
    <div class="h-full w-full rounded-md bg-white"></div>
  </div>

  <input
    type="text"
    class="relative h-12 w-full rounded-lg border-transparent bg-transparent px-4 text-gray-900 placeholder-gray-500 focus:outline-none"
    bind:value
    {placeholder}
    disabled={isPending}
  />
</div>
