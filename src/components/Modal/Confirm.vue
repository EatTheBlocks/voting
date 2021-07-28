<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>Confirm vote</template>
    <div class="text-center font-semibold text-lg">
      <p>Are you sure you want to vote "{{ _shorten(proposal.choices[choice - 1], 32) }}"?</p>
      <p>This action cannot be undone</p>
    </div>
    <div class="mt-5 border border-main-border rounded-md p-6 space-y-1">
      <div class="flex justify-between">
        <div class="font-semibold">Option</div>
        <div class="text-main-link">{{ _shorten(proposal.choices[choice - 1], 32) }}</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Snapshot</div>
        <div class="text-main-link">
          <a :href="_explorer(56, proposal.snapshot, 'block')" target="_blank" class="flex items-center">
            {{ _n(proposal.snapshot, '0,0') }}
            <ExternalLinkIcon class="ml-1 h-4 w-4"/>
          </a>
        </div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Your voting power</div>
        <div class="text-main-link">{{ _n(power) }} {{ $store.state.space.token }}</div>
      </div>
    </div>
    <template #footer>
      <div class="flex space-x-5">
        <UiButton @click="$emit('close')">Cancel</UiButton>
        <UiButton @click="vote" :disabled="!canVote">Vote</UiButton>
      </div>
    </template>
  </UiModal>
</template>

<script>
import {ref, computed, onMounted} from 'vue'
import {useStore} from 'vuex'
import {useRouter} from 'vue-router'
import {ExternalLinkIcon} from '@heroicons/vue/outline'
import axios from 'axios'
import {provider} from '@/store/modules/web3'

export default {
  name: 'Confirm',
  components: {
    ExternalLinkIcon
  },
  props: {
    open: Boolean,
    proposal: Object,
    choice: Number,
  },
  emits: ['close', 'voted'],
  setup(props, {emit}) {
    const store = useStore()
    const router = useRouter()

    const power = ref(0)

    const address = computed(() => {
      //return "0x9891d990a6608c8501145ab064b6c14a2e4f0e5e"
      return store.state.web3.address
    })

    async function votingPower() {
      if (!store.state.web3.connected) {
        return 0
      }
      const response = await axios.post(`${process.env.VUE_APP_HUB_URL}/score`, {
        snapshot: props.proposal.snapshot,
        addresses: [address.value],
      })

      power.value = response.data.result.scores[0][address.value]
    }

    const canVote = computed(() => {
      return power.value > 0
    })

    async function vote() {
      const vote = {
        author: store.state.web3.address,
        proposal: props.proposal.id,
        choice: props.choice,
        timestamp: Date.now(),
      }

      const signer = provider.getSigner()
      const signature = await signer.signMessage(JSON.stringify(vote))

      axios.post(`${process.env.VUE_APP_HUB_URL}/vote`, {
        signature: signature,
        vote: vote,
      }).then(() => {
        router.push({name: 'Proposal', params: {id: props.proposal.id}})
        emit('close')
        emit('voted')
      }).catch((error) => {
        store.dispatch('notify', ["Oops, " + error.response.data, 'bg-red-500'])
        console.error(error)
        emit('close')
      })
    }

    onMounted(votingPower)

    return {
      power,
      canVote,
      vote,
    }
  }
}
</script>
