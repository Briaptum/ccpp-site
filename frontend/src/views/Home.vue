<template>
  <div>
    <!-- Hero Section with Parallax Background -->
    <div
      ref="heroSection"
      class="relative min-h-screen flex items-center justify-center overflow-hidden"
    >
      <!-- Parallax Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="parallaxStyle"
      >
        <!-- Multiple gradient overlays for professional look and text readability -->
        <!-- Navbar blue radial gradient from center -->
        <div class="absolute inset-0" style="background-image: radial-gradient(circle at center, rgba(34, 112, 130, 0.30), rgba(34, 112, 130, 0.18), transparent);"></div>
        <!-- Navbar blue overlay from center spreading outward -->
        <div class="absolute inset-0 bg-gradient-to-br from-brand-blue/20 via-brand-blue/8 to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-tl from-brand-blue/20 via-brand-blue/8 to-transparent"></div>
        <!-- Additional navbar blue accent from bottom center -->
        <div class="absolute inset-0 bg-gradient-to-t from-brand-blue/15 via-transparent to-transparent"></div>
        <!-- Center radial effect using multiple layers -->
        <div class="absolute inset-0 bg-gradient-to-b from-transparent via-black/15 to-black/50"></div>
        <!-- Bottom gradient for text area readability -->
        <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/45 to-black/15"></div>
        <!-- Top gradient for balance -->
        <div class="absolute inset-0 bg-gradient-to-b from-black/30 via-transparent to-transparent"></div>
        <!-- Diagonal accent with brand orange from top-left -->
        <div class="absolute inset-0 bg-gradient-to-br from-brand-orange/12 via-brand-orange/4 to-transparent"></div>
        <!-- Side gradients for framing and depth -->
        <div class="absolute inset-0 bg-gradient-to-r from-black/25 via-transparent to-black/25"></div>
        <!-- Additional center focus gradient -->
        <div class="absolute inset-0 bg-gradient-to-t from-transparent via-transparent to-black/20"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12 items-center min-h-[70vh] py-12">
          <!-- Left Side: Mission Statement -->
          <div class="text-center lg:text-left">
            <div class="mb-6">
              <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold mb-4 text-white leading-tight drop-shadow-2xl">
                <span class="text-brand-orange uppercase">CALVARY CHAPEL</span><br />
                <span class="text-white text-2xl md:text-3xl lg:text-4xl uppercase font-light tracking-[0.3em]">PHNOM PENH</span>
              </h1>
              <div class="w-24 h-1 bg-brand-orange mx-auto lg:mx-0 mb-6 shadow-lg"></div>
            </div>
            
            <p class="text-xl md:text-2xl text-white font-light drop-shadow-xl mb-4 leading-relaxed">
              Studying the Word verse by verse, chapter by chapter
            </p>
            
            <p class="text-base md:text-lg text-white/90 drop-shadow-md leading-relaxed font-light mb-8 max-w-2xl mx-auto lg:mx-0">
              A community of believers committed to teaching through the Bible systematically, 
              verse by verse, chapter by chapter, book by book.
            </p>
          </div>
          
          <!-- Right Side: Service Times -->
          <div class="flex justify-center lg:justify-end">
            <div class="p-0 md:p-2 text-white max-w-md w-full">
              <h2 class="text-2xl font-semibold text-white mb-6 text-center">Gather With Us</h2>

              <div class="space-y-8 mb-8">
                <div>
                  <p class="text-sm uppercase tracking-[0.4em] text-white/80 mb-2">
                    Sunday Worship · <span class="font-semibold tracking-normal">9:00 AM</span>
                  </p>
                  <p class="text-base text-white/85 leading-relaxed">
                    Verse-by-verse teaching, worship, and fellowship every Sunday morning.
                  </p>
                </div>

                <div>
                  <p class="text-sm uppercase tracking-[0.4em] text-white/80 mb-2">
                    Wednesday Bible Study · <span class="font-semibold tracking-normal">7:00 PM</span>
                  </p>
                  <p class="text-base text-white/85 leading-relaxed">
                    Mid-week Bible study, discussion, and community prayer.
                  </p>
                </div>
              </div>

              <router-link
                to="/contact"
                class="inline-flex items-center justify-center w-full text-sm font-semibold tracking-wide text-brand-orange hover:text-brand-blue transition-colors"
              >
                <span class="underline underline-offset-4 decoration-2">Get Directions</span>
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Upcoming Section -->
    <div class="py-24 bg-brand-blue">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-white">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Upcoming</p>
          <h2 class="text-4xl md:text-5xl font-bold mb-4">What’s ahead this season</h2>
          <p class="text-white/80 max-w-2xl mx-auto">
            A glimpse at the moments, gatherings, and opportunities we’re preparing for together.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 2</p>
            <h3 class="text-2xl font-semibold mb-3">City Prayer Night</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              An evening of prayer and worship for Phnom Penh, hosted in the main sanctuary at 6:30 PM.
            </p>
            <router-link
              to="/events"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Details
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 10</p>
            <h3 class="text-2xl font-semibold mb-3">Foundations Class</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              Four weeks exploring what we believe, how to grow, and where to serve. Sundays at 11:15 AM.
            </p>
            <router-link
              to="/ministries"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Save a seat
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 16</p>
            <h3 class="text-2xl font-semibold mb-3">Serve Phnom Penh</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              Neighborhood outreach day providing meals, prayer, and practical care. Teams launch at 8:00 AM.
            </p>
            <router-link
              to="/ministries/outreaches"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Volunteer
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="py-24 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-gray-500 mb-3">Latest Teachings</p>
          <h2 class="text-4xl md:text-5xl font-bold text-gray-800 mb-4">Recent messages from our pulpit</h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto font-light">
            Teachings captured each week to keep you connected to the verse-by-verse journey through Scripture.
          </p>
        </div>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="teaching in latestTeachings"
            :key="teaching.id"
            class="bg-white rounded-2xl border border-gray-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden flex flex-col"
          >
            <div class="relative h-48 overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover transition-transform duration-500 hover:scale-105"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/40 via-transparent to-transparent"></div>
              <div class="absolute bottom-4 left-4 inline-flex items-center px-3 py-1 bg-white/90 rounded-full text-xs font-semibold tracking-wide text-gray-900 shadow">
                Teaching Series
              </div>
            </div>
            <div class="p-6 flex flex-col flex-1">
              <h3 class="text-lg font-semibold text-gray-800 mb-2 leading-snug line-clamp-2">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-600 mb-5 leading-relaxed line-clamp-3 flex-1">
                {{ teaching.description }}
              </p>
              <div class="flex items-center justify-between border-t border-gray-100 pt-4">
                <span class="text-xs uppercase tracking-[0.3em] text-gray-500">YouTube</span>
                <a
                  :href="teaching.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center text-sm font-semibold text-brand-blue hover:text-gray-900 transition-colors"
                >
                  Watch now
                  <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import backgroundImage from '@/assets/background/bg-hero4.jpg'
