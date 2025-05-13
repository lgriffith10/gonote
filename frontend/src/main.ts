import '@/assets/main.css'

import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/plugins/router'
import { upfetch } from './lib/upfetch'

const app = createApp(App)

app.use(router)
app.mount('#app')
