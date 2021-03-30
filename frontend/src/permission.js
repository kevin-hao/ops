import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'
const _import = require('./router/_import_' + process.env.NODE_ENV) // 获取组件的方法
import Layout from "@/layout";
NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)

  // determine whether the user has logged in
  const hasToken = getToken()

  if (hasToken) {
    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      next({ path: '/' })
      NProgress.done()
    } else {
      const hasGetUserInfo = store.getters.name
      if (hasGetUserInfo) {
        next()
      } else {
        try {
          // get user info
          await store.dispatch('user/getInfo') //触发获取用户信息的接口
          await store.dispatch('user/getRouter') //触发获取路由表的接口
          if (store.getters.menus.length < 1){
            global.antRouter = []
            next()
          }
          const menus = filterAsyncRouter(store.getters.menus) //过滤路由

          router.addRoutes(menus) //动态添加路由
          global.antRouter = menus //将路由数据传递给全局变量，做侧边栏渲染工作
          next({...to, replace:true})
        } catch (error) {

          // remove token and go to login page to re-login
          await store.dispatch('user/resetToken')
          Message.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    /* has no token*/

    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})

//  遍历后台传来的路由字符串，转换为组件对象
function filterAsyncRouter(asyncRouterMap) {
  const accessedRouters = asyncRouterMap.filter(route => {
    if (route.component) {
      if (route.component === 'Layout') {
        route.component = Layout
      } else {
        try {
          route.component = _import(route.component) // 导入组件
        } catch (error) {
          route.component = _import('system/menu/index') // 导入组件
          Message.error('请查看控制台修改或者删除不存在的组件路径') //弹出异常('请修改或者删除不存在的组件路径')
        }
      }

    }
    
    if (route.children  && route.children.length) {

      route.children = filterAsyncRouter(route.children)
    }

    return true
  })

  return accessedRouters
}
