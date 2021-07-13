<template>
  <div>
    <ProposalCard v-for="proposal in proposals" :key="proposal.ipfsHash"
                  :ipfsHash="proposal.ipfsHash"
                  :address="proposal.address"
                  :label="proposal.label"
                  :title="proposal.title"
                  :description="proposal.description"
                  :startDate="proposal.startDate"
                  :endDate="proposal.endDate"
                  :status="proposal.status"
    />
  </div>
</template>

<script>
import ProposalCard from './ProposalCard'

let proposals = [
  {
    ipfsHash: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    address: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    label: {color: "blue", text: "Core"},
    title: "Test proposal",
    description: "This is a test proposal for ETB",
    startDate: Date.now(),
    endDate: Date.now() - 24 * 60 * 60 * 1000,
    status: "Closed",
  },
  {
    ipfsHash: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    address: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    label: {color: "red", text: "Core"},
    title: "Test proposal #2",
    description: "This is a test proposal for ETB",
    startDate: Date.now(),
    endDate: Date.now() - 24 * 60 * 60 * 1000,
    status: "Closed",
  }
]

export default {
  name: 'ProposalList',
  components: {
    ProposalCard,
  },
  setup() {
    return {
      proposals
    }
  }
}
</script>

<style scoped>

</style>
