<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">About Us</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="prose max-w-none">
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            <span class="font-bold text-primary-600">Calvary Chapel Phnom Penh</span> is a community of believers dedicated to serving God and sharing His love with the people of Cambodia. We are a part of the Calvary Chapel movement, which emphasizes verse-by-verse, chapter-by-chapter teaching through the Bible.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            Our church family is diverse, welcoming people from all walks of life. We believe in creating a warm, loving environment where everyone can grow in their relationship with Jesus Christ.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            Whether you're new to faith or have been walking with Christ for years, you're welcome here. We'd love to get to know you and have you join us as we seek to follow Jesus together.
          </p>
        </div>
        
        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-primary-600 text-white font-semibold rounded-lg hover:bg-secondary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base border-2 border-primary-600 text-primary-600 font-semibold rounded-lg hover:bg-main-400 transition-colors"
          >
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>
