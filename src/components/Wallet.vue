<template>
  <button
    class="border border-white px-6 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary"
    @click="modalAccountOpen = true"
  >
    <span v-if="!connected">Connect Wallet</span>
    <User :address="address" :popover="false" v-else/>
  </button>

  <ModalAccount
    :open="modalAccountOpen"
    @close="modalAccountOpen= false"
  />
</template>

<script>
import {ref, computed} from 'vue'
import {useStore} from 'vuex'

export default {
  name: 'Wallet',
  setup() {
    const store = useStore()
    const modalAccountOpen = ref(false)

    const connected = computed(() => {
      return store.state.web3.connected
    })

    const address = computed(() => {
      return store.state.web3.address
    })

    return {
      modalAccountOpen,
      connected,
      address,
    }
  }
}
</script>

<style scoped>

</style>
