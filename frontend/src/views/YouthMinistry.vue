<template>
  <div>
    <!-- Content Section -->
    <div class="bg-background py-32 md:py-40">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-center text-gray-900">Youth Ministry</h1>
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
              <span class="text-primary-600 mr-2">•</span>
              <span>Interactive Bible teaching that speaks to everyday life</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Authentic worship led by and for students</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Service projects that put faith into action</span>
            </li>
          </ul>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-8 py-3 bg-primary-600 text-white font-semibold rounded-lg hover:bg-primary-700 transition-colors shadow-lg"
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
  name: 'YouthMinistry',
};
</script>

