<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <nav class="columns">
        <div class="column">
          <div class="field is-grouped">
            <p class="control is-expanded">
              <input class="input is-small" type="text" v-model="SearchKey" placeholder="输入卡号" @input="Search">
            </p>
            <p class="control">
              <button class="button is-small is-info" :disabled="SearchKey.length > 0 ? false : true" @click="Clean">
                清空结果
              </button>
            </p>
          </div>
        </div>
      </nav>
      <div class="card events-card">
        <header class="card-header">
          <p class="card-header-title">
            卡号列表
          </p>
        </header>
        <div class="card-content">
          <div class="content has-text-centered	min-heights" style="min-height: 11.3rem">
            <div class="com__box" v-if="loading" :style="loading? 'margin-top:5rem':''">
              <LoadIng></LoadIng>
            </div>
            <div v-else>
              <div v-if="data.length <= 0">
                <EmptyEd></EmptyEd>
              </div>
              <table class="table is-striped is-hoverable is-fullwidth is-narrow has-text-left" v-else>
                <thead>
                  <tr>
                    <td width="35%">卡号</td>
                    <td>剩余时间</td>
                    <td>到期事件</td>
                    <td>创建时间</td>
                    <td width="15%">操作</td>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(item) in data" :key="item.id">
                    <td>{{item.card}}</td>
                    <td>{{surplus(item.updatetime)}}</td>
                    <td><FormaTime :DateTime="item.updatetime"></FormaTime></td>
                    <td><FormaTime :DateTime="item.createdtime"></FormaTime></td>
                    <td>
                      <div class="buttons">
                        <button class="button is-success is-small">续费</button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <PaginAtion v-if="data.length > 0" :total="total" :GetData="GetData"></PaginAtion>
    </div>
    <NotIfication
      :showData="openerr">
    </NotIfication>
  </div>
</template>
<script>
import { reactive, toRefs, onMounted, defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import ManageHeader from '@/components/Other/Header'
import LoadIng from '@/components/Other/Loading'
import EmptyEd from '@/components/Other/Empty'
import NotIfication from "@/components/Other/Notification"
import PaginAtion from '@/components/Other/PaginAtion'
import FormaTime from '@/components/Other/FormaTime'


import Fetch from '@/helper/fetch'
import CheckLogin from '@/helper/checkLogin'
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
export default defineComponent({
  name: 'CardList',
  components: { ManageHeader, LoadIng, EmptyEd, NotIfication, PaginAtion, FormaTime },
  setup() {
    let states = reactive({
      loading: false,
      temp: [],
      data: [],
      total: 0,
      openerr: {
        active: false,
        message: "",
        color: ""
      },
      currentTime: 0,
      SearchKey: ""
    })
    const router = useRouter()
    onMounted(async() => {
      const data = await CheckLogin()
      if (data == 0) {
        GetData()
      }else{
        setStorage(false)
        router.push("/")
      }
    })
    const GetData = async(page = 1) => {
      const token = localStorage.getItem("token")
      const datetime = await Fetch(Config.Api.currentTime, {}, 'GET', token)
      states.currentTime = datetime.datetime
      const d = await Fetch(Config.Api.cardlist, {page:page}, 'GET', token)
      states.loading = true
      if (d.status == 0) {
        states.data = d.data
        states.temp = d.data
        states.total = d.total
        states.loading = false
      }else{
        states.data = []
        states.total = 0
        states.page = []
        states.loading = false
      }
    }
    const surplus = (d) => {
      const t = states.currentTime
      const s = d - t
      if (s > 0) {
        if (s < 86400) {
          return Math.ceil(s / 3600) + "小时"
        }else if(s < 3600){
          return Math.ceil(s / 60) + "分钟"
        }else {
          return Math.ceil(s / 86400) + "天"
        }
      }else {
        return 0
      }
    }
    const Search = () => {
      const search = states.SearchKey.toUpperCase()
      if (search) {
        states.data = states.temp.filter((data) => {
          return Object.keys(data).some(() => {
            return String(data['card']).indexOf(search) !== -1
          })
        })
      }else {
        states.data = states.temp
      }
    }
    const Clean = () => {
      states.SearchKey = ""
      states.data = states.temp
    }

    return {
      ...toRefs(states),
      surplus,
      GetData,
      Search,
      Clean
    }
  },
})
</script>
