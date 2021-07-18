<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-8/12 space-y-5">
      <BackButton :to="{name: 'Home'}"/>
      <input type="text" class="block outline-none text-3xl bg-transparent text-main-link font-semibold" placeholder="Question">
      <textarea name="" id="" placeholder="What is your proposal?" class="block outline-none resize-none bg-transparent text-main-link font-semibold w-full h-24`"></textarea>
      <div class="text-xl font-semibold text-main-heading">Preview</div>
      <div>TODO</div>
      <div class="panel">
        <div class="panel-title">Choices</div>
        <div class="panel-body">YOLO</div>
      </div>
    </div>
    <div class="w-4/12 space-y-5">
      <div class="panel">
        <div class="panel-title">Actions</div>
        <div class="panel-body">YOLO</div>
      </div>
    </div>
  </div>
</template>

<script>
import BackButton from '@/components/BackButton'

export default {
  name: 'NewProposal',
  components: {
    BackButton
  }
}
</script>

<style scoped>

</style>
