<template>
  <NavBar /> 
 <div class="container">
  <div class="row">
  <div class="col-md-12">
   <div class="col-md-3 col-lg-2">
      <MenuBar />
   </div>
   <div class="col-md-8 col-lg-9 px-md-4">
    <main >
     <router-view />
    </main>
    </div>
  </div>
   
  </div>
</div> 

</template>

<script lang="ts" >
import {onMounted, TrackOpTypes} from 'vue';
import NavBar from '@/components/NavBar.vue';
import MenuBar from '@/components/MenuBar.vue';
import {useRouter} from "vue-router"
import axios from 'axios';
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
      } catch (e){
        await router.push('/login')
      }
     
    
     });
  }
}
</script>