<template>
  <router-link :to="to" class="flex items-center font-semibold text-sm">
    <ArrowNarrowLeftIcon class="mr-1 h-4 w-4"/>
    {{ text }}
  </router-link>
</template>

<script>
import {ArrowNarrowLeftIcon} from '@heroicons/vue/outline'

export default {
  name: 'BackButton',
  props: {
    to: Object,
    text: {
      type: String,
      default: 'Back'
    },
  },
  components: {
    ArrowNarrowLeftIcon
  }
}
</script>

<style scoped>

</style>
