<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Grace Church</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          We spread the message of God's amazing grace throughout Cambodia, gathering communities to worship, grow, and extend grace to others.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Grace Church represents our commitment to spreading the message of God's amazing grace throughout Cambodia. Grace is at the heart of the Gospel—the unmerited favor of God extended to all who believe in Jesus Christ.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            We cultivate communities where people encounter God's grace, grow in faith, and extend that same grace to others. Our gatherings focus on worship, fellowship, and faithful teaching of God's Word.
          </p>
          <p class="text-lg text-gray-700 mb-4">Our focus includes:</p>
          <ul class="space-y-3 text-gray-700 mb-8">
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Cultivating worship gatherings centered on Jesus</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Building gospel-centered communities across Cambodia</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Equipping believers to extend grace to their neighbors</span>
            </li>
          </ul>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-custom-orange text-white font-semibold rounded-lg hover:bg-custom-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
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
          Glimpses of our Grace Church gatherings and community
        </p>
        
        <div>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            <div 
              v-for="(image, index) in paginatedImages" 
              :key="image.src"
              @click="openLightbox(getFullImageIndex(index))"
              class="relative overflow-hidden rounded-lg shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group bg-gray-200"
            >
              <img
                :src="image.src"
                :alt="image.alt"
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
                    ? 'bg-custom-orange text-white font-semibold'
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
import { ref, computed } from 'vue'
import image1 from '@/assets/gallery/548188348_1236675751827416_5183731636046995004_n.jpg'
import image2 from '@/assets/gallery/548044046_1236675671827424_5262319447121551375_n.jpg'
import image3 from '@/assets/gallery/547686117_1236675835160741_3886842454015531546_n.jpg'
import image4 from '@/assets/gallery/547653470_1236675908494067_1828648105188893790_n.jpg'
import image5 from '@/assets/gallery/547588178_1236675705160754_8393862384233332823_n.jpg'
import image6 from '@/assets/gallery/547541557_1236675618494096_195308602630672666_n.jpg'
import image7 from '@/assets/gallery/547212625_1236675858494072_4425998334324483137_n.jpg'
import image8 from '@/assets/gallery/514261004_1275851211243203_2877388550521442118_n.jpg'

export default {
  name: 'GraceChurch',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const currentPage = ref(1)
    const itemsPerPage = 6

    const galleryImages = ref([
      { src: image1, alt: 'Grace Church 1' },
      { src: image2, alt: 'Grace Church 2' },
      { src: image3, alt: 'Grace Church 3' },
      { src: image4, alt: 'Grace Church 4' },
      { src: image5, alt: 'Grace Church 5' },
      { src: image6, alt: 'Grace Church 6' },
      { src: image7, alt: 'Grace Church 7' },
      { src: image8, alt: 'Grace Church 8' },
    ])

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

    return {
      lightboxOpen,
      currentImageIndex,
      galleryImages,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage,
      currentPage,
      itemsPerPage,
      totalPages,
      paginatedImages,
      getFullImageIndex,
      goToPage
    }
  },
};
</script>
