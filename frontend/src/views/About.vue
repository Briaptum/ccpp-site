<template>
  <div>
    <!-- Hero Section -->
    <div class="relative min-h-[80vh] md:min-h-[90vh] flex items-center overflow-hidden pt-16 md:pt-24">
      <div
        class="absolute inset-0 bg-cover bg-center"
        :style="{ backgroundImage: `url(${aboutHeroImage})` }"
      ></div>
      <div class="absolute inset-0 bg-gradient-to-r from-black/80 via-black/60 to-black/40"></div>
      <div class="relative z-10 w-full max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <p class="text-sm uppercase tracking-[0.4em] text-white/70 mb-4">About Calvary Chapel Phnom Penh</p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight mb-6">
          A worshiping community rooted in Scripture and present for Phnom Penh.
        </h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light mb-10">
          We gather to teach verse by verse through the Bible, cultivate authentic relationships,
          and serve the city with the hope of Jesus.
        </p>
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/what-we-believe"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors"
          >
            What We Believe
          </router-link>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-white/10 text-white font-semibold rounded-md border border-white/20 hover:bg-white/20 transition-colors backdrop-blur-sm"
          >
            Plan a Visit
          </router-link>
        </div>
      </div>
    </div>

    <!-- Story Section -->
    <section class="py-24 bg-white">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 lg:grid-cols-2 gap-12">
        <div>
          <p class="text-sm uppercase tracking-[0.3em] text-brand-blue mb-4">Our story</p>
          <h2 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Calvary Chapel Phnom Penh is a verse-by-verse church with an open Bible and open doors.
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
        <div class="bg-gray-50 rounded-3xl border border-gray-100 p-8 space-y-6">
          <div class="flex items-start space-x-4">
            <div class="flex-shrink-0 w-12 h-12 rounded-full bg-brand-orange/15 text-brand-orange flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2-1.343-2-3-2zm0 0V4m0 8v8m4-6h4m-8 0H4" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold text-gray-900 mb-1">Scripture-centered</h3>
              <p class="text-gray-600">Teaching verse by verse through books of the Bible so that the Word shapes everything.</p>
            </div>
          </div>
          <div class="flex items-start space-x-4">
            <div class="flex-shrink-0 w-12 h-12 rounded-full bg-brand-blue/15 text-brand-blue flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold text-gray-900 mb-1">Authentic discipleship</h3>
              <p class="text-gray-600">Smaller communities and mentors walking together in prayer, care, and mission.</p>
            </div>
          </div>
          <div class="flex items-start space-x-4">
            <div class="flex-shrink-0 w-12 h-12 rounded-full bg-gray-900/10 text-gray-900 flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8m-9 5v9" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold text-gray-900 mb-1">City presence</h3>
              <p class="text-gray-600">Serving Phnom Penh through mercy initiatives, youth mentorship, and partnerships.</p>
            </div>
          </div>
        </div>
      </div>
    </section>


    <!-- Gallery Section -->
    <div class="py-24 bg-brand-blue text-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
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
