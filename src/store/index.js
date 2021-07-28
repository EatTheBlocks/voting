import {createStore} from 'vuex'
import VuexPersistence from "vuex-persist";
import web3 from './modules/web3'

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
    }
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
    }
  },
  actions: {},
  modules: {
    web3
  },
  plugins: [
    vuexLocal.plugin
  ]
})
