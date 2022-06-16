<template>
  <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="sidebar-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/NewNotes" active-class="active" class="nav-link">Dashboard</router-link>
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