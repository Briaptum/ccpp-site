<template>
  <div>
    <!-- Content Section -->
    <div class="bg-background py-32 md:py-40">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-center text-gray-900">Ministries</h1>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Every ministry is designed to meet people where they are and help them grow in their relationship with Jesus.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            At Calvary Chapel Phnom Penh, ministry is about serving others and sharing the love of Christ. From the youngest to the oldest, we are committed to creating spaces where people encounter Jesus and are formed into disciples.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Whether you're looking to get involved, receive support, or grow in your faith, we have ministries that can help you on your journey. Each ministry is committed to serving our community with excellence, humility, and love.
          </p>
          <p class="text-lg text-gray-700 mb-4">We offer various ministries including:</p>
          <ul class="space-y-3 text-gray-700 mb-8">
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Children's Ministry:</strong> Nurturing young hearts with age-appropriate Bible teaching, worship, and joyful community (Ages 4-12)</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Worship Ministry:</strong> Leading our church in heartfelt praise through music, prayer, and a posture of surrender</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Bible Study:</strong> Weekly gatherings to explore Scripture deeply and apply it in community (Wednesday Evenings)</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Prayer Ministry:</strong> Standing in the gap for our church family, our city, and the nations through intercession</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Outreach Ministry:</strong> Extending God's love beyond our walls through service, evangelism, and partnerships</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Family Ministry:</strong> Resources and gatherings that strengthen families to walk with Jesus together</span>
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
  name: 'Ministries',
};
</script>

