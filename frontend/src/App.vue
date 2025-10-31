<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <nav class="sticky top-0 z-50 bg-white shadow-md border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-24 md:h-32">
          <!-- Logo -->
          <div class="flex items-center">
            <router-link to="/" class="flex items-center">
              <img 
                src="/src/assets/images/ccpp-logo.png" 
                alt="Calvary Chapel Phnom Penh Logo" 
                class="h-20 md:h-20 w-auto object-contain"
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
              to="/news"
              class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
              :class="{ 'text-primary-600': $route.name === 'News' }"
            >
              News
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
            
            <!-- How to Help Dropdown -->
            <div class="relative" @mouseenter="helpDropdownOpen = true" @mouseleave="helpDropdownOpen = false">
              <div class="flex items-center">
                <router-link
                  to="/how-to-help"
                  class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-base font-medium transition-colors"
                  :class="{ 'text-primary-600': $route.name === 'HowToHelp' }"
                >
                  How to Help
                </router-link>
                <button
                  class="text-gray-700 hover:text-primary-600 px-1 py-2 rounded-md transition-colors"
                  @click="helpDropdownOpen = !helpDropdownOpen"
                >
                  <svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>
              </div>
              <div
                v-show="helpDropdownOpen"
                class="absolute left-0 mt-0 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-50"
              >
                <div class="py-1">
                  <router-link
                    to="/how-to-help#prayer"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-primary-600"
                  >
                    Prayer
                  </router-link>
                  <router-link
                    to="/how-to-help#donation"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-primary-600"
                  >
                    Donation
                  </router-link>
                  <router-link
                    to="/how-to-help#needs-list"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-primary-600"
                  >
                    Needs List
                  </router-link>
                  <router-link
                    to="/how-to-help#testimonial"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-primary-600"
                  >
                    Testimonial
                  </router-link>
                  <router-link
                    to="/how-to-help#service"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-primary-600"
                  >
                    Service
                  </router-link>
                </div>
              </div>
            </div>
            
            <router-link
              to="/donate"
              class="text-gray-700 hover:text-primary-700 px-3 py-2 rounded-md text-base font-medium transition-colors bg-primary-500 hover:bg-primary-600 text-white hover:text-white"
              :class="{ 'bg-primary-700': $route.name === 'Donate' }"
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
            to="/news"
            @click="mobileMenuOpen = false"
            class="block px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
            :class="{ 'text-primary-600 bg-primary-50': $route.name === 'News' }"
          >
            News
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
          
          <!-- Mobile How to Help Dropdown -->
          <div>
            <div class="flex items-center">
              <router-link
                to="/how-to-help"
                @click="mobileMenuOpen = false"
                class="flex-1 px-3 py-3 rounded-md text-base font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
                :class="{ 'text-primary-600 bg-primary-50': $route.name === 'HowToHelp' }"
              >
                How to Help
              </router-link>
              <button
                @click="mobileHelpDropdownOpen = !mobileHelpDropdownOpen"
                class="px-3 py-3 rounded-md text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                <svg 
                  class="h-5 w-5 transition-transform"
                  :class="{ 'rotate-180': mobileHelpDropdownOpen }"
                  xmlns="http://www.w3.org/2000/svg" 
                  fill="none" 
                  viewBox="0 0 24 24" 
                  stroke="currentColor"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </button>
            </div>
            <div v-show="mobileHelpDropdownOpen" class="pl-4 space-y-1">
              <router-link
                to="/how-to-help#prayer"
                @click="mobileMenuOpen = false; mobileHelpDropdownOpen = false"
                class="block px-3 py-2 rounded-md text-sm text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                Prayer
              </router-link>
              <router-link
                to="/how-to-help#donation"
                @click="mobileMenuOpen = false; mobileHelpDropdownOpen = false"
                class="block px-3 py-2 rounded-md text-sm text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                Donation
              </router-link>
              <router-link
                to="/how-to-help#needs-list"
                @click="mobileMenuOpen = false; mobileHelpDropdownOpen = false"
                class="block px-3 py-2 rounded-md text-sm text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                Needs List
              </router-link>
              <router-link
                to="/how-to-help#testimonial"
                @click="mobileMenuOpen = false; mobileHelpDropdownOpen = false"
                class="block px-3 py-2 rounded-md text-sm text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                Testimonial
              </router-link>
              <router-link
                to="/how-to-help#service"
                @click="mobileMenuOpen = false; mobileHelpDropdownOpen = false"
                class="block px-3 py-2 rounded-md text-sm text-gray-600 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              >
                Service
              </router-link>
            </div>
          </div>
          
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
      mobileMenuOpen: false,
      helpDropdownOpen: false,
      mobileHelpDropdownOpen: false
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
