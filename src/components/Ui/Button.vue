<template>
  <button
    class="w-full py-3 border border-main-border rounded-full text-center text-main-link font-semibold hover:border-main-link cursor-pointer"
    :class="disabled?'hover:border-main-border cursor-default bg-main-border':''"
    :disabled="disabled"
  >
    <slot/>
  </button>
</template>

<script>
export default {
  name: 'Button',
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
  }
}
</script>
