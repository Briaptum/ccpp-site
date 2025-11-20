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
import araksatImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'
import steunmeancheyImage from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'
import preyVengImage from '@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg'
import image4 from '@/assets/gallery/548188348_1236675751827416_5183731636046995004_n.jpg'
import image5 from '@/assets/gallery/548044046_1236675671827424_5262319447121551375_n.jpg'
import image6 from '@/assets/gallery/547686117_1236675835160741_3886842454015531546_n.jpg'
import image7 from '@/assets/gallery/547653470_1236675908494067_1828648105188893790_n.jpg'
import image8 from '@/assets/gallery/547588178_1236675705160754_8393862384233332823_n.jpg'

export default {
  name: 'Outreaches',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const currentPage = ref(1)
    const imagesPerPage = 6

    const araksatImageData = araksatImage
    const steunmeancheyImageData = steunmeancheyImage
    const preyVengImageData = preyVengImage

    const galleryImages = ref([
      { src: araksatImage, alt: 'Outreach 1' },
      { src: steunmeancheyImage, alt: 'Outreach 2' },
      { src: preyVengImage, alt: 'Outreach 3' },
      { src: image4, alt: 'Outreach 4' },
      { src: image5, alt: 'Outreach 5' },
      { src: image6, alt: 'Outreach 6' },
      { src: image7, alt: 'Outreach 7' },
      { src: image8, alt: 'Outreach 8' },
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
      araksatImage: araksatImageData,
      steunmeancheyImage: steunmeancheyImageData,
      preyVengImage: preyVengImageData,
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

