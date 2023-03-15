import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persist'
import { router } from './router'

import App from './App.vue'

import i18n from './locale'

import './api/interceptor'
import 'uno.css'
import './assets/css/transition.css'



const pinia = createPinia()
pinia.use(piniaPersist)

createApp(App).use(pinia).use(router).use(i18n).mount('#app')
