<template>
  <div  class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 mt-3 border-bottom">
    <div class="btn-toolbar mb-2 mb-md-0 ml-3"  >
      <router-link to="/groups/create" class="btn btn-sm btn-outline-secondary">Add</router-link>
    </div>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th>#</th>
        <th>Todo's Group Name</th>
       
      </tr>
      </thead>
      <tbody>
      <tr v-for="group in groups" :key="group.id">
        <td>{{ group.id }}</td>
        <td>{{ group.name }} </td>  
        <td>
          <div class="btn-group mr-2" >
            <router-link :to="`/groups/${group.id}/edit`" class="btn btn-sm btn-outline-secondary">Edit
            </router-link>
            <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(group.id)">Delete</a>
          </div>
        </td>
      </tr>
      </tbody>
    </table>
  </div>

</template>

<script lang="ts">
import {ref, onMounted, computed} from 'vue';
import axios from 'axios';
import {Entity} from "@/interfaces/entity";
import {useStore} from "vuex";
export default {
  name: "TodoGroups",
  components: {
    
  },
  setup() {
    const groups = ref([]);
    // const lastPage = ref(0);
    // const store = useStore();
    // const authenticatedUser = computed(() => store.state.User.user);
    // const load = async (page = 1) => {
    //   const response = await axios.get(`groups`);
    //   groups.value = response.data.data;
    //   lastPage.value = response.data.meta.last_page;
    // }
     const del = async (id: number) => {
      if (confirm('Are you sure you want to delete this record?')) {
         await axios.delete(`group/${id}`);
         groups.value = groups.value.filter((e: Entity) => e.id !== id);
       }
    }
    onMounted(async ()=>{
         const {data} = await axios.get(`groups`);
         debugger
           groups.value = data;
    })
    return {
      groups, 
       del
    //   load
    }
  }
}
</script>