<template>
  <div>
    <!-- Hero Section with Background Image -->
    <div
      ref="heroSection"
      class="relative bg-main text-white py-40 md:py-48 overflow-hidden"
    >
      <!-- Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat will-change-transform"
        :style="parallaxBackgroundStyle"
      >
        <!-- Gradient overlay for better text readability and visual appeal -->
        <div class="absolute inset-0 bg-gradient-to-br from-main/85 via-main/75 to-secondary/90"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-2xl text-white">Welcome to Calvary Chapel Phnom Penh</h1>
        <p class="text-xl md:text-2xl mb-8 drop-shadow-lg text-primary">Statement of Faith</p>
        <p class="text-xl mb-8 max-w-2xl mx-auto drop-shadow-md text-gray-100">    
          Studying the Word verse by verse, chapter by chapter.
        </p>
        <div class="space-x-4 flex flex-wrap justify-center gap-4">
          <router-link to="/about" class="px-8 py-3 bg-white text-main font-semibold rounded-lg hover:bg-primary hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg">
            Learn More
          </router-link>
          <router-link to="/contact" class="px-8 py-3 border-2 border-white text-white font-semibold rounded-lg hover:bg-white hover:text-main hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg">
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

    <!-- Services Section -->
    <div class="relative py-16 md:py-20 overflow-hidden">
      <!-- Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="{ backgroundImage: `url(${joinUsBgImage})` }"
      >
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl md:text-4xl font-bold text-center text-main mb-4">Please join us for:</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mt-10">
          <div class="card text-center bg-white overflow-hidden p-0 shadow-xl hover:shadow-2xl transform hover:-translate-y-1 transition-all duration-300 rounded-xl">
            <div class="w-full h-64 mb-4 overflow-hidden">
              <img 
                :src="sundayWorshipImage" 
                alt="Sunday Worship" 
                class="w-full h-full object-cover hover:scale-110 transition-transform duration-500"
              />
            </div>
            <div class="p-6">
              <h3 class="text-2xl font-semibold mb-3 text-main">Sunday Worship</h3>
              <p class="text-gray-700 leading-relaxed">Join us every Sunday at 9:00 AM for inspiring worship and fellowship.</p>
            </div>
          </div>
          
          <div class="card text-center bg-white overflow-hidden p-0 shadow-xl hover:shadow-2xl transform hover:-translate-y-1 transition-all duration-300 rounded-xl">
            <div class="w-full h-64 mb-4 overflow-hidden">
              <img 
                :src="bibleStudyImage" 
                alt="Wednesday Bible Study" 
                class="w-full h-full object-cover hover:scale-110 transition-transform duration-500"
              />
            </div>
            <div class="p-6">
              <h3 class="text-2xl font-semibold mb-3 text-main">Wednesday Bible Study</h3>
              <p class="text-gray-700 leading-relaxed">Wednesday night Bible Study at 7:00 p.m.</p>
            </div>
          </div>
          
        </div>
        <div class="text-center mt-10">
          <router-link to="/contact" class="inline-block px-10 py-4 bg-main text-white font-semibold rounded-lg hover:bg-secondary hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg text-lg">
            Join Now
          </router-link>
        </div>
      </div>
    </div>

    <!-- Picture Gallery Section -->
    <div class="bg-gradient-to-b from-primary via-white to-primary py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-main mb-4">Picture Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Glimpses of our community, worship services, and ministry events
        </p>
        
        <!-- Loading State -->
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"></div>
          <p class="text-gray-700 font-medium">Loading gallery...</p>
        </div>

        <!-- Error Message -->
        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-600 mb-4">{{ error }}</p>
        </div>

        <!-- Gallery Grid -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          <div 
            v-for="(image, index) in paginatedImages" 
            :key="index"
            @click="openLightbox(getActualImageIndex(index))"
            class="relative overflow-hidden rounded-xl shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group transform hover:-translate-y-1"
          >
            <div class="aspect-w-16 aspect-h-12 bg-gradient-to-br from-gray-200 to-gray-300">
              <img
                :src="image.src"
                :alt="image.alt"
                class="w-full h-64 object-cover group-hover:scale-110 transition-transform duration-500"
              />
            </div>
            <div class="absolute inset-0 bg-gradient-to-t from-main/60 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-end justify-center pb-4">
              <svg class="w-12 h-12 text-white opacity-90" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
              </svg>
            </div>
          </div>
        </div>

        <!-- Pagination Controls -->
        <div v-if="totalPages > 1" class="flex justify-center items-center space-x-4 mt-10">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="px-5 py-2 bg-main text-white rounded-lg hover:bg-secondary hover:shadow-lg disabled:bg-gray-300 disabled:cursor-not-allowed disabled:hover:bg-gray-300 transition-all duration-200 font-medium"
          >
            Previous
          </button>
          
          <div class="flex space-x-2">
            <button
              v-for="page in visiblePages"
              :key="page"
              @click="goToPage(page)"
              :class="[
                'px-4 py-2 rounded-lg transition-all duration-200 font-medium',
                page === currentPage 
                  ? 'bg-main text-white shadow-lg scale-110' 
                  : 'bg-white text-main border-2 border-main hover:bg-main hover:text-white hover:shadow-md'
              ]"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-5 py-2 bg-main text-white rounded-lg hover:bg-secondary hover:shadow-lg disabled:bg-gray-300 disabled:cursor-not-allowed disabled:hover:bg-gray-300 transition-all duration-200 font-medium"
          >
            Next
          </button>
        </div>
        
        <!-- Page Info -->
        <div v-if="totalPages > 1" class="text-center mt-6 text-gray-600 font-medium">
          Showing {{ startIndex + 1 }}-{{ endIndex }} of {{ galleryImages.length }} images
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="bg-gradient-to-b from-custom-orange via-custom-orange/95 to-custom-orange/90 py-20">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-white mb-4">Latest Teachings</h2>
        <p class="text-lg text-white/90 text-center mb-16 max-w-2xl mx-auto">
          Catch up on recent messages and stay encouraged throughout the week.
        </p>
        <div
          v-if="latestTeachings.length > 0"
          class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"
        >
          <div
            v-for="teaching in paginatedTeachings"
            :key="teaching.id"
            class="bg-white rounded-xl shadow-xl overflow-hidden flex flex-col hover:shadow-2xl transform hover:-translate-y-1 transition-all duration-300"
          >
            <div class="relative overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-40 object-cover transition-transform duration-500 hover:scale-110"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-custom-orange/60 to-transparent"></div>
              <div class="absolute inset-0 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                <div class="bg-white/20 backdrop-blur-sm rounded-full p-4 hover:bg-white/30 transition-colors">
                  <svg class="w-12 h-12 text-white" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z" />
                  </svg>
                </div>
              </div>
            </div>
            <div class="p-6 flex flex-col flex-1">
              <h3 class="text-lg font-semibold text-main mb-3">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-700 flex-1 mb-5 leading-relaxed">{{ teaching.description }}</p>
              <a
                :href="teaching.url"
                target="_blank"
                rel="noopener"
                class="inline-flex items-center justify-center px-5 py-2.5 bg-custom-orange text-white font-semibold rounded-lg hover:bg-custom-orange/90 hover:shadow-lg transform hover:-translate-y-0.5 transition-all duration-300 text-sm"
              >
                Watch Now
              </a>
            </div>
          </div>
        </div>
        
        <!-- Video Pagination Controls -->
        <div v-if="totalVideoPages > 1" class="flex justify-center items-center space-x-4 mt-12">
          <button
            @click="previousVideoPage"
            :disabled="currentVideoPage === 1"
            class="px-5 py-2 bg-white text-custom-orange rounded-lg hover:bg-primary hover:shadow-lg disabled:bg-gray-400 disabled:cursor-not-allowed disabled:text-gray-600 transition-all duration-200 font-medium"
          >
            Previous
          </button>
          
          <div class="flex space-x-2">
            <button
              v-for="page in visibleVideoPages"
              :key="page"
              @click="goToVideoPage(page)"
              :class="[
                'px-4 py-2 rounded-lg transition-all duration-200 font-medium',
                page === currentVideoPage 
                  ? 'bg-white text-custom-orange shadow-lg scale-110'
                  : 'bg-white/20 text-white border-2 border-white/50 hover:bg-white/30 hover:shadow-md'
              ]"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="nextVideoPage"
            :disabled="currentVideoPage === totalVideoPages"
            class="px-5 py-2 bg-white text-custom-orange rounded-lg hover:bg-primary hover:shadow-lg disabled:bg-gray-400 disabled:cursor-not-allowed disabled:text-gray-600 transition-all duration-200 font-medium"
          >
            Next
          </button>
        </div>
        
        <p v-else class="text-center text-white/90 text-lg">
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
import joinUsBgImage from '@/assets/background/join-us-bg.jpg'
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
      joinUsBgImage,
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
