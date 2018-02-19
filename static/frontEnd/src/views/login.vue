<template>
  <div>
      <el-input v-model="userInfo.UserName" placeholder="请输入用户名" class="input"></el-input>
      <el-input type="password" v-model="userInfo.Password" placeholder="请输入密码" class="input"></el-input>
      <el-button round @click="login">登录</el-button>
      <div>
        {{ msg }}
      </div>
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
      msg: ''
    }
  },
  http: {
    root: '/root',
    headers: {
      Authorization: 'Basic YXBpOnBhc3N3b3Jk'
    }
  },
  methods: {
    login: function () {
      var that = this
      axios.post('http://127.0.0.1:8080/webapi/login_user', qs.stringify(that.userInfo)).then(res => {
        that.msg = res.data
      })
    }
  }
}
</script>

<style>
.input{
  width: 200px;
}
</style>
