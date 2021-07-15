<template>
  <div class="py-4 bg-main-primary text-white border-b border-main-border">
    <div class="max-w-[1012px] mx-auto flex justify-between items-center">
      <div class="text-xl font-semibold">{{ $AppName }} / voting app</div>
      <div class="flex space-x-2 items-center">
        <div>
          <Switch
            v-model="enabled"
            :class="enabled ? 'bg-gray-400' : 'bg-gray-700'"
            class="relative inline-flex items-center h-8 rounded-full w-14 mr-5"
            @click="$emit('theme', enabled)"
          >
          <span
            :class="enabled ? 'translate-x-7' : 'translate-x-1'"
            class="inline-block w-6 h-6 transform bg-white rounded-full"
          />
          </Switch>
        </div>
        <button class="border border-white px-6 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">Connect Wallet</button>
        <button class="border border-white px-5 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">?</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { Switch } from '@headlessui/vue'

export default {
  name: 'Header',
  components: {
    Switch
  },
  setup() {
    const enabled = ref(false)

    return { enabled }
  },
  emits: ['theme'],
}
</script>

<style scoped>

</style>
