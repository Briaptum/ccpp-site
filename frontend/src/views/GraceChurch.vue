<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Grace Church</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
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
              <span class="text-primary-600 mr-2">•</span>
              <span>Cultivating worship gatherings centered on Jesus</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Building gospel-centered communities across Cambodia</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Equipping believers to extend grace to their neighbors</span>
            </li>
          </ul>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-primary-600 text-white font-semibold rounded-lg hover:bg-secondary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'GraceChurch',
};
</script>
