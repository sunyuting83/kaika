<template>
  <div>
    <ManageHeader></ManageHeader>
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
        <CreateCard></CreateCard>

        <div class="column is-6">
          <div class="card events-card">
            <header class="card-header">
              <p class="card-header-title">
                新卡记录
              </p>
            </header>
            <div class="card-content">
              <div class="content has-text-centered	min-heights">
                <img :src="nothong" />
                <!-- <table class="table is-striped is-clipped">
                  <tbody>
                    <tr>
                      <td>73A90ACAAE2B1CCC0E969709665BC62F</td>
                    </tr>
                    <tr>
                      <td>73A90ACAAE2B1CCC0E969709665BC62F</td>
                    </tr>
                  </tbody>
                </table> -->
              </div>
            </div>
            <!-- <footer class="card-footer">
              <a class="card-footer-item is-size-7" @click="copyAccount">
                <span class="level-item" v-show="check">
                  <span class="icon is-small"><i class="fa fa-check"></i></span>
                </span>
                复制到剪切板
              </a>
            </footer> -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Fetch from '@/helper/fetch'
import CheckLogin from '@/helper/checkLogin'
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'

import ManageHeader from './Header'
import CreateCard from './CreateCard'

import useClipboard from 'vue-clipboard3'
const { toClipboard } = useClipboard()
export default {
  name: 'AppManage',
  components: { ManageHeader, CreateCard },
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
  async mounted(){
    const data = await CheckLogin()
    if (data == 0) {
      this.ip = Config.IP
      this.logo = Config.images[2]
      this.nothong = Config.images[3]
      this.GetJT()
    }else{
      setStorage(false)
      this.$router.push("/")
    }
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
.min-heights {
  min-height: 14rem;
}
</style>