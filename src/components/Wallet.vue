<template>
  <button
    class="border border-white px-6 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary"
    @click="modalAccountOpen = true"
  >
    Connect Wallet
  </button>

  <ModalAccount
    :open="modalAccountOpen"
    @close="modalAccountOpen= false"
  />
</template>

<script>
import {ref} from 'vue'

export default {
  name: 'Wallet',
  setup() {
    const modalAccountOpen = ref(true)

    return {
      modalAccountOpen,
    }
  }
}
</script>

<style scoped>

</style>
