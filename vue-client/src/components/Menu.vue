<template>
  <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="sidebar-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/dashboard" active-class="active" class="nav-link">Dashboard</router-link>
        </li>
        <li class="nav-item" v-if="user.canView('users')">
          <router-link to="/newNotes" active-class="active" class="nav-link">New Notes</router-link>
        </li>
        <li class="nav-item" v-if="user.canView('users')">
          <router-link to="/completedNotes" active-class="active" class="nav-link">Complated Notes</router-link>
        </li>      
        <li class="nav-item" v-if="user.canView('products')">
          <router-link to="/groups" active-class="active" class="nav-link">Groups</router-link>
        </li>      
      </ul>
    </div>
  </nav>
</template>

<script>
import {computed} from 'vue';
import {useStore} from "vuex";
export default {
  name: "Menu",
  setup() {
    const store = useStore();
    const user = computed(() => store.state.User.user);
    return {
      user
    }
  }
}
</script>