<template>
  <div class="column is-6">
    <div class="card events-card">
      <header class="card-header">
        <p class="card-header-title">
          {{listData.title}}
        </p>
      </header>
      <div class="card-content">
        <div class="content has-text-centered	min-heights" :style="listData.loading || listData.data.length <= 0 ? 'min-height: 14rem':'min-height: 11.3rem'">
          <div class="com__box" v-if="listData.loading" :style="listData.loading? 'margin-top:5rem':''">
            <LoadIng></LoadIng>
          </div>
          <div v-else>
            <div v-if="listData.data.length <= 0">
              <EmptyEd></EmptyEd>
            </div>
            <table class="table is-striped is-clipped has-text-left	" v-else>
              <tbody>
                <tr v-for="(item, index) in listData.data" :key="index">
                  <td>{{item}}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <footer class="card-footer" v-if="listData.data.length > 0">
        <a class="card-footer-item is-size-7" @click="copyAccount">
          <span class="level-item" v-if="check">
            <span class="icon is-small"><i class="fa fa-check"></i></span>
          </span>
          复制到剪切板
        </a>
      </footer>
    </div>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
import useClipboard from 'vue-clipboard3'
const { toClipboard } = useClipboard()
import LoadIng from '@/components/Other/Loading'
import EmptyEd from '@/components/Other/Empty'
export default defineComponent({
  name: 'CreateCard',
  components: { LoadIng, EmptyEd },
  data(){
    return {
      check: false
    }
  },
  props: {
    listData: {
      loading: {
        type: Boolean
      },
      data: {
        type: Array
      },
      title: {
        type: String
      }
    }
  },
  methods:{
    async copyAccount(){
      const _this = this
      let data = ""
      _this.listData.data.forEach(element => {
        data += `${element}\r\n`
      });
      data = data.replace(/(^\r\n+)|(\r\n+$)/g, "")
      await toClipboard(data)
      this.check = true
    }
  }
})
</script>

<style scoped>
.min-heights {
  min-height: 11.3rem;
  height: 11.3rem;
  overflow-y: auto;
}
</style>
