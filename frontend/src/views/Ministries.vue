<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Ministries</h1>
        <p class="text-base md:text-lg text-gray-700 mt-4 max-w-2xl mx-auto">
          At Calvary Chapel Phnom Penh, ministry is about serving others and sharing the love of Christ. Every ministry is designed to meet people where they are and help them grow in their relationship with Jesus.
        </p>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="prose max-w-none">
          <p class="text-base md:text-lg text-gray-900 mb-6">
            Whether you're looking to get involved, receive support, or grow in your faith, we have ministries that can help you on your journey. Each ministry is committed to serving our community with excellence, humility, and love.
          </p>
          <p class="text-base md:text-lg text-gray-900 mb-4">We also offer:</p>
          <ul class="space-y-3 text-base md:text-lg text-gray-900 mb-8">
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span><strong>Children's Ministry:</strong> Nurturing young hearts with age-appropriate Bible teaching, worship, and joyful community (Ages 4-12)</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span><strong>Worship Ministry:</strong> Leading our church in heartfelt praise through music, prayer, and a posture of surrender</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span><strong>Bible Study:</strong> Weekly gatherings to explore Scripture deeply and apply it in community (Wednesday Evenings)</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span><strong>Prayer Ministry:</strong> Standing in the gap for our church family, our city, and the nations through intercession</span>
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span><strong>Family Ministry:</strong> Resources and gatherings that strengthen families to walk with Jesus together</span>
            </li>
          </ul>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
          <!-- Youth Ministry Box -->
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="youthMinistryImage" 
                alt="Youth Ministry" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Youth Ministry</h2>
              <p class="text-gray-700 mb-4">
                A vibrant community where teens encounter God, build lasting friendships, and discover their identity in Christ through engaging teaching, worship, and service.
              </p>
              <router-link
                to="/ministries/youth-ministry"
                class="inline-flex items-center text-main font-semibold hover:opacity-80 transition-colors"
              >
                Learn More →
              </router-link>
            </div>
          </div>
          
          <!-- Grace Church Box -->
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="graceChurchImage" 
                alt="Grace Church" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Grace Church</h2>
              <p class="text-gray-700 mb-4">
                We spread the message of God's amazing grace throughout Cambodia, gathering communities to worship, grow, and extend grace to others.
              </p>
              <router-link
                to="/ministries/grace-church"
                class="inline-flex items-center text-main font-semibold hover:opacity-80 transition-colors"
              >
                Learn More →
              </router-link>
            </div>
          </div>
          
          <!-- Outreaches Box -->
          <div class="bg-primary rounded-lg overflow-hidden shadow-md">
            <div class="w-full h-48 overflow-hidden">
              <img 
                :src="outreachesImage" 
                alt="Outreaches" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <h2 class="text-2xl font-semibold text-gray-900 mb-4">Outreaches</h2>
              <p class="text-gray-700 mb-4">
                We carry the love and message of Jesus into the communities of Cambodia, serving with compassion and sharing hope through practical acts of service.
              </p>
              <router-link
                to="/ministries/outreaches"
                class="inline-flex items-center text-main font-semibold hover:opacity-80 transition-colors"
              >
                Learn More →
              </router-link>
            </div>
          </div>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-custom-orange text-white font-semibold rounded-lg hover:bg-custom-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import youthMinistryImage from '@/assets/gallery/571314951_1275851217909869_307313164703676157_n.jpg'
import graceChurchImage from '@/assets/gallery/548188348_1236675751827416_5183731636046995004_n.jpg'
import outreachesImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'

export default {
  name: 'Ministries',
  data() {
    return {
      youthMinistryImage,
      graceChurchImage,
      outreachesImage,
    };
  },
};
</script>

