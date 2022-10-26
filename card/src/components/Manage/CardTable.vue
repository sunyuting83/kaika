<template>
  <div class="column is-6">
    <div class="card events-card">
      <header class="card-header">
        <p class="card-header-title">
          新卡记录
        </p>
      </header>
      <div class="card-content">
        <div class="content has-text-centered	min-heights" :style="listData.loading || listData.data.length <= 0 ? 'min-height: 14rem':'min-height: 11.3rem'">
          <div class="com__box" v-if="listData.loading" :style="listData.loading? 'margin-top:5rem':''">
            <div class="loading"></div>
          </div>
          <div v-else>
            <img :src="listData.image" alt="nothing" v-if="listData.data.length <= 0" />
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
export default defineComponent({
  name: 'CreateCard',
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
      image: {
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
.loading {
  width: 16px;
  height: 16px;
  box-shadow: 0 30px, 0 -30px;
  border-radius: 4px;
  background: currentColor;
  display: inline-block;
  position: relative;
  color: #ddd;
  left: -30px;
  animation: loading-animation 2s ease infinite;
}

.loading::after,
.loading::before {
  content: "";
  width: 16px;
  height: 16px;
  box-shadow: 0 30px, 0 -30px;
  border-radius: 4px;
  background: currentColor;
  color: #ddd;
  position: absolute;
  left: 30px;
  top: 0;
  animation: loading-animation 2s 0.2s ease infinite;
}

.loading::before {
  animation-delay: 0.4s;
  left: 60px;
}

@keyframes loading-animation {
  0% {
    top: 0;
    color: #ddd;
  }
  50% {
    top: 30px;
    color: rgba(255, 255, 255, 0.2);
  }
  100% {
    top: 0;
    color: #ddd;
  }
}
</style>
