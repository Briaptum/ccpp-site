<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Youth Ministry</h1>
        <p class="text-xl md:text-2xl">Empowering the Next Generation</p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">About Our Youth Ministry</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Our Youth Ministry is dedicated to helping teenagers and young adults discover their identity in Christ 
            and develop a deep, personal relationship with God. We believe in creating an environment where young 
            people can ask questions, explore faith, and build lasting friendships.
          </p>
          <p class="text-gray-700 mb-6">
            Our youth programs combine engaging teaching, authentic worship, fun activities, and meaningful service 
            opportunities to help young people grow in their walk with Jesus Christ.
          </p>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Join Our Youth</h2>
        <p class="text-xl mb-8">
          We'd love to have you be part of our youth community.
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
  name: 'YouthMinistry'
}
</script>

