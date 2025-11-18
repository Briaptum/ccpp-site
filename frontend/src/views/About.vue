<template>
  <div>
    <!-- Hero Section with Background Image -->
    <div
      ref="heroSection"
      class="relative bg-gradient-to-r from-primary-600 to-primary-800 text-white py-40 md:py-48 overflow-hidden"
    >
      <!-- Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat will-change-transform"
        :style="parallaxBackgroundStyle"
      >
        <!-- Dark overlay for better text readability -->
        <div class="absolute inset-0 bg-black bg-opacity-70"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-lg">About Us</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-32 md:py-40">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Calvary Chapel Phnom Penh is a community of believers dedicated to serving God and sharing His love with the people of Cambodia. We are a part of the Calvary Chapel movement, which emphasizes verse-by-verse, chapter-by-chapter teaching through the Bible.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Our church family is diverse, welcoming people from all walks of life. We believe in creating a warm, loving environment where everyone can grow in their relationship with Jesus Christ.
          </p>
          <p class="text-lg text-gray-700 mb-8">
            Whether you're new to faith or have been walking with Christ for years, you're welcome here. We'd love to get to know you and have you join us as we seek to follow Jesus together.
          </p>
        </div>
        
        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-8 py-3 bg-primary-600 text-white font-semibold rounded-lg hover:bg-primary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-8 py-3 border-2 border-primary-600 text-primary-600 font-semibold rounded-lg hover:bg-primary-50 transition-colors"
          >
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import background2Image from '@/assets/images/background2.jpg'

export default {
  name: 'About',
  setup() {
    const heroSection = ref(null)
    const parallaxOffset = ref(0)
    let animationFrameId = null

    const updateParallax = () => {
      if (!heroSection.value) {
        return
      }

      const rect = heroSection.value.getBoundingClientRect()
      parallaxOffset.value = rect.top * -0.3
    }

    const handleScroll = () => {
      if (animationFrameId !== null) {
        return
      }

      animationFrameId = requestAnimationFrame(() => {
        updateParallax()
        animationFrameId = null
      })
    }

    const parallaxBackgroundStyle = computed(() => ({
      backgroundImage: `url(${background2Image})`,
      transform: `translateY(${parallaxOffset.value}px) scale(1.1)`
    }))

    onMounted(() => {
      updateParallax()
      window.addEventListener('scroll', handleScroll, { passive: true })
    })

    onBeforeUnmount(() => {
      window.removeEventListener('scroll', handleScroll)
      if (animationFrameId !== null) {
        cancelAnimationFrame(animationFrameId)
      }
    })

    return {
      background2Image,
      parallaxBackgroundStyle,
      heroSection
    }
  }
}
</script>
