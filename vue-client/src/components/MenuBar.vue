<template>
  <nav id="sidebarMenu" class="col-md-2 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="sidebar-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/Dashboard" active-class="active" class="nav-link">Home</router-link>
          <router-link to="/Groups" active-class="active" class="nav-link">Todo Groups</router-link>
          <router-link to="/NewTodo" active-class="active" class="nav-link">New Todos</router-link>
          <router-link to="/ComplatedTodo" active-class="active" class="nav-link">Complated Todos</router-link>
        </li>
       
      </ul>
    </div>
  </nav>
</template>

<script lang="ts">
// import {computed} from 'vue';
// import {useStore} from "vuex";
export default {
  name: "MenuBar",
  setup() {
    // const store = useStore();
    // const user = computed(() => store.state.User.user);
    // return {
    //   user
    // }
  }
}
</script>