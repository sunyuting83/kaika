<template>
  <div class="modal is-clipped" :class="showData.active ? 'is-active': ''" v-if="showData.active">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">卡号续费</p>
        <button class="delete" aria-label="close" @click="closErr" v-if="loading ? false : true"></button>
      </header>
      <section class="modal-card-body">
        <div class="field has-addons">
          <p class="control is-expanded">
            <input class="input" type="number" v-model="form.createTime" placeholder="续费时间">
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
        </div>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-info" @click="onSubmit" :disabled="form.createTime <= 0 ? true : false" :class="loading ? 'is-loading' : ''">续费</button>
        <button class="button" @click="closErr" :disabled="loading ? true : false" :class="loading ? 'is-loading' : ''">取消</button>
      </footer>
    </div>
  </div>
</template>

<script>
import { reactive, toRefs, defineComponent } from 'vue'
import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
export default defineComponent ({
  name: 'RenewalCard',
  props: {
    showData:{
      active:{
        type: Boolean,
        default: false
      },
      card: {
        type: String
      }
    },
    ShowMessage:Function
  },
  setup(props){
    let _data = reactive({
      loading: false,
      form:{
        createTime: 1,
        dateTime: "m"
      }
    })
    const closErr = () => {
      const _this = props
      _data.loading = false
      _this.showData.message = ""
      _this.showData.active = false
    }
    const createDate =() =>{
      let x = 0
      const { createTime, dateTime } = _data.form
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
    const onSubmit = async() => {
      const { createTime } = _data.form
      if (createTime > 0) {
        _data.loading = true
        const realTime = createDate()
        const params = {
          'card': props.showData.card,
          'datetime': realTime
        }
        const token = localStorage.getItem("token")
        const data = await Fetch(Config.Api.renewa, params, 'PUT', token)
        if (data.status === 0) {
          _data.loading = false
          closErr()
          props.ShowMessage({
            active: true,
            message: data.message,
            color: 'is-success',
            newtime: data.updatetime,
            id: data.id
          })
        }else{
          closErr()
          props.ShowMessage({
            active: true,
            message: "请求失败",
            color: 'is-danger',
            newtime: 0,
            id: 0
          })
        }
      }
    }
    return {
      ...toRefs(_data),
      closErr,
      onSubmit
    }
  }
})
</script>
