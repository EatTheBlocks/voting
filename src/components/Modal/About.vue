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
        <div class="text-main-blockText">{{ network }}</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">IPFS</div>
        <div class="text-main-blockText">{{ ipfsGateway }}</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Hub</div>
        <div class="text-main-blockText">{{ hubUrl }}</div>
      </div>
    </div>
  </UiModal>
</template>

<script>
import networks from '@snapshot-labs/snapshot.js/src/networks.json'

export default {
  name: 'About',
  props: {
    open: Boolean,
  },
  emits: ['close'],
  setup() {
    const network = networks["56"].name
    const ipfsGateway = process.env.VUE_APP_IPFS_GATEWAY
    const hubUrl = process.env.VUE_APP_HUB_URL

    return {
      network,
      ipfsGateway,
      hubUrl
    }
  }
}
</script>
