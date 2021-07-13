<template>
  <router-link :to="{ name: 'Proposal', params: { ipfsHash: ipfsHash }}">
    ICON by {{ ethShortAddress }}
    <div :style="'color: '+label.color">{{ label.text }}</div>
    <h2>{{ title }}</h2>
    <p>{{ description }}</p>
    <div>{{ endDateAgo }}</div>
    {{ status }}
  </router-link>
</template>

<script>
import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en'

TimeAgo.addDefaultLocale(en)
const timeAgo = new TimeAgo('en-US')

export default {
  name: 'ProposalCard',
  props: {
    ipfsHash: String,
    address: String,
    label: Object,
    title: String,
    description: String,
    startDate: Number,
    endDate: Number,
    status: String,
  },
  computed: {
    ethShortAddress() {
      if (this.address !== undefined && this.address.length > 0) {
        return (
          this.address.substring(0, 6) +
          "..." +
          this.address.substring(this.address.length - 4)
        );
      }
      return "unknown address";
    },
    endDateAgo() {
      const now = Date.now()
      const text = timeAgo.format(this.endDate)
      if (this.endDate > now) {
        return "end " + text
      }
      return "ended " + text
    }
  }
}
</script>

<style scoped>

</style>
