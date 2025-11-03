<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Outreaches</h1>
        <p class="text-xl md:text-2xl">Taking the Gospel Beyond Our Walls</p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">Our Outreach Ministries</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Our outreach programs are designed to take the love and message of Jesus Christ beyond our church 
            walls into the communities of Cambodia. We believe that serving others is a vital expression of 
            our faith and a powerful way to share the Gospel.
          </p>
          <p class="text-gray-700 mb-6">
            Through our various outreach initiatives, we partner with local organizations, serve those in need, 
            and create opportunities for people to encounter God's love through practical acts of service.
          </p>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Get Involved</h2>
        <p class="text-xl mb-8">
          Join us in reaching out to our community with God's love.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/ministries" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Back to Ministries
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Outreaches'
}
</script>

