<template>
  <div class="py-4 bg-main-primary text-white border-b border-main-border">
    <div class="max-w-[1012px] mx-auto flex justify-between items-center">
      <router-link :to="{name:'Home'}" class="text-xl font-semibold flex items-center">
        <img src="@/assets/imgs/logo-long.png" alt="etb logo" class="w-72 mr-3">
        / voting app
      </router-link>
      <div class="flex space-x-2 items-center">
        <div class="flex items-center">
          <!--          <component :is="enabled?'SunIcon':'MoonIcon'" class="h-6 w-6 mr-2" />-->
          <Switch
            v-model="enabled"
            :class="enabled ? 'bg-gray-100' : 'bg-gray-100'"
            class="relative inline-flex items-center h-8 rounded-full w-14 mr-4"
            @click="changeTheme(enabled)"
          >
          <span
            :class="enabled ? 'bg-gray-900 translate-x-7' : 'bg-gray-900 translate-x-1'"
            class="inline-block w-6 h-6 transform rounded-full flex transition duration-200 ease-in-out"
          >
            <component :is="enabled?'MoonIcon':'SunIcon'" class="h-4 w-4 self-center mx-auto"
                       :class="enabled?'text-white':'text-white'"/>
          </span>
          </Switch>
        </div>
        <Wallet/>
        <button
          v-if="parseInt($store.state.web3.network, 16) != $store.state.space.network"
          @click="changeNetwork"
          class="px-5 py-3 rounded-full font-semibold text-sm bg-red-500 text-white hover:bg-white hover:text-red-500">
          Incorrect Network
        </button>
        <button
          @click="modalAboutOpen = true"
          class="border border-white px-5 py-3 rounded-full font-semibold text-sm hover:bg-white hover:text-main-primary">
          ?
        </button>
      </div>
    </div>
  </div>

  <ModalAbout
    :open="modalAboutOpen"
    @close="modalAboutOpen= false"
  />
</template>

<script>
import {ref} from 'vue'
import {useStore} from 'vuex'
import {Switch} from '@headlessui/vue'
import {SunIcon, MoonIcon} from '@heroicons/vue/outline'
import {changeNetwork} from '@/store/modules/web3'

export default {
  name: 'Header',
  components: {
    Switch,
    SunIcon,
    MoonIcon,
  },
  setup() {
    const store = useStore()
    store.commit('changeTheme', store.state.darkMode)
    const enabled = ref(store.state.darkMode)

    const modalAboutOpen = ref(false)

    function changeTheme(darkMode) {
      store.commit('changeTheme', darkMode)
    }

    return {
      enabled,
      modalAboutOpen,
      changeTheme,
      changeNetwork,
    }
  },
}
</script>
