import { createRouter, createWebHistory } from 'vue-router'
import AppIndex from '@/components/Index/Index'
import AppManage from '@/components/Admin/Index'

const routerHistory = createWebHistory()

const constantRoutes = [
  {
    path: '/',
    name: 'index',
    component: AppIndex
  },
  {
    path: '/manage',
    name: 'manage',
    component: AppManage
  }
]


let router = createRouter({
  history: routerHistory,
  routes: constantRoutes,
})

export default router