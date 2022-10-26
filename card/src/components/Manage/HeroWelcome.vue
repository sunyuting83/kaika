<template>
  <div>
  <section class="hero is-info welcome is-small">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Hello World.
          </h1>
          <h2 class="subtitle">
            {{jt}}
          </h2>
        </div>
      </div>
    </section>
    <section class="info-tiles">
      <div class="tile is-ancestor has-text-centered">
        <div class="tile is-parent">
          <article class="tile is-child box">
            <p class="title">{{city}}</p>
            <p class="subtitle">城市</p>
          </article>
        </div>
        <div class="tile is-parent">
          <article class="tile is-child box">
            <p class="title">{{now.temperature}}°C</p>
            <p class="subtitle">气温</p>
          </article>
        </div>
        <div class="tile is-parent">
          <article class="tile is-child box">
            <p class="title">{{now.text}}</p>
            <p class="subtitle">气候</p>
          </article>
        </div>
        <div class="tile is-parent">
          <article class="tile is-child box">
            <p class="title">IP</p>
            <p class="subtitle">{{ip}}</p>
          </article>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'HeroWelcome.vue',
  props: {
    ip: String,
    city: String,
    now:{
      temperature: String,
      text: String
    },
    jt: String
  }
})
</script>