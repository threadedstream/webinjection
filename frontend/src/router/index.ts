import { createRouter, createWebHistory } from 'vue-router'
import App from '../App.vue'
import SqlI from '@/components/SqlI.vue'
import Home from '@/components/Home.vue'
import Tasks from '@/components/Tasks.vue'


const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/sqli',
        name: 'SqlI',
        component: SqlI
    },
    {
        path: '/tasks',
        name: 'Tasks',
        component: Tasks
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router