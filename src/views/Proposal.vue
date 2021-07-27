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
        <div v-html="markdown(proposal.body)" class="space-y-5 markdown"></div>
        <BlockCastVote :proposal="proposal"/>
        <BlockVotes :proposal="proposal" :votes="votes" :loaded="true"/>
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
              <a :href="_explorer(56, 12345678, 'block')" target="_blank" class="flex items-center">
                {{ _n(12345678, '0,0') }}
                <ExternalLinkIcon class="ml-1 h-4 w-4"/>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div class="panel">
        <div class="panel-title">Results</div>
        <div class="panel-body">
          <div class="animate-pulse" v-if="loadingResults">
            <div class="space-y-3">
              <div class="h-6 bg-main-text rounded w-5/6"></div>
              <div class="h-6 bg-main-text rounded w-3/6"></div>
            </div>
          </div>
          <div class="space-y-3" v-else>
            <UiProgress :text="proposal.choices[0] + ' 123.25k ' + $TokenName" :percent="40"></UiProgress>
            <UiProgress :text="proposal.choices[1] + ' 853.86k ' + $TokenName" :percent="60"></UiProgress>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ref, onMounted} from 'vue'
import {useRoute} from 'vue-router'
import axios from 'axios'
import {ExternalLinkIcon} from '@heroicons/vue/outline'
import DOMPurify from 'dompurify'
import marked from 'marked'

let votes = [
  {
    id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    voter: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    created: 1625695452,
    choice: 0,
    tokens: 100,
  },
  {
    id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    voter: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    created: 1625695452,
    choice: 1,
    tokens: 200,
  }
]

export default {
  name: 'Proposal',
  components: {
    ExternalLinkIcon,
  },
  setup() {
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

    const route = useRoute()
    const id = route.params.id

    async function getProposals() {
      axios.get(`${process.env.VUE_APP_HUB_URL}/proposal/` + id)
        .then((response) => {
          proposal.value = response.data.proposal
          proposal.value.state = proposal.value.end > Date.now() ? 'active' : 'closed'
        }).catch((error) => {
        console.error(error)
      })
    }

    function markdown(text) {
      return marked(DOMPurify.sanitize(text))
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

    onMounted(load)

    return {
      loading,
      loadingResults,
      proposal,
      votes: votes,
      markdown,
      capitalize,
    }
  },
}
</script>
