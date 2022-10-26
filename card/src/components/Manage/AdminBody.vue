<template>
  <div class="columns">
    <div class="column is-6">
      <CreateCard :CallBack="callBack"></CreateCard>
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
                <input class="input is-large" type="text" placeholder="">
                <span class="icon is-medium is-left">
                  <i class="fa fa-search"></i>
                </span>
                <span class="icon is-medium is-right">
                  <i class="fa fa-check"></i>
                </span>
              </div>
              <div class="control">
                <a class="button is-large is-info">
                  搜索
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <CardTable :listData="list"></CardTable>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
import Config from '@/helper/config'
import CreateCard from '@/components/Manage/CreateCard'
import CardTable from '@/components/Manage/CardTable'
export default defineComponent({
  name: 'AdminBody',
  components: { CreateCard, CardTable },
  data(){
    return {
      list:{
        loading: false,
        data: [],
        image: Config.images[3],
      }
    }
  },
  methods:{
    callBack(e){
      this.list = e
    }
  }
})
</script>
