<template lang="html">
  <div>
    <h1>test</h1>
    <FoodItem />
  </div>
</template>

<script>
import FoodItem from "./../../features/foods/FoodList";
export default {
  name: "foodmenu",
  components: {
    FoodItem
  }
};
</script>