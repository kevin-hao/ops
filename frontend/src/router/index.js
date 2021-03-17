import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '@/views/Login'
import Home from '@/views/Home'
import Userlist from '@/views/Userlist'
import Userinfo from '@/views/Userinfo'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/logout',
    redirect: '/login'
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    children: [
      {
        path:'/userlist',
        name:'Userlist',
        component:Userlist
      },
      {
        path: '/userifo/:id',
        name: 'Userinfo',
        component: Userinfo
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
