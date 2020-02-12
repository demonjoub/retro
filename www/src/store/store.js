import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import foods from './modules/foods'

Vue.use(Vuex, axios)
export const store = new Vuex.Store({
  modules: {
    foods
  }
})
// export const store = new Vuex.Store({
//   state: {
//     food: 'หมูปิ้ง',
//     sport: 'ฟุตบอล'
//   }, 
//   mutations: {
//     // method (update, edit => state)
//     setFood(state, food) {
//       state.food = food // <= assign
//     },
//     setSport(state, sport) {
//       state.sport = sport
//     }
//   },
//   getters: {
//     food: state => state.food,
//     sport: state => state.sport
//   }
// })
