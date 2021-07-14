import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'tailwindcss/tailwind.css'
import './assets/css/index.styl'
import helpers from './mixins/helpers'

createApp(App).use(store).use(router).mixin(helpers).mount('#app')
