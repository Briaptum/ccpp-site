<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"></div>
      <div class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"></div>
    </div>

    <section class="relative flex min-h-screen flex-col px-4 py-20 sm:px-8 lg:px-12">
      <div class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"></div>
        <div class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"></div>
      </div>

      <div class="w-full text-slate-100">
        <div class="w-full">
          <div class="text-center">
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Grace Church</p>
            <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">A community shaped by grace.</h1>
            <p class="mt-6 text-base text-slate-100/80">
              We spread the message of God’s amazing grace throughout Cambodia, gathering communities to worship, grow, and extend grace to others.
            </p>
          </div>

          <div class="mt-12 grid gap-10 lg:grid-cols-[1.1fr,0.9fr]">
            <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-10 shadow-xl backdrop-blur">
              <h2 class="text-2xl font-semibold text-white">About Grace Church</h2>
              <div class="mt-6 space-y-5 text-base text-slate-100/85">
                <p>
                  Grace Church represents our commitment to spreading the message of God's amazing grace throughout Cambodia. Grace is at the heart of the Gospel—the unmerited favor of God extended to all who believe in Jesus Christ.
                </p>
                <p>
                  We cultivate communities where people encounter God's grace, grow in faith, and extend that same grace to others. Our gatherings focus on worship, fellowship, and faithful teaching of God's Word.
                </p>
              </div>
            </article>

            <div class="rounded-3xl border border-primary-300/30 bg-primary-500/10 p-10 shadow-xl backdrop-blur">
              <h3 class="text-lg font-semibold uppercase tracking-[0.3em] text-primary-200">Our Focus</h3>
              <ul class="mt-6 space-y-4 text-sm text-slate-100/80">
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Cultivating worship gatherings centered on Jesus.</span>
                </li>
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Building gospel-centered communities across Cambodia.</span>
                </li>
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Equipping believers to extend grace to their neighbors.</span>
                </li>
              </ul>
            </div>
          </div>

          <div class="mt-16">
            <h2 class="text-center text-2xl font-semibold text-white">Grace Church Photo Gallery</h2>
            <p class="mt-4 text-center text-base text-slate-100/80">
              Moments from our Grace Church community gatherings and ministries.
            </p>

            <div class="mt-12 rounded-3xl border border-white/10 bg-slate-900/70 p-6 shadow-xl backdrop-blur">
              <div v-if="loading" class="flex flex-col items-center justify-center py-12">
                <div class="mb-4 inline-block h-12 w-12 animate-spin rounded-full border-b-2 border-primary-300"></div>
                <p class="text-sm text-slate-100/70">Loading gallery...</p>
              </div>

              <div v-else-if="error" class="py-8 text-center">
                <p class="text-sm font-semibold text-red-300">{{ error }}</p>
              </div>

              <div v-else-if="galleryImages.length === 0" class="py-12 text-center">
                <p class="text-sm text-slate-100/70">Gallery images coming soon.</p>
              </div>

              <div v-else class="space-y-8">
                <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
                  <div
                    v-for="(image, index) in paginatedImages"
                    :key="index"
                    @click="openLightbox(getActualImageIndex(index))"
                    class="group relative overflow-hidden rounded-2xl border border-white/10 bg-slate-950/40 shadow-lg ring-1 ring-white/5 transition hover:-translate-y-1 hover:border-primary-200/50 hover:ring-primary-300/40"
                  >
                    <div class="aspect-[4/3] bg-slate-900/70">
                      <img
                        :src="image.src"
                        :alt="image.alt"
                        class="h-full w-full object-cover transition-all duration-300 group-hover:scale-105 group-hover:brightness-110"
                      />
                    </div>
                    <div class="pointer-events-none absolute inset-0 flex items-center justify-center bg-slate-950/0 transition group-hover:bg-slate-950/40">
                      <svg class="h-12 w-12 text-white opacity-0 transition-opacity group-hover:opacity-100" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <div v-if="totalPages > 1" class="space-y-4 text-center">
                  <div class="flex flex-wrap items-center justify-center gap-3">
                    <button
                      @click="previousPage"
                      :disabled="currentPage === 1"
                      class="inline-flex items-center rounded-full border border-white/20 px-4 py-2 text-xs font-semibold uppercase tracking-[0.2em] text-slate-100 transition hover:-translate-y-0.5 hover:border-primary-200 hover:text-white disabled:border-white/10 disabled:text-slate-500 disabled:hover:translate-y-0"
                    >
                      Previous
                    </button>
                    <div class="flex flex-wrap items-center justify-center gap-2">
                      <button
                        v-for="page in visiblePages"
                        :key="page"
                        @click="goToPage(page)"
                        :class="[
                          'inline-flex items-center justify-center rounded-full px-3 py-1 text-xs font-semibold uppercase tracking-[0.2em] transition',
                          page === currentPage
                            ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/30'
                            : 'border border-white/15 text-slate-100 hover:-translate-y-0.5 hover:border-primary-200 hover:text-white'
                        ]"
                      >
                        {{ page }}
                      </button>
                    </div>
                    <button
                      @click="nextPage"
                      :disabled="currentPage === totalPages"
                      class="inline-flex items-center rounded-full border border-white/20 px-4 py-2 text-xs font-semibold uppercase tracking-[0.2em] text-slate-100 transition hover:-translate-y-0.5 hover:border-primary-200 hover:text-white disabled:border-white/10 disabled:text-slate-500 disabled:hover:translate-y-0"
                    >
                      Next
                    </button>
                  </div>
                  <p class="text-xs font-semibold uppercase tracking-[0.3em] text-slate-100/60">
                    Showing {{ startIndex + 1 }}-{{ endIndex }} of {{ galleryImages.length }} images
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div
            v-if="lightboxOpen"
            @click="closeLightbox"
            class="fixed inset-0 z-50 flex items-center justify-center bg-black/90 p-4"
          >
            <button
              @click="closeLightbox"
              class="absolute right-4 top-4 text-white transition hover:text-slate-200"
            >
              <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>

            <button
              v-if="currentImageIndex > 0"
              @click.stop="prevImage"
              class="absolute left-4 text-white transition hover:text-slate-200"
            >
              <svg class="h-12 w-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
              </svg>
            </button>

            <div class="w-full" @click.stop>
              <img
                :src="galleryImages[currentImageIndex].src"
                :alt="galleryImages[currentImageIndex].alt"
                class="max-h-[80vh] w-full rounded-3xl object-contain"
              />
            </div>

            <button
              v-if="currentImageIndex < galleryImages.length - 1"
              @click.stop="nextImage"
              class="absolute right-4 text-white transition hover:text-slate-200"
            >
              <svg class="h-12 w-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
          </div>

          <div class="mt-16 rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
            <h2 class="text-2xl font-semibold text-white">Join Our Community</h2>
            <p class="mt-4 text-base text-slate-100/80">Experience God’s grace with us and discover your place in the family.</p>
            <div class="mt-6 flex flex-wrap justify-center gap-4">
              <router-link
                to="/contact"
                class="inline-flex items-center justify-center rounded-full border border-primary-200 px-6 py-3 text-sm font-semibold text-primary-100 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
              >
                Contact Us
              </router-link>
              <router-link
                to="/ministries"
                class="inline-flex items-center justify-center rounded-full border border-white/20 px-6 py-3 text-sm font-semibold text-slate-100 transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
              >
                Back to Ministries
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
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

