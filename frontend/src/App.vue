<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <nav class="sticky top-0 z-50 shadow-md border-b" style="background-color: #265267;">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-24 md:h-32">
          <!-- Logo -->
          <div class="flex items-center">
            <router-link to="/" class="flex items-center">
              <img 
                src="/src/assets/images/logo3.png" 
                alt="Calvary Chapel Phnom Penh Logo" 
                class="h-20 md:h-20 w-auto object-contain"
              />
            </router-link>
          </div>
          
          <!-- Desktop Navigation -->
          <div class="hidden lg:flex items-center space-x-8">
            <router-link
              to="/"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'Home' }"
            >
              Home
            </router-link>
            <router-link
              to="/about"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'About' }"
            >
              About Us
            </router-link>
            <router-link
              to="/ministries"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'Ministries' }"
            >
              Ministries
            </router-link>
            <router-link
              to="/events"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'Events' }"
            >
              Events
            </router-link>
            <router-link
              to="/resources"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'Resources' }"
            >
              Resources
            </router-link>
            <router-link
              to="/contact"
              class="text-white hover:text-gray-200 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-gray-200 font-semibold': $route.name === 'Contact' }"
            >
              Contact
            </router-link>
          </div>

          <!-- Mobile menu button -->
          <div class="flex items-center lg:hidden">
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="inline-flex items-center justify-center p-2 rounded-md text-white hover:text-gray-200 hover:bg-opacity-20 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white transition-colors"
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
        <div class="px-2 pt-2 pb-3 space-y-1 border-t" style="background-color: #265267;">
          <router-link
            to="/"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Home' }"
          >
            Home
          </router-link>
          <router-link
            to="/about"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'About' }"
          >
            About Us
          </router-link>
          <router-link
            to="/ministries"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Ministries' }"
          >
            Ministries
          </router-link>
          <router-link
            to="/events"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Events' }"
          >
            Events
          </router-link>
          <router-link
            to="/resources"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Resources' }"
          >
            Resources
          </router-link>
          <router-link
            to="/contact"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Contact' }"
          >
            Contact
          </router-link>
          <router-link
            to="/users"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-white hover:text-gray-200 hover:bg-opacity-20 transition-colors"
            :class="{ 'text-gray-200 font-semibold bg-opacity-20': $route.name === 'Users' }"
          >
            Admin
          </router-link>
        </div>
      </div>
    </nav>
    
    <main :class="{ 'max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8': $route.name !== 'Home' }">
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
