<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-md border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-24 md:h-32">
          <!-- Logo -->
          <div class="flex items-center">
            <router-link to="/" class="flex items-center">
              <img 
                src="/src/assets/images/Group 1.png" 
                alt="Calvary Chapel Phnom Penh Logo" 
                class="h-40 md:h-44 w-auto object-contain"
              />
            </router-link>
          </div>
          
          <!-- Desktop Navigation -->
          <div class="hidden lg:flex items-center space-x-8">
            <router-link
              to="/"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-primary-600': $route.name === 'Home' }"
            >
              Home
            </router-link>
            <router-link
              to="/about"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-primary-600': $route.name === 'About' }"
            >
              About
            </router-link>
            <router-link
              to="/missionary"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-primary-600': $route.name === 'Missionary' }"
            >
              Missionary
            </router-link>
            <router-link
              to="/contact"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-primary-600': $route.name === 'Contact' }"
            >
              Contact
            </router-link>
            <router-link
              to="/donate"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors bg-blue-500 hover:bg-primary-600 text-white hover:text-white"
              :class="{ 'bg-blue-700': $route.name === 'Donate' }"
            >
              Donate
            </router-link>
          </div>

          <!-- Mobile menu button -->
          <div class="flex items-center lg:hidden">
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="inline-flex items-center justify-center p-2 rounded-md text-gray-700 hover:text-primary-600 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500 transition-colors"
              aria-expanded="false"
            >
              <span class="sr-only">Open main menu</span>
              <!-- Hamburger icon -->
              <svg
                v-if="!mobileMenuOpen"
                class="block h-8 w-8"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
              <!-- Close icon -->
              <svg
                v-else
                class="block h-8 w-8"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile menu -->
      <div v-show="mobileMenuOpen" class="lg:hidden">
        <div class="px-2 pt-2 pb-3 space-y-1 bg-white border-t">
          <router-link
            to="/"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'Home' }"
          >
            Home
          </router-link>
          <router-link
            to="/about"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'About' }"
          >
            About
          </router-link>
          <router-link
            to="/missionary"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'Missionary' }"
          >
            Missionary
          </router-link>
          <router-link
            to="/contact"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'Contact' }"
          >
            Contact
          </router-link>
          <router-link
            to="/donate"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'Donate' }"
          >
            Donate
          </router-link>
          <router-link
            to="/users"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'Users' }"
          >
            Admin
          </router-link>
        </div>
      </div>
    </nav>
    
    <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      mobileMenuOpen: false
    }
  },
  watch: {
    $route() {
      // Close mobile menu when route changes
      this.mobileMenuOpen = false
    }
  }
}
</script>
