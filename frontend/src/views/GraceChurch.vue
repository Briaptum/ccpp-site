<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#101c33] via-[#152a47] to-[#1b3a5c] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 18% 25%, rgba(255,255,255,0.2), transparent 45%), radial-gradient(circle at 72% 15%, rgba(255,255,255,0.15), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Grace Church</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          House gatherings and Sunday services proclaiming the grace of Jesus across Phnom Penh.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Word & Table
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            House Gatherings
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Phnom Penh
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Grace Culture
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">What is Grace Church?</h2>
          <div class="space-y-4 text-left text-lg text-gray-700 leading-relaxed max-w-3xl mx-auto">
            <p>
              Grace Church gathers believers around <span class="font-semibold text-brand-orange">gospel-centered teaching</span>, worship, and community across Phnom Penh neighborhoods.
            </p>
            <p>
              We plant house churches, host Sunday services, and cultivate leaders who extend God’s grace to their friends, families, and co-workers.
            </p>
            <p>
              Every gathering keeps Scripture, communion, and Spirit-led prayer at the center so Jesus shapes our culture and mission.
            </p>
          </div>
        </section>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Gather</p>
            <h3 class="text-2xl font-semibold text-gray-900">Word & worship</h3>
            <p class="text-gray-600">
              Verse-by-verse teaching and gospel-saturated liturgy anchor every house gathering and Sunday service.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Community</p>
            <h3 class="text-2xl font-semibold text-gray-900">Grace culture</h3>
            <p class="text-gray-600">
              Hospitality, shared meals, and pastoral care create safe spaces to confess sin, celebrate redemption, and grow together.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Mission</p>
            <h3 class="text-2xl font-semibold text-gray-900">City-wide vision</h3>
            <p class="text-gray-600">
              Training leaders and planting new communities ensures Phnom Penh hears the message of grace house by house.
            </p>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            Visit Grace Church
          </router-link>
        </div>
      </div>
    </section>

    <!-- Gallery Section -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
        <div class="text-center space-y-3">
          <h2 class="text-4xl font-bold text-main">Gallery</h2>
          <p class="text-lg text-gray-700 max-w-2xl mx-auto">
            Scroll through images from Grace Church house gatherings, baptisms, and Sunday services.
          </p>
        </div>

        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"></div>
          <p class="text-gray-700">Loading gallery...</p>
        </div>
        <div v-else-if="galleryImages.length === 0" class="text-center text-gray-500 py-12">
          No gallery images yet.
        </div>
        <div v-else-if="galleryImages.length === 0" class="text-center text-gray-500 py-12">
          No gallery images yet.
        </div>
        <div v-else class="relative max-w-5xl mx-auto">
          <div
            ref="galleryScroll"
            class="flex gap-4 overflow-x-auto rounded-3xl shadow-2xl bg-gray-200 px-4 py-4 scroll-smooth"
          >
            <div
              v-for="image in galleryImages"
              :key="image.id || image.src"
              class="w-64 h-44 md:w-80 md:h-52 flex-shrink-0 rounded-2xl overflow-hidden bg-gray-200"
            >
              <img
                :src="image.src"
                :alt="image.alt"
                @error="handleImageError"
                class="w-full h-full object-cover"
              />
            </div>
          </div>
          <button
            @click="scrollGallery(-1)"
            class="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 hover:bg-white text-gray-700 rounded-full p-3 shadow-lg transition hidden md:flex"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button
            @click="scrollGallery(1)"
            class="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 hover:bg-white text-gray-700 rounded-full p-3 shadow-lg transition hidden md:flex"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>
    </section>

  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { galleryService } from '@/services/galleryService'

export default {
  name: 'GraceChurch',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const galleryImages = ref([])
    const galleryScroll = ref(null)

    const getImageUrl = (path) => {
      if (path.startswith('http')) {
        return path
      }
      return path.startsWith('/') ? path : `/${path}`
    }

    const fetchGalleryImages = async () => {
      loading.value = true
      error.value = ''
      try {
        const response = await galleryService.getGalleries('grace-church')
        const galleries = response.galleries || []
        galleryImages.value = galleries.map((gallery) => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `Grace Church Image ${gallery.id}`
        }))
        loading.value = false
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
        loading.value = false
      }
    }

    const handleImageError = (event) => {
      event.target.src =
        'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E'
    }

    const scrollGallery = (direction) => {
      const el = galleryScroll.value
      if (!el) return
      const amount = el.clientWidth * 0.8 * direction
      el.scrollBy({ left: amount, behavior: 'smooth' })
    }

    onMounted(() => {
      fetchGalleryImages()
    })

    return {
      galleryImages,
      loading,
      error,
      handleImageError,
      galleryScroll,
      scrollGallery
    }
  }
}
</script>
