<template>
  <div class="fixed inset-x-0 bottom-0 text-center z-50">
    <div class="mb-4 space-y-2">
      <div v-for="(item, i) in items" :key="i">
        <div class="inline-block py-2 px-5 rounded-full text-center text-white font-semibold cursor-pointer"
             :class="item.color"
             v-if="now<item.timestamp + duration && !item.hide"
             @click="item.hide = true"
        >
          {{ item.message }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ref, computed} from 'vue'
import {useStore} from 'vuex'

export default {
  name: 'Notifications',
  setup() {
    const store = useStore()
    const now = ref(Date.now())
    const duration = ref(40000)

    const items = computed(() => store.state.notifs)

    setInterval(() => (now.value = Date.now()), 1000)

    return {
      items,
      now,
      duration,
    }
  }
}
</script>
