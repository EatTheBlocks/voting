<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>Receipt</template>
    <div class="flex justify-between border border-main-border p-5 !rounded-lg">
      <div class="font-semibold">Author</div>
      <div class="text-main-blockText">
        <a :href="_ipfsUrl(ipfsHash)" target="_blank" class="flex items-center">
          #{{ ipfsHash.substring(0, 7) }}
          <ExternalLinkIcon class="ml-1 h-4 w-4"/>
        </a>
      </div>
    </div>
  </UiModal>
</template>

<script>
import {ExternalLinkIcon} from '@heroicons/vue/outline'

export default {
  name: 'Receipt',
  props: {
    open: Boolean,
    ipfsHash: String,
  },
  emits: ['close'],
  components: {
    ExternalLinkIcon,
  }
}
</script>
