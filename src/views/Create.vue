<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <input type="text" class="block outline-none text-3xl bg-transparent text-main-link font-semibold"
             placeholder="Question">
      <TextareaAutosize
        v-model="form.body"
        placeholder="What is your proposal?"
        class="block outline-none resize-none bg-transparent text-main-link font-semibold w-full h-16 tracking-wider"></TextareaAutosize>
      <div v-if="form.body.length > 0" class="space-y-5">
        <div class="text-xl font-semibold text-main-heading">Preview</div>
        <div v-html="markdown" class="space-y-5 markdown"></div>
      </div>
      <div class="panel">
        <div class="panel-title">Choices</div>
        <div class="panel-body">YOLO</div>
      </div>
    </div>
    <div class="w-4/12 space-y-5">
      <div class="panel">
        <div class="panel-title">Actions</div>
        <div class="panel-body space-y-3">
          <UiButton
            @click="(modalSelectDateOpen = true), (selectedDate = 'start')"
          >
            <span v-if="!form.start">Select start date</span>
            <span v-else>{{ $d(form.start * 1e3, 'short') }}</span>
          </UiButton>
          <UiButton
            @click="(modalSelectDateOpen = true), (selectedDate = 'end')"
          >
            <span v-if="!form.end">Select end date</span>
            <span v-else>{{ $d(form.end * 1e3, 'short') }}</span>
          </UiButton>
          <UiButton :disabled="true">Publish</UiButton>
        </div>
      </div>
    </div>
  </div>

  <ModalSelectDate
    :open="modalSelectDateOpen"
    :value="form[selectedDate]"
    :selectedDate="selectedDate"
    @close="modalSelectDateOpen= false"
    @input="setDate"
  />
</template>

<script>
import {ref} from 'vue'
import DOMPurify from 'dompurify'
import marked from 'marked'

export default {
  name: 'Create',
  setup() {
    const selectedDate = ref('')
    const form = ref({
      name: '',
      body: '',
      choices: [],
      start: 0,
      end: 0,
    })

    const modalSelectDateOpen = ref(false)

    function setDate(ts) {
      if (selectedDate.value) {
        form.value[selectedDate.value] = ts;
      }
    }

    return {
      form,
      selectedDate,
      modalSelectDateOpen,
      setDate,
    }
  },
  computed: {
    markdown() {
      return marked(DOMPurify.sanitize(this.body))
    }
  }
}
</script>

<style scoped>

</style>
