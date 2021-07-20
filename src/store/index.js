import {createStore} from 'vuex'
import VuexPersistence from "vuex-persist";

const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})

export default createStore({
  state: {
    darkMode: true
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
    }
  },
  actions: {},
  modules: {},
  plugins: [
    vuexLocal.plugin
  ]
})
