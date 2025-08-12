import { createRouter, createWebHistory } from 'vue-router'
import Login from './Login.vue'
import Chat from './Chat.vue'

const routes = [
  { path: '/', component: Login },     // početna stranica (login/registracija)
  { path: '/chat', component: Chat } // nova chat stranica
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
