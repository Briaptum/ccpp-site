<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Youth Ministry</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          A vibrant community where teens encounter God, build lasting friendships, and discover their identity in Christ.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Our Youth Ministry helps teenagers and young adults discover who they are in Christ and develop a deep, personal relationship with God. We cultivate an environment where young people can ask questions, explore faith, and build authentic community.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Weekly gatherings combine engaging teaching, heartfelt worship, fun activities, and meaningful service opportunities so students can grow as disciples of Jesus.
          </p>
          <p class="text-lg text-gray-700 mb-4">What to expect:</p>
          <ul class="space-y-3 text-gray-700 mb-8">
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Interactive Bible teaching that speaks to everyday life</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Authentic worship led by and for students</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span>Service projects that put faith into action</span>
            </li>
          </ul>
        </div>

        <!-- Gallery Section -->
        <div class="mt-12">
          <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Gallery</h2>
          
          <!-- Gallery Grid -->
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-8">
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
              class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
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
                    ? 'bg-main text-white' 
                    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
                ]"
              >
                {{ page }}
              </button>
            </div>
            
            <button
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
            >
              Next
            </button>
          </div>
          
          <!-- Page Info -->
          <div v-if="totalPages > 1" class="text-center mt-4 text-gray-600">
            Showing {{ startIndex + 1 }}-{{ endIndex }} of {{ galleryImages.length }} images
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

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-main text-white font-semibold rounded-lg hover:opacity-90 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { ref, computed } from 'vue'
import image1 from '@/assets/gallery/571314951_1275851217909869_307313164703676157_n.jpg'
import image2 from '@/assets/gallery/571323032_1275851334576524_3846026136297630090_n.jpg'
import image3 from '@/assets/gallery/572406336_1275851351243189_5572256207364945942_n.jpg'
import image4 from '@/assets/gallery/573067084_1275851234576534_1576992884320795404_n.jpg'
import image5 from '@/assets/gallery/560041125_1275851251243199_3196792532920309856_n.jpg'
import image6 from '@/assets/gallery/548625714_1236675678494090_1776087346866151382_n.jpg'
import image7 from '@/assets/gallery/548273506_1236675898494068_1283754074671780680_n.jpg'
import image8 from '@/assets/gallery/548222744_1236675851827406_3818316444684279685_n.jpg'

export default {
  name: 'YouthMinistry',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const currentPage = ref(1)
    const imagesPerPage = 6

    const galleryImages = ref([
      { src: image1, alt: 'Youth Ministry 1' },
      { src: image2, alt: 'Youth Ministry 2' },
      { src: image3, alt: 'Youth Ministry 3' },
      { src: image4, alt: 'Youth Ministry 4' },
      { src: image5, alt: 'Youth Ministry 5' },
      { src: image6, alt: 'Youth Ministry 6' },
      { src: image7, alt: 'Youth Ministry 7' },
      { src: image8, alt: 'Youth Ministry 8' },
    ])

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
    }
  },
};
</script>

