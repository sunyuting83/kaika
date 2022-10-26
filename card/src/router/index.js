import { createRouter, createWebHistory } from 'vue-router'
import AppIndex from '@/components/Index/Index'
import AppManage from '@/components/Manage/Index'
import AdminList from '@/components/Admin/AdminList'
import CardList from '@/components/Card/CardList'

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
  },
  {
    path: '/adminlist',
    name: 'adminlist',
    component: AdminList
  },
  {
    path: '/cardlist',
    name: 'cardlist',
    component: CardList
  }
]


let router = createRouter({
  history: routerHistory,
  routes: constantRoutes,
})

export default router