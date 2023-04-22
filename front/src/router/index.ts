import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'


const constantRoutes: Array<RouteRecordRaw> = [

    {
      path: '/login',
      component: () => import('@/views/login/index.vue'),
      meta: { hidden: true }
    }

]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 })
})


export default router
