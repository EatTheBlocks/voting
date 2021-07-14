<template>
  <div>
    back
    ICON by {{ ethShortAddress(proposal.author) }}
    <div :style="'color: '+proposal.label.color">{{ proposal.label.text }}</div>
    <h2>{{ proposal.title }}</h2>
    <p>{{ proposal.body }}</p>
    <div>{{ endDateAgo(proposal.end) }}</div>
    {{ proposal.state }}

    <div>
      Votes
      <div v-for="vote in votes" :key="vote.author">
        BLOCKIE {{ vote.author }} {{ vote.choice }} {{ vote.tokens }} {{ vote.receipt }}
      </div>
      <div>
        <h3>Information</h3>
        Author BLOCKIE {{ ethShortAddress(proposal.author) }} {{ proposal.label.text }}
        IPFS #{{ proposal.id.substring(0, 7) }} icon new page
        Start date {{ proposal.start }}
        End date {{ proposal.end }}
      </div>

      <div>
        <h3>Results</h3>
        YES x TOKENS x%
        NO x TOKENS x%
      </div>
    </div>
  </div>
</template>

<script>
let proposal = {
  id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
  label: {color: "blue", text: "Core"},
  title: "Test proposal",
  body: "This is a test proposal for ETB",
  start: Date.now(),
  end: Date.now() - 24 * 60 * 60 * 1000,
  state: "Closed",
}

let votes = [
  {
    author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    timestamp: 1625695452,
    choice: 0,
    tokens: 100,
    receipt: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  },
  {
    author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    timestamp: 1625695452,
    choice: 1,
    tokens: 200,
    receipt: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  }
]

export default {
  name: 'Proposal',
  data() {
    return {
      proposal: proposal,
      votes: votes,
    }
  },
}
</script>

<style scoped>

</style>
