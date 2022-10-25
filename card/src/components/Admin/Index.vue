<template>
  <div>
    <nav class="navbar is-white">
      <div class="container">
        <div class="navbar-brand">
          <span class="navbar-item brand-text">
            <img :src="logo">Card Manage
          </span>
        </div>
        <div id="navMenu" class="navbar-menu">
          <div class="navbar-start">
            <a class="navbar-item" href="admin.html">
              首页
            </a>
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
                    <a class="navbar-item">
                      退出登陆
                    </a>
                  </div>
                </div>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <!-- END NAV -->
    <div class="container">
      <section class="hero is-info welcome is-small">
        <div class="hero-body">
          <div class="container">
            <h1 class="title">
              Hello World.
            </h1>
            <h2 class="subtitle">
              {{jt}}
            </h2>
          </div>
        </div>
      </section>
      <section class="info-tiles">
        <div class="tile is-ancestor has-text-centered">
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">{{city}}</p>
              <p class="subtitle">城市</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">{{now.temperature}}°C</p>
              <p class="subtitle">气温</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">{{now.text}}</p>
              <p class="subtitle">气候</p>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">IP</p>
              <p class="subtitle">{{ip}}</p>
            </article>
          </div>
        </div>
      </section>
      <div class="columns">
        <div class="column is-6">
          <div class="card">
            <header class="card-header">
              <p class="card-header-title">
                开新卡
              </p>
            </header>
            <div class="card-content">
              <div class="content">
                <div class="field has-addons">
                  <p class="control is-expanded">
                    <input class="input" type="number" placeholder="开卡时间">
                  </p>
                  <p class="control">
                    <span class="select">
                      <select>
                        <option value="m">月</option>
                        <option value="d">天</option>
                        <option value="y">年</option>
                      </select>
                    </span>
                  </p>
                  <p class="control is-expanded">
                    <input class="input" type="number" placeholder="开卡数量">
                  </p>
                  <p class="control">
                    <a class="button is-info">
                      开卡
                    </a>
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card mt-3">
            <header class="card-header">
              <p class="card-header-title">
                搜索卡
              </p>
            </header>
            <div class="card-content">
              <div class="content">
                <div class="control has-icons-left has-icons-right">
                  <input class="input is-large" type="text" placeholder="">
                  <span class="icon is-medium is-left">
                    <i class="fa fa-search"></i>
                  </span>
                  <span class="icon is-medium is-right">
                    <i class="fa fa-check"></i>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="column is-6">
          <div class="card events-card">
            <header class="card-header">
              <p class="card-header-title">
                新卡记录
              </p>
            </header>
            <div class="card-content">
              <div class="content">
                <table class="table is-striped is-clipped">
                  <tbody>
                    <tr>
                      <td>73A90ACAAE2B1CCC0E969709665BC62F</td>
                    </tr>
                    <tr>
                      <td>73A90ACAAE2B1CCC0E969709665BC62F</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
            <footer class="card-footer">
              <a class="card-footer-item is-size-7" @click="copyAccount">
                <span class="level-item" v-show="check">
                  <span class="icon is-small"><i class="fa fa-check"></i></span>
                </span>
                复制到剪切板
              </a>
            </footer>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
import useClipboard from 'vue-clipboard3'
const { toClipboard } = useClipboard()
export default {
  name: 'AppManage',
  data(){
    return {
      jt: "破罐子不能破摔，得使劲摔！",
      city: '',
      now: {
        code: "0",
        temperature: "0",
        text: ""
      },
      ip:"",
      check: false
    }
  },
  mounted(){
    this.ip = Config.IP
    this.logo = Config.images[2]
    this.GetJT()
  },
  methods:{
    async GetJT(){
      const tianqi = await Fetch(Config.Api.tianqi, {})
      const tq = tianqi.results[0]
      this.city = tq.location.name
      this.now = tq.now
      const data = await Fetch(Config.Api.jt,{})
      if (data.code === 1) {
        this.jt = data.content
      }
    },
    async copyAccount(){
      const _this = this
      await toClipboard(_this.jt)
      this.onCopy()
      this.check = true
    }
  },
}
</script>
<style scoped>
nav.navbar {
  border-top: 4px solid #276cda;
}
</style>