<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Outreaches</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          We carry the love and message of Jesus into the communities of Cambodia, serving with compassion and sharing hope.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Our outreach programs are designed to take the love and message of Jesus Christ beyond our church walls into the communities of Cambodia. Serving others is a vital expression of our faith and a powerful way to share the Gospel.
          </p>
          <p class="text-lg text-gray-700 mb-12">
            Through these initiatives we partner with local organizations, serve those in need, and create opportunities for people to encounter God's love through practical acts of service.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="araksatImage" 
                alt="Araksat Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Araksat</h2>
              <p class="text-gray-700">
                Our outreach to the Araksat community focuses on sharing the Gospel and serving the people with love and compassion.
              </p>
            </div>
          </div>
          
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="steunmeancheyImage" 
                alt="Steunmeanchey Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Steunmeanchey</h2>
              <p class="text-gray-700">
                We serve the Steunmeanchey community through practical acts of service and sharing the hope of Jesus Christ.
              </p>
            </div>
          </div>
          
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="preyVengImage" 
                alt="Prey Veng Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Prey Veng</h2>
              <p class="text-gray-700">
                Our Prey Veng outreach brings the love of Christ to this community through ministry, service, and discipleship.
              </p>
            </div>
          </div>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-lg hover:bg-brand-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

    <!-- Gallery Section -->
    <div class="py-20 bg-gray-200">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-main mb-4">Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Glimpses of our outreach programs and community service
        </p>
        
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"></div>
          <p class="text-gray-700">Loading gallery...</p>
        </div>

        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-600">{{ error }}</p>
        </div>

        <div v-else>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            <div 
              v-for="(image, index) in paginatedImages" 
              :key="image.id || image.src"
              @click="openLightbox(getFullImageIndex(index))"
              class="relative overflow-hidden rounded-lg shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group bg-gray-200"
            >
              <img
                :src="image.src"
                :alt="image.alt"
                @error="handleImageError"
                class="w-full h-64 object-cover group-hover:scale-110 transition-transform duration-500"
              />
              <div class="absolute inset-0 bg-black/0 group-hover:bg-black/30 transition-all duration-300 flex items-center justify-center">
                <svg class="w-12 h-12 text-white opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- Pagination Controls -->
          <div v-if="totalPages > 1" class="flex items-center justify-center gap-2 mt-8">
            <button
              @click="goToPage(currentPage - 1)"
              :disabled="currentPage === 1"
              class="px-4 py-2 bg-main text-white rounded-lg hover:bg-secondary transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Previous
            </button>
            
            <div class="flex gap-2">
              <button
                v-for="page in totalPages"
                :key="page"
                @click="goToPage(page)"
                :class="[
                  'px-4 py-2 rounded-lg transition-colors',
                  currentPage === page
                    ? 'bg-brand-orange text-white font-semibold'
                    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
                ]"
              >
                {{ page }}
              </button>
            </div>
            
            <button
              @click="goToPage(currentPage + 1)"
              :disabled="currentPage === totalPages"
              class="px-4 py-2 bg-main text-white rounded-lg hover:bg-secondary transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Lightbox Modal -->
    <div
      v-if="lightboxOpen"
      @click="closeLightbox"
      class="fixed inset-0 bg-black/90 z-50 flex items-center justify-center p-4"
    >
      <button
        @click="closeLightbox"
        class="absolute top-4 right-4 text-white hover:text-gray-300 transition-colors z-10"
      >
        <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>

      <button
        v-if="currentImageIndex > 0"
        @click.stop="prevImage"
        class="absolute left-4 text-white hover:text-gray-300 transition-colors z-10"
      >
        <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
        </svg>
      </button>
      
      <div class="max-w-5xl w-full" @click.stop>
        <img
          :src="galleryImages[currentImageIndex]?.src"
          :alt="galleryImages[currentImageIndex]?.alt"
          class="w-full h-auto max-h-[80vh] object-contain rounded-lg"
        />
      </div>

      <button
        v-if="currentImageIndex < galleryImages.length - 1"
        @click.stop="nextImage"
        class="absolute right-4 text-white hover:text-gray-300 transition-colors z-10"
      >
        <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
        </svg>
      </button>
    </div>

  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { galleryService } from '@/services/galleryService'

export default {
  name: 'Outreaches',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const currentPage = ref(1)
    const itemsPerPage = 6
    const loading = ref(true)
    const error = ref('')
    const galleryImages = ref([])

    const getImageUrl = (path) => {
      if (path.startsWith('http')) {
        return path
      }
      return path.startsWith('/') ? path : `/${path}`
    }

    const fetchGalleryImages = async () => {
      loading.value = true
      error.value = ''
      try {
        const response = await galleryService.getGalleries('outreaches')
        const galleries = response.galleries || []
        galleryImages.value = galleries.map(gallery => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `Outreach Image ${gallery.id}`
        }))
        loading.value = false
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
        loading.value = false
      }
    }

    const handleImageError = (event) => {
      event.target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E'
    }

    const totalPages = computed(() => {
      return Math.ceil(galleryImages.value.length / itemsPerPage)
    })

    const paginatedImages = computed(() => {
      const start = (currentPage.value - 1) * itemsPerPage
      const end = start + itemsPerPage
      return galleryImages.value.slice(start, end)
    })

    const getFullImageIndex = (paginatedIndex) => {
      return (currentPage.value - 1) * itemsPerPage + paginatedIndex
    }

    const goToPage = (page) => {
      if (page >= 1 && page <= totalPages.value) {
        currentPage.value = page
        window.scrollTo({ top: 0, behavior: 'smooth' })
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

    onMounted(() => {
      fetchGalleryImages()
    })

    return {
      lightboxOpen,
      currentImageIndex,
      galleryImages,
      loading,
      error,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage,
      currentPage,
      itemsPerPage,
      totalPages,
      paginatedImages,
      getFullImageIndex,
      goToPage,
      handleImageError
    }
  },
};
</script>

