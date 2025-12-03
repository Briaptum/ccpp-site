<template>
  <div>
    <!-- Hero Section -->
    <div class="relative min-h-[62vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-20">
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
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <div class="text-center">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Gallery</p>
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">Life around Calvary Chapel Phnom Penh</h2>
          <p class="text-lg text-white/80 max-w-2xl mx-auto font-light">
            Scenes from worship, discipleship, and everyday ministry in Phnom Penh.
          </p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
          <article
            v-for="highlight in highlightSections"
            :key="highlight.id"
            class="rounded-[32px] border border-white/15 bg-white/10 backdrop-blur-lg overflow-hidden flex flex-col md:flex-row"
          >
            <div class="md:w-1/2 relative">
              <img
                :src="highlight.image"
                :alt="highlight.title"
                class="w-full h-64 md:h-full object-cover"
              />
              <div class="absolute inset-0 bg-gradient-to-tr from-black/50 via-transparent to-transparent"></div>
            </div>
            <div class="md:w-1/2 p-8 flex flex-col gap-4">
              <p class="text-xs uppercase tracking-[0.35em] text-white/70">{{ highlight.tag }}</p>
              <h3 class="text-2xl font-semibold leading-tight">{{ highlight.title }}</h3>
              <p class="text-white/80 text-base leading-relaxed">
                {{ highlight.description }}
              </p>
              <p class="text-white/60 text-sm leading-relaxed">
                {{ highlight.detail }}
              </p>
            </div>
          </article>
        </div>

        <div class="space-y-6">
          <div class="flex flex-wrap items-end justify-between gap-4">
            <div>
              <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-2">Rolling banner</p>
              <h3 class="text-3xl font-semibold">A moving glimpse of community moments</h3>
            </div>
            <span class="text-white/70 text-sm">Updates automatically as new gallery photos are added.</span>
          </div>
          <div class="relative overflow-hidden rounded-[40px] border border-white/15 bg-white/5 px-4 py-10">
            <div v-if="loading" class="flex flex-col items-center justify-center py-10 text-white/80">
              <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-white border-t-transparent mb-4"></div>
              <p>Loading gallery...</p>
            </div>
            <div v-else-if="error" class="text-center py-10 text-red-200">
              {{ error }}
            </div>
            <div v-else class="marquee-mask">
              <div class="marquee-track">
                <div
                  v-for="(image, index) in marqueeImages"
                  :key="`${image.src}-${index}`"
                  class="marquee-item"
                >
                  <img
                    :src="image.src"
                    :alt="image.alt"
                    class="h-56 w-auto object-cover rounded-2xl border border-white/15 shadow-lg"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import aboutHeroImage from '@/assets/background/bg-hero4.jpg'
import storyImage from '@/assets/background/background2.jpg'
import highlightWorship from '@/assets/gallery/547212625_1236675858494072_4425998334324483137_n.jpg'
import highlightOutreach from '@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg'
import highlightYouth from '@/assets/gallery/547541557_1236675618494096_195308602630672666_n.jpg'
import highlightPrayer from '@/assets/gallery/548044046_1236675671827424_5262319447121551375_n.jpg'

export default {
  name: 'About',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const galleryImages = ref([])

    const highlightSections = [
      {
        id: 'worship',
        tag: 'Gathered worship',
        title: 'Sundays rooted in Scripture',
        description: 'A bilingual family of believers worshiping, opening the Bible verse by verse, and sharing life together each week.',
        detail: 'Captured during our 9:00 AM service inside the Phnom Penh sanctuary.',
        image: highlightWorship
      },
      {
        id: 'outreach',
        tag: 'City outreach',
        title: 'Serving Phnom Penh neighborhoods',
        description: 'Teams distribute meals, pray with neighbors, and offer practical care throughout the city.',
        detail: 'Moments from Serve Phnom Penh outreach mornings.',
        image: highlightOutreach
      },
      {
        id: 'youth',
        tag: 'Next generation',
        title: 'Youth discovering Jesus together',
        description: 'Students gather midweek for teaching, discussion, and discipleship that equips them to follow Jesus boldly.',
        detail: 'From our Friday youth ministry evenings.',
        image: highlightYouth
      },
      {
        id: 'prayer',
        tag: 'Prayer & community',
        title: 'House gatherings & prayer nights',
        description: 'Circles of prayer, encouragement, and hospitality happening in homes across the city.',
        detail: 'Snapshots from our monthly city prayer night.',
        image: highlightPrayer
      }
    ]

    const fallbackBannerImages = highlightSections.map((section) => ({
      src: section.image,
      alt: section.title
    }))

    const fetchGalleryImages = async () => {
      loading.value = true
      try {
        const imageModules = import.meta.glob('@/assets/gallery/*.jpg', { eager: true })
        const images = Object.keys(imageModules).map((path) => {
          const imageName = path.split('/').pop().replace('.jpg', '')
          return {
            src: imageModules[path].default,
            alt: `Gallery Image ${imageName}`
          }
        })
        images.sort((a, b) => a.alt.localeCompare(b.alt))
        galleryImages.value = images
        error.value = ''
      } catch (err) {
        console.error('Error loading gallery images:', err)
        error.value = 'Failed to load gallery images'
        galleryImages.value = []
      } finally {
        loading.value = false
      }
    }

    const marqueeImages = computed(() => {
      const source = galleryImages.value.length ? galleryImages.value : fallbackBannerImages
      const base = source.slice(0, Math.min(source.length, 8))
      return base.length ? [...base, ...base] : []
    })

    onMounted(() => {
      fetchGalleryImages()
    })

    return {
      aboutHeroImage,
      storyImage,
      highlightSections,
      loading,
      error,
      marqueeImages
    }
  }
}
</script>

<style scoped>
.marquee-mask {
  overflow: hidden;
  mask-image: linear-gradient(to right, transparent, black 10%, black 90%, transparent);
  -webkit-mask-image: linear-gradient(to right, transparent, black 10%, black 90%, transparent);
}

.marquee-track {
  display: flex;
  width: max-content;
  gap: 1.5rem;
  animation: marquee 45s linear infinite;
}

.marquee-item {
  flex: 0 0 auto;
}

@keyframes marquee {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-50%);
  }
}
</style>
