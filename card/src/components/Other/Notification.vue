<template>
  <div class="notification is-danger error" :style="{'top':hoverTop+'px'}" v-if="props.active">
    <button class="delete" @click="props.closErr"></button>
    <p>{{props.message}}</p>
  </div>
</template>

<script>
import { reactive, toRefs, watch, defineComponent } from 'vue'
export default defineComponent ({
  name: 'NotIfication',
  props: {
    active:{
      type: Boolean,
      default: false
    },
    message: {
      type: String,
      default: ""
    },
    closErr: {
      type: Function
    }
  },
  emits: ["update:active","update:message"],
  setup(props, context){
    let _data = reactive({
      hoverTop: 0
    })
    watch(() => props.active,(newValue) => {
      context.emit("update:active", newValue)
      openError()
    })
    watch(() => props.message,(newValue) => {
      context.emit("update:message", newValue)
    })
    watch(() => _data.hoverTop,(newValue) => {
      _data.hoverTop = newValue
    })
    const openError = () => {
      const scrollTop = document.documentElement.scrollTop || window.pageYOffset || document.body.scrollTop;
      // context.emit("update:hoverTop", scrollTop)
      _data.hoverTop = scrollTop
      const c = props.closErr
      setTimeout(function(){
        c()
      },1500)
    }
    return {
      ...toRefs(_data),
      props,
    }
  }
})
</script>

<style scoped>
.error {
  position:absolute;
  right: 0px;
  width: 240px;
  z-index: 9999999999;
}
</style>