import joinUsBgImage from '@/assets/background/join-us-bg.jpg'
import sundayWorshipImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'
import bibleStudyImage from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'

export default {
  name: 'Home',
  setup() {
    const heroSection = ref(null)
    const scrollY = ref(0)
    let animationFrameId = null
    
    const parallaxStyle = computed(() => {
      const yPos = scrollY.value * 0.5
      return {
        backgroundImage: `url(${backgroundImage})`,
        transform: `translateY(${yPos}px)`,
        willChange: 'transform'
      }
    })

    const handleScroll = () => {
      if (animationFrameId !== null) {
        return
      }

      animationFrameId = requestAnimationFrame(() => {
        if (heroSection.value) {
          const rect = heroSection.value.getBoundingClientRect()
          scrollY.value = window.scrollY
        }
        animationFrameId = null
      })
    }

    onMounted(() => {
      window.addEventListener('scroll', handleScroll, { passive: true })
      handleScroll()
    })

    onBeforeUnmount(() => {
      window.removeEventListener('scroll', handleScroll)
      if (animationFrameId !== null) {
        cancelAnimationFrame(animationFrameId)
      }
    })

    const latestTeachings = [
      {
        id: 'IrTzMEBtbaI',
        title: 'Latest Teaching 1',
        description: 'Watch our latest message and be encouraged.',
        url: 'https://www.youtube.com/watch?v=IrTzMEBtbaI'
      },
      {
        id: 'AdDxynFrn8E',
        title: 'Latest Teaching 2',
        description: 'Join us for this week\'s teaching.',
        url: 'https://www.youtube.com/watch?v=AdDxynFrn8E'
      },
      {
        id: 'FC0Tm4bNPkk',
        title: 'Latest Teaching 3',
        description: 'Stay connected with our community through God\'s Word.',
        url: 'https://www.youtube.com/watch?v=FC0Tm4bNPkk'
      }
    ]

    return {
      heroSection,
      parallaxStyle,
      joinUsBgImage,
      sundayWorshipImage,
      bibleStudyImage,
      latestTeachings
    }
  }
}
</script>
