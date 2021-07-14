<template>
  <h2>Proposals</h2>
  <div class="space-y-5">
    <div v-for="proposal in proposals" :key="proposal.id" class="p-5 border border-gray-300 rounded-md">
      <router-link class="space-y-2" :to="{ name: 'Proposal', params: { id: proposal.id }}">
        <div class="flex justify-between">
          <div class="flex space-x-2 items-center">
            <div class="h-5 w-5 bg-red-500"></div>
            <div>{TokenName} by {{ ethShortAddress(proposal.author) }}</div>
            <div class="border border-gray-200 rounded-full px-2 text-sm" :style="'color: '+proposal.label.color">{{
                proposal.label.text
              }}
            </div>
          </div>
          <div class="bg-blue-400 rounded-full px-2 py-1 text-sm text-white">{{ proposal.state }}</div>
        </div>
        <h2 class="text-xl">{{ proposal.title }}</h2>
        <p class="">{{ proposal.body }}</p>
        <div class="text-sm">{{ endDateAgo(proposal.end) }}</div>
      </router-link>
    </div>
  </div>
</template>

<script>
let proposals = [
  {
    id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    label: {color: "blue", text: "Core"},
    title: "Test proposal",
    body: "This is a test proposal for ETB",
    start: Date.now(),
    end: Date.now() - 24 * 60 * 60 * 1000,
    state: "Closed",
  },
  {
    id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    label: {color: "red", text: "Core"},
    title: "Test proposal #2",
    body: "This is a test proposal for ETB",
    start: Date.now(),
    end: Date.now() - 24 * 60 * 60 * 1000,
    state: "Closed",
  }
]

export default {
  name: 'Home',
  data() {
    return {
      proposals: proposals
    }
  }
}
</script>
