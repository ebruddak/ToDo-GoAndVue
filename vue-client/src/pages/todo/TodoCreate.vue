<template>
 <h5 class="pr-2 mt-4"> New Todo Item </h5> 
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
    <div class="form-group mt-4">
      <label> Due Date</label>
      <input type="date" class="form-control" name="dueDate" v-model="dueDate"/>
    </div>
    <button class="btn btn-outline-secondary">Save</button>
  </form>
</template>

<script lang="ts" >
import {ref,onMounted} from 'vue';
import axios from 'axios';
import {useRouter} from "vue-router";
import moment from 'moment'

export default {
  name: "TodoCreate",
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
   
    onMounted(async () => {
        debugger
      const response = await axios.get('groups');
      groups.value = response.data;
    });

debugger
    const submit = async () => {
      await axios.post('todo', {
        
        title: title.value,
        content: content.value,
        priority: priority.value,
        state: true,
        groupId: groupId.value,
        userId: ''
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
      priority
    }
  }
}
</script>