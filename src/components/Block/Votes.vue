<template>
  <div class="panel">
    <div class="flex panel-title items-center">
      Votes
      <UiLabel class="badge-core ml-2 mt-0.5 bg-main-bg">{{ votes ? votes.length : 0 }}</UiLabel>
    </div>
    <div class="panel-body p-0 divide-y divide-main-border">
      <div class="flex p-4" v-for="vote in visibleVotes" :key="vote.id">
        <div class="flex-none flex items-center w-36">
          <User :address="vote.author"/>
        </div>
        <div class="flex-grow flex items-center justify-center">{{ proposal.choices[vote.choice - 1] }}</div>
        <div class="flex-none flex items-center justify-end w-36">
          {{ _n(vote.balance) }} {{ $store.state.space.token }}
          <BadgeCheckIcon class="ml-1 h-5 w-5 cursor-pointer text-green-500" @click="openReceiptModal(vote)"/>
        </div>
      </div>
    </div>
    <div class="panel-footer cursor-pointer"
         @click="showAllVotes = true"
         v-if="!showAllVotes && votes.length > 10">See more</div>
  </div>

  <ModalReceipt
    :open="modalReceiptOpen"
    @close="modalReceiptOpen = false"
    :ipfsHash="ipfsHash"
  />
</template>

<script>
import {computed, ref} from 'vue'
import {BadgeCheckIcon} from '@heroicons/vue/outline'
import {useStore} from "vuex";

export default {
  name: 'Votes',
  components: {
    BadgeCheckIcon,
  },
  props: {
    proposal: Object,
    votes: Array,
  },
  setup(props) {
    const store = useStore()

    const modalReceiptOpen = ref(false)
    const ipfsHash = ref('')
    const showAllVotes = ref(false)

    function openReceiptModal(vote) {
      ipfsHash.value = vote.id
      modalReceiptOpen.value = true
    }

    const visibleVotes = computed(() => {
      return showAllVotes.value ? sortVotesUserFirst() : sortVotesUserFirst().slice(0, 10)
    })

    const address = computed(() => store.state.web3.address);

    function sortVotesUserFirst() {
      const votes = props.votes;
      if (votes.map(vote => vote.author).includes(address.value)) {
        votes.unshift(
          votes.splice(
            votes.findIndex(item => item.author === address.value),
            1
          )[0]
        );
        return votes
      }
      return votes
    }

    return {
      modalReceiptOpen,
      openReceiptModal,
      showAllVotes,
      ipfsHash,
      visibleVotes,
    }
  }
}
</script>
