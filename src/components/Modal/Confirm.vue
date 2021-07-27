<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>Confirm vote</template>
    <div class="text-center font-semibold text-lg">
      <p>Are you sure you want to vote "{{ _shorten(proposal.choices[choice-1], 32) }}"?</p>
      <p>This action cannot be undone</p>
    </div>
    <div class="mt-5 border border-main-border rounded-md p-6 space-y-1">
      <div class="flex justify-between">
        <div class="font-semibold">Option</div>
        <div class="text-main-link">{{ _shorten(proposal.choices[choice-1], 32) }}</div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Snapshot</div>
        <div class="text-main-link">
          <a :href="_explorer(56, 12345678, 'block')" target="_blank" class="flex items-center">
            {{ _n(12345678, '0,0') }}
            <ExternalLinkIcon class="ml-1 h-4 w-4"/>
          </a>
        </div>
      </div>
      <div class="flex justify-between">
        <div class="font-semibold">Your voting power</div>
        <div class="text-main-link">0 {{ $TokenName }}</div>
      </div>
    </div>
    <template #footer>
      <div class="flex space-x-5">
        <UiButton @click="$emit('close')">Cancel</UiButton>
        <UiButton @click="handleSubmit" :disabled="!canVote">Vote</UiButton>
      </div>
    </template>
  </UiModal>
</template>

<script>
import {computed} from 'vue'
import {ExternalLinkIcon} from '@heroicons/vue/outline'

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
  emits: ['close'],
  setup() {
    const canVote = computed(() => {
      return false
    })

    function handleSubmit() {

    }

    return {
      canVote,
      handleSubmit,
    }
  }
}
</script>
