<template>
  <nav id="sidebarMenu" class="col-md-2 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="sidebar-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/todo/create" active-class="active" class="nav-link">Create New Todo</router-link>
          <router-link :to="`/groups`" active-class="active" class="nav-link">Todo Groups</router-link>
          <router-link :to="`/NewTodos`" active-class="active" class="nav-link">Active Todos</router-link>
          <router-link :to="`/CompetedTodos`" active-class="active" class="nav-link">Completed Todos</router-link>
        </li>
       
      </ul>
    </div>
  </nav>
</template>

<script lang="ts">
 import {computed,onMounted,ref} from 'vue';
 import {useRouter} from "vue-router";
 import axios from 'axios';
export default {
  name: "MenuBar",
  setup() {
        const router = useRouter();
    const username=ref('')
    const userId=ref('')
     onMounted(async () => {
      const {data} =await axios.get('user')
       username.value=data.userName ; 
       userId.value=data.id ;     
     });
      const logout = async () => {
       await axios.post('logout', {});
            router.push('/login')

     }
   return {
    username,logout,userId
  }
}}
</script>