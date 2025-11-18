<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Cambodia</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="prose max-w-none">
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            <span class="font-bold text-primary-600">Cambodia</span>, officially the Kingdom of Cambodia, is a Southeast Asian nation bordered by Thailand to the northwest, Laos to the north, Vietnam to the east, and the Gulf of Thailand to the southwest. Covering approximately 181,035 square kilometers, the country features a central plain dominated by the Mekong River and Tonlé Sap Lake, Southeast Asia's largest freshwater lake.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            As of 2024, Cambodia's population is estimated at around 17.6 million people, with the majority being ethnically Khmer. The capital and largest city is <span class="font-bold text-primary-600">Phnom Penh</span>, situated at the confluence of the Tonlé Sap and Mekong rivers, where Calvary Chapel Phnom Penh serves the community.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            Historically, Cambodia was home to the <span class="font-bold">Khmer Empire</span>, which flourished from the 9th to the 15th centuries and is renowned for constructing the magnificent Angkor Wat temple complex. In the 19th century, Cambodia became a French protectorate, gaining independence in 1953. The nation endured significant turmoil during the Khmer Rouge regime from 1975 to 1979, resulting in the deaths of an estimated 1.7 to 2.2 million people—a period that deeply shaped the nation's resilience and character.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            In recent years, Cambodia has experienced rapid economic growth, with the economy primarily driven by the textile industry, tourism, and agriculture. Tourism plays a vital role, with millions of international visitors drawn to the country's rich cultural heritage, including traditional arts, dance, and cuisine. The Khmer language is the official language, and Buddhism is the predominant religion.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            Cambodia's diverse ecosystems, including tropical forests and the Tonlé Sap Biosphere Reserve, support a wide array of wildlife and plant species. The country's rich cultural heritage and resilient spirit, having overcome tremendous challenges throughout its history, make it a nation of great promise and potential.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-4">Our ministry in Cambodia is focused on:</p>
          <ul class="space-y-3 text-base md:text-lg text-gray-900 mb-8">
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
          <p class="text-base md:text-lg text-gray-900 mb-8 leading-relaxed">
            Cambodia is a nation open to the Gospel. The people are friendly, welcoming, and hungry for truth. Whether you're a Cambodian national or an expatriate living here, we invite you to join us as we share the love of Jesus Christ with all who call Cambodia home.
          </p>
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
  name: 'Cambodia',
};
</script>

