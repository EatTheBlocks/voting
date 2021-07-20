<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-3/12">
      <Menu/>
    </div>
    <div class="w-9/12 space-y-5">
      <div class="text-2xl font-semibold text-main-heading">{{ $AppName }}</div>
      <div class="panel">
        <div class="panel-body space-y-3">
          <h4 class="text-xl">About</h4>
          <div>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque vulputate mi nunc. Mauris posuere, nunc a vestibulum luctus, magna nibh ultricies enim, ut dignissim ex lorem nec diam.</div>
          <h4 class="text-xl">Network</h4>
          <div>Binance Smart Chain</div>
          <h4 class="text-xl">Token Address</h4>
          <div class="flex items-center">
            0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a
            <ExternalLinkIcon class="ml-1 h-4 w-4"/>
          </div>
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
import {ExternalLinkIcon} from '@heroicons/vue/outline'

const members = [
  {address: "0x4e48c12cf0abef413a2e8994b4a6a743c3f2d296"},
  {address: "0x2509ec5907f6e265939c750c80548bbda3c08d2c"},
  {address: "0xb3ee966ff8c19522a3603d6e889d9440fb4b63a9"},
]

export default {
  name: 'About',
  components: {
    ExternalLinkIcon
  },
  data() {
    return {
      members: members
    }
  }
}
</script>

<style scoped>

</style>
