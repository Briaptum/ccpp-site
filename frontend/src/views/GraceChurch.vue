<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Grace Church</h1>
        <p class="text-xl md:text-2xl">A Community of Grace</p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">About Grace Church</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Grace Church represents our commitment to spreading the message of God's amazing grace throughout 
            Cambodia. We believe that grace is at the heart of the Gospel - the unmerited favor of God extended 
            to all who believe in Jesus Christ.
          </p>
          <p class="text-gray-700 mb-6">
            Through Grace Church, we seek to create communities where people experience God's grace, grow in 
            faith, and extend that same grace to others. Our gatherings focus on worship, fellowship, and 
            teaching from God's Word.
          </p>
        </div>
      </div>
    </div>

    <!-- Picture Gallery Section -->
    <div class="bg-gray-100 py-12">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl font-bold text-center text-gray-900">Grace Church Photo Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-8 max-w-2xl mx-auto">
          Moments from our Grace Church community gatherings and ministries
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

        <!-- Empty State -->
        <div v-else-if="galleryImages.length === 0" class="text-center py-12">
          <p class="text-gray-600">Gallery images coming soon.</p>
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
            class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
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
            class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
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

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Join Our Community</h2>
        <p class="text-xl mb-8">
          Experience God's grace with us.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/ministries" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Back to Ministries
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'

export default {
  name: 'GraceChurch',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const loading = ref(true)
    const error = ref('')
    const currentPage = ref(1)
    const imagesPerPage = 6

    const galleryImages = ref([])

    const buildImageList = (modules, labelPrefix) => {
      const images = []
      for (const path in modules) {
        const imageName = path.split('/').pop().replace('.jpg', '')
        images.push({
          src: modules[path].default,
          alt: `${labelPrefix} ${imageName}`,
          title: `${labelPrefix} ${imageName}`
        })
      }
      images.sort((a, b) => a.title.localeCompare(b.title))
      return images
    }

    const fetchGalleryImages = async () => {
      loading.value = true
      try {
        let imageModules = import.meta.glob('@/assets/gallery/grace-church/*.jpg', { eager: true })
        let images = buildImageList(imageModules, 'Grace Church Gallery Image')

        if (images.length === 0) {
          const fallbackModules = import.meta.glob('@/assets/gallery/*.jpg', { eager: true })
          images = buildImageList(fallbackModules, 'Gallery Image')
        }

        galleryImages.value = images
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
      } finally {
        loading.value = false
      }
    }

    const totalPages = computed(() => Math.ceil(galleryImages.value.length / imagesPerPage))
    const startIndex = computed(() => (currentPage.value - 1) * imagesPerPage)
    const endIndex = computed(() => Math.min(startIndex.value + imagesPerPage, galleryImages.value.length))
    const paginatedImages = computed(() => galleryImages.value.slice(startIndex.value, endIndex.value))

    const visiblePages = computed(() => {
      const pages = []
      const start = Math.max(1, currentPage.value - 2)
      const end = Math.min(totalPages.value, start + 4)
      for (let i = start; i <= end; i++) {
        pages.push(i)
      }
      return pages
    })

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

    const getActualImageIndex = (paginatedIndex) => startIndex.value + paginatedIndex

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
      getActualImageIndex
    }
  }
}
</script>

