<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <div class="border border-main-border rounded-md px-6 py-5 w-full flex items-center" v-if="!isAdmin">
        <InformationCircleIcon class="w-4 h-4 mr-2"/>
        You need to be in the admin list in order to submit a proposal.
      </div>
      <input v-model="form.title" type="text"
             class="block w-full outline-none text-3xl bg-transparent text-main-link font-semibold"
             placeholder="Question">
      <TextareaAutosize
        v-model="form.body"
        placeholder="What is your proposal?"
        class="block outline-none resize-none bg-transparent text-main-link font-semibold w-full h-16 tracking-wider"/>
      <div v-if="form.body.length > bodyLimit" class="text-red-500">
        -{{ _n(-(bodyLimit - form.body.length)) }}
      </div>
      <div v-if="!!form.body" class="space-y-5">
        <div class="text-xl font-semibold text-main-heading">Preview</div>
        <UiMarkdown :body="form.body" class="space-y-5"></UiMarkdown>
      </div>
      <div class="panel">
        <div class="panel-title">Choices</div>
        <div class="panel-body">
          <div class="flex-col space-y-2">
            <div
              class="border border-main-border hover:border-main-link rounded-full text-main-link flex justify-between items-center py-1 px-5 font-semibold cursor-pointer"
              v-for="(choice, i) in form.choices" :key="i">
              <div>{{ i + 1 }}</div>
              <div class="w-full px-2"><input type="text" v-model="form.choices[i]"
                                              class="outline-none w-full bg-transparent text-center h-9 font-semibold text-main-link">
              </div>
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
          <UiButton>
            <input v-model="form.snapshot"
                   class="outline-none h-6 text-center font-semibold text-main-link bg-transparent">
          </UiButton>
          <UiButton @click="publish" :disabled="!isValid">Publish</UiButton>
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
import {ref, computed, onMounted} from 'vue'
import {useStore} from 'vuex'
import {useRouter} from 'vue-router'
import axios from 'axios'
import {XIcon, SparklesIcon, InformationCircleIcon} from '@heroicons/vue/outline'
import {provider} from '@/store/modules/web3'

export default {
  name: 'Create',
  components: {
    XIcon, SparklesIcon, InformationCircleIcon
  },
  setup() {
    const store = useStore()
    const router = useRouter()

    const selectedDate = ref('')
    const form = ref({
      title: '',
      body: '',
      choices: ['', ''],
      start: 0,
      end: 0,
      snapshot: 0,
    })
    const bodyLimit = ref(1e4)

    const modalSelectDateOpen = ref(false)

    function setDate(ts) {
      if (selectedDate.value) {
        form.value[selectedDate.value] = ts;
      }
    }

    function addChoice() {
      form.value.choices.push('')
    }

    function removeChoice(i) {
      form.value.choices.splice(i, 1);
    }

    const isAdmin = computed(() => {
      if (!store.state.web3.connected) {
        return false
      }
      return store.state.space.admins.map(admin => admin.toLowerCase()).includes(store.state.web3.address.toLowerCase())
    })

    const isValid = computed(() => {
      return isAdmin.value &&
        form.value.title &&
        form.value.body &&
        form.value.body.length <= bodyLimit.value &&
        form.value.choices.length >= 2 &&
        !form.value.choices.some(a => a === '') &&
        form.value.start > 0 &&
        form.value.end > 0 &&
        form.value.end > form.value.start &&
        form.value.snapshot > 0
    })

    async function publish() {
      const proposal = {
        author: store.state.web3.address,
        title: form.value.title,
        body: form.value.body,
        choices: form.value.choices,
        created: Date.now(),
        start: form.value.start * 1e3,
        end: form.value.end * 1e3,
        snapshot: form.value.snapshot,
      }

      const signer = provider.getSigner()
      const signature = await signer.signMessage(JSON.stringify(proposal))

      axios.post(`${process.env.VUE_APP_HUB_URL}/proposal`, {
        signature: signature,
        proposal: proposal,
      }).then((response) => {
        router.push({name: 'Proposal', params: {id: response.data}})
      }).catch((error) => {
        store.dispatch('notify', ["Oops, " + error.response.data, 'bg-red-500'])
        console.error(error)
      })
    }

    onMounted(async () => {
      setInterval(async () => {
        if (form.value.snapshot == 0) {
          form.value.snapshot = await provider.getBlockNumber()
        }
      }, 1000)
      form.value.snapshot = await provider.getBlockNumber()
    })

    return {
      isAdmin,
      form,
      selectedDate,
      modalSelectDateOpen,
      setDate,
      addChoice,
      removeChoice,
      isValid,
      bodyLimit,
      publish,
    }
  }
}
</script>
