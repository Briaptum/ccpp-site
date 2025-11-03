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
        <h2 class="text-3xl text-light text-gray-900 mb-6">Who We Are</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Calvary Chapel Phnom Penh is a community of believers dedicated to serving God and sharing His love with the people of Cambodia. We are a part of the Calvary Chapel movement, which emphasizes verse-by-verse, chapter-by-chapter teaching through the Bible.
          </p>
          <p class="text-gray-700 mb-6">
            Our church family is diverse, welcoming people from all walks of life. We believe in creating a warm, loving environment where everyone can grow in their relationship with Jesus Christ.
          </p>
          <p class="text-gray-700 mb-6">
            We are committed to:
          </p>
          <ul class="text-gray-700 mb-6 list-disc list-inside space-y-2">
            <li>Teaching the Word of God faithfully and accurately</li>
            <li>Fostering genuine fellowship and community</li>
            <li>Serving our local community in Phnom Penh and throughout Cambodia</li>
            <li>Making disciples who make disciples</li>
            <li>Sharing the Gospel with those who don't know Jesus</li>
          </ul>
          <p class="text-gray-700 mb-6">
            Whether you're new to faith or have been walking with Christ for years, you're welcome here. We'd love to get to know you and have you join us as we seek to follow Jesus together.
          </p>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl text-light mb-4">Join Our Family</h2>
        <p class="text-xl mb-8">
          We'd love to welcome you into our community of faith and fellowship.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/about" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Learn More
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'WhoWeAre'
}
</script>

