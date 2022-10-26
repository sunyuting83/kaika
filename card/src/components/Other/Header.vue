<template>
  <nav class="navbar is-white">
    <div class="container">
      <div class="navbar-brand">
        <span class="navbar-item brand-text">
          <img :src="logo">Card Manage
        </span>
      </div>
      <div id="navMenu" class="navbar-menu">
        <div class="navbar-start">
          <router-link class="navbar-item" to="manage">
            首页
          </router-link>
          <a class="navbar-item" href="admin.html">
            卡开记录
          </a>
        </div>
      </div>
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="field is-grouped">
            <div class="navbar-item has-dropdown is-hoverable">
              <span class="navbar-link">
                <span class="icon">
                  <i class="fa fa-user-circle-o"></i>
                </span>
                <span>
                用户中心
                </span>
              </span>
          
              <div class="navbar-dropdown">
                <div class="dropdown-item">
                  <p>当前登陆用户 <strong>{{openModal.username}}</strong> <br /><span class="is-size-7">可以开卡 努力賺錢.</span></p>
                </div>
                <hr class="dropdown-divider">
                <a class="navbar-item" @click="showModel">
                  修改密码
                </a>
                <router-link class="navbar-item" to="adminlist">
                  管理员管理
                </router-link>
                <hr class="navbar-divider">
                <a class="navbar-item" @click="LogOut">
                  退出登陆
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <ChangePassword
      :showData="openModal"
      :ShowMessage="ShowMessage">
    </ChangePassword>
    <NotIfication
      :showData="openerr">
    </NotIfication>
  </nav>
</template>
<script>
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
import ChangePassword from "./ChangePassword"
import NotIfication from "@/components/Other/Notification"
export default {
  name: 'ManageHeader',
  components: { ChangePassword, NotIfication },
  data(){
    return {
      logo: "",
      openModal:{
        active: false,
        username: ""
      },
      openerr: {
        active: false,
        message: ""
      },
    }
  },
  mounted(){
    this.logo = Config.images[2]
    this.openModal.username = localStorage.getItem('user')
  },
  methods:{
    async LogOut(){
      setStorage(false)
      this.$router.push("/")
    },
    showModel(){
      this.openModal.active = true
    },
    ShowMessage(e){
      this.openerr = e
    }
  }
}
</script>

<style scoped>
nav.navbar {
  border-top: 4px solid #276cda;
}
</style>