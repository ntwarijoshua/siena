import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './styles.scss'
import 'bootstrap'
import {library} from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {faEnvelope, faBell } from '@fortawesome/free-regular-svg-icons';
import {faCaretDown, faSearch} from '@fortawesome/free-solid-svg-icons'
library.add(
  faEnvelope,
  faBell,
  faCaretDown,
  faSearch
)
Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
