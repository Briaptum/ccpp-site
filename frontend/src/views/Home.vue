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
        <div class="absolute inset-0 bg-black/60"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold mb-6 drop-shadow-2xl">
          <div class="text-white">Welcome to,</div>
          <div class="text-custom-orange relative inline-block">
            Calvary Chapel Phnom Penh
            <span class="absolute bottom-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-custom-orange to-transparent transform -skew-x-12"></span>
          </div>
        </h1>
        <p class="text-xl md:text-2xl mb-4 text-white font-semibold drop-shadow-lg">
          Statement of Faith
        </p>
        <p class="text-lg md:text-xl mb-10 text-white drop-shadow-md max-w-2xl mx-auto">
          Studying the Word verse by verse, chapter by chapter.
        </p>
      </div>
    </div>

    <!-- Services Section -->
    <div class="py-20 bg-primary">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-main mb-12">Join Us</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-shadow duration-300">
            <div class="h-64 overflow-hidden">
              <img 
                :src="sundayWorshipImage" 
                alt="Sunday Worship" 
                class="w-full h-full object-cover hover:scale-110 transition-transform duration-500"
              />
            </div>
            <div class="p-6">
              <h3 class="text-2xl font-semibold mb-3 text-main">Sunday Worship</h3>
              <p class="text-gray-700 mb-4">Join us every Sunday at 9:00 AM for inspiring worship and fellowship.</p>
              <router-link
                to="/contact"
                class="inline-block px-6 py-2 bg-custom-orange text-white font-semibold rounded-lg hover:bg-custom-orange/90 hover:shadow-lg transition-all duration-300"
              >
                Learn More
              </router-link>
            </div>
          </div>
          
          <div class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-shadow duration-300">
            <div class="h-64 overflow-hidden">
              <img 
                :src="bibleStudyImage" 
                alt="Wednesday Bible Study" 
                class="w-full h-full object-cover hover:scale-110 transition-transform duration-500"
              />
            </div>
            <div class="p-6">
              <h3 class="text-2xl font-semibold mb-3 text-main">Wednesday Bible Study</h3>
              <p class="text-gray-700 mb-4">Wednesday night Bible Study at 7:00 p.m.</p>
              <router-link
                to="/contact"
                class="inline-block px-6 py-2 bg-custom-orange text-white font-semibold rounded-lg hover:bg-custom-orange/90 hover:shadow-lg transition-all duration-300"
              >
                Learn More
          </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="relative py-20 overflow-hidden">
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="{ backgroundImage: `url(${joinUsBgImage})` }"
      >
        <div class="absolute inset-0 bg-black/40"></div>
      </div>
      <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-white mb-4">Latest Teachings</h2>
        <p class="text-lg text-white/90 text-center mb-12 max-w-2xl mx-auto">
          Catch up on recent messages and stay encouraged throughout the week.
        </p>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="teaching in latestTeachings"
            :key="teaching.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-2xl transition-shadow duration-300"
          >
            <div class="relative h-40 overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover"
              />
              <div class="absolute inset-0 flex items-center justify-center bg-black/20">
                <svg class="w-12 h-12 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z" />
                </svg>
              </div>
            </div>
            <div class="p-6 text-center">
              <h3 class="text-lg font-semibold text-main mb-2">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-700 mb-4">{{ teaching.description }}</p>
              <a
                :href="teaching.url"
                target="_blank"
                rel="noopener"
                class="inline-block px-5 py-2 bg-main text-white font-semibold rounded-lg hover:bg-secondary transition-colors text-sm"
              >
                Watch Now
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
import backgroundImage from '@/assets/background/background2.jpg'
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
