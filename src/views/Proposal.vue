<template>
  <div>
    <div class="flex space-x-10">
      <div class="w-9/12">
        <div>Back</div>
        <div>{{ proposal.title }}</div>
        <div>{{ proposal.state }}</div>
        <div v-html="markdown(proposal.body)"></div>
        <div class="panel">
          <div class="panel-title">Votes</div>
          <div class="panel-body" v-for="vote in votes" :key="vote.author">
            BLOCKIE {{ vote.author }} {{ vote.choice }} {{ vote.tokens }} {{ vote.receipt }}
          </div>
        </div>
      </div>
      <div class="w-3/12 space-y-5">
        <div class="panel">
          <div class="panel-title">Information</div>
          <div class="panel-body table">
            <div>
              <div>Author</div>
              <div class="flex items-center">
                <Blockie class="rounded-full h-4 w-4 mr-2" :opts="{seed:proposal.author}"></Blockie>
                {{ ethShortAddress(proposal.author) }}
                <span class="ml-2 px-1.5 py-0.5 rounded-full border border-gray-200 text-xs">{{ proposal.label.text }}</span>
              </div>
            </div>
            <div>
              <div>IPFS</div>
              <div class="flex items-center">
                #{{ proposal.id.substring(0, 7) }}<ExternalLinkIcon class="ml-1 h-4 w-4"/>
              </div>
            </div>
            <div>
              <div>Start date</div>
              <div>{{ proposal.start }}</div>
            </div>
            <div>
              <div>End date</div>
              <div>{{ proposal.end }}</div>
            </div>
            <div>
              <div>Snapshot</div>
              <div class="flex items-center">
                123456<ExternalLinkIcon class="ml-1 h-4 w-4"/>
              </div>
            </div>
          </div>
        </div>
        <div class="panel">
          <div class="panel-title">Results</div>
          <div class="panel-body">
            YES x TOKENS x%
            NO x TOKENS x%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Blockie from '@/components/Blockie'
import {ExternalLinkIcon} from '@heroicons/vue/outline'
import DOMPurify from 'dompurify'
import marked from 'marked'

let proposal = {
  id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
  label: {color: "blue", text: "Core"},
  title: "Test proposal",
  body: "# Summary\n\nThis proposal is to approve an upgrade for the compensation of Byterose for the next 3 months.\n\n# Abstract\n\n@e: Byterose",
  start: Date.now(),
  end: Date.now() - 24 * 60 * 60 * 1000,
  state: "Closed",
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
    Blockie,
    ExternalLinkIcon,
  },
  data() {
    return {
      proposal: proposal,
      votes: votes,
    }
  },
  methods: {
    markdown(text) {
      return marked(DOMPurify.sanitize(text))
    }
  }
}
</script>

<style lang="stylus">
h1
  @apply text-xl text-gray-800

.panel
  @apply border border-gray-300 rounded-md

  &-title
    @apply border-b border-gray-300 bg-gray-100 rounded-t-md font-bold text-lg

  &-title, &-body
    @apply p-4

  &-body
    @apply w-full

    &.table
      @apply space-y-1
      > div
        @apply flex justify-between w-full

        div:nth-child(1)
          @apply font-semibold
</style>
