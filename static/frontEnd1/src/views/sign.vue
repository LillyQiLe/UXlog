<template>
  <div class="sign">
    <div class="main">
      <h4 class="title">
        <router-link to="login">登录</router-link>
        <b>.</b>
        <router-link to="register">注册</router-link>
      </h4>
      <div class="container">
        <router-view></router-view>
      </div>
    </div>
    
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      userInfo: {
        UserName: '',
        Password: ''
      },
      msg: 'I,m '
    }
  },
  http: {
    root: '/root',
    headers: {
      Authorization: 'zhongli'
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
  }
}
</script>

<style>
.sign{
  width: 100%;
  height: 100%;
  min-height: 750px;
  text-align: center;
  font-size: 14px;
  background-color: #f1f1f1;
}
.sign .main {
    width: 400px;
    margin: 60px auto 0;
    padding: 50px 50px 30px;
    background-color: #fff;
    border-radius: 4px;
    box-shadow: 0 0 8px rgba(0,0,0,.1);
    vertical-align: middle;
    display: inline-block;
}
.sign .main .container{
    margin: 40px auto;
}
.sign .main .title b,a{
  padding: 10px;
  font-weight: 700;
  color: #969696;
  line-height: 1.1;
}
.sign .main .title b{
    padding: 10px;
}
.sign .main a:active{
  font-weight: 700;
  color: #ea6f5a;
  border-bottom: 2px solid #ea6f5a;
}
@media all and (max-width: 700px){
  .sign {
    height: auto;
    min-height: 0;
    background-color: transparent;
  }
  .sign .main {
    width: 400px;
    margin: 60px auto 0;
    padding: 50px 50px 30px;
    background-color: #fff;
    border-radius: 0;
    box-shadow: none;
    vertical-align: middle;
    display: inline-block;
  }
}
</style>
