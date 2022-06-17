<template>
<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
  <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#">Todo App (Go-Vue3)</a>
  <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="navbar-nav">
    <div class="nav-item text-nowrap">
      <a class="p-2 text-white text-decoration-none" href="#">{{username}}</a>
      <router-link to="/login" class="nav-link px-3" @click="logout">Sign out</router-link>
    </div>
  </div>
</header>
</template>

<script lang="ts">
 import {computed,onMounted,ref} from 'vue';
 import {useRouter} from "vue-router";
 import axios from 'axios';
export default {
  name: "NavBar",
  
  setup() {
        const router = useRouter();
    const username=ref('')
     onMounted(async () => {
      const {data} =await axios.get('user')
      console.log(data)
       username.value=data.userName ;   
     });
      const logout = async () => {
       await axios.post('logout', {});
            router.push('/login')

     }
   return {
    username,logout
  }
  }

  }

</script>