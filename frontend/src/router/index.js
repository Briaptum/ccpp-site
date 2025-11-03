import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Ministries from '@/views/Ministries.vue'
import Events from '@/views/Events.vue'
import Resources from '@/views/Resources.vue'
import Contact from '@/views/Contact.vue'
import Users from '@/views/Users.vue'
import Missionary from '@/views/Missionary.vue'
import Donate from '@/views/Donate.vue'
import HowToHelp from '@/views/HowToHelp.vue'
import News from '@/views/News.vue'

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
    path: '/ministries',
    name: 'Ministries',
    component: Ministries
  },
  {
    path: '/events',
    name: 'Events',
    component: Events
  },
  {
    path: '/resources',
    name: 'Resources',
    component: Resources
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact
  },
  {
    path: '/users',
    name: 'Users',
    component: Users
  },
  {
    path: '/missionary',
    name: 'Missionary',
    component: Missionary
  },
  {
    path: '/donate',
    name: 'Donate',
    component: Donate
  },
  {
    path: '/how-to-help',
    name: 'HowToHelp',
    component: HowToHelp
  },
  {
    path: '/news',
    name: 'News',
    component: News
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
