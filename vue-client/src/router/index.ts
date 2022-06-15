import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';

import Dashboard from '@/secure/dashboard/Dashboard.vue';


const routes: Array<RouteRecordRaw> = [
    {path: '', component: Dashboard},
    {path: '/dashboard', component: Dashboard}

]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router