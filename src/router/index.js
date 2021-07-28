import {createRouter, createWebHashHistory} from 'vue-router'
import Home from '@/views/Home.vue'
import Proposal from '@/views/Proposal.vue'
import Create from '@/views/Create.vue'
import About from '@/views/About.vue'
import NotFound from '@/views/NotFound.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/proposal/:id',
    name: 'Proposal',
    component: Proposal
  },
  {
    path: '/create',
    name: 'Create',
    component: Create
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },

  {
    path: '/404',
    alias: '/:catchAll(.*)',
    component: NotFound
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach(() => {
  window.scrollTo(0, 0)
})

export default router
