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
import {onMounted} from 'vue';
import NavBar from '@/components/NavBar.vue';
import MenuBar from '@/components/MenuBar.vue';

import axios from 'axios';
// import * as c3 from 'c3';
export default {
  name: "AdminWrapper",
  components:{
    NavBar,MenuBar
  },
  setup() {
     onMounted(async () => {
    //   const chart = c3.generate({
    //     bindto: '#chart',
    //     data: {
    //       x: 'x',
    //       columns: [
    //         ['x'],
    //         ['Sales']
    //       ],
    //       types: {
    //         Sales: 'bar'
    //       }
    //     },
    //     axis: {
    //       x: {
    //         type: 'timeseries',
    //         tick: {
    //           format: '%Y-%m-%d'
    //         }
    //       }
    //     }
    //   });
    //   const response = await axios.get('chart');
    //   const records = response.data.data;
    //   chart.load({
    //     columns: [
    //       ['x', ...records.map(r => r.date)],
    //       ['Sales', ...records.map(r => r.sum)]
    //     ]
    //   });
     });
  }
}
</script>