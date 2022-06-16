<template>
  <form class="form-signin" @submit.prevent="submit">
    <h1 class="h3 mb-3 font-weight-normal">Todo App Register</h1>

    <label for="username" class="sr-only">Username</label>
    <input type="text" id="username" class="form-control" placeholder="Username" v-model="username" required >


    <label for="inputEmail" class="sr-only">Email address</label>
    <input type="email" id="inputEmail" class="form-control" placeholder="Email address" v-model="email"  required >

    <label for="inputPassword" class="sr-only">Password</label>
    <input type="password" id="inputPassword" class="form-control" placeholder="Password" required v-model="password">
  

    <button class="btn btn-lg btn-primary btn-block" type="submit">Register</button>
  </form>
</template>

<script lang="ts">
import {ref} from 'vue';
import axios from 'axios';
import {useRouter} from "vue-router";
export default {
  name: "RegisterUser",
  setup() {
    const username = ref('');
    const email = ref('');
    const password = ref('');
    const router = useRouter();
    const submit = async () => {
      debugger;
       await axios.post('register', {
        
         username: username.value,
         email: email.value,
         password: password.value,
       });
       await router.push('/login');
    }
     return {
      username,
       email,
       password,
      submit
     }
   }
}
</script>

<style scoped>
html,
body {
  height: 100%;
}
body {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
}
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>