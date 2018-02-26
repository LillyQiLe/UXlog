<template>
    <div class="register">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="username" prop="name">
                <el-input v-model="ruleForm.UserName"></el-input>
            </el-form-item>
            <el-form-item label="sex" prop="sex">
                <el-radio-group v-model="ruleForm.UserSex">
                    <el-radio :label="0">girl</el-radio>
                    <el-radio :label="1">boy</el-radio>
                    <el-radio :label="2">跨性别</el-radio>
                    <el-radio :label="3">未知</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="email" prop="email">
                <el-input v-model="ruleForm.UserEmail" placeholder="请输入邮箱"></el-input>
            </el-form-item>
            <el-form-item label="password" prop="password">
                    <el-input type="password" v-model="ruleForm.Password" placeholder="请输入密码" class="input"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="register('ruleForm')">立即创建</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
        <p>{{ msg }}</p>
    </div>
</template>

<script>
import axios from 'axios'
import qs from 'qs'
export default {
  data () {
    return {
      msg: '',
      ruleForm: {
        UserName: '',
        UserSex: '',
        UserEmail: '',
        Password: '',
        PicUrl: 'xx'
      },
      rules: {
        UserName: [
          { required: true, message: '请输入活动名称', trigger: 'blur' },
          { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
        ],
        UserSex: [
          { required: true, message: '请选择性别', trigger: 'change' }
        ],
        Password: [
          { required: true, message: '请填写密码', trigger: 'blur' }
        ],
        UserEmail: [
          { required: true, message: '请填写活动形式', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    register (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          var that = this
          axios.post('http://127.0.0.1:8080/webapi/account', qs.stringify(that.ruleForm)).then(res => {
            that.msg = res.data
          })
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

<style>
.register{
    width: 60%;
}
</style>
