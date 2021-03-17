<template>
    <div class="login-box">
        <el-form ref="form" :rules="rules" :model="form" label-width="80px">
            <h3 class="login-title">欢迎登录</h3>
            <el-form-item label="用户名" prop="username">
                <el-input v-model="form.username" placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input type="password" v-model="form.password" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button size="medium" type="primary" @click="onSubmit('form')">登录</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
  export default {
    data() {
      return {
        form: {
          username: '',
          password: ''
        },
        rules:{
          username:[
            {
              required: true, message: '请输用户名', trigger: 'blur' },
              { min: 5, max: 18, message: '长度在 6 到 18 个字符', trigger: 'blur' }
          ],
          password: [
              { required: true, message: '请输入密码', trigger: 'change' },
              { min: 6, max: 18, message: '长度在 8 到 18 个字符', trigger: 'blur' }
           ]
        }
      }
    },
    methods: {
        onSubmit(formName){
          this.$refs[formName].validate((valid) => {
              var that = this
              if (valid) {

                 this.axios({
                   method: "get",
                   url: "http://127.0.0.1:8080/auth/login?username="+that.form.username+"&&password="+that.form.password,
                   // data:{
                   //   username: "kevin",
                   //   password: "123456"
                   // }
                 }).then(function(res){
                   if (res.data.code == 0){
                     that.$router.push("/home")
                   }
                 })
              } else {
                    this.$message.error('用户名或密码错误');
                    return false;
               }
           });
        }
    }
  }
</script>

<style scoped>
    .login-box{
        width: 350px;
        height: 300px;
        border: 1px solid #DCEFE6;
        margin: 150px auto;
        padding: 20px 50px 20px 30px;
        border-radius: 20px;
        box-shadow: 0px 0px 20px #DCEFE6;
    }
    .login-title {
        text-align: center;
        margin-bottom: 40px;
    }
</style>
