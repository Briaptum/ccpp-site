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
        <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-lg">Welcome to Calvary Chapel Phnom Penh</h1>
        <p class="text-xl md:text-2xl mb-8 drop-shadow-md">Statement of Faith</p>
        <p class="text-xl mb-8 max-w-2xl mx-auto drop-shadow-md">    
          Studying the Word verse by verse, chapter by chapter.
        </p>
        <div class="space-x-4">
          <router-link to="/about" class="btn-primary bg-white text-primary-600 hover:bg-gray-100 shadow-lg">
            Learn More
          </router-link>
          <router-link to="/contact" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600 shadow-lg">
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

    <!-- Services Section -->
    <div class="relative py-16 md:py-20 overflow-hidden">
      <!-- Blurred Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="{ backgroundImage: `url(${background2Image})` }"
      >
        <div class="absolute inset-0 backdrop-blur-md bg-background/60"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-2xl font-bold text-center text-gray-900 mb-6">Please join us for:</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="card text-center bg-white overflow-hidden p-6">
            <div class="w-full h-64 mb-4 overflow-hidden">
              <img 
                :src="sundayWorshipImage" 
                alt="Sunday Worship" 
                class="w-full h-full object-cover"
              />
            </div>
            <h3 class="text-xl font-semibold mb-3">Sunday Worship</h3>
            <p class="text-gray-600">Join us every Sunday at 9:00 AM for inspiring worship and fellowship.</p>
          </div>
          
          <div class="card text-center bg-white overflow-hidden p-6">
            <div class="w-full h-64 mb-4 overflow-hidden">
              <img 
                :src="bibleStudyImage" 
                alt="Wednesday Bible Study" 
                class="w-full h-full object-cover"
              />
            </div>
            <h3 class="text-xl font-semibold mb-3">Wednesday Bible Study</h3>
            <p class="text-gray-600">Wednesday night Bible Study at 7:00 p.m.</p>
          </div>
          
        </div>
        <div class="text-center mt-6">
          <router-link to="/contact" class="inline-block px-8 py-3 bg-primary-600 text-white font-semibold rounded-lg hover:bg-secondary-700 transition-colors shadow-lg">
            Join Now
          </router-link>
        </div>
      </div>
    </div>

    <!-- Picture Gallery Section -->
    <div class="bg-background-100 py-12">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl font-bold text-center text-gray-900">Picture Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-8 max-w-2xl mx-auto">
          Glimpses of our community, worship services, and ministry events
        </p>
        
        <!-- Loading State -->
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mb-4"></div>
          <p class="text-gray-600">Loading gallery...</p>
        </div>

        <!-- Error Message -->
        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-600 mb-4">{{ error }}</p>
        </div>

        <!-- Gallery Grid -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-8">
          <div 
            v-for="(image, index) in paginatedImages" 
            :key="index"
            @click="openLightbox(getActualImageIndex(index))"
            class="relative overflow-hidden rounded-lg shadow-md hover:shadow-xl transition-shadow cursor-pointer group"
          >
            <div class="aspect-w-16 aspect-h-12 bg-gray-200">
              <img
                :src="image.src"
                :alt="image.alt"
                class="w-full h-64 object-cover group-hover:scale-110 transition-transform duration-300"
              />
            </div>
            <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-30 transition-opacity flex items-center justify-center">
              <svg class="w-12 h-12 text-white opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
              </svg>
            </div>
          </div>
        </div>

        <!-- Pagination Controls -->
        <div v-if="totalPages > 1" class="flex justify-center items-center space-x-4 mt-8">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-secondary-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
          >
            Previous
          </button>
          
          <div class="flex space-x-2">
            <button
              v-for="page in visiblePages"
              :key="page"
              @click="goToPage(page)"
              :class="[
                'px-3 py-2 rounded-lg transition-colors',
                page === currentPage 
                  ? 'bg-primary-600 text-white' 
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              ]"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-secondary-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
          >
            Next
          </button>
        </div>
        
        <!-- Page Info -->
        <div v-if="totalPages > 1" class="text-center mt-4 text-gray-600">
          Showing {{ startIndex + 1 }}-{{ endIndex }} of {{ galleryImages.length }} images
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="bg-primary-800 py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl font-bold text-center text-white">Latest Teachings</h2>
        <p class="text-lg text-white text-center mb-12 max-w-2xl mx-auto">
          Catch up on recent messages and stay encouraged throughout the week.
        </p>
        <div
          v-if="latestTeachings.length > 0"
          class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"
        >
          <div
            v-for="teaching in paginatedTeachings"
            :key="teaching.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden flex flex-col"
          >
            <div class="relative overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-40 object-cover transition-transform duration-300 hover:scale-105"
              />
              <div class="absolute inset-0 bg-black bg-opacity-25"></div>
              <div class="absolute inset-0 flex items-center justify-center">
                <svg class="w-12 h-12 text-white opacity-80" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z" />
                </svg>
              </div>
            </div>
            <div class="p-4 flex flex-col flex-1">
              <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-600 flex-1 mb-4">{{ teaching.description }}</p>
              <a
                :href="teaching.url"
                target="_blank"
                rel="noopener"
                class="inline-flex items-center justify-center px-4 py-2 bg-primary-600 text-white font-semibold rounded-lg hover:bg-secondary-700 transition-colors shadow-md text-sm"
              >
                Watch Now
              </a>
            </div>
          </div>
        </div>
        
        <!-- Video Pagination Controls -->
        <div v-if="totalVideoPages > 1" class="flex justify-center items-center space-x-4 mt-8">
          <button
            @click="previousVideoPage"
            :disabled="currentVideoPage === 1"
            class="px-4 py-2 bg-white text-primary-600 rounded-lg hover:bg-gray-100 disabled:bg-gray-300 disabled:cursor-not-allowed disabled:text-gray-500 transition-colors"
          >
            Previous
          </button>
          
          <div class="flex space-x-2">
            <button
              v-for="page in visibleVideoPages"
              :key="page"
              @click="goToVideoPage(page)"
              :class="[
                'px-3 py-2 rounded-lg transition-colors',
                page === currentVideoPage 
                  ? 'bg-primary-600 text-white' 
                  : 'bg-white text-primary-600 hover:bg-gray-100'
              ]"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="nextVideoPage"
            :disabled="currentVideoPage === totalVideoPages"
            class="px-4 py-2 bg-white text-primary-600 rounded-lg hover:bg-gray-100 disabled:bg-gray-300 disabled:cursor-not-allowed disabled:text-gray-500 transition-colors"
          >
            Next
          </button>
        </div>
        
        <p v-else class="text-center text-white">
          Latest teachings will appear here once available.
        </p>
      </div>
    </div>

    <!-- Lightbox Modal -->
    <div
      v-if="lightboxOpen"
      @click="closeLightbox"
      class="fixed inset-0 bg-black bg-opacity-90 z-50 flex items-center justify-center p-4"
    >
      <button
        @click="closeLightbox"
        class="absolute top-4 right-4 text-white hover:text-gray-300 transition-colors"
      >
        <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
      
      <button
        v-if="currentImageIndex > 0"
        @click.stop="prevImage"
        class="absolute left-4 text-white hover:text-gray-300 transition-colors"
      >
        <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
        </svg>
      </button>
      
      <div class="max-w-5xl w-full" @click.stop>
        <img
          :src="galleryImages[currentImageIndex].src"
          :alt="galleryImages[currentImageIndex].alt"
          class="w-full h-auto max-h-[80vh] object-contain rounded-lg"
        />
      </div>
      
      <button
        v-if="currentImageIndex < galleryImages.length - 1"
        @click.stop="nextImage"
        class="absolute right-4 text-white hover:text-gray-300 transition-colors"
      >
        <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
        </svg>
      </button>
    </div>

  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import background2Image from '@/assets/background/background2.jpg'
import sundayWorshipImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'
import bibleStudyImage from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'

export default {
  name: 'Home',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const loading = ref(true)
    const error = ref('')
    const currentPage = ref(1)
    const imagesPerPage = 6
    const currentVideoPage = ref(1)
    const videosPerPage = 3
    const heroSection = ref(null)
    const latestTeachings = [
      {
        id: 'dQw4w9WgXcQ',
        title: 'Service Highlight: Trust and Hope',
        description: 'Reflect on how trust in Christ shapes our response to trials.',
        url: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ'
      },
      {
        id: 'fLexgOxsZu0',
        title: 'Midweek Devotional: Walking in Faith',
        description: 'A midweek encouragement to walk boldly in faith.',
        url: 'https://www.youtube.com/watch?v=fLexgOxsZu0'
      },
      {
        id: '3GwjfUFyY6M',
        title: 'Sunday Sermon Recap: Grace Abounds',
        description: 'Grace meets us where we are and leads us forward.',
        url: 'https://www.youtube.com/watch?v=3GwjfUFyY6M'
      },
      {
        id: 'a3Z7zEc7AXQ',
        title: 'Youth Gathering: Rooted in Scripture',
        description: 'Highlights from our youth ministry Bible study.',
        url: 'https://www.youtube.com/watch?v=a3Z7zEc7AXQ'
      },
      {
        id: '2vjPBrBU-TM',
        title: 'Testimony Night: Stories of Renewal',
        description: 'Members share how God is renewing their lives.',
        url: 'https://www.youtube.com/watch?v=2vjPBrBU-TM'
      },
      {
        id: 'V-_O7nl0Ii0',
        title: 'Prayer Focus: Seeking His Presence',
        description: 'Join us as we focus our hearts in prayer.',
        url: 'https://www.youtube.com/watch?v=V-_O7nl0Ii0'
      }
    ]
    const parallaxOffset = ref(0)
    let animationFrameId = null
    
    // Gallery images
    const galleryImages = ref([])

    // Load local gallery images
    const fetchGalleryImages = async () => {
      loading.value = true
      
      try {
        // Import all gallery images dynamically
        const imageModules = import.meta.glob('@/assets/gallery/*.jpg', { eager: true })
        
        const images = []
        for (const path in imageModules) {
          const imageName = path.split('/').pop().replace('.jpg', '')
          images.push({
            src: imageModules[path].default,
            alt: `Gallery Image ${imageName}`,
            title: `Gallery Image ${imageName}`
          })
        }
        
        // Sort images by filename for consistent ordering
        images.sort((a, b) => a.title.localeCompare(b.title))
        
        galleryImages.value = images
        loading.value = false
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
        loading.value = false
      }
    }

    // Pagination computed properties
    const totalPages = computed(() => Math.ceil(galleryImages.value.length / imagesPerPage))
    
    const startIndex = computed(() => (currentPage.value - 1) * imagesPerPage)
    const endIndex = computed(() => Math.min(startIndex.value + imagesPerPage, galleryImages.value.length))
    
    const paginatedImages = computed(() => {
      return galleryImages.value.slice(startIndex.value, endIndex.value)
    })
    
    const visiblePages = computed(() => {
      const pages = []
      const start = Math.max(1, currentPage.value - 2)
      const end = Math.min(totalPages.value, start + 4)
      
      for (let i = start; i <= end; i++) {
        pages.push(i)
      }
      return pages
    })

    // Pagination methods
    const goToPage = (page) => {
      if (page >= 1 && page <= totalPages.value) {
        currentPage.value = page
      }
    }
    
    const nextPage = () => {
      if (currentPage.value < totalPages.value) {
        currentPage.value++
      }
    }
    
    const previousPage = () => {
      if (currentPage.value > 1) {
        currentPage.value--
      }
    }
    
    const getActualImageIndex = (paginatedIndex) => {
      return startIndex.value + paginatedIndex
    }

    // Video pagination computed properties
    const totalVideoPages = computed(() => Math.ceil(latestTeachings.length / videosPerPage))
    
    const videoStartIndex = computed(() => (currentVideoPage.value - 1) * videosPerPage)
    const videoEndIndex = computed(() => Math.min(videoStartIndex.value + videosPerPage, latestTeachings.length))
    
    const paginatedTeachings = computed(() => {
      return latestTeachings.slice(videoStartIndex.value, videoEndIndex.value)
    })
    
    const visibleVideoPages = computed(() => {
      const pages = []
      const start = Math.max(1, currentVideoPage.value - 2)
      const end = Math.min(totalVideoPages.value, start + 4)
      
      for (let i = start; i <= end; i++) {
        pages.push(i)
      }
      return pages
    })

    // Video pagination methods
    const goToVideoPage = (page) => {
      if (page >= 1 && page <= totalVideoPages.value) {
        currentVideoPage.value = page
      }
    }
    
    const nextVideoPage = () => {
      if (currentVideoPage.value < totalVideoPages.value) {
        currentVideoPage.value++
      }
    }
    
    const previousVideoPage = () => {
      if (currentVideoPage.value > 1) {
        currentVideoPage.value--
      }
    }

    const openLightbox = (index) => {
      currentImageIndex.value = index
      lightboxOpen.value = true
    }

    const closeLightbox = () => {
      lightboxOpen.value = false
    }

    const nextImage = () => {
      if (currentImageIndex.value < galleryImages.value.length - 1) {
        currentImageIndex.value++
      }
    }

    const prevImage = () => {
      if (currentImageIndex.value > 0) {
        currentImageIndex.value--
      }
    }

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

    // Fetch images when component mounts
    onMounted(() => {
      fetchGalleryImages()
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
      sundayWorshipImage,
      bibleStudyImage,
      lightboxOpen,
      currentImageIndex,
      loading,
      error,
      galleryImages,
      currentPage,
      totalPages,
      startIndex,
      endIndex,
      paginatedImages,
      visiblePages,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage,
      goToPage,
      nextPage,
      previousPage,
      getActualImageIndex,
      parallaxBackgroundStyle,
      heroSection,
      latestTeachings,
      currentVideoPage,
      totalVideoPages,
      paginatedTeachings,
      visibleVideoPages,
      goToVideoPage,
      nextVideoPage,
      previousVideoPage
    }
  }
}
</script>
