<template>
  <div  class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 mt-3 border-bottom">
    <div class="btn-toolbar mb-2 mb-md-0 ml-3"  >
    <h5 class="pr-2">My New Todo List </h5>  <router-link to="/todo/create" class="btn btn-sm btn-outline-secondary">Add</router-link>
    </div>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th>#</th>
        <th>Name</th>
        <th>Content</th>
        <th>Priority</th>
        <th>Actions</th>

      </tr>
      </thead>
      <tbody>
      <tr v-for="todo in todos" :key="todo.id">
        <td>{{ todo.id }}</td>
        <td >{{ todo.title }}</td>
        <td>{{ todo.content }} </td>  
        <td>{{ todo.priority }} </td>  
        <td>
          <div class="btn-group mr-2" >
            <router-link :to="`/todo/${todo.id}/edit`" class="btn btn-sm btn-outline-secondary">Edit
            </router-link>
             <a href="javascript:void(0)" class="btn btn-sm btn-outline-primary" @click="comp(todo.id)">Complete</a>
            <a href="javascript:void(0)" class="btn btn-sm btn-outline-danger" @click="del(todo.id)">Delete</a>
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
  name: "NewTodos",
  components: {
    
  },
  setup() {
    const todos = ref([]);

    const del = async (id: number) => {
      if (confirm('Are you sure you want to delete this record?')) {
         await axios.delete(`todo/${id}`);
         todos.value = todos.value.filter((e: Entity) => e.id !== id);
       }
       
    }
    const comp = async (id: number) => {
      if (confirm('Are you sure you want to complete this record?')) {
         await axios.put(`complete/${id}`);
         todos.value = todos.value.filter((e: Entity) => e.id !== id);
       }
       
     }
    onMounted(async ()=>{
         const {data} = await axios.get(`todos`);
         debugger
           todos.value = data;
    })
    return {
      todos, 
       del,
       comp
    }
  }
}

</script>