<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <router-link :to="{name: 'Home'}" class="flex items-center font-semibold text-sm">
        <ArrowNarrowLeftIcon class="mr-1 h-4 w-4"/> Back
      </router-link>
      <h1 class="mt-5 text-3xl font-bold">{{ proposal.title }}</h1>
      <div class="inline-block badge-state">{{ proposal.state }}</div>
      <div v-html="markdown(proposal.body)" class="space-y-5 markdown"></div>
      <div class="panel">
        <div class="panel-title">Votes</div>
        <div class="panel-body p-0 divide-y divide-main-border">
          <div class="flex justify-between p-4" v-for="vote in votes" :key="vote.id">
            <div class="flex items-center">
              <Blockie class="mr-2" :opts="{seed:vote.voter, size:16}"/>
              {{ ethShortAddress(vote.voter) }}
            </div>
            <div>{{ vote.choice }}</div>
            <div class="flex items-center">
              {{ vote.tokens }} {{ $TokenName }}
              <BadgeCheckIcon class="ml-1 h-5 w-5 cursor-pointer" @click="setIsOpen(true, vote.id)"/>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="w-4/12 space-y-5">
      <div class="panel">
        <div class="panel-title">Information</div>
        <div class="panel-body table">
          <div>
            <div>Author</div>
            <div class="flex items-center">
              <Blockie class="mr-2" :opts="{seed:proposal.author, size:16}"/>
              {{ ethShortAddress(proposal.author) }}
              <span class="badge-core ml-2">
                {{ proposal.label.text }}
              </span>
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
            <div>{{ proposal.start }}</div>
          </div>
          <div>
            <div>End date</div>
            <div>{{ proposal.end }}</div>
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
          YES x TOKENS x%
          NO x TOKENS x%
        </div>
      </div>
    </div>
  </div>
  <Dialog :open="isOpen" @close="setIsOpen" class="fixed inset-0 z-10 overflow-hidden flex items-center">
    <DialogOverlay class="fixed inset-0 bg-black opacity-40"/>
    <div class="flex flex-col w-full max-w-[440px] overflow-hidden z-20 mx-auto">
      <div class="panel rounded-lg relative">
        <div class="absolute top-O right-0 p-5">
          <button class="outline-none">
            <XIcon class="h-5 w-5 cursor-pointer" @click="setIsOpen(false)"/>
          </button>
        </div>
        <div class="panel-title rounded-t-lg text-center">Receipt</div>
        <div class="panel-body bg-white rounded-b-lg">
          Author {{ modalId }}
        </div>
      </div>
    </div>
  </Dialog>
</template>

<script>
import Blockie from '@/components/Blockie'
import {ArrowNarrowLeftIcon, ExternalLinkIcon, BadgeCheckIcon, XIcon} from '@heroicons/vue/outline'
import {ref} from 'vue'
import {Dialog, DialogOverlay} from '@headlessui/vue'
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
    ArrowNarrowLeftIcon, ExternalLinkIcon, BadgeCheckIcon, XIcon,
    Dialog, DialogOverlay
  },
  data() {
    return {
      proposal: proposal,
      votes: votes,
    }
  },
  setup() {
    let isOpen = ref(false)
    let modalId = ref("")

    return {
      isOpen,
      modalId,
      setIsOpen(value, id) {
        isOpen.value = value
        modalId.value = id
      },
    };
  },
  methods: {
    markdown(text) {
      return marked(DOMPurify.sanitize(text))
    }
  }
}
</script>

<style lang="stylus">
.panel
  @apply border border-main-border rounded-md divide-y divide-main-border

  &-title, &-body
    @apply px-4

  &-title
    @apply bg-main-block rounded-t-md font-bold text-lg text-main-heading py-3

  &-body
    @apply w-full py-5

    &.table
      @apply space-y-1

      > div
        @apply flex justify-between w-full

        div:nth-child(1)
          @apply font-semibold

        div:nth-child(2)
          @apply text-main-blockText
</style>
