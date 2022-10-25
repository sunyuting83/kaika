<template>
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
            <input class="input" type="number" v-model="form.createTime" placeholder="开卡时间">
          </p>
          <p class="control">
            <span class="select">
              <select v-model="form.dateTime">
                <option selected value="m">月</option>
                <option value="d">天</option>
                <option value="y">年</option>
              </select>
            </span>
          </p>
          <p class="control is-expanded">
            <input class="input" v-model="form.createTotal" type="number" placeholder="开卡数量">
          </p>
          <p class="control">
            <button class="button is-info"
              @click="onSubmit"
              :disabled="form.createTime <= 0 || form.createTotal <= 0 ? true : false"
              :class="loading ? 'is-loading' : ''">
              开卡
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
export default defineComponent({
  name: 'CreateCard',
  props: {
    CallBack: Function
  },
  data(){
    return {
      form:{
        createTime: 1,
        dateTime: "m",
        createTotal: 1
      },
      loading: false
    }
  },
  methods:{
    async onSubmit(){
      const { createTime, createTotal } = this.form
      if (createTime > 0 && createTotal > 0) {
        this.loading = true
        const realTime = this.createDate()
        const params = {
          'card': createTotal,
          'datetime': realTime
        }
        const token = localStorage.getItem("token")
        const d = {
          data: [],
          loading: true
        }
        this.CallBack(d)
        const data = await Fetch(Config.Api.creatdcard, params, 'POST', token)
        if (data.status === 0) {
          const d = {
            data: data.data,
            loading: false
          }
          this.loading = false
          this.CallBack(d)
        }
      }
    },
    createDate(){
      let x = 0
      const { createTime, dateTime } = this.form
      switch (dateTime) {
        case 'm':
          x = createTime * 2678400
          break;
        case 'y':
          x = createTime * 32140800
          break;
        case 'd':
          x = createTime * 86400
          break;
        default:
          x = createTime * 86400
          break;
      }
      return x
    }
  }
})
</script>
