import Vue from "vue";
import App from "./App.vue";
import VueRouter from "vue-router";
import Page1 from "./components/pages/page1/page1.vue";
import Page2 from "./components/pages/page2/page2.vue";
import Menu from "./components/pages/menu/menu.vue";
import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import { store } from "./store/store";

Vue.config.productionTip = false;
Vue.use(VueMaterial);
Vue.use(VueRouter);

const routes = [
  { path: "/page1", component: Page1 },
  { path: "/page2", component: Page2 },
  { path: "/menu", component: Menu }
];
const router = new VueRouter({ routes });

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
