<template>
  <NavBar /> 
 <div class="container-fluid">
  <div class="row">
    <MenuBar />
    <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
     <router-view />
    </main>
  </div>
</div> 

</template>

<script lang="ts" >
import {onMounted, TrackOpTypes} from 'vue';
import NavBar from '@/components/NavBar.vue';
import MenuBar from '@/components/MenuBar.vue';
import {useRouter} from "vue-router"
import axios from 'axios';
// import * as c3 from 'c3';
export default {
  name: "AdminWrapper",
  components:{
    NavBar,MenuBar
  },
  setup() {
    const router= useRouter();
     onMounted(async () => {
      try{
      const {data} =await axios.get('user')
      console.log(data)
      } catch (e){
        await router.push('/login')
      }
     
    
     });
  }
}
</script>