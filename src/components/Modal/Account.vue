<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>Connect Wallet</template>

    <UiButton class="flex justify-center items-center"
              @click="connect"
    >
      <img src="@/assets/imgs/connectors/metamask.png" class="w-7 h-7 mr-2">
      <span>MetaMask</span>
    </UiButton>
  </UiModal>
</template>

<script>
//import {computed} from 'vue'
import {ethers} from 'ethers'

export default {
  name: 'Account',
  props: {
    open: Boolean,
  },
  emits: ['close'],
  setup(props, {emit}) {
    const provider = new ethers.providers.Web3Provider(window.ethereum)

    async function connect() {
      await provider.send("eth_requestAccounts", [])
      const signer = provider.getSigner()
      console.log("Account:", await signer.getAddress())
      emit('close')
    }

    return {
      connect,
    }
  }
}
</script>

<style scoped>

</style>
