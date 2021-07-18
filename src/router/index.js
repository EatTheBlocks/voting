import {createRouter, createWebHistory} from 'vue-router'
import Home from '@/views/Home.vue'
import Proposal from '@/views/Proposal.vue'
import NewProposal from '@/views/NewProposal.vue'
import About from '@/views/About.vue'

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
    name: 'NewProposal',
    component: NewProposal
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach(() => {
  window.scrollTo(0, 0)
})

export default router
