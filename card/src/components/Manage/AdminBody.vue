<template>
  <div class="columns">
    <div class="column is-6">
      <CreateCard :CallBack="callBack"></CreateCard>
      <SearchCard :CallBack="callBack"></SearchCard>
    </div>
    <CardTable :listData="list"></CardTable>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
import CreateCard from '@/components/Manage/CreateCard'
import CardTable from '@/components/Manage/CardTable'
import SearchCard from '@/components/Manage/SearchCard'
export default defineComponent({
  name: 'AdminBody',
  components: { CreateCard, CardTable, SearchCard },
  data(){
    return {
      list:{
        loading: false,
        data: [],
        title: "记录显示板"
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
