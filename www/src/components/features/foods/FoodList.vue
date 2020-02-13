<template lang="html">
  <div>
    <h1>รายการอาหาร ทั้งหมด</h1>
  </div>
</template>

<script>
export default {
  mounted() {
    this.setup()
  },
  methods: {
    setup() {
      this.getFoodsList()
    },
    getFoodsList() {
      
    }
  }
}
</script>