import {ethers} from 'ethers'

export const provider = new ethers.providers.Web3Provider(window.ethereum)

// initial state
const state = () => ({
  initialized: false,
  connected: false,
  error: null,
  user: '',
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
    commit('address', await signer.getAddress())
    commit('connected', true)
  },

  logout({commit}) {
    commit('address', '')
    commit('error', '')
    commit('user', '')
    commit('network', '')
    commit('ens', '')
    commit('connected', false)
  }
}

// mutations
const mutations = {
  initialized: function (state, value) {
    state.initialized = value
  },
  connected: function (state, value) {
    state.connected = value
  },
  error: function (state, value) {
    state.error = value
  },
  user: function (state, value) {
    state.user = value
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
