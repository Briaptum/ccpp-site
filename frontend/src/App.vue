<template>
  <div id="app" class="min-h-screen bg-background">
    <nav
      :class="[
        'sticky top-0 z-50 shadow-md border-b-4 border-primary-500 bg-main-700 transition-all duration-300',
        isScrolled ? 'py-1' : 'py-2'
      ]"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div
          :class="[
            'flex justify-between transition-all duration-300',
            isScrolled ? 'h-16 md:h-20' : 'h-20 md:h-24'
          ]"
        >
          <!-- Logo -->
          <div class="flex items-center">
            <router-link to="/" class="flex items-center">
              <img 
                src="/src/assets/images/logo3.png" 
                alt="Calvary Chapel Phnom Penh Logo" 
                :class="[
                  'w-auto object-contain transition-all duration-300',
                  isScrolled ? 'h-12 md:h-12' : 'h-16 md:h-16'
                ]"
              />
            </router-link>
          </div>
          
          <!-- Desktop Navigation -->
          <div class="hidden lg:flex items-center space-x-8">
            <router-link
              to="/"
              class="nav-link"
            >
              Home
            </router-link>
            <!-- About Us Dropdown -->
            <div 
              class="relative"
              @mouseenter="aboutDropdownOpen = true"
              @mouseleave="aboutDropdownOpen = false"
            >
              <router-link
                to="/about"
                class="nav-link flex items-center"
                :class="{ 'text-white': aboutDropdownOpen }"
              >
                About Us
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="aboutDropdownOpen"
                class="absolute top-full left-0 pt-1 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                <router-link
                  to="/about/who-we-are"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="aboutDropdownOpen = false"
                >
                  Who We Are
                </router-link>
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
                  What We Believe
                </router-link>
                <router-link
                  to="/about/calvary-chapel"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="aboutDropdownOpen = false"
                >
                  Calvary Chapel
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
                class="nav-link flex items-center"
                :class="{ 'text-white': ministriesDropdownOpen }"
              >
                Ministries
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="ministriesDropdownOpen"
                class="absolute top-full left-0 pt-1 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                <router-link
                  to="/ministries/youth-ministry"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="ministriesDropdownOpen = false"
                >
                  Youth Ministry
                </router-link>
                <router-link
                  to="/ministries/grace-church"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="ministriesDropdownOpen = false"
                >
                  Grace Church
                </router-link>
                <router-link
                  to="/ministries/outreaches"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="ministriesDropdownOpen = false"
                >
                  Outreaches
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
                class="nav-link flex items-center"
                :class="{ 'text-white': eventsDropdownOpen }"
              >
                Events
                <svg class="ml-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </router-link>
              <!-- Dropdown Menu -->
              <div 
                v-show="eventsDropdownOpen"
                class="absolute top-full left-0 pt-1 w-48 z-50"
              >
                <div class="bg-white rounded-md shadow-lg py-2 border border-gray-200">
                <router-link
                  to="/events/calendar"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="eventsDropdownOpen = false"
                >
                  Calendar
                </router-link>
                <router-link
                  to="/events/upcoming-events"
                  class="block px-4 py-2 text-gray-800 hover:bg-gray-100 transition-colors"
                  @click="eventsDropdownOpen = false"
                >
                  Upcoming Events
                </router-link>
                </div>
              </div>
            </div>
            <router-link
              to="/resources"
              class="nav-link"
            >
              Resources
            </router-link>
            <router-link
              to="/contact"
              class="nav-link"
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
            class="nav-link-mobile"
          >
            Home
          </router-link>
          <!-- About Us Dropdown Mobile -->
          <div>
            <div class="flex items-center gap-2">
              <router-link
                to="/about"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile flex-1 text-left"
              >
                About Us
              </router-link>
              <button
                @click="aboutDropdownMobileOpen = !aboutDropdownMobileOpen"
                class="p-2 text-white hover:text-gray-200 transition-colors"
                aria-label="Toggle About Us menu"
              >
                <svg 
                  class="h-5 w-5 transition-transform"
                  :class="{ 'rotate-180': aboutDropdownMobileOpen }"
                  fill="none" 
                  stroke="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
            </div>
            <div 
              v-show="aboutDropdownMobileOpen"
              class="pl-4 mt-1 space-y-1"
            >
              <router-link
                to="/about/who-we-are"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Who We Are
              </router-link>
              <router-link
                to="/about/leadership"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Leadership
              </router-link>
              <router-link
                to="/about/what-we-believe"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                What We Believe
              </router-link>
              <router-link
                to="/about/calvary-chapel"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Calvary Chapel
              </router-link>
              <router-link
                to="/about/cambodia"
                @click="mobileMenuOpen = false; aboutDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Cambodia
              </router-link>
            </div>
          </div>
          <!-- Ministries Dropdown Mobile -->
          <div>
            <button
              @click="ministriesDropdownMobileOpen = !ministriesDropdownMobileOpen"
              class="nav-link-mobile w-full text-left flex items-center justify-between"
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
                to="/ministries/youth-ministry"
                @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Youth Ministry
              </router-link>
              <router-link
                to="/ministries/grace-church"
                @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Grace Church
              </router-link>
              <router-link
                to="/ministries/outreaches"
                @click="mobileMenuOpen = false; ministriesDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Outreaches
              </router-link>
            </div>
          </div>
          <!-- Events Dropdown Mobile -->
          <div>
            <button
              @click="eventsDropdownMobileOpen = !eventsDropdownMobileOpen"
              class="nav-link-mobile w-full text-left flex items-center justify-between"
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
                to="/events/calendar"
                @click="mobileMenuOpen = false; eventsDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Calendar
              </router-link>
              <router-link
                to="/events/upcoming-events"
                @click="mobileMenuOpen = false; eventsDropdownMobileOpen = false"
                class="nav-link-mobile block"
              >
                Upcoming Events
              </router-link>
            </div>
          </div>
          <router-link
            to="/resources"
            @click="mobileMenuOpen = false"
            class="nav-link-mobile"
          >
            Resources
          </router-link>
          <router-link
            to="/contact"
            @click="mobileMenuOpen = false"
            class="nav-link-mobile"
          >
            Contact
          </router-link>
          <router-link
            to="/users"
            @click="mobileMenuOpen = false"
            class="nav-link-mobile"
          >
            Admin
          </router-link>
        </div>
      </div>
    </nav>
    
    <main class="flex-1">
      <router-view />
    </main>
    <footer class="bg-gray-900 text-white py-12">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center text-gray-400">
          <p>&copy; calvary chapel phnom penh</p>
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
      isScrolled: false,
      lastScrollY: 0,
      scrollBreakpoint: 100,
      mobileMenuOpen: false,
      aboutDropdownOpen: false,
      aboutDropdownMobileOpen: false,
      ministriesDropdownOpen: false,
      ministriesDropdownMobileOpen: false,
      eventsDropdownOpen: false,
      eventsDropdownMobileOpen: false
    }
  },
  mounted() {
    this.lastScrollY = window.scrollY
    this.handleScroll()
    window.addEventListener('scroll', this.handleScroll, { passive: true })
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleScroll() {
      const currentScrollY = window.scrollY
      const scrollDirection = currentScrollY > this.lastScrollY ? 'down' : 'up'
      
      // Use hysteresis: shrink at 100px when scrolling down, expand at 50px when scrolling up
      if (scrollDirection === 'down') {
        this.isScrolled = currentScrollY > this.scrollBreakpoint
      } else {
        this.isScrolled = currentScrollY > 50
      }
      
      this.lastScrollY = currentScrollY
    }
  },
  watch: {
    $route() {
      // Close mobile menu and dropdowns when route changes
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
