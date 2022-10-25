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
                  <a class="navbar-item">
                    修改密码
                  </a>
                  <a class="navbar-item">
                    管理员管理
                  </a>
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
  </nav>
</template>
<script>
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
export default {
  name: 'ManageHeader',
  data(){
    return {
      logo: ""
    }
  },
  mounted(){
    this.logo = Config.images[2]
  },
  methods:{
    async LogOut(){
      setStorage(false)
      this.$router.push("/")
    }
  }
}
</script>