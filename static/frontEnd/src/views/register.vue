<template>
  <div>
    <el-input v-model="username" placeholder="请输入用户名" class="input"></el-input>
    <el-radio-group v-model="sex">
      <el-radio :label="0">女</el-radio>
      <el-radio :label="1">男</el-radio>
      <el-radio :label="2">跨性别</el-radio>
      <el-radio :label="3">未知</el-radio>
    </el-radio-group>
    <el-input v-model="email" placeholder="请输入邮箱" class="input"></el-input>
    <el-input type="password" v-model="password" placeholder="请输入密码" class="input"></el-input>
    <el-button round @click="register">注册</el-button>
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
      userInfo: {
        username: '',
        sex: '',
        email: '',
        pic: 'xx',
        password: '',
        msg: ''
      }
    }
  },
  methods: {
    register: function () {
      var that = this
      axios.post('http://127.0.0.1:8080/webapi/account', {
        UserName: this.username,
        UserSex: this.sex,
        UserEmail: this.email,
        PicUrl: 'xx',
        Password: this.password
      }).then(res => {
        that.msg = res.data
      })
    }
  }
}
</script>

<style>
.input {
  width: 200px;
}
</style>
