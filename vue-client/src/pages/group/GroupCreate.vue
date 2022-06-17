<template>
  <form @submit.prevent="submit">
    <div class="form-group mt-4">
      <label>Todo Group Name</label>
      <input type="text" class="form-control" name="group_name" v-model="name"/>
    </div>
   
    <button class="btn btn-outline-secondary">Save</button>
  </form>
</template>

<script lang="ts" >
import {ref,onMounted} from 'vue';
import axios from 'axios';
import {useRouter} from "vue-router";
export default {
  name: "GroupCreate",
  setup() {
    const name = ref(''); 
    const router = useRouter();
    
    const userId = ref(''); 
    onMounted(async () => {
      const {data} =await axios.get('user')
      console.log(data)
       userId.value=data.id ;   
     });
    const submit = async () => {
      await axios.post('group', {
        name: name.value,
        userId: userId.value
      });
      await router.push('/Groups');
    }
    return {
      name,
      submit,userId
    }
  }
}
</script>