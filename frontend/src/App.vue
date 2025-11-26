<template>
  <div id="app" class="min-h-screen bg-primary">
    <!-- Navigation Bar -->
    <nav
      v-if="!isAdminRoute"
      class="fixed top-0 left-0 right-0 z-50 shadow-md"
      style="background-color: #227082;"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-20">
          <!-- Logo -->
          <router-link to="/" class="flex items-center">
            <img 
              src="/src/assets/background/ccpp-logo.png" 
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
        <div v-show="mobileMenuOpen" class="md:hidden pb-4" style="background-color: #227082;">
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
    <main :class="isAdminRoute ? '' : 'pt-20'">
      <router-view />
    </main>

    <!-- Footer -->
    <footer v-if="!isAdminRoute" class="bg-gray-800 text-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-8">
          <!-- About Section -->
          <div>
            <h3 class="text-lg font-bold mb-4 text-white">About Us</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-4">
              A community of believers committed to teaching through the Bible systematically, 
              verse by verse, chapter by chapter, book by book.
            </p>
            <div class="flex items-center space-x-4 mt-4">
              <a
                href="https://www.facebook.com"
                target="_blank"
                rel="noopener noreferrer"
                class="text-white/80 hover:text-white transition-colors duration-200"
                aria-label="Facebook"
              >
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                </svg>
              </a>
              <a
                href="https://www.instagram.com"
                target="_blank"
                rel="noopener noreferrer"
                class="text-white/80 hover:text-white transition-colors duration-200"
                aria-label="Instagram"
              >
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
                </svg>
              </a>
              <a
                href="https://www.youtube.com"
                target="_blank"
                rel="noopener noreferrer"
                class="text-white/80 hover:text-white transition-colors duration-200"
                aria-label="YouTube"
              >
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>
                </svg>
              </a>
            </div>
          </div>

          <!-- Quick Links -->
          <div>
            <h3 class="text-lg font-bold mb-4 text-white">Quick Links</h3>
            <ul class="space-y-2">
              <li>
                <router-link to="/about" class="text-white/80 hover:text-white transition-colors duration-200 text-sm">
                  About Us
                </router-link>
              </li>
              <li>
                <router-link to="/about/leadership" class="text-white/80 hover:text-white transition-colors duration-200 text-sm">
                  Leadership
                </router-link>
              </li>
              <li>
                <router-link to="/ministries" class="text-white/80 hover:text-white transition-colors duration-200 text-sm">
                  Ministries
                </router-link>
              </li>
              <li>
                <router-link to="/events" class="text-white/80 hover:text-white transition-colors duration-200 text-sm">
                  Events
                </router-link>
              </li>
              <li>
                <router-link to="/contact" class="text-white/80 hover:text-white transition-colors duration-200 text-sm">
                  Contact
                </router-link>
              </li>
            </ul>
          </div>

          <!-- Service Times -->
          <div>
            <h3 class="text-lg font-bold mb-4 text-white">Service Times</h3>
            <ul class="space-y-2 text-sm">
              <li class="text-white/80">
                <span class="font-medium text-white">Sunday Worship:</span><br>
                <span>9:00 AM</span>
              </li>
              <li class="text-white/80">
                <span class="font-medium text-white">Wednesday Bible Study:</span><br>
                <span>7:00 PM</span>
              </li>
            </ul>
          </div>

          <!-- Contact Info -->
          <div>
            <h3 class="text-lg font-bold mb-4 text-white">Contact Us</h3>
            <ul class="space-y-3 text-sm text-white/80">
              <li class="flex items-start">
                <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <span>123 Faith Street<br>Hope City, HC 12345<br>United States</span>
              </li>
              <li class="flex items-center">
                <svg class="w-5 h-5 mr-2 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                </svg>
                <a href="tel:+15551234567" class="hover:text-white transition-colors duration-200">(555) 123-4567</a>
              </li>
              <li class="flex items-center">
                <svg class="w-5 h-5 mr-2 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                <a href="mailto:info@ccpp.org" class="hover:text-white transition-colors duration-200">info@ccpp.org</a>
              </li>
            </ul>
          </div>
        </div>

        <!-- Copyright -->
        <div class="border-t border-gray-700 pt-6 mt-8">
          <p class="text-center text-white/70 text-sm">
            &copy; {{ new Date().getFullYear() }} Calvary Chapel Phnom Penh. All rights reserved.
          </p>
        </div>
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
  computed: {
    isAdminRoute() {
      return this.$route.path.startsWith('/admin')
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
