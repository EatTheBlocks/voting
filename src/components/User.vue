<template>
<span>
    <UiPopover :options="{ offset: [0, 12], placement: 'bottom-start' }">
      <template #item>
        <a class="whitespace-nowrap cursor-pointer">
          <UiAvatar class="mr-2" :opts="{seed:address, size:14}"/>
          {{ ethShortAddress(address) }}
        </a>
      </template>
      <template #content>
        <div class="p-4 text-center">
          <UiAvatar class="mb-4" :opts="{seed:address, size:64}"/>
          <h3>{{ ethShortAddress(address) }}</h3>
        </div>
      </template>
    </UiPopover>
  </span>
</template>

<script>
export default {
  name: 'User',
  props: {
    address: String,
  }
}
</script>

<style scoped>

</style>
