<template>
  <div id="page1">
    <h1>Page 1 {{ message }}</h1>
    <Foods />
    <Sports />
    <Display />
  </div>
</template>

<script>
import Foods from "./../../features/Foods";
import Sports from "./../../features/Sports";
import Display from "./../../features/Display";

export default {
  name: "page1",
  components: {
    Foods,
    Sports,
    Display
  },
  data() {
    return {
      message: "Welcome to vuejs page"
    };
  }
};
</script>