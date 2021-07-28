<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-3/12">
      <Menu/>
    </div>
    <div class="w-9/12">
      <div class="text-sm font-semibold">{{ $store.state.space.name }}</div>
      <h2 class="font-semibold text-2xl">Proposals</h2>
      <div class="animate-pulse" v-if="loading">
        <div class="space-y-6 mt-4">
          <div class="h-8 bg-main-text rounded w-5/6"></div>
          <div class="h-8 bg-main-text rounded w-4/6"></div>
        </div>
      </div>
      <div class="space-y-6 mt-4" v-else>
        <div v-for="proposal in proposals" :key="proposal.id" class="p-5 border border-main-border rounded-md">
          <router-link class="space-y-2 hover:text-main-link" :to="{ name: 'Proposal', params: { id: proposal.id }}">
            <div class="flex justify-between">
              <div class="flex space-x-2 items-center mb-1">
                <img src="@/assets/imgs/token-etb.png" alt="logo" class="h-5 w-5">
                <div class="font-semibold text-sm">{{ $store.state.space.token }} by {{ _shorten(proposal.author) }}
                </div>
                <UiLabel class="badge-core ml-2">Core</UiLabel>
              </div>
              <UiLabel class="badge-state" :class="'badge-'+proposal.state">{{ capitalize(proposal.state) }}</UiLabel>
            </div>
            <h2 class="text-2xl font-semibold">{{ proposal.title }}</h2>
            <p class="text-lg">{{ _shorten(removeMarkdown(proposal.body), 140) }}</p>
            <div>{{ _dateAgo(proposal.end) }}</div>
          </router-link>
        </div>
        <div class="border border-main-border rounded-md p-6 font-semibold text-center w-full" v-if="proposals.length === 0">Oops, we can't find any results</div>
      </div>
    </div>
  </div>
</template>

<script>
import {ref, onMounted} from 'vue'
import axios from 'axios'
import removeMd from 'remove-markdown'

export default {
  name: 'Home',
  setup() {
    const loading = ref(true)
    const proposals = ref([])

    async function getProposals() {
      axios.get(`${process.env.VUE_APP_HUB_URL}/proposals`)
        .then((response) => {
          proposals.value = response.data.proposals || []
          proposals.value.forEach(proposal => {
            proposal.state = proposal.end > Date.now() ? 'active' : 'closed'
          })
        }).catch((error) => {
        console.error(error)
      })
    }

    function capitalize(value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    }

    async function load() {
      loading.value = true
      await getProposals()
      loading.value = false
    }

    function removeMarkdown(body) {
      return removeMd(body)
    }

    onMounted(load)

    return {
      loading,
      proposals,
      capitalize,
      removeMarkdown,
    }
  }
}
</script>
