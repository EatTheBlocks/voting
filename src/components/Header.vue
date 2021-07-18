<template>
  <div class="py-4 bg-main-primary text-white border-b border-main-border">
    <div class="max-w-[1012px] mx-auto flex justify-between items-center">
      <router-link :to="{name:'Home'}" class="text-xl font-semibold">{{ $AppName }} / voting app</router-link>
      <div class="flex space-x-2 items-center">
        <div class="flex items-center">
          <!--          <component :is="enabled?'SunIcon':'MoonIcon'" class="h-6 w-6 mr-2" />-->
          <Switch
            v-model="enabled"
            :class="enabled ? 'bg-gray-100' : 'bg-gray-700'"
            class="relative inline-flex items-center h-8 rounded-full w-14 mr-4"
            @click="changeTheme(enabled)"
          >
          <span
            :class="enabled ? 'bg-gray-900 translate-x-7' : 'bg-gray-100 translate-x-1'"
            class="inline-block w-6 h-6 transform rounded-full flex transition duration-200 ease-in-out"
          >
            <component :is="enabled?'MoonIcon':'SunIcon'" class="h-4 w-4 self-center mx-auto"
                       :class="enabled?'text-white':'text-black'"/>
          </span>
          </Switch>
        </div>
        <button
          class="border border-white px-6 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">
          Connect Wallet
        </button>
        <button
          class="border border-white px-5 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">
          ?
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {ref} from 'vue'
import {Switch} from '@headlessui/vue'
import {SunIcon, MoonIcon} from '@heroicons/vue/outline'

export default {
  name: 'Header',
  components: {
    Switch,
    SunIcon,
    MoonIcon,
  },
  setup() {
    const enabled = ref(false)

    return {enabled}
  },
  methods: {
    changeTheme(darkMode) {
      if (darkMode) {
        document.body.classList.add('dark')
        document.body.classList.remove('light')
      } else {
        document.body.classList.remove('dark')
        document.body.classList.add('light')
      }
    }
  }
  ,
}
</script>

<style scoped>

</style>
