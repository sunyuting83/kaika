<template>
  
</template>
<script>
import { defineComponent } from 'vue'

export default defineComponent({
  setup(props) {
    let states = reactive({
      loading: false
    })
    return {
      ...toRefs(states)
    }
  },
})
</script>
