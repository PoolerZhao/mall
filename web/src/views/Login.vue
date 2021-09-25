<template>
  <el-row>
    <el-col :span="6" :offset="9">
      <el-card class="box-card"
               shadow="never"
               style="border: none;">
        <img src="../assets/logo.png" class="logo-img" alt="logo"/><br>
        <p class="logo-family">Welcome To Mall</p><br>
        <el-form :model="loginForm"
                 status-icon
                 :rules="rules"
                 ref="loginForm"
                 class="demo-ruleForm">
          <el-form-item prop="username">
            <el-input type="text" v-model="loginForm.username" autocomplete="off" prefix-icon="el-icon-user">
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" v-model="loginForm.password" autocomplete="off" prefix-icon="el-icon-lock">
            </el-input>
          </el-form-item>
          <el-form-item prop="captchaValue" style="width: 380px;">
            <el-input v-model="loginForm.captchaValue" prefix-icon="el-icon-circle-check"
                      style="width: 150px; float: left" maxlength="5"></el-input>
            <el-image :src="captchaImg" class="captchaImg" @click="getCaptcha"/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       :style="{ width: '100%'}"
                       @click="submitForm('loginForm')">登 录
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>

export default {
  name: "Login",
  data() {
    const validateUsercode = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入账号'));
      } else {
        callback();
      }
    };
    const validatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        callback();
      }
    };
    return {
      show: true,
      loginForm: {
        username: 'admin',
        password: '123',
        captchaId: '',
        captchaValue: '',
      },
      captchaImg: null,
      rules: {
        username: [
          {validator: validateUsercode, trigger: 'blur'}
        ],
        password: [
          {validator: validatePassword, trigger: 'blur'}
        ],
        captchaValue: [
          {required: true, message: '请输入验证码', trigger: 'blur'},
          {min: 4, max: 4, message: '长度为 5 个字符', trigger: 'blur'}
        ],
      }
    };
  },
  mounted() {
    this.getCaptcha();
  },
  updated() {
    this.getCaptcha();
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$axios.post('/login', {
            username: this.loginForm.username,
            password: this.loginForm.password,
            captchaId: this.loginForm.captchaId,
            captchaValue: this.loginForm.captchaValue,
          }).then((response) => {
            localStorage.setItem("token", response.data.data.token)
            this.$store.commit('SET_TOKEN', response.data.data.token)
            console.log(response.data.data.token)
            this.$router.push('/home');
          }).catch((error) => { console.log(error); })
        }
      });
      this.getCaptcha();
    },
    getCaptcha() {
      this.$axios.get('/captcha').then(response => {
        this.loginForm.captchaId = response.data.data.captchaId
        this.captchaImg = response.data.data.captchaImg
        this.loginForm.captchaValue = ''
      })
    }
  }
}
</script>

<style scoped>
.logo-img {
  width: 80px;
  height: 80px;
}

.box-card {
  margin-top: 150px;
  text-align: center;
}

.logo-family {
  font-size: 16px;
  color: #409eff;
  font-weight: bold;
}

.captchaImg {
  float: left;
  margin-left: 8px;
  border-radius: 4px;
}
</style>