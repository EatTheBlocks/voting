<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>About</template>
    <div class="flex mx-auto">
<!--      TODO social icons-->
    </div>
    <div class="border border-main-border p-5 !rounded-lg space-y-1">
      <div class="flex justify-between">
        <div class="font-semibold">Version</div>
        <div class="text-main-blockText">0.1.0</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">License</div>
        <div class="text-main-blockText">MIT</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Network</div>
        <div class="text-main-blockText">Binance Smart Chain</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">IPFS</div>
        <div class="text-main-blockText">cloudflare-ipfs.com</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Hub</div>
        <div class="text-main-blockText">https://hub.SERVER</div>
      </div>
    </div>
  </UiModal>
</template>

<script>
export default {
  name: 'About',
  props: {
    open: Boolean,
  },
  emits: ['close'],
}
</script>
