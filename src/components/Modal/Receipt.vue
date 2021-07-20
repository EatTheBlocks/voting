<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>Receipts</template>
    Author {{ ipfsHash }}
  </UiModal>
</template>

<script>
export default {
  name: 'Receipt',
  props: {
    open: Boolean,
    ipfsHash: String,
  },
  emits: ['close']
}
</script>
