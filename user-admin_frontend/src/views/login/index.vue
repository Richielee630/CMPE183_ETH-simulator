<template>
  <div class="login-container">
    <el-form
      v-show="isLogin"
      ref="loginForm"
      :model="loginForm"
      :rules="loginRules"
      class="login-form"
      auto-complete="on"
      label-position="left"
    >
      <div class="title-container">
        <h3 class="title">{{ title }}</h3>
      </div>
      <el-form-item prop="user">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <!-- 用户名 -->
        <el-input
          ref="user"
          v-model="loginForm.user"
          placeholder="User Name"
          name="user"
          type="text"
          tabindex="1"
          auto-complete="on"
        />
      </el-form-item>
      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <!-- 密码 -->
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          placeholder="Password"
          name="password"
          tabindex="2"
          auto-complete="on"
          @keyup.enter.native="handleLogin"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon
            :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"
          />
        </span>
      </el-form-item>
      <!-- 登 录 -->
      <el-button
        :loading="loading"
        type="primary"
        style="width: 100%; margin-bottom: 30px"
        @click.native.prevent="handleLogin"
      >Login</el-button>
      <div class="registe">
        <!-- 注 册 -->
        <span @click="isLogin=false">Register</span>
      </div>
    </el-form>

    <el-form
      v-show="!isLogin"
      ref="registeForm"
      :model="registeForm"
      :rules="registeRules"
      class="registe-form"
      auto-complete="on"
      label-position="left"
    >
      <div class="title-container">
        <h3 class="title">{{ title }}</h3>
      </div>
      <el-form-item prop="user">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          ref="user"
          v-model="registeForm.user"
          placeholder="User Name"
          name="user"
          type="text"
          tabindex="1"
          auto-complete="on"
        />
      </el-form-item>

      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="registeForm.password"
          :type="passwordType"
          placeholder="Password"
          name="password"
          tabindex="2"
          auto-complete="on"
          @keyup.enter.native="handleRegiste"
        />
      </el-form-item>
      <el-form-item prop="confirmPwd">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <!-- 确认密码 -->
        <el-input
          :key="passwordType"
          ref="confirmPwd"
          v-model="registeForm.confirmPwd"
          :type="passwordType"
          placeholder="Confirm Password"
          name="confirmPwd"
          tabindex="3"
          auto-complete="on"
          @keyup.enter.native="handleRegiste"
        />
      </el-form-item>
      <el-form-item prop="confirmPassword">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <!-- 请确认密码 -->
        <el-input
          :key="passwordType"
          ref="confirmPassword"
          v-model="registeForm.password"
          :type="passwordType"
          placeholder="Please confirm the password"
          name="confirmPassword"
          tabindex="2"
          auto-complete="on"
          @keyup.enter.native="handleRegiste"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon
            :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"
          />
        </span>
      </el-form-item>

      <el-button
        :loading="loading"
        type="primary"
        style="width: 100%; margin-bottom: 30px"
        @click.native.prevent="handleRegiste"
      >Register</el-button>
      <div class="registe">
        <span @click="isLogin=true">Login</span>
      </div>
    </el-form>
  </div>
</template>

<script>
import { register } from 'api/user'

export default {
  name: 'Login',
  data() {
    const validateUser = (rule, value, callback) => {
      if (value.length < 6) {
        // 用户名不能少于6位
        callback(new Error('User name can not be less than 6 bits'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        // 密码不能少于6位
        callback(new Error('Password cannot be less than 6 digits'))
      } else {
        callback()
      }
    }
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== this.registeForm.password) {
        // 密码输入不匹配,请重新输入
        callback(new Error('Password input does not match, please re-enter'))
      } else {
        callback()
      }
    }
    return {
      // 用户管理系统
      title: 'User Management System',
      loginForm: {
        user: '',
        password: ''
      },
      loginRules: {
        user: [{ validator: validateUser, required: true, trigger: 'blur' }],
        password: [{ validator: validatePassword, required: true, trigger: 'blur' }]
      },
      registeForm: {
        user: '',
        password: '',
        confirmPwd: ''
      },
      registeRules: {
        user: [{ validator: validateUser, required: true, trigger: 'blur' }],
        password: [{ validator: validatePassword, required: true, trigger: 'blur' }],
        confirmPwd: [{ validator: validateConfirmPassword, required: true, trigger: 'blur' }]
      },
      loading: false,
      isLogin: true,
      passwordType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: (route) => {
        try {
          this.redirect = route.query && route.query.redirect
        } catch (error) {
          console.log(error)
        }
      },
      immediate: true
    }
  },
  methods: {
    // 显示密码
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      // 异步执行 操作dom
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    // 处理登录
    handleLogin() {
      const vm = this
      // 给loginForm 组件添加校验函数
      vm.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store
            .dispatch('user/login', this.loginForm)
            .then((response) => {
              console.log(response)
              const uri = '/user'
              this.$router.push({ path: this.redirect || uri })
              this.loading = false
            })
            .catch(() => {
              this.loading = false
            })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    // 处理注册
    handleRegiste() {
      if (this.registeForm.password !== this.registeForm.confirmPwd) {
        this.$message({
          // 确认密码不一致,请重新输入!
          message: 'Confirm the password is not the same, please re-enter!',
          type: 'warning'
        })
        return
      }
      this.loading = true
      register({ user: this.registeForm.user, password: this.registeForm.password }).then((response) => {
        this.loading = false
        console.log(response)
        this.$message({
          // 注册成功!
          message: 'Registration is successful!',
          type: 'success'
        })
        this.registeForm = {
          user: '',
          password: '',
          confirmPwd: ''
        }
        this.isLogin = true
      }).catch(function(error) {
        this.loading = false
        console.log(error)
        this.$message({
          // 注册失败
          message: 'Registration failed!',
          type: 'error'
        })
      })
    }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0;
      -webkit-appearance: none;
      border-radius: 0;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, .1);
    background: rgba(0, 0, 0, .1);
    border-radius: 5px;
    color: #454545;
  }
}

</style>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;

  .login-form, .registe-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0 auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .registe {
    text-align: right;
    span {
      color: #409eff;
      cursor: pointer;
    }
  }
}

</style>
