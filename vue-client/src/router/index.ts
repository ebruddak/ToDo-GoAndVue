import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';

 import NewNotes from '@/pages/NewNotes.vue';
 import RegisterUser from '@/pages/RegisterUser.vue';
 import LoginUser from '@/pages/LoginUser.vue';
 import AdminWrapper from '@/pages/AdminWrapper.vue';


const routes: Array<RouteRecordRaw> = [
    {path: '/Login', component: LoginUser},
     {path: '/Register', component: RegisterUser},
     {
        path:'',
        component: AdminWrapper,
        children:[
            {path: '', component: NewNotes},
            {path: '/NewNotes', component: NewNotes},
        ]
     }

]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router