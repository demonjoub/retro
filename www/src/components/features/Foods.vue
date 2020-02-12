<template lang="html">
  <div>
    <h2>อาหารที่ชอบ</h2>
    <input type="text" name="food" @input="changeFood" value="">
  </div>
</template>

<script>
export default {
  methods: {
    changeFood (event) {
      this.$store.commit('setFood', event.target.value)
    }
  }
}
</script>