<template>
  <Dialog :open="open" @close="$emit('close')" class="fixed inset-0 z-10 overflow-hidden flex items-center">
    <DialogOverlay class="fixed inset-0 bg-black opacity-40"/>
    <div class="flex flex-col w-full max-w-[440px] overflow-hidden z-20 mx-auto">
      <div class="panel overflow-hidden">
        <div class="panel-title text-center relative text-xl py-5">
          <slot name="header"/>
          <div class="absolute top-0 right-0 p-6">
            <button class="outline-none">
              <XIcon class="h-5 w-5 cursor-pointer" @click="$emit('close')"/>
            </button>
          </div>
        </div>
        <div class="panel-body bg-main-bg text-main-text">
          <slot/>
        </div>
      </div>
    </div>
  </Dialog>
</template>

<script>
import {XIcon} from '@heroicons/vue/outline'
import {Dialog, DialogOverlay} from '@headlessui/vue'

export default {
  name: 'Modal',
  props: {
    open: {
      type: Boolean,
      required: true
    }
  },
  emits: ['close'],
  components: {
    XIcon,
    Dialog, DialogOverlay
  }
}
</script>
