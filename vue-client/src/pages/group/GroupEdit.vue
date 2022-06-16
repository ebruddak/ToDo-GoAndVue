<template>
  <form @submit.prevent="submit">
    <div class="form-group mt-4">
      <label>Edit Todo Group Name</label>
      <input type="text" class="form-control" name="group_name" v-model="name"/>
    </div>
   
    <button class="btn btn-outline-secondary">Save</button>
  </form>
</template>

<script lang="ts" >
import {ref,onMounted} from 'vue';
import axios from 'axios';
import {useRoute, useRouter} from "vue-router";
import {Group} from "@/classes/group";

export default {
  name: "GroupCreate",
  setup() {
    const name = ref(''); 
    const userId = ref(''); 

    const router = useRouter();
    const {params} = useRoute();



   onMounted(async () => {
     debugger
      const groupCall = await axios.get(`group/${params.id}`);
      const group: Group = groupCall.data;
      name.value = group.name;
      userId.value = group.userId;

    });

    const submit = async () => {
       await axios.put('group', {
        name: name.value,
        id:params.id
      });
      await router.push('/groups');
    }

    return {
      name,
      submit
    }
  }
}
</script>