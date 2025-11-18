<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Outreaches</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          We carry the love and message of Jesus into the communities of Cambodia, serving with compassion and sharing hope.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Our outreach programs are designed to take the love and message of Jesus Christ beyond our church walls into the communities of Cambodia. Serving others is a vital expression of our faith and a powerful way to share the Gospel.
          </p>
          <p class="text-lg text-gray-700 mb-12">
            Through these initiatives we partner with local organizations, serve those in need, and create opportunities for people to encounter God's love through practical acts of service.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
          <div class="bg-background-50 rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="araksatImage" 
                alt="Araksat Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Araksat</h2>
              <p class="text-gray-700">
                Our outreach to the Araksat community focuses on sharing the Gospel and serving the people with love and compassion.
              </p>
            </div>
          </div>
          
          <div class="bg-background-50 rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="steunmeancheyImage" 
                alt="Steunmeanchey Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Steunmeanchey</h2>
              <p class="text-gray-700">
                We serve the Steunmeanchey community through practical acts of service and sharing the hope of Jesus Christ.
              </p>
            </div>
          </div>
          
          <div class="bg-background-50 rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="preyVengImage" 
                alt="Prey Veng Outreach" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Prey Veng</h2>
              <p class="text-gray-700">
                Our Prey Veng outreach brings the love of Christ to this community through ministry, service, and discipleship.
              </p>
            </div>
          </div>
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
import araksatImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'
import steunmeancheyImage from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'
import preyVengImage from '@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg'

export default {
  name: 'Outreaches',
  data() {
    return {
      araksatImage,
      steunmeancheyImage,
      preyVengImage,
    };
  },
};
</script>

