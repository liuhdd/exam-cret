import { createApp } from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from '@/router'

import { setupStore } from '@/store'
import 'virtual:svg-icons-register';

// 国际化
import i18n from '@/lang/index'
// 样式
import 'element-plus/theme-chalk/dark/css-vars.css'
import '@/styles/index.scss';
import 'uno.css';

const app = createApp(App)

setupStore(app)
app.use(router)
app.use(ElementPlus)
app.use(i18n)
app.mount('#app')
