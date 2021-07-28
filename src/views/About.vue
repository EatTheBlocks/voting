<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-3/12">
      <Menu/>
    </div>
    <div class="w-9/12 space-y-5">
      <div class="text-2xl font-semibold text-main-heading">{{ $AppName }}</div>
      <div class="panel">
        <div class="panel-body space-y-3">
          <h4 class="text-xl font-semibold">About</h4>
          <p>The #1 school for Blockchain Development<br>Up-to-date &amp; easy-to-follow tutorials to learn Blockchain Development: Ethereum, Solidity, Web3, DeFi</p>
          <h4 class="text-xl font-semibold">Network</h4>
          <div>{{ network }}</div>
          <h4 class="text-xl font-semibold">Token Address</h4>
          <a :href="_explorer(56, tokenAddress, 'token')" target="_blank">
            <div class="flex items-center mt-3">
              {{ tokenAddress }}
              <ExternalLinkIcon class="ml-1 h-4 w-4"/>
            </div>
          </a>
        </div>
      </div>
      <div class="panel">
        <div class="panel-title">Members</div>
        <div class="panel-body divide-y divide-main-border p-0">
          <div class="flex p-4" v-for="member in members" :key="member.address">
            <User :address="member.address"></User>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import networks from '@/helpers/networks.json'
import {ExternalLinkIcon} from '@heroicons/vue/outline'

const members = [
  {address: "0x4e48c12cf0abef413a2e8994b4a6a743c3f2d296"},
]

export default {
  name: 'About',
  components: {
    ExternalLinkIcon
  },
  setup() {
    const network = networks[process.env.VUE_APP_NETWORK_ID].name
    const tokenAddress = process.env.VUE_APP_TOKEN_ADDRESS

    return {
      network,
      tokenAddress,
      members: members,
    }
  }
}
</script>

<style scoped>

</style>
