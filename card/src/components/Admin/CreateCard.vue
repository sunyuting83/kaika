<template>
  <div class="column is-6">
    <div class="card">
      <header class="card-header">
        <p class="card-header-title">
          开新卡
        </p>
      </header>
      <div class="card-content">
        <div class="content">
          <div class="field has-addons">
            <p class="control is-expanded">
              <input class="input" type="number" placeholder="开卡时间">
            </p>
            <p class="control">
              <span class="select">
                <select>
                  <option value="m">月</option>
                  <option value="d">天</option>
                  <option value="y">年</option>
                </select>
              </span>
            </p>
            <p class="control is-expanded">
              <input class="input" type="number" placeholder="开卡数量">
            </p>
            <p class="control">
              <a class="button is-info">
                开卡
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="card mt-3">
      <header class="card-header">
        <p class="card-header-title">
          搜索卡
        </p>
      </header>
      <div class="card-content">
        <div class="content">
          <div class="control has-icons-left has-icons-right">
            <input class="input is-large" type="text" placeholder="">
            <span class="icon is-medium is-left">
              <i class="fa fa-search"></i>
            </span>
            <span class="icon is-medium is-right">
              <i class="fa fa-check"></i>
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'CreateCard',
  setup() {
    
  },
})
</script>
