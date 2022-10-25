<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <HeroWelcome
        :ip="ip"
        :city="city"
        :now="now"
        :jt="jt"></HeroWelcome>
      <AdminBody></AdminBody>
    </div>
  </div>
</template>

<script>
import Fetch from '@/helper/fetch'
import CheckLogin from '@/helper/checkLogin'
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
import returnCitySN from 'returnCitySN'

import ManageHeader from '@/components/Admin/Header'
import AdminBody from '@/components/Admin/AdminBody'
import HeroWelcome from '@/components/Admin/HeroWelcome'

export default {
  name: 'AppManage',
  components: { ManageHeader, AdminBody, HeroWelcome },
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
    localStorage.setItem("ip",returnCitySN.cip)
    const data = await CheckLogin()
    if (data == 0) {
      this.ip = returnCitySN.cip
      this.logo = Config.images[2]
      this.GetJT()
    }else{
      setStorage(false)
      this.$router.push("/")
    }
  },
  methods:{
    async GetJT(){
      const data = await Fetch(Config.Api.jt,{})
      if (data.code === 1) {
        this.jt = data.content
      }
      const tianqi = await Fetch(Config.makeIP(returnCitySN.cip), {})
      if (tianqi.status !== 'status_code') {
        const tq = tianqi.results[0]
        this.city = tq.location.name
        this.now = tq.now
      }
    }
  },
}
</script>
<style scoped>
nav.navbar {
  border-top: 4px solid #276cda;
}
</style>