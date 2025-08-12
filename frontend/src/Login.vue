<template>
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8 bg-gray-900 text-white">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="mt-10 text-center text-2xl font-bold tracking-tight text-white">
        {{ showRegister ? 'Registriraj se' : 'Prijavi se' }}
      </h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <!-- Toggle buttons -->
      <div class="flex justify-center mb-6 space-x-4">
        <button
          @click="showRegister = false"
          :class="showRegister ? 'text-gray-400' : 'text-indigo-400 font-semibold'"
          class="px-4 py-1 rounded"
        >
          Prijava
        </button>
        <button
          @click="showRegister = true"
          :class="showRegister ? 'text-indigo-400 font-semibold' : 'text-gray-400'"
          class="px-4 py-1 rounded"
        >
          Registracija
        </button>
      </div>

      <!-- Login form -->
      <form v-if="!showRegister" class="space-y-6" @submit.prevent="handleLogin">
        <div>
          <label for="email" class="block text-sm font-medium text-gray-100">Email address</label>
          <div class="mt-2">
            <input
              id="email"
              type="email"
              v-model="loginForm.email"
              autocomplete="email"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>

        <div>
          <label for="password" class="block text-sm font-medium text-gray-100">Password</label>
          <div class="mt-2">
            <input
              id="password"
              type="password"
              v-model="loginForm.password"
              autocomplete="current-password"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-indigo-500 px-3 py-1.5 text-sm font-semibold text-white hover:bg-indigo-400 focus:outline-indigo-500"
          >
            Prijavi se
          </button>
        </div>
      </form>

      <!-- Register form -->
      <form v-if="showRegister" class="space-y-6" @submit.prevent="handleRegister">
        <div>
          <label for="name" class="block text-sm font-medium text-gray-100">Ime</label>
          <div class="mt-2">
            <input
              id="name"
              type="text"
              v-model="registerForm.name"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>
        <div>
          <label for="lastname" class="block text-sm font-medium text-gray-100">Prezime</label>
          <div class="mt-2">
            <input
              id="lastname"
              type="text"
              v-model="registerForm.lastname"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>
        <div>
          <label for="phone" class="block text-sm font-medium text-gray-100">Telefon</label>
          <div class="mt-2">
            <input
              id="phone"
              type="tel"
              v-model="registerForm.phone"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>
        <div>
          <label for="emailReg" class="block text-sm font-medium text-gray-100">Email</label>
          <div class="mt-2">
            <input
              id="emailReg"
              type="email"
              v-model="registerForm.email"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>
        <div>
          <label for="passwordReg" class="block text-sm font-medium text-gray-100">Lozinka</label>
          <div class="mt-2">
            <input
              id="passwordReg"
              type="password"
              v-model="registerForm.password"
              required
              class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white placeholder:text-gray-500 focus:outline-indigo-500"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-indigo-500 px-3 py-1.5 text-sm font-semibold text-white hover:bg-indigo-400 focus:outline-indigo-500"
          >
            Registriraj se
          </button>
        </div>
        
      </form>

      <!-- Poruke -->
      <p class="mt-4 text-center text-sm text-red-400" v-if="showRegister">{{ registerMessage }}</p>
      <p class="mt-4 text-center text-sm text-red-400" v-if="!showRegister">{{ loginMessage }}</p>
    </div>
  </div>
</template>

<script>
import { registerUser, loginUser } from './api.js'

export default {
  data() {
    return {
      showRegister: false,
      registerForm: {
        name: '',
        lastname: '',
        phone: '',
        email: '',
        password: '',
      },
      loginForm: {
        email: '',
        password: '',
      },
      registerMessage: '',
      loginMessage: '',
    }
  },
  methods: {
    async handleRegister() {
      try {
        await registerUser(this.registerForm)
        this.registerMessage = 'Registracija uspješna!'

        // redirect na chat stranicu
        this.$router.push('/chat')
      } catch (error) {
        this.registerMessage = 'Greška kod registracije: ' + (error.response?.data?.error || error.message)
      }
    },
    async handleLogin() {
      try {
        const response = await loginUser(this.loginForm)
        this.loginMessage = 'Prijava uspješna! Dobrodošli, ' + response.data.user.name

        // redirect na chat stranicu
        this.$router.push('/chat').then(() => {
  console.log('Navigacija na /chat uspješna');
}).catch(err => {
  console.error('Navigacija nije uspjela:', err);
});

      } catch (error) {
        this.loginMessage = 'Greška kod prijave: ' + (error.response?.data?.error || error.message)
      }
    },
  },
}
</script>