import Vue from 'vue'
import Router from 'vue-router'
import Login from './../views/login.vue'
import Register from './../views/register.vue'
import Main from './../views/main'
import Uxlog from '@/views/uxlog'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
      // children: [
      //   {
      //     path: 'login',
      //     name: 'login',
      //     component: Login
      //   },
      //   {
      //     path: 'register',
      //     name: 'register',
      //     component: Register
      //   }
      // ]
    },
    {
      path: '/uxlog',
      name: 'uxlog',
      component: Uxlog
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    }
  ]
})
