import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('bearer-token')

  if (to.meta.isPublic && token) {
    return next({ name: 'Home' })
  }

  if (!to.meta.isPublic && !token) {
    return next({ name: 'Login' })
  }

  return next()
})

export default router
