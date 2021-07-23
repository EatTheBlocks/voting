<template>
  <div class="border border-main-border rounded-md divide-y divide-main-border overflow-hidden">
    <div class="p-5 text-center bg-main-block">
      <img src="@/assets/imgs/logo-etb.png" alt="logo"
           class="rounded-full border border-main-border h-20 w-20 mx-auto mb-2">
      <span class="text-lg font-semibold text-main-heading">{{ $AppName }}</span>
    </div>
    <div class="py-3 text-main-link menu">
      <router-link :to="{name:'Home'}" class="block py-2 px-5">Proposals</router-link>
      <router-link :to="{name:'Create'}" class="block py-2 px-5">New proposal</router-link>
      <router-link :to="{name:'About'}" class="block py-2 px-5">About</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Menu'
}
</script>

<style lang="stylus">
.menu a
  border-left 3px solid transparent

  &.router-link-active
    @apply border-main-link

</style>
