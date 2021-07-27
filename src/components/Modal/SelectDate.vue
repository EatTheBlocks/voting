<template>
  <UiModal :open="open" @close="$emit('close')">
    <template #header>
      <span v-if="step == 0">Select {{ selectedDate }} date</span>
      <span v-else>Select {{ selectedDate }} time</span>
    </template>

    <UiCalendar v-model="input" v-if="step == 0"/>

    <div v-else class="mx-auto w-36">
      <div class="border border-main-border rounded-full flex py-3">
        <input v-model="form.h" max="24" class="w-16 border-0 bg-transparent text-main-link text-center outline-none">
        <span>:</span>
        <input v-model="form.m" max="60" class="w-16 border-0 bg-transparent text-main-link text-center outline-none">
      </div>
    </div>

    <template #footer>
      <div class="flex space-x-5">
        <UiButton @click="$emit('close')">Cancel</UiButton>
        <UiButton @click="handleSubmit" :disabled="!input">
          <span v-if="step==0">Next</span>
          <span v-else>Select</span>
        </UiButton>
      </div>
    </template>
  </UiModal>
</template>

<script>
import {ref, toRefs, watch} from 'vue'

export default {
  name: 'SelectDate',
  props: {
    open: Boolean,
    value: Number,
    selectedDate: String
  },
  emits: ['input', 'close'],
  setup(props, { emit }) {
    const {open} = toRefs(props)
    const step = ref(0)
    const input = ref('')
    const form = ref({
      h: '12',
      m: '00'
    })

    function handleSubmit() {
      if (step.value === 0) return (step.value = 1)
      const [year, month, day] = input.value.split('-')
      let timestamp = new Date(
        year,
        month - 1,
        day,
        form.value.h,
        form.value.m,
        0
      )
      timestamp = new Date(timestamp).getTime() / (1e3).toFixed()
      emit('input', timestamp)
      emit('close')
    }

    function formatDate(date) {
      const output = {h: '12', m: '00', dateString: ''}
      if (!date) return output
      const dateObject = new Date(date * 1000)
      const offset = dateObject.getTimezoneOffset()
      const data = new Date(dateObject.getTime() - offset * 60 * 1000)
      output.dateString = data.toISOString().split('T')[0]
      output.h = ('0' + dateObject.getHours().toString()).slice(-2)
      output.m = ('0' + dateObject.getMinutes().toString()).slice(-2)
      return output
    }

    watch(open, () => {
      const {dateString, h, m} = formatDate(props.value)
      step.value = 0
      form.value = {h, m}
      input.value = dateString
    });

    return {
      step,
      input,
      form,
      handleSubmit,
    }
  }
}
</script>

<style scoped>

</style>
