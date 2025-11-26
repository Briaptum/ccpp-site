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
        <!-- Center radial effect using multiple layers -->
        <div class="absolute inset-0 bg-gradient-to-b from-transparent via-black/20 to-black/60"></div>
        <!-- Bottom gradient for text area readability -->
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/50 to-black/20"></div>
        <!-- Top gradient for balance -->
        <div class="absolute inset-0 bg-gradient-to-b from-black/40 via-transparent to-transparent"></div>
        <!-- Diagonal accent with brand color from top-left -->
        <div class="absolute inset-0 bg-gradient-to-br from-custom-orange/20 via-custom-orange/5 to-transparent"></div>
        <!-- Side gradients for framing and depth -->
        <div class="absolute inset-0 bg-gradient-to-r from-black/30 via-transparent to-black/30"></div>
        <!-- Additional center focus gradient -->
        <div class="absolute inset-0 bg-gradient-to-t from-transparent via-transparent to-black/30"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <!-- Main Heading -->
        <div class="mb-10">
          <h1 class="text-5xl md:text-6xl lg:text-7xl xl:text-8xl font-bold mb-4 text-white leading-tight">
            <span class="block drop-shadow-2xl">Calvary Chapel</span>
            <span class="block text-custom-orange drop-shadow-2xl mt-2">Phnom Penh</span>
          </h1>
        </div>
        
        <!-- Tagline -->
        <div class="mb-8">
          <div class="w-32 h-1 bg-custom-orange mx-auto mb-6 shadow-lg"></div>
          <p class="text-2xl md:text-3xl lg:text-4xl text-white font-light drop-shadow-xl tracking-wide mb-4">
            Studying the Word
          </p>
          <p class="text-xl md:text-2xl text-white/95 font-light drop-shadow-lg">
            Verse by verse, chapter by chapter
          </p>
        </div>
        
        <!-- Description -->
        <div class="max-w-3xl mx-auto mt-12">
          <p class="text-lg md:text-xl text-white/90 drop-shadow-md leading-relaxed font-light">
            A community of believers committed to teaching through the Bible systematically, 
            verse by verse, chapter by chapter, book by book.
          </p>
        </div>
      </div>
    </div>

    <!-- Services Section -->
    <div class="py-24 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl md:text-5xl font-bold text-main mb-4">Join Us for Worship</h2>
          <div class="w-20 h-1 bg-custom-orange mx-auto mb-4"></div>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto font-light">
            We welcome you to join our community for meaningful worship and Bible study
          </p>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-10 lg:gap-12">
          <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 border border-gray-100">
            <div class="h-72 overflow-hidden bg-gray-100">
              <img 
                :src="sundayWorshipImage" 
                alt="Sunday Worship Service" 
                class="w-full h-full object-cover transition-transform duration-700 hover:scale-105"
              />
            </div>
            <div class="p-8">
              <h3 class="text-2xl font-bold mb-4 text-main">Sunday Worship Service</h3>
              <p class="text-gray-600 mb-2 font-medium">Every Sunday at 9:00 AM</p>
              <p class="text-gray-700 mb-6 leading-relaxed">
                Join us for our weekly worship service featuring verse-by-verse Bible teaching, 
                heartfelt worship, and fellowship with our church family.
              </p>
              <router-link
                to="/contact"
                class="inline-flex items-center px-6 py-3 bg-custom-orange text-white font-medium rounded-md hover:bg-custom-orange/90 transition-colors duration-200 shadow-sm hover:shadow-md"
              >
                Get Directions
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>
          
          <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 border border-gray-100">
            <div class="h-72 overflow-hidden bg-gray-100">
              <img 
                :src="bibleStudyImage" 
                alt="Wednesday Bible Study" 
                class="w-full h-full object-cover transition-transform duration-700 hover:scale-105"
              />
            </div>
            <div class="p-8">
              <h3 class="text-2xl font-bold mb-4 text-main">Wednesday Bible Study</h3>
              <p class="text-gray-600 mb-2 font-medium">Every Wednesday at 7:00 PM</p>
              <p class="text-gray-700 mb-6 leading-relaxed">
                Mid-week Bible study providing deeper exploration of Scripture through 
                systematic teaching and discussion in a more intimate setting.
              </p>
              <router-link
                to="/contact"
                class="inline-flex items-center px-6 py-3 bg-custom-orange text-white font-medium rounded-md hover:bg-custom-orange/90 transition-colors duration-200 shadow-sm hover:shadow-md"
              >
                Get Directions
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="relative py-24 overflow-hidden bg-gradient-to-b from-gray-900 via-gray-700 to-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">Latest Teachings</h2>
          <div class="w-20 h-1 bg-custom-orange mx-auto mb-4"></div>
          <p class="text-lg text-white/90 max-w-2xl mx-auto font-light">
            Watch our recent messages and stay connected with God's Word throughout the week
          </p>
        </div>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="teaching in latestTeachings"
            :key="teaching.id"
            class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 border border-gray-100 group"
          >
            <div class="relative h-48 overflow-hidden bg-gray-900">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
              />
              <div class="absolute inset-0 flex items-center justify-center bg-black/30 group-hover:bg-black/20 transition-colors duration-300">
                <div class="w-16 h-16 bg-white/90 rounded-full flex items-center justify-center group-hover:bg-white transition-colors duration-300">
                  <svg class="w-6 h-6 text-main ml-1" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z" />
                  </svg>
                </div>
              </div>
            </div>
            <div class="p-6">
              <h3 class="text-lg font-bold text-main mb-3 leading-tight">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-600 mb-5 leading-relaxed">{{ teaching.description }}</p>
              <a
                :href="teaching.url"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center px-5 py-2.5 bg-main text-white font-medium rounded-md hover:bg-secondary transition-colors duration-200 text-sm shadow-sm hover:shadow-md"
              >
                Watch Now
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import backgroundImage from '@/assets/background/hero-bg.jpg'
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
