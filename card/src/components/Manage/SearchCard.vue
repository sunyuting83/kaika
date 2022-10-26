<template>
  <div class="card mt-3">
    <header class="card-header">
      <p class="card-header-title">
        搜索卡
      </p>
    </header>
    <div class="card-content">
      <div class="content">
        <div class="field has-addons">
          <div class="control has-icons-left has-icons-right is-expanded">
            <input class="input is-large" v-model="form.search" type="text" placeholder="请输入卡号">
            <span class="icon is-medium is-left">
              <i class="fa fa-search"></i>
            </span>
            <span class="icon is-medium is-right" v-if="form.search.length == 32">
              <i class="fa fa-check"></i>
            </span>
          </div>
          <div class="control">
            <button class="button is-large is-info" :disabled="form.search.length == 32 ? false : true" @click="onSubmit">
              搜索
            </button>
          </div>
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
  name: 'SearchCard',
  props: {
    CallBack: Function
  },
  data(){
    return {
      form:{
        search: ""
      },
      loading: false
    }
  },
  methods:{
    async onSubmit(){
      const { search } = this.form
      if (search.length == 32) {
        this.loading = true
        const params = {
          'card': search
        }
        const token = localStorage.getItem("token")
        const d = {
          data: [],
          loading: true,
          title: "搜索记录"
        }
        this.CallBack(d)
        const data = await Fetch(Config.Api.search, params, 'POST', token)
        if (data.status === 0) {
          const d = {
            data: [data.data],
            loading: false,
          title: "搜索记录"
          }
          this.loading = false
          this.CallBack(d)
        }else{
          const d = {
            data: [],
            loading: false,
          title: "搜索记录"
          }
          this.loading = false
          this.CallBack(d)
        }
      }
    }
  }
})
</script>
