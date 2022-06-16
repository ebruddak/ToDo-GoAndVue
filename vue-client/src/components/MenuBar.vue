<template>
  <nav id="sidebarMenu" class="col-md-2 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="sidebar-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/todo/create" active-class="active" class="nav-link">Create New Todo</router-link>
          <router-link to="/Groups" active-class="active" class="nav-link">Todo Groups</router-link>
          <router-link to="/NewTodos" active-class="active" class="nav-link">New Todos</router-link>
          <router-link to="/CompetedTodos" active-class="active" class="nav-link">Completed Todos</router-link>
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