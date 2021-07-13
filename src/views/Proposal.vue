<template>
  <div>
    back
    ICON by {{ ethShortAddress(proposal.address) }}
    <div :style="'color: '+proposal.label.color">{{ proposal.label.text }}</div>
    <h2>{{ proposal.title }}</h2>
    <p>{{ proposal.description }}</p>
    <div>{{ endDateAgo(proposal.endDate) }}</div>
    {{ proposal.status }}

    <div>
      Votes
      <VoteCard v-for="vote in votes" :key="vote.address"
                :address="vote.address"
                :timestamp="vote.timestamp"
                :choice="vote.choice"
                :tokens="vote.tokens"
                :receipt="vote.receipt"
      />
    </div>
    <div>
      <h3>Information</h3>
      Author BLOCKIE {{ ethShortAddress(proposal.address) }} {{ proposal.label.text }}
      IPFS #{{ proposal.ipfsHash.substring(0, 7) }} icon new page
      Start date {{ proposal.startDate }}
      End date {{ proposal.endDate }}
    </div>

    <div>
      <h3>Results</h3>
      YES x TOKENS x%
      NO x TOKENS x%
    </div>
  </div>
</template>

<script>
import VoteCard from '@/components/VoteCard'
import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en'

TimeAgo.addDefaultLocale(en)
const timeAgo = new TimeAgo('en-US')

let proposal = {
  ipfsHash: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  address: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
  label: {color: "blue", text: "Core"},
  title: "Test proposal",
  description: "This is a test proposal for ETB",
  startDate: Date.now(),
  endDate: Date.now() - 24 * 60 * 60 * 1000,
  status: "Closed",
}

let votes = [
  {
    address: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    timestamp: 1625695452,
    choice: 0,
    tokens: 100,
    receipt: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  },
  {
    address: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    timestamp: 1625695452,
    choice: 1,
    tokens: 200,
    receipt: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
  }
]

export default {
  name: 'Proposal',
  components: {
    VoteCard,
  },
  data() {
    return {
      proposal: proposal,
      votes: votes,
    }
  },
  methods: {
    ethShortAddress(address) {
      if (address !== undefined && address.length > 0) {
        return (
          address.substring(0, 6) +
          "..." +
          address.substring(address.length - 4)
        );
      }
      return "unknown address";
    },
    endDateAgo(endDate) {
      const now = Date.now()
      const text = timeAgo.format(endDate)
      if (endDate > now) {
        return "end " + text
      }
      return "ended " + text
    }
  }
}
</script>

<style scoped>

</style>
