import { createApp } from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from 'Q@/router'

import { setupStore } from '@/store'

import './assets/main.css'

const app = createApp(App)

setupStore(app)
app.use(router)
app.use(ElementPlus)

app.mount('#app')
