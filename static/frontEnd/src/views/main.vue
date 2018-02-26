<template>
  <div class="container">
    <header>
      <button @click="jump" class="uxButton">
        UXlog
      </button>
      <el-breadcrumb separator="/" class="option">
        <el-breadcrumb-item :to="{ path: 'uxlog' }">主页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: 'sign/login'}">登录</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: 'sign/register'}">注册</el-breadcrumb-item>
      </el-breadcrumb>
      <div>
        <router-view></router-view>
      </div>
    </header>
    <section class="firstBox">
      <div class="uxRow">
        <div class="theFirst"> UXlog</div>
        <flip-box front="简单" back="哈"></flip-box>
        <flip-box front="易用" back="哈"></flip-box>
      </div>
      <div class="uxRow">
        <flip-box front="专注" back="哈"></flip-box>
        <flip-box front="轻量化" back="哈"></flip-box>
        <flip-box front="公众化" back="哈"></flip-box>
        <flip-box front="的博客系统" back="哈"></flip-box>
      </div>
    </section>
    <div>
      <p>{{ msg }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
axios.defaults.timeout = 5000
export default {
  data () {
    return {
      msg: 'hello'
    }
  },
  methods: {
    jump () {
      this.$router.push('uxlog')
    },
    init () {
      axios.get('http://127.0.0.1:8080/webapi/login_user').then(function (res) {
        console.log(res.statusCode)
        // this.msg = res
      }).catch(function () {
        console.log('catch')
      })
    }
  },
  beforeMount: function () {
    this.init()
  },
  mounted: function () {
    axios.interceptors.request.use(config => {
      console.log('request init.')
      return config
    }, error => {
      return Promise.reject(error)
    })
    axios.interceptors.response.use(data => {
      console.log('response init.')
    }, error => {
      return Promise.reject(error)
    })
  },
  components: {
    'flip-box': {
      props: ['front', 'back'],
      template: '<div class="flipBox"><div class="flip"><div class="front">{{ front }}</div><div class="back">{{ back }}</div></div></div>'
    }
  }
}
</script>

<style>
.container {
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
  /* TODO */
  /* background-color:skyblue;  */
  height: 900px;
}
.uxButton {
  -moz-appearance: none;
  -webkit-appearance: none;
  outline: none;
  height: 65px;
  padding: 0px 10px;
  background-color: black;
  border: none;
  color: aliceblue;
  font-size: 2rem;
  font-family: sans-serif;
}
.option {
  float: right;
  margin-top: 22px;
  margin-right: 14px;
}

.flipBox {
  perspective: 1000;
}
.flip {
  position: relative;
  transform-style: preserve-3d;
  transition: 0.6s;
  margin: 1px;
}
.front,
.back {
  position: absolute;
  width: 100%;
  backface-visibility: hidden; /* 避免在实现动画效果时露出背面 */
}
.front {
  z-index: 2;
}
.flipBox:nth-child(odd) .flip .back {
  z-index: 1;
  transform: rotateY(-180deg); /* 最开始就翻转180度，以背面示人 */
  color: black;
  background-color:white;
}

.flipBox:nth-child(even) .flip .back{
  color: black;
  background-color:white;
  transform: rotateX(180deg)
}
.flipBox:nth-child(even) .flip .front{
  color: snow;
  background-color:black;
}
.flipBox:nth-child(odd) .flip .front{
  color: black;
  background-color:pink;
}
.flipBox:nth-child(even):hover .flip{
  transform: rotateX(180deg)
}
.flipBox:hover .flip {
  transform: rotateY(180deg);
}

.theFirst{
  margin: 1px;
  background-color: rgba(0, 0, 0, 0.12);
  color: white;
}

@media all and (min-width: 1000px) {
  .firstBox {
    width: 80%;
    margin: 35px auto;
    height: 400px;
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
  }
  .uxRow {
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
    flex: 1;
  }
  .uxRow:first-child div:first-child {
    flex: 2;
  }
  .uxRow div {
    flex: 1;
  }

  .front,.back,.uxRow div:first-child{
    height: 196px;
  }
  .flipBox:nth-child(even) .flip{
    transform-origin: 40% 98px;
  }
}

@media all and (min-width: 601px) and (max-width: 999px) {
  /* 翻转Box */
  .firstBox {
    width: 80%;
    margin: 80px auto;
    height: 250px;
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
  }
  .uxRow {
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
    flex: 1;
  }
  .uxRow:first-child div:first-child {
    flex: 2;
  }
  .uxRow div {
    flex: 1;
  }

  .front,.back,.uxRow div:first-child{
    height: 121px;
  }
  .flipBox:nth-child(even) .flip{
    transform-origin: 0 61px;
  }
}

@media all and (max-width: 600px) {
  .firstBox {
    width: 100%;
    margin-top: 1vmax;
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
  }
  .uxRow {
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
    flex: 1;
  }
  .uxRow div {
    height: 17vmax;
    flex: 1;
  }

  .front,.back,.uxRow div:first-child{
    height: 16vmax;
  }
}
</style>
