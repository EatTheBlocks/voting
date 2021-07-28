<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-3/12">
      <Menu/>
    </div>
    <div class="w-9/12 space-y-5">
      <div class="text-2xl font-semibold text-main-heading">{{ $store.state.space.name }}</div>
      <div class="panel">
        <div class="panel-body space-y-3">
          <h4 class="text-xl font-semibold">About</h4>
          <p>The #1 school for Blockchain Development<br>Up-to-date &amp; easy-to-follow tutorials to learn Blockchain Development: Ethereum, Solidity, Web3, DeFi</p>
          <h4 class="text-xl font-semibold">Network</h4>
          <div>{{ network }}</div>
          <h4 class="text-xl font-semibold">Token Address</h4>
          <a :href="_explorer(56, $store.state.space.address, 'token')" target="_blank">
            <div class="flex items-center mt-3">
              {{ $store.state.space.address }}
              <ExternalLinkIcon class="ml-1 h-4 w-4"/>
            </div>
          </a>
        </div>
      </div>
      <div class="panel">
        <div class="panel-title">Admins</div>
        <div class="panel-body divide-y divide-main-border p-0">
          <div class="flex p-4" v-for="(admin, i) in $store.state.space.admins" :key="i">
            <User :address="admin"></User>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import networks from '@/helpers/networks.json'
import {ExternalLinkIcon} from '@heroicons/vue/outline'
import {useStore} from 'vuex'
import {computed} from "vue";

export default {
  name: 'About',
  components: {
    ExternalLinkIcon
  },
  setup() {
    const store = useStore()

    const network = computed(() => {
      return networks[store.state.space.network].name
    })

    return {
      network,
    }
  }
}
</script>

<style scoped>

</style>
