<template>
  <div class="py-4 bg-main-primary text-white border-b border-main-border">
    <div class="max-w-[1012px] mx-auto flex justify-between items-center">
      <div class="text-xl font-semibold">{{ $AppName }} / voting app</div>
      <div class="space-x-2">
        <button class="border border-white px-6 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">Connect Wallet</button>
        <button class="border border-white px-5 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">?</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Header"
}
</script>

<style scoped>

</style>
