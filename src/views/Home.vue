<template>
  <div class="flex max-w-[1012px] mx-auto space-x-10">
    <div class="w-3/12">
      <div class="border border-main-border rounded-md divide-y divide-main-border overflow-hidden">
        <div class="p-5 text-center bg-main-block">
          <img src="@/assets/imgs/logo-etb.png" alt="logo" class="rounded-full border border-main-border h-20 w-20 mx-auto mb-2">
          <span class="text-lg font-semibold text-main-heading">{{ $AppName }}</span>
        </div>
        <div class="py-3 text-main-link">
          <div class="py-2 px-5 border-l-[3px] border-gray-800">Proposals</div>
          <div class="py-2 px-5">New proposal</div>
          <div class="py-2 px-5">About</div>
        </div>
      </div>
    </div>
    <div class="w-9/12">
      <div class="text-sm font-semibold">{{ $AppName }}</div>
      <h2 class="font-semibold text-2xl">Proposals</h2>
      <div class="space-y-6 mt-4">
        <div v-for="proposal in proposals" :key="proposal.id" class="p-5 border border-main-border rounded-md">
          <router-link class="space-y-2 hover:text-main-link" :to="{ name: 'Proposal', params: { id: proposal.id }}">
            <div class="flex justify-between">
              <div class="flex space-x-2 items-center mb-1">
                <img src="@/assets/imgs/logo-etb.png" alt="logo" class="rounded-full h-6 w-6">
                <div class="font-semibold text-sm">{{ $TokenName }} by {{ ethShortAddress(proposal.author) }}
                </div>
                <div class="badge-core ml-2">
                  {{ proposal.label.text }}
                </div>
              </div>
              <div class="badge-state">{{ proposal.state }}</div>
            </div>
            <h2 class="text-2xl font-semibold">{{ proposal.title }}</h2>
            <p class="text-lg">{{ proposal.body }}</p>
            <div>{{ endDateAgo(proposal.end) }}</div>
          </router-link>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
let proposals = [
  {
    id: "QmYTcx9abcdY5RkFrD15yCvFD5eMxwdsfhSgSbdB2UxNJgd",
    author: "0x7ac64008fa000bfdc4494e0bfcc9f4eff3d51d2a",
    label: {color: "blue", text: "Core"},
    title: "YIP-62: Change Two Multisig Signers",
    body: "Authors\n" +
      "@banteg @lehnberg @milkyklim @tracheopteryx\n" +
      "\n" +
      "Summary\n" +
      "A proposal to exercise the 'Change Multisig Signers' power held by YFI holders...",
    created: Date.now(),
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
    created: Date.now(),
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
