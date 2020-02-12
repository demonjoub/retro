<!-- <template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template> -->
<template>
  <div id="app">
    <!-- <router-link to="page1">Page1</router-link>
    <router-link to="page2">Page1</router-link>
    <router-view></router-view> -->
    <md-bottom-bar class="md-accent" md-shift>
      <md-bottom-bar-item md-icon="address-book" @click="page1()">Page1</md-bottom-bar-item>
      <md-bottom-bar-item md-icon="plus" @click="page2()">Page2</md-bottom-bar-item>
    </md-bottom-bar>
    <router-view></router-view>
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
  },
  methods: {
    page1() {
      this.$router.push({ path: "/page1"})
    },
    page2() {
      this.$router.push({ path: "/page2"})
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
