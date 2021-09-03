<template>
  <div class="content flex flex-col min-h-screen bg-main-bg">
    <Header></Header>
    <div class="flex-grow p-5">
      <UiLoading v-if="loading" class="overlay big" />
      <router-view v-else/>
    </div>
    <Footer></Footer>
    <div id="modal"/>
    <Notifications />
  </div>
</template>

<script>

import {ref, onMounted} from 'vue'
import {useStore} from 'vuex'
import axios from 'axios'

export default {
  name: 'app',
  setup() {
    const store = useStore()

    const loading = ref(true)
    const loadError = ref(false)

    async function getSpace() {
      try {
        const response = await axios.get(`${process.env.VUE_APP_HUB_URL}`)
        store.commit('setSpace', response.data)
      } catch (error) {
        await store.dispatch('notify', ["Oops, " + error.response.data.message, 'bg-red-500'])
        loadError.value = error
      }
    }

    async function load() {
      loading.value = true
      await getSpace()
      loading.value = false
    }

    store.commit('emptyNotifs')

    onMounted(load)

    return {
      loading,
      loadError,
    }
  }
};
</script>
