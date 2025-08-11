<template>
  <div class="max-w-md mx-auto mt-10 p-6 border rounded shadow">
    <div class="flex justify-center mb-6 space-x-4">
      <button 
        :class="activeForm === 'login' ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'" 
        class="px-4 py-2 rounded transition"
        @click="activeForm = 'login'"
      >
        Prijava
      </button>
      <button 
        :class="activeForm === 'register' ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'" 
        class="px-4 py-2 rounded transition"
        @click="activeForm = 'register'"
      >
        Registracija
      </button>
    </div>

    <div>
      <form v-if="activeForm === 'login'" @submit.prevent="handleLogin" class="space-y-4">
        <input v-model="loginForm.email" type="email" placeholder="Email" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <input v-model="loginForm.password" type="password" placeholder="Lozinka" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <button type="submit" class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition">
          Prijavi se
        </button>
        <p class="text-red-600 mt-2">{{ loginMessage }}</p>
      </form>

      <form v-if="activeForm === 'register'" @submit.prevent="handleRegister" class="space-y-4">
        <input v-model="registerForm.name" type="text" placeholder="Ime" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <input v-model="registerForm.lastname" type="text" placeholder="Prezime" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <input v-model="registerForm.phone" type="text" placeholder="Telefon" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <input v-model="registerForm.email" type="email" placeholder="Email" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <input v-model="registerForm.password" type="password" placeholder="Lozinka" required
          class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <button type="submit" class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition">
          Registriraj se
        </button>
        <p class="text-red-600 mt-2">{{ registerMessage }}</p>
      </form>
    </div>
  </div>
</template>

<script>
import { registerUser, loginUser } from './api.js'

export default {
  data() {
    return {
      activeForm: 'login',
      registerForm: {
        name: '',
        lastname: '',
        phone: '',
        email: '',
        password: ''
      },
      loginForm: {
        email: '',
        password: ''
      },
      registerMessage: '',
      loginMessage: ''
    }
  },
  methods: {
    async handleRegister() {
      this.registerMessage = ''
      try {
        await registerUser(this.registerForm)
        this.registerMessage = 'Registracija uspješna!'
        // Očisti formu nakon registracije
        this.registerForm = { name: '', lastname: '', phone: '', email: '', password: '' }
      } catch (error) {
        this.registerMessage = 'Greška kod registracije: ' + (error.response?.data?.error || error.message)
      }
    },
    async handleLogin() {
      this.loginMessage = ''
      try {
        const response = await loginUser(this.loginForm)
        this.loginMessage = 'Prijava uspješna! Dobrodošli, ' + response.data.user.name
        // Očisti formu nakon prijave
        this.loginForm = { email: '', password: '' }
      } catch (error) {
        this.loginMessage = 'Greška kod prijave: ' + (error.response?.data?.error || error.message)
      }
    }
  }
}
</script>


