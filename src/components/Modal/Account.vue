<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>
      <span v-if="!connected">Connect Wallet</span>
      <span v-else>Account</span>
    </template>

    <div v-if="!connected">
      <UiButton class="flex justify-center items-center"
                @click="connect"
      >
        <img src="@/assets/imgs/connectors/metamask.png" class="w-7 h-7 mr-2">
        <span>MetaMask</span>
      </UiButton>
    </div>
    <div class="space-y-3" v-else>
      <UiButton class="flex justify-center items-center">
        <User :address="address" :popover="false" class=""/>
      </UiButton>
      <UiButton class="flex justify-center items-center text-red-500"
                @click="logout"
      >
        Log out
      </UiButton>
    </div>
  </UiModal>
</template>

<script>
import {computed} from 'vue'
import {useStore} from 'vuex'

export default {
  name: 'Account',
  props: {
    open: Boolean,
  },
  emits: ['close'],
  setup(props, {emit}) {
    const store = useStore()

    async function connect() {
      await store.dispatch('web3/connect')
      emit('close')
    }

    async function logout() {
      await store.dispatch('web3/logout')
      emit('close')
    }

    const connected = computed(() => {
      return store.state.web3.connected
    })

    const address = computed(() => {
      return store.state.web3.address
    })

    return {
      connect,
      logout,
      connected,
      address,
    }
  }
}
</script>

<style scoped>

</style>
