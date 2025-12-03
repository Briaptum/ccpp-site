<template>
  <div>
    <!-- Hero Section -->
    <div class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16">
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="{ backgroundImage: `linear-gradient(120deg, rgba(17, 39, 84, 0.25), rgba(17, 39, 84, 0.45)), url(${aboutHeroImage})` }"
        ></div>
        <div class="absolute inset-0 bg-gradient-to-br from-black/70 via-black/55 to-black/65"></div>
        <div class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-blue/35 via-brand-blue/15 to-transparent blur-[120px] opacity-60"></div>
        <div class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-orange/30 via-brand-orange/12 to-transparent blur-[180px] opacity-65"></div>
        <div class="absolute -top-24 left-1/2 w-[420px] h-[420px] -translate-x-1/2 bg-white/15 blur-[180px] opacity-60"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/80 via-black/60 to-transparent"></div>
        <div
          class="absolute inset-0 opacity-15 mix-blend-screen"
          style="background-image: radial-gradient(circle at 20% 28%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 10%, rgba(255,255,255,0.22), transparent 30%), radial-gradient(circle at 62% 78%, rgba(255,255,255,0.18), transparent 32%);"
        ></div>
      </div>
      <div class="relative z-10 w-full max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-2">
        <h1 class="text-xl sm:text-2xl md:text-3xl font-semibold text-white leading-snug">
          About Calvary Chapel Phnom Penh
        </h1>
        <p class="text-xs sm:text-sm md:text-base text-white/80 font-light">
          A simple Calvary Chapel fellowship teaching the Bible verse by verse, welcoming Phnom Penh to know and follow Jesus together.
        </p>
      </div>
    </div>

    <!-- Story Section -->
    <section class="py-24 bg-white">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 lg:grid-cols-2 gap-12">
        <div>
          <p class="text-sm uppercase tracking-[0.3em] text-brand-blue mb-4">Our story</p>
          <h2 class="text-4xl md:text-5xl font-bold text-brand-orange mb-6">
            What is Calvary Chapel Phnom Penh?
          </h2>
          <p class="text-lg text-gray-700 leading-relaxed mb-4">
            Beginning with a handful of families studying Scripture together, the church has grown into a diverse
            community anchored in the whole counsel of God. We value expository preaching, Spirit-led worship,
            and discipleship for every generation.
          </p>
          <p class="text-lg text-gray-700 leading-relaxed">
            From Sunday gatherings to mid-week studies, outreach projects, and leadership development, our aim is to
            equip people to know Jesus, love one another, and serve Phnom Penh with humility and hope.
          </p>
        </div>
        <div class="relative rounded-3xl overflow-hidden border border-gray-100 shadow-2xl bg-gray-900/5">
          <img
            :src="storyImage"
            alt="Calvary Chapel Phnom Penh community"
            class="w-full h-full object-cover"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/45 via-transparent to-transparent"></div>
        </div>
      </div>
    </section>


    <!-- Gallery Section -->
    <div class="relative py-24 text-white bg-brand-blue overflow-hidden">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 bg-gradient-to-br from-brand-blue via-[#0a1f3e] to-[#153866]"></div>
        <div class="absolute inset-y-[-25%] left-[-20%] w-2/3 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[160px] opacity-60"></div>
        <div class="absolute inset-y-[-30%] right-[-15%] w-1/2 bg-gradient-to-l from-white/15 via-transparent to-transparent blur-[140px] opacity-70"></div>
        <div class="absolute inset-x-0 bottom-0 h-44 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="background-image: radial-gradient(circle at 25% 20%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 70% 10%, rgba(255,255,255,0.22), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"
        ></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Gallery</p>
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">Life around Calvary Chapel Phnom Penh</h2>
          <p class="text-lg text-white/80 max-w-2xl mx-auto font-light">
            Glimpses of our community, worship services, and ministry events
          </p>
        </div>
        
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-white border-t-transparent mb-4"></div>
          <p class="text-white/80">Loading gallery...</p>
        </div>

        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-200">{{ error }}</p>
        </div>

        <div v-else>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-10">
            <div 
              v-for="(image, index) in paginatedImages" 
              :key="image.src"
              @click="openLightbox(getFullImageIndex(index))"
              class="relative overflow-hidden rounded-2xl border border-white/15 shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group bg-white/10 backdrop-blur-md"
            >
              <img
                :src="image.src"
                :alt="image.alt"
                class="w-full h-64 object-cover group-hover:scale-105 transition-transform duration-500 brightness-90"
              />
              <div class="absolute inset-0 bg-black/10 group-hover:bg-black/30 transition-all duration-300 flex items-center justify-center">
                <svg class="w-10 h-10 text-white opacity-0 group-hover:opacity-100 transition-opacity" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- Pagination Controls -->
          <div v-if="totalPages > 1" class="flex items-center justify-center gap-2">
            <button
              @click="goToPage(currentPage - 1)"
              :disabled="currentPage === 1"
              class="px-5 py-2.5 bg-white/10 blur-0 text-white border border-white/20 rounded-md hover:bg-white/20 transition-colors disabled:opacity-50 disabled:cursor-not-allowed font-medium"
            >
              Previous
            </button>
            
            <div class="flex gap-2">
              <button
                v-for="page in totalPages"
                :key="page"
                @click="goToPage(page)"
                :class="[
                  'px-4 py-2.5 rounded-md transition-colors font-medium border border-white/20',
                  currentPage === page
                    ? 'bg-brand-orange text-white'
                    : 'bg-white/10 text-white hover:bg-white/20'
                ]"
              >
                {{ page }}
              </button>
            </div>
            
            <button
              @click="goToPage(currentPage + 1)"
              :disabled="currentPage === totalPages"
              class="px-5 py-2.5 bg-white/10 text-white border border-white/20 rounded-md hover:bg-white/20 transition-colors disabled:opacity-50 disabled:cursor-not-allowed font-medium"
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
import { ref, onMounted, computed } from 'vue'
import aboutHeroImage from '@/assets/background/bg-hero4.jpg'
import storyImage from '@/assets/background/background2.jpg'

export default {
  name: 'About',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const loading = ref(true)
    const error = ref('')
    const galleryImages = ref([])
    const currentPage = ref(1)
    const itemsPerPage = 6

    const fetchGalleryImages = async () => {
      loading.value = true
      try {
        const imageModules = import.meta.glob('@/assets/gallery/*.jpg', { eager: true })
        const images = []
        for (const path in imageModules) {
          const imageName = path.split('/').pop().replace('.jpg', '')
          images.push({
            src: imageModules[path].default,
            alt: `Gallery Image ${imageName}`
          })
        }
        images.sort((a, b) => a.alt.localeCompare(b.alt))
        galleryImages.value = images
        loading.value = false
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
        loading.value = false
      }
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
      aboutHeroImage,
      storyImage,
      loading,
      error,
      galleryImages,
      lightboxOpen,
      currentImageIndex,
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
  }
}
</script>
