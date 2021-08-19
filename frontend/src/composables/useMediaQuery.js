import { onBeforeMount, onBeforeUnmount, ref, computed } from 'vue';
// Breakpoints from TailwindCSS
export function useMediaQuery() {
  const width = ref(window.innerWidth);

  const isXSmallScreen = computed(() => width.value < 370);
  const isSmallScreen = computed(() => width.value < 640);
  const isMediumScreen = computed(() => width.value < 768);
  const isLargeScreen = computed(() => width.value < 1024);
  const isXLargeScreen = computed(() => width.value < 1280);
  const is2XLargeScreen = computed(() => width.value < 1536);

  function updateWidth() {
    width.value = window.innerWidth;
  }

  onBeforeMount(() => window.addEventListener('resize', updateWidth));
  onBeforeUnmount(() => window.removeEventListener('resize', updateWidth));

  return {
    isXSmallScreen,
    isSmallScreen,
    isMediumScreen,
    isLargeScreen,
    isXLargeScreen,
    is2XLargeScreen
  };
}
