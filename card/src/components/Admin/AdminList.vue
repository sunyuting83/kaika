<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <div class="card events-card">
        <header class="card-header">
          <p class="card-header-title">
            管理员列表
          </p>
          <button class="card-header-icon">
            <button class="button is-link is-small" @click="showAddModel">
              <span class="icon is-small">
                <i class="fa fa-plus"></i>
              </span>
              <span>添加管理员</span>
            </button>
          </button>
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
              <table class="table is-striped is-hoverable is-fullwidth is-narrow has-text-left	" v-else>
                <thead>
                  <tr>
                    <td>用户名</td>
                    <td>状态</td>
                    <td>创建时间</td>
                    <td>操作</td>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(item) in data" :key="item.id">
                    <td width="35%">{{item.username}}</td>
                    <td>
                      <span class="has-text-success" v-if="item.status === 0">正常</span>
                      <span class="has-text-danger" v-else>锁定</span>
                    </td>
                    <td>{{FormatTime(item.createdtime)}}</td>
                    <td width="35%">
                      <div class="buttons">
                        <button class="button is-success is-small" @click="showModel(item.username)">修改密码</button>
                        <PopoButton
                          message="锁定" color="is-info" :callBack="lockIt" v-if="item.username !== username"></PopoButton>
                        <PopoButton
                          message="删除" color="is-danger" :callBack="deleteIt" v-if="item.username !== username"></PopoButton>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <hr />
      <nav class="pagination" role="navigation" aria-label="pagination">
        <a class="pagination-previous" @click="setPreviousPage" title="This is the first page" :disabled="current !== 1 && page.length > 1?'true':'false'">上一页</a>
        <a class="pagination-next" @click="setNextPage" :disabled="current !== page.length && page.length > 1?'true':'false'">下一页</a>
        <ul class="pagination-list">
          <li v-for="(item) in page" :key="item">
            <a class="pagination-link" @click="setPage(item)" :class="item === current ? 'is-current': ''" :aria-label="'Page '+item" :aria-current="item === current ? 'page': ''">{{item}}</a>
          </li>
        </ul>
      </nav>
    </div>
    <ChangePassword
      :showData="openModal"
      :ShowMessage="ShowMessage">
    </ChangePassword>
    <AddAdmin
      :showData="openAddModal"
      :ShowMessage="ShowMessage">
    </AddAdmin>
    <NotIfication
      :showData="openerr">
    </NotIfication>
  </div>
</template>
<script>
import { reactive, toRefs, onMounted, defineComponent } from 'vue'
import ManageHeader from '@/components/Other/Header'
import LoadIng from '@/components/Other/Loading'
import EmptyEd from '@/components/Other/Empty'
import ChangePassword from "@/components/Other/ChangePassword"
import NotIfication from "@/components/Other/Notification"
import AddAdmin from "@/components/Admin/AddAdmin"
import PopoButton from '@/components/Other/PopoButton'

import Fetch from '@/helper/fetch'
import CheckLogin from '@/helper/checkLogin'
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
export default defineComponent({
  name: 'AdminList',
  components: { ManageHeader, LoadIng, EmptyEd, ChangePassword, NotIfication, AddAdmin, PopoButton },
  setup() {
    let states = reactive({
      loading: false,
      data: [],
      total: 0,
      current: 1,
      page: [],
      username: "",
      openModal:{
        active: false,
        username: ""
      },
      openAddModal:{
        active: false
      },
      openerr: {
        active: false,
        message: "",
        color: ""
      }
    })

    onMounted(async() => {
      const data = await CheckLogin()
      if (data == 0) {
        const username = localStorage.getItem('user')
        states.username = username
        GetData()
      }else{
        setStorage(false)
        this.$router.push("/")
      }
    })

    const makePage = (t) => {
      let x = []
      const p = Math.ceil(t/100)
      for (let i = 0; i < p; i++) {
        x = [...x, i + 1]
      }
      return x
    }

    const setPage = (p) =>{
      if (p !== states.current && p >= 1) {
        states.current = p
        GetData()
      }
    }
    const setNextPage = () =>{
      if (this.states !== states.page.length) {
        states.current = states.current + 1
        GetData()
      }
    }
    const setPreviousPage = () =>{
      if (states.current > 1) {
        states.current = states.current - 1
        GetData()
      }
    }
    const GetData = async() => {
      const token = localStorage.getItem("token")
      const d = await Fetch(Config.Api.adminlist, {}, 'GET', token)
      states.loading = true
      if (d.status == 0) {
        states.data = d.data
        states.total = d.total
        states.page = makePage(d.total)
        states.loading = false
      }else{
        states.data = []
        states.total = 0
        states.page = []
        states.loading = false
      }
    }
    const FormatTime = (timestamp) => {
      var date = new Date(timestamp * 1000);
      var Y = date.getFullYear() + '-'
      var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-'
      var D = date.getDate() + ' '
      return Y+M+D
    }
    /**
     * 
     * @param {*} e message用到的值
     * @param {*} status 0默认不传参 1添加-加入列表 2锁定-替换列表值 3删除-filter值
     */
    const ShowMessage = (e, status = 0) => {
      states.openerr.active = e.active
      states.openerr.message = e.message
      states.openerr.color = e.color
      switch (status) {
        case 1:
          states.data = [...states.data, e.data]
          break;
        default:
          break;
      }
    }
    const showModel = (username) => {
      states.openModal.active = true
      states.openModal.username = username
    }
    const showAddModel = () => {
      states.openAddModal.active = true
    }
    const lockIt = () => {
      console.log("e")
    }
    const deleteIt = () => {
      console.log("e")
    }

    return {
      ...toRefs(states),
      setPage,
      setNextPage,
      setPreviousPage,
      FormatTime,
      ShowMessage,
      showModel,
      showAddModel,
      lockIt,
      deleteIt
    }
  },
})
</script>
