<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <div class="animate-pulse" :class="loader?'':'hidden'">
        <div class="space-y-5">
          <div class="h-8 bg-main-text rounded w-5/6"></div>
          <div class="h-8 bg-main-text rounded w-1/6"></div>
          <div class="h-8 bg-main-text rounded w-2/6"></div>
        </div>
      </div>
      <div class="space-y-5" :class="loader?'hidden':''">
        <h1 class="mt-5 text-3xl font-bold">{{ proposal.title }}</h1>
        <UiLabel class="inline-block badge-state">{{ proposal.state }}</UiLabel>
        <div v-html="markdown(proposal.body)" class="space-y-5 markdown"></div>
        <BlockVotes :proposal="proposal" :votes="votes" :loaded="true" ></BlockVotes>
      </div>
    </div>
    <div class="w-4/12 space-y-5" :class="loader?'hidden':''">
      <div class="panel">
        <div class="panel-title">Information</div>
        <div class="panel-body table">
          <div>
            <div>Author</div>
            <div class="flex items-center">
              <User :address="proposal.author"/>
              <UiLabel class="badge-core ml-2">{{ proposal.label.text }}</UiLabel>
            </div>
          </div>
          <div>
            <div>IPFS</div>
            <div class="flex items-center">
              #{{ proposal.id.substring(0, 7) }}
              <ExternalLinkIcon class="ml-1 h-4 w-4"/>
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
            <div class="flex items-center">
              12,345,678
              <ExternalLinkIcon class="ml-1 h-4 w-4"/>
            </div>
          </div>
        </div>
      </div>
      <div class="panel">
        <div class="panel-title">Results</div>
        <div class="panel-body">
          <div class="animate-pulse" :class="loaderResults?'':'hidden'">
            <div class="space-y-3">
              <div class="h-6 bg-main-text rounded w-5/6"></div>
              <div class="h-6 bg-main-text rounded w-3/6"></div>
            </div>
          </div>
          <div class="space-y-3" :class="loaderResults?'hidden':''">
            <UiProgress :text="proposal.choices[0] + ' 123.25k ' + $TokenName" :percent="40"></UiProgress>
            <UiProgress :text="proposal.choices[1] + ' 853.86k ' + $TokenName" :percent="60"></UiProgress>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ExternalLinkIcon} from '@heroicons/vue/outline'
import DOMPurify from 'dompurify'
import marked from 'marked'

let proposal = {
  id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
  label: {color: "blue", text: "Core"},
  title: "YIP-62: Change Two Multisig Signers",
  body: "# Summary\n\nThis proposal is to approve an upgrade for the compensation of Byterose for the next 3 months.\n\n# Abstract\n\n@e: Byterose",
  start: Date.now(),
  end: Date.now() - 24 * 60 * 60 * 1000,
  state: "Closed",
  choices: [
    "Approve",
    "Reject"
  ]
}

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
  data() {
    return {
      proposal: proposal,
      votes: votes,
      loader: false,
      loaderResults: true,
    }
  },
  mounted() {
    setTimeout(() => {
      this.loader = false;
    }, 500)

    setTimeout(() => {
      this.loaderResults = false;
    }, 500)
  },
  methods: {
    markdown(text) {
      return marked(DOMPurify.sanitize(text))
    }
  }
}
</script>
