import Vue from 'vue'
import Router from 'vue-router'
import Login from './../views/login.vue'
// import Register from './../views/register.vue'
import Main from './../views/main'
import Uxlog from '@/views/uxlog'
import Register from '@/views/newregister'
import sign from '@/views/sign'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
    },
    {
      path: '/uxlog',
      name: 'uxlog',
      component: Uxlog
    },
    {
      path: '/sign',
      name: 'sign',
      component: sign,
      children: [
        {
          path: 'login',
          name: 'login',
          component: Login
        },
        {
          path: 'register',
          name: 'register',
          component: Register
        }
      ]
    }
  ]
})
