<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <input type="text" class="block w-full outline-none text-3xl bg-transparent text-main-link font-semibold"
             placeholder="Question">
      <TextareaAutosize
        v-model="form.body"
        placeholder="What is your proposal?"
        class="block outline-none resize-none bg-transparent text-main-link font-semibold w-full h-16 tracking-wider"></TextareaAutosize>
      <div v-if="!!form.body" class="space-y-5">
        <div class="text-xl font-semibold text-main-heading">Preview</div>
        <div v-html="markdown" class="space-y-5 markdown"></div>
      </div>
      <div class="panel">
        <div class="panel-title">Choices</div>
        <div class="panel-body">
          <div class="flex-col space-y-2">
            <div
              class="border border-main-border hover:border-main-link rounded-full text-main-link flex justify-between items-center py-1 px-5 font-semibold cursor-pointer"
              v-for="(choice, i) in choices" :key="i">
              <div>{{ i + 1 }}</div>
              <div class="w-full px-2"><input type="text" v-model="choices[i]" class="outline-none w-full bg-transparent text-center h-9 font-semibold text-main-link"></div>
              <div>
                <XIcon @click="removeChoice(i)" class="w-4 h-4"/>
              </div>
            </div>
            <div>
              <UiButton @click="addChoice" class="mt-3">Add choice</UiButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="w-4/12 space-y-5">
      <div class="panel">
        <div class="panel-title flex justify-between items-center">
          <span>Actions</span>
          <SparklesIcon class="w-4 h-4"/>
        </div>
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
import {ref, computed} from 'vue'
import DOMPurify from 'dompurify'
import marked from 'marked'
import {XIcon, SparklesIcon} from '@heroicons/vue/outline'

export default {
  name: 'Create',
  components: {
    XIcon, SparklesIcon,
  },
  setup() {
    const selectedDate = ref('')
    const form = ref({
      name: '',
      body: '',
      choices: [],
      start: 0,
      end: 0,
    })

    const markdown = computed(() => {
      return marked(DOMPurify.sanitize(form.value.body))
    })

    const choices = ref(['', ''])

    const modalSelectDateOpen = ref(false)

    function setDate(ts) {
      if (selectedDate.value) {
        form.value[selectedDate.value] = ts;
      }
    }

    function addChoice() {
      choices.value.push('')
    }

    function removeChoice(i) {
      choices.value.splice(i, 1);
    }

    return {
      form,
      markdown,
      selectedDate,
      modalSelectDateOpen,
      setDate,
      choices,
      addChoice,
      removeChoice,
    }
  }
}
</script>

<style scoped>

</style>
