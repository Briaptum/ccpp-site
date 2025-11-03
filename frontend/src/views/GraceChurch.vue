<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Grace Church</h1>
        <p class="text-xl md:text-2xl">A Community of Grace</p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">About Grace Church</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Grace Church represents our commitment to spreading the message of God's amazing grace throughout 
            Cambodia. We believe that grace is at the heart of the Gospel - the unmerited favor of God extended 
            to all who believe in Jesus Christ.
          </p>
          <p class="text-gray-700 mb-6">
            Through Grace Church, we seek to create communities where people experience God's grace, grow in 
            faith, and extend that same grace to others. Our gatherings focus on worship, fellowship, and 
            teaching from God's Word.
          </p>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Join Our Community</h2>
        <p class="text-xl mb-8">
          Experience God's grace with us.
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
  name: 'GraceChurch'
}
</script>

