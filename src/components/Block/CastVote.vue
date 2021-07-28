<template>
  <div class="panel">
    <div class="panel-title">Cast your vote</div>
    <div class="panel-body space-y-2">
      <UiButton
        v-for="(choice, i) in proposal.choices"
        :key="i"
        @click="selectChoice(i+1)"
        :class="selectedChoice == i +1 && 'active'"
      >
        {{ _shorten(choice, 32) }}
      </UiButton>
      <div>
        <UiButton :disabled="selectedChoice == 0" @click="modalConfirmOpen = true" class="mt-3">Vote</UiButton>
      </div>
    </div>
  </div>

  <ModalConfirm
    :open="modalConfirmOpen"
    :proposal="proposal"
    :choice="selectedChoice"
    @close="modalConfirmOpen= false"
    @voted="$emit('voted')"
  />
</template>

<script>
import {ref} from 'vue'

export default {
  name: 'CastVote',
  props: {
    proposal: {
      type: Object,
      required: true,
    }
  },
  emits: ['voted'],
  setup(_, {emit}) {
    const selectedChoice = ref(0)
    const modalConfirmOpen = ref(false)

    function selectChoice(i) {
      selectedChoice.value = i
      emit('selectedChoice', i)
    }

    return {
      selectedChoice,
      selectChoice,
      modalConfirmOpen,
    }
  }
}
</script>

<style lang="stylus" scoped>
.active
  @apply border-main-link
</style>
