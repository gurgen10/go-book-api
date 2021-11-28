import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";

// Important !!!!!!!!!!!!!!!!!
// for running the app - type on the terminal
// google-chrome --disable-site-isolation-trials --disable-web-security --user-data-dir="/tmp"

// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

// Import Bootstrap an BootstrapVue CSS files (order is important)
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
