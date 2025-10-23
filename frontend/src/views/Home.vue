<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section with Background Image -->
    <div class="relative bg-gradient-to-r from-primary-600 to-primary-800 text-white py-24 mb-8 overflow-hidden">
      <!-- Background Image -->
      <div 
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="{ backgroundImage: `url(${backgroundImage})` }"
      >
        <!-- Dark overlay for better text readability -->
        <div class="absolute inset-0 bg-black bg-opacity-50"></div>
      </div>
      
      <!-- Content -->
      <div class="relative z-10 max-w-4xl mx-auto text-center px-4">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-lg">Welcome to Calvary Chapel Phnom Penh</h1>
        <p class="text-xl md:text-2xl mb-8 drop-shadow-md">Statement of Faith</p>
        <p class="text-lg mb-8 max-w-2xl mx-auto drop-shadow-md">    
WE BELIEVE that there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.
        </p>
        <div class="space-x-4">
          <router-link to="/about" class="btn-primary bg-white text-primary-600 hover:bg-gray-100 shadow-lg">
            Learn More
          </router-link>
          <router-link to="/contact" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600 shadow-lg">
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

    <!-- Services Section -->
    <div class="max-w-6xl mx-auto mb-12">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Please join us for:</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="card text-center">
          <div class="text-4xl mb-4">⛪</div>
          <h3 class="text-xl font-semibold mb-3">Sunday Worship</h3>
          <p class="text-gray-600">Join us every Sunday at 9:00 AM for inspiring worship and fellowship.</p>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">📚</div>
          <h3 class="text-xl font-semibold mb-3">Wednesday Bible Study</h3>
          <p class="text-gray-600">Wednesday night Bible Study at 7:00 p.m. 

</p>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">🤝</div>
          <h3 class="text-xl font-semibold mb-3">Children Program</h3>
          <p class="text-gray-600">​
Sunday School is provided for children under the age of 12. <br>
Don't forget to bring a friend!</p>
        </div>
      </div>
    </div>

    <!-- Picture Gallery Section -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Picture Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-8 max-w-2xl mx-auto">
          Glimpses of our community, worship services, and ministry events
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

        <!-- Gallery Grid -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-8">
          <div 
            v-for="(image, index) in galleryImages" 
            :key="index"
            @click="openLightbox(index)"
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
            <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black to-transparent p-4">
              <p class="text-white text-sm font-medium">{{ image.title }}</p>
            </div>
          </div>
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
        <p class="text-white text-center mt-4 text-lg">{{ galleryImages[currentImageIndex].title }}</p>
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

    <!-- Quick Links -->
    <div class="max-w-4xl mx-auto">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Quick Links</h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <router-link to="/about" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">👥</div>
          <span class="font-medium">About Us</span>
        </router-link>
        
        <router-link to="/missionary" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">🌍</div>
          <span class="font-medium">Missionary</span>
        </router-link>
        
        <router-link to="/contact" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">📞</div>
          <span class="font-medium">Contact</span>
        </router-link>
        
        <router-link to="/donate" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">💝</div>
          <span class="font-medium">Donate</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import backgroundImage from '@/assets/images/background.jpg'

export default {
  name: 'Home',
  setup() {
    const lightboxOpen = ref(false)
    const currentImageIndex = ref(0)
    const loading = ref(true)
    const error = ref('')
    
    // Gallery images
    const galleryImages = ref([])

    // Load static gallery images
    const fetchGalleryImages = async () => {
      loading.value = true
      
      // Use static placeholder images
      galleryImages.value = [
        {
          src: 'https://images.unsplash.com/photo-1438032005730-c779502df39b?w=800',
          alt: 'Sunday Worship Service',
          title: 'Sunday Worship Service'
        },
        {
          src: 'https://images.unsplash.com/photo-1511632765486-a01980e01a18?w=800',
          alt: 'Community Fellowship',
          title: 'Community Fellowship'
        },
        {
          src: 'https://images.unsplash.com/photo-1507692049790-de58290a4334?w=800',
          alt: 'Bible Study Group',
          title: 'Bible Study Group'
        },
        {
          src: 'https://images.unsplash.com/photo-1529070538774-1843cb3265df?w=800',
          alt: 'Youth Ministry',
          title: 'Youth Ministry'
        },
        {
          src: 'https://images.unsplash.com/photo-1504052434569-70ad5836ab65?w=800',
          alt: 'Community Outreach',
          title: 'Community Outreach'
        },
        {
          src: 'https://images.unsplash.com/photo-1519834785169-98be25ec3f84?w=800',
          alt: 'Prayer Meeting',
          title: 'Prayer Meeting'
        },
        {
          src: 'https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=800',
          alt: 'Church Community',
          title: 'Church Community'
        },
        {
          src: 'https://images.unsplash.com/photo-1556075798-4825dfaaf498?w=800',
          alt: 'Sunday School',
          title: 'Sunday School'
        },
        {
          src: 'https://images.unsplash.com/photo-1511593358241-7eea1f3c84e5?w=800',
          alt: 'Worship Music',
          title: 'Worship Music'
        }
      ]
      
      loading.value = false
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

    // Fetch images when component mounts
    onMounted(() => {
      fetchGalleryImages()
    })

    return {
      backgroundImage,
      lightboxOpen,
      currentImageIndex,
      loading,
      error,
      galleryImages,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage
    }
  }
}
</script>
