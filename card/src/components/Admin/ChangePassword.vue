<template>
  <div class="modal is-clipped" :class="this.showData.active ? 'is-active': ''" v-if="this.showData.active">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">{{title}}</p>
        <button class="delete" aria-label="close" @click="closErr"></button>
      </header>
      <section class="modal-card-body">
        <!-- Content ... -->
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success">修改</button>
        <button class="button" @click="closErr">取消</button>
      </footer>
    </div>
  </div>
</template>

<script>
import { reactive, toRefs, watch, defineComponent } from 'vue'
export default defineComponent ({
  name: 'ChangePassword',
  props: {
    showData:{
      active:{
        type: Boolean,
        default: false
      },
      title: {
        type: String,
        default: ""
      }
    }
  },
  emits: ["update:active","update:message","update:title"],
  setup(props, context){
    let _data = reactive({
      loading: false
    })
    watch(() => props.showData.active,(newValue) => {
      context.emit("update:active", newValue)
      openError()
    })
    watch(() => props.showData.title,(newValue) => {
      context.emit("update:title", newValue)
    })
    watch(() => _data.loading,(newValue) => {
      _data.loading = newValue
    })
    const closErr = () => {
      const _this = props
      _this.showData.message = ""
      _this.showData.active = false
    }
    return {
      ...toRefs(_data),
      closErr
    }
  }
})
</script>
