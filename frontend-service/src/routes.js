import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    {
      path: '/',
      name: 'login',
      component: () => import('./pages/Login.vue')
    },
    {
        path: '/home',
        name: 'home',
        component: () => import('./pages/Home.vue')
      }
  ]

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})