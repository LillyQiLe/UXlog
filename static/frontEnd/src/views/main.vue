<template>
  <header>
    <button @click="jump" class="uxButton">
      UXlog
    </button>
    <el-breadcrumb separator="/" class="option">
      <el-breadcrumb-item :to="{ path: 'uxlog' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: 'login'}">登录</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: 'register'}">注册</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- <button @click="jump">button跳转</button> -->
    <div>
      <router-view></router-view>
    </div>
  </header>
</template>

<script>
export default {
  data () {
    return {
      msg: 'hello'
    }
  },
  methods: {
    jump () {
      this.$router.push('uxlog')
    }
  }
}
</script>

<style>
.option{
  float: right;
  margin-top: 22px;
  margin-right: 14px;
}
.uxButton{
  appearance: none;
  -moz-appearance: none;
  -webkit-appearance: none;
  outline: none;
  width: 15%;
  height: 65px;
  background-color: black;
  border: none;
  color: aliceblue;
  font-size: 2rem;
  font-family: sans-serif;
}
</style>
