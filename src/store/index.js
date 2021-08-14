import {createStore} from 'vuex'
import VuexPersistence from "vuex-persist";
import web3, {defaultProvider} from './modules/web3'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})

export default createStore({
  state: {
    darkMode: true,
    space: {
      name: "",
      token: "",
      address: "",
      network: "",
      admins: [],
    },
    notifs: [],
    ens: {},
  },
  mutations: {
    changeTheme(state, darkMode) {
      state.darkMode = darkMode

      if (darkMode) {
        document.body.classList.add('dark')
        document.body.classList.remove('light')
      } else {
        document.body.classList.remove('dark')
        document.body.classList.add('light')
      }
    },
    setSpace(state, space) {
      state.space = space
    },
    notify(state, payload) {
      state.notifs.push({...payload, timestamp: Date.now()})
    },
    emptyNotifs(state) {
      state.notifs = []
    },
    saveENS(state, payload) {
      state.ens[payload[0]] = payload[1]
    }
  },
  actions: {
    notify({commit}, payload) {
      Array.isArray(payload)
        ? commit('notify', {message: payload[0], color: payload[1]})
        : commit('notify', {message: payload, color: 'bg-green-500'});
    },
    async getENS({commit}, payload) {
      const ensName = await defaultProvider.lookupAddress(payload)
      if (ensName) {
        const ensAddr = await defaultProvider.resolveName(ensName)
        if (ensAddr.toLowerCase() === payload.toLowerCase()) {
          commit('saveENS', [payload.toLowerCase(), ensName])
          return
        }
      }
      commit('saveENS', [payload.toLowerCase(), ''])
    }
  },
  modules: {
    web3
  },
  plugins: [
    vuexLocal.plugin
  ]
})
