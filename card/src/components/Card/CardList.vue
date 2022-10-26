<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <table class="table is-striped is-hoverable is-fullwidth is-narrow">
        <thead>
          <tr>
            <td>卡号</td>
            <td>剩余时间</td>
            <td>到期时间</td>
            <td>开卡时间</td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
import { reactive, toRefs, defineComponent } from 'vue'
import ManageHeader from '@/components/Other/Header.vue'

export default defineComponent({
  name: 'CardList',
  components: { ManageHeader },
  setup() {
    let states = reactive({
      loading: false,
      data: []
    })
    return {
      ...toRefs(states)
    }
  },
})
</script>
