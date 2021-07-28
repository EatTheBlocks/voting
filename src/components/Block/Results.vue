<template>
  <div class="panel">
    <div class="panel-title">Results</div>
    <div class="panel-body">
      <div class="animate-pulse" v-if="loading">
        <div class="space-y-3">
          <div class="h-6 bg-main-text rounded w-5/6"></div>
          <div class="h-6 bg-main-text rounded w-3/6"></div>
        </div>
      </div>
      <div class="space-y-3" v-else>
        <UiProgress v-for="(choice, i) in proposal.choices" :key="i"
                    :percent="100*results.resultsByVoteBalance[i]/results.sumOfResultsBalance">
          {{ _shorten(choice, 'choice') }} {{ _n(results.resultsByVoteBalance[i]) }} {{ $TokenName }}
        </UiProgress>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Results',
  props: {
    loading: Boolean,
    proposal: Object,
    votes: Array,
    results: Object,
  }
}
</script>

<style scoped>

</style>
