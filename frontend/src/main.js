import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import i18n from './i18n'
import helpers from './mixins/helpers'
import upperFirst from 'lodash/upperFirst'
import camelCase from 'lodash/camelCase'
import './assets/css/index.styl'
import 'tailwindcss/tailwind.css'

// if (process.env.NODE_ENV === 'production' && location.protocol !== 'https:') {
//   location.replace(`https:${location.href.substring(location.protocol.length)}`);
// }

const app = createApp(App)
  .use(store)
  .use(router)
  .use(i18n)
  .mixin(helpers)

const requireComponent = require.context('@/components', true, /[\w-]+\.vue$/);
requireComponent.keys().forEach(fileName => {
  const componentConfig = requireComponent(fileName);
  const componentName = upperFirst(
    camelCase(fileName.replace(/^\.\//, '').replace(/\.\w+$/, ''))
  );
  app.component(componentName, componentConfig.default || componentConfig);
});

app.mount('#app')
