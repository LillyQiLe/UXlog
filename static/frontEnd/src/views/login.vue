<template>
  <div class="login">
      <el-input v-model="userInfo.UserName" placeholder="请输入用户名" class="input"></el-input>
      <el-input type="password" v-model="userInfo.Password" placeholder="请输入密码" class="input"></el-input>
      <el-button round @click="login">登录</el-button>
      {{ msg }}
  </div>
</template>

<script>
import axios from 'axios'
import qs from 'qs'
export default {
  data () {
    return {
      userInfo: {
        UserName: '',
        Password: ''
      },
      msg: 'I am login '
    }
  },
  http: {
    root: '/root',
    headers: {
      Authorization: 'Basic YXBpOnBhc3N3b3Jk'
    }
  },
  mounted: function () {
    axios.interceptors.request.use(config => {
      console.log('request init.')
      return config
    }, error => {
      return Promise.reject(error)
    })
    axios.interceptors.response.use(response => {
      console.log('response init.')
    }, error => {
      return Promise.reject(error)
    })
  },
  methods: {
    login: function () {
      var that = this
      axios.post('http://127.0.0.1:8080/webapi/login_user', qs.stringify(that.userInfo)).then(res => {
        that.msg = res
      })
    }
  }
}
</script>

<style>
/* @media all and (max-width: 768px){
} */
</style>
