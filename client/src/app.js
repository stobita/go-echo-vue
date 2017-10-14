import Vue from 'vue'
import axios from 'axios'

var app = new Vue({
  el: '#main-flame',
  date: {
    login: false
  }
})

var app = new Vue({
  el: '#login-form',
  data: {
    email: ''
  },
  methods: {
    createUser: function() {
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
