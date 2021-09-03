<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <div class="animate-pulse" v-if="loading">
        <div class="space-y-5">
          <div class="h-8 bg-main-text rounded w-5/6"></div>
          <div class="h-8 bg-main-text rounded w-1/6"></div>
          <div class="h-8 bg-main-text rounded w-2/6"></div>
        </div>
      </div>
      <div class="space-y-5" v-else>
        <h1 class="mt-5 text-3xl font-bold">{{ proposal.title }}</h1>
        <UiLabel class="inline-block badge-state" :class="'badge-'+proposal.state">{{
            capitalize(proposal.state)
          }}
        </UiLabel>
        <UiMarkdown :body="proposal.body" class="space-y-5"></UiMarkdown>
        <BlockCastVote
          :proposal="proposal"
          @voted="load"
          v-if="proposal.state === 'active' && !alreadyVoted"
        />
        <BlockVotes :proposal="proposal" :votes="votes" v-if="!loadingResults && votes.length > 0"/>
      </div>
    </div>
    <div class="w-4/12 space-y-5" v-if="!loading">
      <div class="panel">
        <div class="panel-title">Information</div>
        <div class="panel-body table">
          <div>
            <div>Author</div>
            <div class="flex items-center">
              <User :address="proposal.author"/>
              <UiLabel class="badge-core ml-2">Core</UiLabel>
            </div>
          </div>
          <div>
            <div>IPFS</div>
            <div>
              <a :href="_ipfsUrl(proposal.id)" target="_blank" class="flex items-center">
                #{{ proposal.id.substring(0, 7) }}
                <ExternalLinkIcon class="ml-1 h-4 w-4"/>
              </a>
            </div>
          </div>
          <div>
            <div>Start date</div>
            <div>{{ $d(proposal.start, 'short') }}</div>
          </div>
          <div>
            <div>End date</div>
            <div>{{ $d(proposal.end, 'short') }}</div>
          </div>
          <div>
            <div>Snapshot</div>
            <div>
              <a :href="_explorer(56, proposal.snapshot, 'block')" target="_blank" class="flex items-center">
                {{ _n(proposal.snapshot, '0,0') }}
                <ExternalLinkIcon class="ml-1 h-4 w-4"/>
              </a>
            </div>
          </div>
        </div>
      </div>
      <BlockResults
        :loading="loadingResults"
        :proposal="proposal"
        :votes="votes"
        :results="results"
      />
    </div>
  </div>
</template>

<script>
import {ref, computed, onMounted} from 'vue'
import {useStore} from 'vuex'
import {useRoute} from 'vue-router'
import axios from 'axios'
import {ExternalLinkIcon} from '@heroicons/vue/outline'

export default {
  name: 'Proposal',
  components: {
    ExternalLinkIcon,
  },
  setup() {
    const store = useStore()

    const loading = ref(true)
    const loadingResults = ref(true)
    const proposal = ref({
      id: '',
      title: '',
      body: '',
      choices: [],
      created: 0,
      start: 0,
      end: 0,
    })
    const votes = ref([])
    const alreadyVoted = ref(false)
    const results = ref({})

    const route = useRoute()
    const id = route.params.id

    async function getProposal() {
      try {
        const response = await axios.get(`${process.env.VUE_APP_HUB_URL}/proposal/` + id)
        proposal.value = response.data.proposal
        proposal.value.state = proposal.value.end > Date.now() ? 'active' : 'closed'
        votes.value = response.data.votes || []
      } catch (error) {
        await store.dispatch('notify', ["Oops, " + error.response.data.message, 'bg-red-500'])
        console.error(error)
      }
    }

    async function getScores() {
      try {
        const responseScore = await axios.post(`${process.env.VUE_APP_HUB_URL}/score`, {
          snapshot: proposal.value.snapshot,
          addresses: votes.value.map(v => v.author),
        })

        votes.value = votes.value.map(vote => {
          vote.balance = responseScore.data.scores[vote.author] || 0
          return vote
        })
          .sort((a, b) => b.balance - a.balance)
        //.filter(vote => vote.balance > 0)

        const sumOfResultsBalance = votes.value.reduce((a, b) => a + b.balance, 0)
        const resultsByVoteBalance = proposal.value.choices.map((choice, i) => {
          return votes.value.filter(vote => vote.choice === i + 1).reduce((a, b) => a + b.balance, 0)
        })

        results.value = {
          sumOfResultsBalance: sumOfResultsBalance,
          resultsByVoteBalance: resultsByVoteBalance,
        }
      } catch (error) {
        await store.dispatch('notify', ["Oops, " + error.response.data.message, 'bg-red-500'])
        console.error(error)
      }
    }

    function capitalize(value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    }

    const address = computed(() => {
      return store.state.web3.address
    })

    async function load() {
      loading.value = true
      loadingResults.value = true

      await getProposal()
      if (votes.value) {
        alreadyVoted.value = votes.value.some(v => v.author === store.state.web3.address)
      }
      loading.value = false

      await getScores()
      loadingResults.value = false
    }

    onMounted(load)

    return {
      loading,
      loadingResults,
      load,
      proposal,
      votes,
      alreadyVoted,
      results,
      capitalize,
      address,
    }
  },
}
</script>
