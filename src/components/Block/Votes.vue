<template>
  <div class="panel">
    <div class="flex panel-title items-center">
      Votes
      <UiLabel class="badge-core ml-2 mt-0.5 bg-main-bg">125</UiLabel>
    </div>
    <div class="panel-body p-0 divide-y divide-main-border">
      <div class="flex justify-between p-4" v-for="vote in votes" :key="vote.id">
        <div class="flex items-center">
          <User :address="vote.voter"/>
        </div>
        <div>{{ proposal.choices[vote.choice] }}</div>
        <div class="flex items-center">
          {{ vote.tokens }} {{ $TokenName }}
          <BadgeCheckIcon class="ml-1 h-5 w-5 cursor-pointer text-green-500" @click="openReceiptModal(vote)"/>
        </div>
      </div>
    </div>
    <div class="panel-footer cursor-pointer">See more</div>
  </div>

  <ModalReceipt
    :open="modalReceiptOpen"
    @close="modalReceiptOpen = false"
    :ipfsHash="ipfsHash"
  />
</template>

<script>
import {ref} from 'vue'
import {BadgeCheckIcon} from '@heroicons/vue/outline'

export default {
  name: 'Votes',
  props: {
    proposal: Object,
    votes: Object,
    loaded: Boolean,
  },
  setup() {
    const modalReceiptOpen = ref(false)
    const ipfsHash = ref('')

    return {
      modalReceiptOpen,
      ipfsHash,
    }
  },
  components: {
    BadgeCheckIcon,
  },
  methods: {
    openReceiptModal(vote) {
      this.ipfsHash = vote.id
      this.modalReceiptOpen = true
    }
  }
}
</script>
