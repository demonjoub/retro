<template lang="html">
  <div>
    <h2>อาหารที่ชอบ {{ like }}</h2>
    <h2>อาหารที่ไม่ชอบ {{ unlike }}</h2>
    <hr>
    อาหารที่ชอบ :<input type="text" name="like" @input="changeFood" value="">
    <hr>
    อาหารที่ไม่ชอบ :<input type="text" name="unlike" @input="changeFood" value="">
    <hr>
    <button type="button" @click="onClick">Click Me!</button>
    <hr>
    <h2 v-for="(item, index) in posts" :key="index">{{ item.title }}</h2>
  </div>
</template>
<script>
import { mapState } from 'vuex'

export default {
  methods: {
    changeFood (event) {
      let payload = { }
      if (event.target.name === "like") {
        payload = {
          actions: 'like',
          msg : event.target.value
        }
      } else {
        payload = {
          actions: 'unlike',
          msg : event.target.value
        }
      }
      this.$store.commit('SET_STORE', payload)
    },
    onClick (event) {
      this.$store.dispatch('localPost', event)
    }
  },
  computed: {
    ...mapState({
      posts: state => state.foods.posts,
      like: state => state.foods.like,
      unlike: state => state.foods.unlike
    })
  }
}
</script>