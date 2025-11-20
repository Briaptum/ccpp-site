<template>
  <div class="min-h-screen bg-main flex items-center justify-center px-4">
    <div class="max-w-md w-full bg-white rounded-lg shadow-xl p-8">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Admin Login</h1>
        <p class="text-gray-600">Calvary Chapel Phnom Penh</p>
      </div>
      
      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
            Email
          </label>
          <input
            id="email"
            v-model="email"
            type="email"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-transparent"
            placeholder="admin@example.com"
          />
        </div>
        
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
            Password
          </label>
          <input
            id="password"
            v-model="password"
            type="password"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-transparent"
            placeholder="Enter your password"
          />
        </div>
        
        <button
          type="submit"
          class="w-full bg-main text-white py-3 rounded-lg font-semibold hover:opacity-90 transition-colors shadow-lg"
        >
          Sign In
        </button>
      </form>
      
      <div class="mt-6 text-center">
        <router-link to="/" class="text-sm text-white hover:opacity-80">
          ← Back to Website
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'AdminLogin',
  setup() {
    const router = useRouter()
    const email = ref('')
    const password = ref('')
    
    const handleLogin = () => {
      // TODO: Implement login logic
      console.log('Login attempt:', email.value)
      // For now, just redirect to dashboard
      router.push('/admin/dashboard')
    }
    
    return {
      email,
      password,
      handleLogin
    }
  }
}
</script>

