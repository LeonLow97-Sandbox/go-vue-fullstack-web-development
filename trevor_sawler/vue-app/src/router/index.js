import { createRouter, createWebHistory } from 'vue-router'
import Body from '@/components/Body.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Body
    }
  ]
})

export default router
