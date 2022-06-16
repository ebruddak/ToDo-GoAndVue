import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';

 import NewNotes from '@/pages/NewNotes.vue';
 import RegisterUser from '@/pages/RegisterUser.vue';
 import LoginUser from '@/pages/LoginUser.vue';
 import AdminWrapper from '@/pages/AdminWrapper.vue';
 import TodoGroups from '@/pages/group/TodoGroups.vue';
 import GroupCreate from '@/pages/group/GroupCreate.vue';
 import GroupEdit from '@/pages/group/GroupEdit.vue';


const routes: Array<RouteRecordRaw> = [
    {path: '/Login', component: LoginUser},
     {path: '/Register', component: RegisterUser},
     {
        path:'',
        component: AdminWrapper,
        children:[
            {path: '', component: NewNotes},
            {path: '/NewNotes', component: NewNotes},
            {path: '/Groups', component: TodoGroups},
            {path: '/Groups/Create', component: GroupCreate},
            {path: '/Groups/:id/edit', component: GroupEdit},
        ]
     }

]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router