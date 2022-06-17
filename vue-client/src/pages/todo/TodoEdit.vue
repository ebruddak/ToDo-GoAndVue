<template>
 <h5 class="pr-2 mt-4"> Edit Todo Item </h5> 
  <form @submit.prevent="submit">
    <div class="form-group">
      <label> Title</label>
      <input type="text" class="form-control" name="title" v-model="title"/>
    </div>
    <div class="form-group">
      <label>Content</label>
      <textarea class="form-control" name="content" v-model="content"></textarea>
    </div>
    <div class="form-group">
      <label>Group</label>
      <select name="groupId" class="form-control" v-model="groupId">
        <option value="0">Select Group</option>
        <option v-for="group in groups" :key="group.id" :value="group.id">
          {{ group.name }}
        </option>
      </select>
    </div>
     <div class="form-group">
      <label>Priority</label>
    <select v-model="priority" class="form-control">
    <option value="Low">Low</option>
    <option value="High">High</option>
    <option value="Very High">Very High</option>
  </select>
    </div>  
    <button class="btn btn-outline-secondary">Save</button>
  </form>
</template>

<script lang="ts" >
import {ref,onMounted} from 'vue';
import axios from 'axios';
import {useRouter,useRoute} from "vue-router";
import {Todo} from "@/classes/todo";



export default {
  name: "TodoEdit",
  setup() {
     const title = ref(''); 
    const state = ref(''); 
    const content = ref(''); 
    const priority = ref(''); 
    const userId = ref(''); 
    const groupId = ref(''); 
    const dueDate = ref(''); 
    const groups = ref([]);
    const router = useRouter();
    const {params} = useRoute();


   onMounted(async () => {
      userId.value= await (await axios.get('user')).data.id
      const todoCall = await axios.get(`todo/${params.id}`);
      const todo: Todo = todoCall.data;
      title.value = todo.title;
      userId.value = todo.userId;
      groupId.value=todo.groupId;
      priority.value=todo.priority;
      content.value=todo.content;
       const response = await axios.get('groups/${userId.value}');
      groups.value = response.data;

    });

    const submit = async () => {
       await axios.put('todo', {
        title: title.value,
        content: content.value,
        priority: priority.value,
        groupId: groupId.value,
        id:params.id,
        userId:userId.value
      });
      await router.push('/NewTodos');
    }

     return {
      title,
      content,
      dueDate,
      groupId,
      submit,
      groups,
      priority,
      userId
    }
  }
}


</script>