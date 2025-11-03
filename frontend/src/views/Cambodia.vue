<template>
  <div class="px-4 sm:px-0">
    <!-- Main Content -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <img 
          src="/src/assets/images/logo3.png" 
          alt="Calvary Chapel Phnom Penh Logo" 
          class="h-20 md:h-20 w-auto object-contain mb-6"
        />
        <h2 class="text-3xl text-light text-gray-900 mb-6">Cambodia</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Cambodia, a beautiful Southeast Asian nation, is home to over 16 million people. The country has a rich cultural heritage and a resilient spirit, having overcome tremendous challenges throughout its history.
          </p>
          <p class="text-gray-700 mb-6">
            Located in Phnom Penh, the capital city, Calvary Chapel Phnom Penh serves a diverse community of both Cambodians and expatriates. We are committed to being a light in this nation and sharing the hope of the Gospel with all people.
          </p>
          <h3 class="text-2xl text-light text-gray-900 mb-4 mt-8">Our Heart for Cambodia</h3>
          <p class="text-gray-700 mb-6">
            We love Cambodia and its people. Our ministry is focused on:
          </p>
          <ul class="text-gray-700 mb-6 list-disc list-inside space-y-2">
            <li>Sharing the Gospel with Cambodians in their heart language</li>
            <li>Supporting local communities through various outreach programs</li>
            <li>Training and equipping Cambodian believers for ministry</li>
            <li>Partnering with local ministries and organizations</li>
            <li>Praying for spiritual awakening across the nation</li>
          </ul>
          <h3 class="text-2xl text-light text-gray-900 mb-4 mt-8">Why Cambodia?</h3>
          <p class="text-gray-700 mb-6">
            Cambodia is a nation open to the Gospel. The people are friendly, welcoming, and hungry for truth. As we serve here, we see God moving in powerful ways, transforming lives and communities.
          </p>
          <p class="text-gray-700 mb-6">
            Whether you're a Cambodian national or an expatriate living here, we invite you to join us as we seek to be a blessing to this nation and share the love of Jesus Christ with all who call Cambodia home.
          </p>
        </div>
      </div>
    </div>

    <!-- Cambodia Facts -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-4xl mx-auto">
        <h2 class="text-3xl text-light text-center text-gray-900 mb-8">About Cambodia</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">🇰🇭</div>
            <h3 class="text-xl text-light mb-3">Kingdom of Cambodia</h3>
            <p class="text-gray-600">Rich in history and culture</p>
          </div>
          
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">👥</div>
            <h3 class="text-xl text-light mb-3">16+ Million People</h3>
            <p class="text-gray-600">Diverse and welcoming population</p>
          </div>
          
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">🏛️</div>
            <h3 class="text-xl text-light mb-3">Phnom Penh</h3>
            <p class="text-gray-600">Capital city where we serve</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl text-light mb-4">Serving Cambodia Together</h2>
        <p class="text-xl mb-8">
          Join us as we serve and share God's love in this beautiful nation.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Get Involved
          </router-link>
          <router-link to="/how-to-help" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            How to Help
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Cambodia'
}
</script>

