<template>
  <div class="hero is-fullheight is-primary">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-8 is-offset-2">
          <h3 class="title has-text-white">后台管理系统</h3>
          <hr class="login-hr">
          <p class="subtitle has-text-white">感谢今天正在努力的自己!</p>
          <div class="box">
            <div class="box">
              <img :src="logo">
            </div>
            <div class="title has-text-grey is-5">请输入用户名和密码.</div>
          </div>
          <div class="field">
            <p class="control has-icons-left has-icons-right">
              <input class="input" type="username" v-model="form.username" placeholder="用户名" :onBlur="checkUsername">
              <span class="icon is-small is-left">
                <i class="fa fa-user-circle-o"></i>
              </span>
              <span class="icon is-small is-right" v-if="form.usernameErr !== '1' && form.usernameErr.length > 1">
                <i class="fa fa-exclamation-triangle"></i>
              </span>
            </p>
            <p class="help has-text-left is-white" v-if="form.usernameErr !== '1' && form.usernameErr.length > 1">{{form.usernameErr}}</p>
          </div>
          <div class="field">
            <p class="control has-icons-left has-icons-right">
              <input class="input" type="password" v-model="form.password" placeholder="密码" :onBlur="checkPassword">
              <span class="icon is-small is-left">
                <i class="fa fa-lock"></i>
              </span>
              <span class="icon is-small is-right" v-if="form.passwordErr !== '1' && form.passwordErr.length > 1">
                <i class="fa fa-exclamation-triangle"></i>
              </span>
            </p>
            <p class="help has-text-left is-white" v-if="form.passwordErr !== '1' && form.passwordErr.length > 1">{{form.passwordErr}}</p>
          </div>
          <div class="field has-addons">
            <div class="control is-expanded">
              <div class="control has-icons-left has-icons-right">
                <input class="input" type="code" placeholder="请输入验证码" @input="checkCode" />
                <span class="icon is-small is-left">
                  <i class="fa fa-qrcode"></i>
                </span>
                <span class="icon is-small is-right" v-if="form.codeErr !== '1' && form.codeErr.length == 4">
                  <i class="fa fa-exclamation-triangle"></i>
                </span>
              </div>
              <p class="help has-text-left is-white" v-if="form.codeErr !== '1' && form.codeErr.length > 1">{{form.codeErr}}</p>
            </div>
            <div class="control">
              <Identify :identifyCode="identifyCode" @click="refreshCode" />
            </div>
          </div>
          <button class="button is-block is-danger is-large is-fullwidth" :disabled="form.usernameErr !== '1' && form.usernameErr.length === 0  && form.passwordErr !== '1' && form.passwordErr.length === 0  && form.codeErr !== '1' && form.codeErr.length === 0 ? false : true" @click="onSubmit">登陆</button>
        </div>
      </div>
    </div>
    <NotIfication
      :message="openerr.message"
      :active="openerr.active"
      :closErr="closErr">
    </NotIfication>
  </div>
</template>

<script>
import Config from '@/helper/config'
import Fetch from '@/helper/fetch'
import Identify from "@/components/Other/Identify"
import NotIfication from "@/components/Other/Notification"
export default {
  name: 'AppIndex',
  components: { NotIfication, Identify },
  data() {
    return {
      identifyCodes: "1234567890",
      identifyCode: "",
      logo: Config.images[1],
      form: {
        username: "",
        password: "",
        code: "",
        usernameErr: "1",
        passwordErr: "1",
        codeErr: "1",
      },
      openerr: {
        active: false,
        message: ""
      }
    };
  },
  mounted() {
    // 初始
    this.identifyCode = "";
    this.makeCode(this.identifyCodes, 4);
  },
  methods: {
    // 确认
    async onSubmit() {
      const { username, password } = this.form
      if (username.length >= 4 && password.length >= 8 && this.validateCode()) {
        const data = {
          username: username,
          password: password
        }
        const d = await Fetch(Config.Api.login, data, "post")
        if (d.status === 0) {
          this.$router.push("/manage")
        }else{
          this.openerr.message = d.message
          this.openerr.active = true
          this.form.codeErr = "1"
          this.form.code = ""
          this.refreshCode()
        }
      }
    },
    checkUsername(){
      const username = this.form.username
      if (username.length < 4) {
        this.form.usernameErr = "用户名不能小于4位"
      }else{
        this.form.usernameErr = ""
      }
    },
    checkPassword(){
      const password = this.form.password
      if (password.length < 8) {
        this.form.passwordErr = "用户名不能小于4位"
      }else{
        this.form.passwordErr = ""
      }
    },
    validateCode() {
      const value = this.form.code
      if (value === "") {
        this.form.codeErr = "验证码为空"
        return false;
      } else if (value !== this.identifyCode) {
        this.form.codeErr = "验证码不正确"
        return false;
      } else {
        this.form.codeErr = ""
        return true
      }
    },
    // 验证码相关
    randomNum(min, max) {
      return Math.floor(Math.random() * (max - min) + min);
    },
    // 点击刷新验证码
    refreshCode() {
      this.identifyCode = "";
      this.makeCode(this.identifyCodes, 4);
    },
    // 生成验证码
    makeCode(o, l) {
      for (let i = 0; i < l; i++) {
        this.identifyCode +=
          this.identifyCodes[this.randomNum(0, this.identifyCodes.length)];
      }
    },
    closErr(){
      this.openerr.active = false
    },
    checkCode(e){
      let code = this.form.code
      const l = code.length
      if (e.data == null) {
        this.form.code = code.slice(0, l - 1)
      }else{
        this.form.code = code + e.data
      }
      this.validateCode()
    }
  }
}
</script>