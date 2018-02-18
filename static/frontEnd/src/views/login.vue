<template>
  <div>
      <el-input v-model="username" placeholder="请输入用户名" class="input"></el-input>
      <el-input type="password" v-model="password" placeholder="请输入密码" class="input"></el-input>
      <el-button round @click="login">圆形按钮</el-button>
      <div>
        {{ msg }}
      </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      username: '',
      password: '',
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
      axios.post('http://127.0.0.1:8080/webapi/login_user', {
        UserName: this.username,
        Password: this.password
      }).then(res => {
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
