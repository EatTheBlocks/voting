import {ethers} from 'ethers'

const networkRequired = 56

export const provider = new ethers.providers.Web3Provider(window.ethereum)

export function changeNetwork() {
  if (!window.ethereum) return

  window.ethereum.request({
    method: "wallet_addEthereumChain",
    params: [
      {
        chainId: "0x38",
        chainName: "Binance Smart Chain",
        nativeCurrency: {
          name: "BNB",
          symbol: "BNB",
          decimals: 18,
        },
        rpcUrls: ["https://bsc-dataseed.binance.org/"],
        blockExplorerUrls: ["https://bscscan.com/"],
      },
    ],
  })
}

// initial state
const state = () => ({
  connected: false,
  error: null,
  address: '',
  network: '',
  ens: null
})

// getters
const getters = {}

// actions
const actions = {
  async connect({commit}) {
    await provider.send("eth_requestAccounts", [])
    const signer = provider.getSigner()
    const network = await provider.getNetwork()
    commit('address', await signer.getAddress())
    commit('network', network.chainId)
    commit('connected', true)
  },

  logout({commit}) {
    commit('address', '')
    commit('error', '')
    commit('network', '')
    commit('ens', '')
    commit('connected', false)
  },

  init({commit, dispatch}) {
    const ethereum = window.ethereum
    if (!ethereum) return

    ethereum.on('chainChanged', (chainId) => {
      commit('network', chainId)

      if (chainId != networkRequired) {
        setTimeout(() => {
          changeNetwork()
        }, 1000)
      }
    })

    ethereum.on('accountsChanged', async (accounts) => {
      if (accounts.length === 0) {
        dispatch('logout')
      } else {
        const signer = provider.getSigner(accounts[0])
        const network = await provider.getNetwork()
        commit('address', await signer.getAddress())
        commit('network', network.chainId)
        commit('connected', true)
      }
    })

    ethereum.on('connect', (connectInfo) => {
      if (connectInfo.chainId != networkRequired) {
        setTimeout(() => {
          changeNetwork()
        }, 1000)
      }
    })

    ethereum.on('disconnect', async (error) => {
      console.error(error)
      dispatch('logout')
    })
  }
}

// mutations
const mutations = {
  connected: function (state, value) {
    state.connected = value
  },
  error: function (state, value) {
    state.error = value
  },
  address: function (state, value) {
    state.address = value
  },
  network: function (state, value) {
    state.network = value
  },
  ens: function (state, value) {
    state.ens = value
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
