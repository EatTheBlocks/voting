import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/css/index.styl'
import 'tailwindcss/tailwind.css'
import helpers from './mixins/helpers'
import i18n from './i18n'

const app = createApp(App)

app.config.globalProperties.$AppName = process.env.VUE_APP_NAME || 'Voting App'
app.config.globalProperties.$TokenName = process.env.VUE_APP_TOKEN || 'TokenName'

app.use(store).use(router).use(i18n).mixin(helpers).mount('#app')
