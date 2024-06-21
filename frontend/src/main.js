import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(createPinia())
app.use(ArcoVue)
app.use(router)
app.config.devtools = false;
app.config.warnHandler = () => null;
app.mount('#app')
