import axios from "axios";

const state = {
  like: "ข้าวเหนียวหมูปิ้ง",
  unlike: "ข้าวมันไก่",
  posts: []
}

const getters = {
  foods: state => state.food,
}

const actions = {
  localPost (context) {
    axios.get("https://jsonplaceholder.typicode.com/posts")
    .then(res => {
      const posts = res.data
      context.commit("SET_POST", posts)
      // commit("SET_POST", posts)
    })
    .catch(error => { console.log(error) })
    .finally(()=> console.log("load success"))
  }
}

const mutations = {
  SET_POST(state, payload) {
    state.posts = payload
  },
  SET_STORE(state, payload) {
    if (payload.actions === 'like') {
      state.like = payload.msg
    } else if (payload.actions === 'unlike') {
      state.unlike = payload.msg
    }
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
