<template>
  <div id="app" class="min-h-screen bg-primary">
    <!-- Navigation Bar -->
    <nav
      class="fixed top-0 left-0 right-0 z-50 shadow-md"
      style="background-color: #121010;"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-20">
          <!-- Logo -->
          <router-link to="/" class="flex items-center">
            <img 
              src="/src/assets/background/logo3.png" 
              alt="Calvary Chapel Phnom Penh Logo" 
              class="h-14 w-auto"
            />
          </router-link>
          
          <!-- Navigation Links -->
          <div class="hidden md:flex items-center space-x-10">
            <!-- About Dropdown -->
            <div 
              class="relative"
              @mouseenter="aboutDropdownOpen = true"
              @mouseleave="aboutDropdownOpen = false"
            >
              <router-link
                to="/about"
                class="text-white hover:text-custom-orange transition-colors text-lg font-medium flex items-center"
              >
                About
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="aboutDropdownOpen"
                class="absolute top-full left-0 pt-2 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                  <router-link
                    to="/about/leadership"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="aboutDropdownOpen = false"
                  >
                    Leadership
                  </router-link>
                  <router-link
                    to="/about/what-we-believe"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="aboutDropdownOpen = false"
                  >
                    What we believe
                  </router-link>
                  <router-link
                    to="/about/calvary-chapel"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="aboutDropdownOpen = false"
                  >
                    About Calvary Chapel
                  </router-link>
                  <router-link
                    to="/about/cambodia"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="aboutDropdownOpen = false"
                  >
                    Cambodia
                  </router-link>
                </div>
              </div>
            </div>
            <!-- Ministries Dropdown -->
            <div 
              class="relative"
              @mouseenter="ministriesDropdownOpen = true"
              @mouseleave="ministriesDropdownOpen = false"
            >
              <router-link
                to="/ministries"
                class="text-white hover:text-custom-orange transition-colors text-lg font-medium flex items-center"
              >
                Ministries
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="ministriesDropdownOpen"
                class="absolute top-full left-0 pt-2 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                  <router-link
                    to="/ministries/outreaches"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="ministriesDropdownOpen = false"
                  >
                    Outreaches
                  </router-link>
                  <router-link
                    to="/ministries/grace-church"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="ministriesDropdownOpen = false"
                  >
                    Grace Church
                  </router-link>
                  <router-link
                    to="/ministries/youth-ministry"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="ministriesDropdownOpen = false"
                  >
                    Youth Ministry
                  </router-link>
                </div>
              </div>
            </div>
            <!-- Events Dropdown -->
            <div 
              class="relative"
              @mouseenter="eventsDropdownOpen = true"
              @mouseleave="eventsDropdownOpen = false"
            >
              <router-link
                to="/events"
                class="text-white hover:text-custom-orange transition-colors text-lg font-medium flex items-center"
              >
                Events
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="eventsDropdownOpen"
                class="absolute top-full left-0 pt-2 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                  <router-link
                    to="/events/upcoming-events"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="eventsDropdownOpen = false"
                  >
                    Upcoming Events
                  </router-link>
                  <router-link
                    to="/events/calendar"
                    class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                    @click="eventsDropdownOpen = false"
                  >
                    Calendar
                  </router-link>
                </div>
              </div>
            </div>
            <router-link
              to="/contact"
              class="text-white hover:text-custom-orange transition-colors text-lg font-medium"
            >
              Contact
            </router-link>
          </div>

          <!-- Mobile Menu Button -->
          <button
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="md:hidden text-white hover:text-custom-orange transition-colors"
          >
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Mobile Menu -->
        <div v-show="mobileMenuOpen" class="md:hidden pb-4" style="background-color: #121010;">
          <div class="flex flex-col space-y-2">
            <!-- About Dropdown Mobile -->
            <div>
              <button
                @click="aboutDropdownMobileOpen = !aboutDropdownMobileOpen"
                class="w-full text-left text-white hover:text-custom-orange transition-colors py-2 flex items-center justify-between"
              >
                About
                <svg 
                  class="h-4 w-4 transition-transform"
                  :class="{ 'rotate-180': aboutDropdownMobileOpen }"
                  fill="none" 
                  stroke="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
              <div 
                v-show="aboutDropdownMobileOpen"
                class="pl-4 mt-1 space-y-1"
              >
                <router-link
                  to="/about/leadership"
                  @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Leadership
                </router-link>
                <router-link
                  to="/about/what-we-believe"
                  @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  What we believe
                </router-link>
                <router-link
                  to="/about/calvary-chapel"
                  @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  About Calvary Chapel
                </router-link>
                <router-link
                  to="/about/cambodia"
                  @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Cambodia
                </router-link>
              </div>
            </div>
            <!-- Ministries Dropdown Mobile -->
            <div>
              <button
                @click="ministriesDropdownMobileOpen = !ministriesDropdownMobileOpen"
                class="w-full text-left text-white hover:text-custom-orange transition-colors py-2 flex items-center justify-between"
              >
                Ministries
                <svg 
                  class="h-4 w-4 transition-transform"
                  :class="{ 'rotate-180': ministriesDropdownMobileOpen }"
                  fill="none" 
                  stroke="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
              <div 
                v-show="ministriesDropdownMobileOpen"
                class="pl-4 mt-1 space-y-1"
              >
                <router-link
                  to="/ministries/outreaches"
                  @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Outreaches
                </router-link>
                <router-link
                  to="/ministries/grace-church"
                  @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Grace Church
                </router-link>
                <router-link
                  to="/ministries/youth-ministry"
                  @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Youth Ministry
                </router-link>
              </div>
            </div>
            <!-- Events Dropdown Mobile -->
            <div>
              <button
                @click="eventsDropdownMobileOpen = !eventsDropdownMobileOpen"
                class="w-full text-left text-white hover:text-custom-orange transition-colors py-2 flex items-center justify-between"
              >
                Events
                <svg 
                  class="h-4 w-4 transition-transform"
                  :class="{ 'rotate-180': eventsDropdownMobileOpen }"
                  fill="none" 
                  stroke="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
              <div 
                v-show="eventsDropdownMobileOpen"
                class="pl-4 mt-1 space-y-1"
              >
                <router-link
                  to="/events/upcoming-events"
                  @click="mobileMenuOpen = false; eventsDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Upcoming Events
                </router-link>
                <router-link
                  to="/events/calendar"
                  @click="mobileMenuOpen = false; eventsDropdownMobileOpen = false"
                  class="block text-white hover:text-custom-orange transition-colors py-2"
                >
                  Calendar
                </router-link>
              </div>
            </div>
            <router-link
              to="/contact"
              @click="mobileMenuOpen = false"
              class="text-white hover:text-custom-orange transition-colors py-2"
            >
              Contact
            </router-link>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="pt-20">
      <router-view />
    </main>

    <!-- Footer -->
    <footer class="bg-gradient-to-b from-custom-orange to-white text-dark py-8">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <p class="text-dark/90">&copy; {{ new Date().getFullYear() }} Calvary Chapel Phnom Penh</p>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      mobileMenuOpen: false,
      aboutDropdownOpen: false,
      aboutDropdownMobileOpen: false,
      ministriesDropdownOpen: false,
      ministriesDropdownMobileOpen: false,
      eventsDropdownOpen: false,
      eventsDropdownMobileOpen: false
    }
  },
  watch: {
    $route() {
      this.mobileMenuOpen = false
      this.aboutDropdownOpen = false
      this.aboutDropdownMobileOpen = false
      this.ministriesDropdownOpen = false
      this.ministriesDropdownMobileOpen = false
      this.eventsDropdownOpen = false
      this.eventsDropdownMobileOpen = false
    }
  }
}
</script>
