<template>
  <div class="space-y-2">
    <div class="flex justify-between">
      <div><slot/></div>
      <div>{{ _n(percent, '0.[00]') }}%</div>
    </div>
    <div class="bg-main-border w-full h-2 rounded-full relative flex overflow-hidden">
      <div class="bg-main-progress h-full rounded-full" :style="'width: '+percent+'%'"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Progress',
  props: {
    percent: Number,
  }
}
</script>

<style scoped>

</style>
