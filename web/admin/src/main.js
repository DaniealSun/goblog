import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugin/authui'
import './plugin/http'
import './assets/CSS/style.css'


// 禁用生产环境提示
Vue.config.productionTip = false


new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
