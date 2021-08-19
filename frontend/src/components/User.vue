<template>
  <UiPopover :options="{ offset: [0, 12], placement: 'bottom-start' }" v-if="popover">
    <template #item>
      <a class="whitespace-nowrap cursor-pointer">
        <UiAvatar class="mr-2 -mt-0.5" :opts="{seed:address, size:14}"/>
        {{ ens ? ens : _shorten(address) }}
      </a>
    </template>
    <template #content>
      <div class="p-4 text-center">
        <UiAvatar class="mb-4" :opts="{seed:address, size:64}"/>
        <h3>{{ ens ? ens : _shorten(address) }}</h3>
        <a :href="_explorer(56, address)" target="_blank">
          <UiButton class="flex items-center justify-center mt-5">
            See on explorer
            <ExternalLinkIcon class="ml-1 h-4 w-4"/>
          </UiButton>
        </a>
      </div>
    </template>
  </UiPopover>
  <div v-else>
    <a class="whitespace-nowrap cursor-pointer">
      <UiAvatar class="mr-2 -mt-0.5" :opts="{seed:address, size:14}"/>
      {{ ens ? ens : _shorten(address) }}
    </a>
  </div>
</template>

<script>
import {computed} from 'vue'
import {useStore} from 'vuex'
import {ExternalLinkIcon} from '@heroicons/vue/outline'

export default {
  name: 'User',
  components: {
    ExternalLinkIcon,
  },
  props: {
    address: String,
    popover: {
      type: Boolean,
      default: true
    }
  },
  setup(props) {
    const store = useStore()

    const ens = computed(() => {
      let ens = store.state.ens[props.address.toLowerCase()]

      if (!ens) {
        store.dispatch('getENS', props.address.toLowerCase())
      }
      return ens
    })

    return {
      ens,
    }
  }
}
</script>
