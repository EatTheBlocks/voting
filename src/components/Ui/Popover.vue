<script>
import { onMounted, ref, watch } from 'vue';
import { createPopper } from '@popperjs/core';
import { useDebounce } from '@/composables/useDebounce';
import { useDetectInput } from '@/composables/useDetectInput';
import { useMediaQuery } from '@/composables/useMediaQuery';
export default {
  props: {
    options: Object
  },
  setup(props) {
    const open = ref(false);
    const popHovered = ref(false);
    const itemref = ref(null);
    const contentref = ref(null);
    const { isTouchScreen } = useDetectInput();
    const { isXLargeScreen } = useMediaQuery();
    function popClose() {
      if (!popHovered.value) open.value = false;
    }
    let popperInstance;
    onMounted(() => {
      popperInstance = createPopper(itemref.value, contentref.value, {
        placement: props.options.placement,
        modifiers: [
          {
            name: 'offset',
            options: {
              offset: props.options.offset
            }
          }
        ]
      });
    });
    watch(open, () => {
      if (!isXLargeScreen.value)
        popperInstance.setOptions({ placement: 'bottom' });
      else popperInstance.setOptions({ placement: 'bottom-start' });
    });
    return {
      open,
      popClose,
      popHovered,
      itemref,
      contentref,
      debounce: useDebounce(),
      isTouchScreen
    };
  }
};
</script>

<template>
  <div
    ref="itemref"
    @mouseenter="isTouchScreen() ? null : debounce(() => (open = true), 800)"
    @mouseleave="debounce(() => popClose(), 300)"
  >
    <slot name="item" />
  </div>
  <div
    ref="contentref"
    v-show="open"
    @mouseenter="popHovered = true"
    @mouseleave="(popHovered = false), debounce(() => (open = false), 300)"
    class="custom-content"
  >
    <slot name="content" />
  </div>
</template>

<style scoped lang="stylus">
.custom-content
  @apply z-50 bg-main-bg border border-main-border rounded-md shadow-md
  min-width 300px
</style>
