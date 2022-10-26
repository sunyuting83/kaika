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
                          :message="item.status == 0?'锁定':'解锁'" color="is-info" :callBack="()=>{lockIt(item.id)}" v-if="item.username !== username"></PopoButton>
                        <PopoButton
                          message="删除" color="is-danger" :callBack="()=>{deleteIt(item.id)}" v-if="item.username !== username"></PopoButton>
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
import PaginAtion from '@/components/Other/PaginAtion'

import Fetch from '@/helper/fetch'
import CheckLogin from '@/helper/checkLogin'
import Config from '@/helper/config'
import setStorage from '@/helper/setStorage'
export default defineComponent({
  name: 'AdminList',
  components: { ManageHeader, LoadIng, EmptyEd, ChangePassword, NotIfication, AddAdmin, PopoButton, PaginAtion },
  setup() {
    let states = reactive({
      loading: false,
      data: [],
      total: 0,
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
    const GetData = async(page = 1) => {
      const token = localStorage.getItem("token")
      const d = await Fetch(Config.Api.adminlist, {page:page}, 'GET', token)
      states.loading = true
      if (d.status == 0) {
        states.data = d.data
        states.total = d.total
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
    const ShowMessage = (e, status = 0, id=0) => {
      states.openerr.active = e.active
      states.openerr.message = e.message
      states.openerr.color = e.color
      switch (status) {
        case 1:
          states.data = [...states.data, e.data]
          break;
        case 2:
          states.data = states.data.map((e)=>{
            if(e.id == id) {
              if(e.status === 0) {
                e.status = 1
              }else {
                e.status = 0
              }
            }
            return e
          })
          break;
        case 3:
          states.data = states.data.filter((e)=>{
            if(e.id !== id) return e
          })
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
    const lockIt = async(id) => {
      const token = localStorage.getItem("token")
      const d = await Fetch(Config.Api.upstatus, {id: id}, 'PUT', token)
      if (d.status == 0) {
        const data = {
          active: true,
          message: d.message,
          color: 'is-success'
        }
        ShowMessage(data, 2, id)
      }else{
        const data = {
          active: true,
          message: d.message,
          color: 'is-danger'
        }
        ShowMessage(data, 0)
      }
    }
    const deleteIt = async(id) => {
      const token = localStorage.getItem("token")
      const d = await Fetch(Config.Api.deladmin, {id: id}, 'DELETE', token)
      if (d.status == 0) {
        const data = {
          active: true,
          message: d.message,
          color: 'is-success'
        }
        ShowMessage(data, 3, id)
      }else{
        const data = {
          active: true,
          message: d.message,
          color: 'is-danger'
        }
        ShowMessage(data, 0)
      }
    }

    return {
      ...toRefs(states),
      FormatTime,
      ShowMessage,
      showModel,
      showAddModel,
      lockIt,
      deleteIt,
      GetData
    }
  },
})
</script>
