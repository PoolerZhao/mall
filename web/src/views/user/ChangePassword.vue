<template>
  <el-row>
    <el-col :span="10" :offset="7">
      <el-card shadow="never" style="margin-top: 100px;">
        <el-form style="width: 450px;" :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="80px" class="demo-ruleForm">
          <el-form-item label="旧密码" prop="oldPass">
            <el-input v-model="ruleForm.oldPass" style="width: 80%;"></el-input>
          </el-form-item>
          <el-form-item label="新密码" prop="newPass">
            <el-input type="password" v-model="ruleForm.newPass" style="width: 80%;" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="checkPass">
            <el-input type="password" v-model="ruleForm.checkPass" style="width: 80%;" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import qs from 'qs'
export default {
  name: "ChangePassword",
  data() {
    let validatePass0 = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('原密码不能为空'));
      }
    };
    let validatePass1 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };
    let validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.newPass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        oldPass: '',
        newPass: '',
        checkPass: ''
      },
      rules: {
        oldPass: [
          { validator: validatePass0, required: true, trigger: 'blur' }
        ],
        newPass: [
          { validator: validatePass1, required: true, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, required: true, trigger: 'blur' }
        ],
      }
    };
  },
  methods: {
    submitForm(formName) {
      console.log(formName)
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$axios.post('/user/pass?' + qs.stringify({
            oldPass: this.ruleForm.oldPass,
            newPass: this.ruleForm.newPass
          })).then(response => {
            console.log(response);
            this.$message({
              message: response.data.message,
              type: 'success'
            });
          })
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>

<style scoped>

</style>