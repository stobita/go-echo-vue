import Vue from 'vue'
import axios from 'axios'

let gemail = ''

// プレビュー表示
Vue.component('preview-email',{
  template: '<p>email:{{ disemail }}</p>',
  props: ['disemail']
})

// 入力フォーム
Vue.component('email-input',{
  template: '<input v-on:input="emailInput" v-model="inemail" placeholder="edit me">',
  data (){
    return {
      inemail: ''
    }
  },
  methods: {
    emailInput: function(){
      this.$emit('input-email', this.inemail)
    }
  }
})

// 送信ボタン
Vue.component('email-submit',{
  template: '<button v-on:click="createUser" >送信</button>',
  methods: {
    createUser: function() {
      this.$emit('submit-email')
    }
  }
})

// root インスタンスを作成する
new Vue({
  el: '#login-form',
  data:{
    email: '',
    sendResult: ''
  },
  methods:{
    emailValue: function(inEmail){
      this.email = inEmail
    },
    sendEmail: function(){
      axios.post('http://localhost:1323/users/',{
        email: this.email
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
    }
  }
})
