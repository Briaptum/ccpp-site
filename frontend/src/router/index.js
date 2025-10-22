import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Missionary from '@/views/Missionary.vue'
import Contact from '@/views/Contact.vue'
import Donate from '@/views/Donate.vue'
import Users from '@/views/Users.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/missionary',
    name: 'Missionary',
    component: Missionary
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact
  },
  {
    path: '/donate',
    name: 'Donate',
    component: Donate
  },
  {
    path: '/users',
    name: 'Users',
    component: Users
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
