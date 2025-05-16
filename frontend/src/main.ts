import '@/assets/main.css'

import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/plugins/router'
import BaseLayout from './layouts/BaseLayout.vue'
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
import { VueQueryPlugin } from '@tanstack/vue-query'

const app = createApp(App)

app.component('BaseLayout', BaseLayout)

app.use(router)
app.use(autoAnimatePlugin)
app.use(VueQueryPlugin)

app.mount('#app')
