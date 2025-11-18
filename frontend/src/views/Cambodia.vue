<template>
  <div>
    <!-- Content Section -->
    <div class="bg-background py-32 md:py-40">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-center text-gray-900">Cambodia</h1>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          We are honored to serve in Phnom Penh—standing with a resilient nation rich in culture, history, and promise.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Cambodia, a beautiful Southeast Asian nation, is home to over 16 million people. The country has a rich cultural heritage and a resilient spirit, having overcome tremendous challenges throughout its history.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Located in Phnom Penh, the capital city, Calvary Chapel Phnom Penh serves a diverse community of Cambodians and expatriates. We are committed to being a light in this nation and sharing the hope of the Gospel with all people.
          </p>
          <p class="text-lg text-gray-700 mb-4">Our ministry in Cambodia is focused on:</p>
          <ul class="space-y-3 text-gray-700 mb-6">
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Sharing the Gospel with Cambodians in their heart language</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Supporting local communities through outreaches and practical care</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Training and equipping Cambodian believers for ministry</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Partnering with local ministries and organizations</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span>Praying for spiritual awakening across the nation</span>
            </li>
          </ul>
          <p class="text-lg text-gray-700 mb-8">
            Cambodia is a nation open to the Gospel. The people are friendly, welcoming, and hungry for truth. Whether you're a Cambodian national or an expatriate living here, we invite you to join us as we share the love of Jesus Christ with all who call Cambodia home.
          </p>
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
  name: 'Cambodia',
};
</script>

