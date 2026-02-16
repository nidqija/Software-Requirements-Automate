import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import ClassDiagramView from '../views/ClassDiagramView.vue'
import RecentView from '../views/RecentView.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: HomeView,
    },
    {
        path: '/about',
        name: 'About',
        component: AboutView,
    },

    {
        path : '/recent',
        name : 'Recent',
        component : RecentView,
    } ,

    {
        path: '/class-diagram',
        name: 'ClassDiagram',
        component: ClassDiagramView,
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
})

export default router
