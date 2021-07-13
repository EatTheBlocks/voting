<template>
  <div>
    BLOCKIE {{ address }} {{ choice }} {{ tokens }} {{ receipt }}
  </div>
</template>

<script>
export default {
  name: 'VoteCard',
  props: {
    address: String,
    timestamp: Number,
    choice: Number,
    tokens: Number,
    receipt: String,
  }
}
</script>

<style scoped>

</style>
