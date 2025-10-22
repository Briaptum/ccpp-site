<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">About CCPP</h1>
        <p class="text-xl md:text-2xl">Christ-Centered People's Place</p>
      </div>
    </div>

    <!-- Our Story Section -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">Our Story</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Founded in faith and built on love, Christ-Centered People's Place (CCPP) has been serving our community 
            for over a decade. We began as a small group of believers who wanted to create a place where people could 
            experience the transforming power of God's love.
          </p>
          <p class="text-gray-700 mb-6">
            Our journey started with a simple vision: to be a beacon of hope in our community, offering spiritual 
            guidance, fellowship, and practical support to all who seek it. Today, we continue to grow and serve, 
            always staying true to our mission of spreading the Gospel and building God's kingdom on earth.
          </p>
        </div>
      </div>
    </div>

    <!-- Our Values Section -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-12">Our Values</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="card text-center">
            <div class="text-4xl mb-4">❤️</div>
            <h3 class="text-xl font-semibold mb-3">Love</h3>
            <p class="text-gray-600">We believe in the power of unconditional love as demonstrated by Christ.</p>
          </div>
          
          <div class="card text-center">
            <div class="text-4xl mb-4">🤝</div>
            <h3 class="text-xl font-semibold mb-3">Community</h3>
            <p class="text-gray-600">Building strong relationships and supporting one another in faith.</p>
          </div>
          
          <div class="card text-center">
            <div class="text-4xl mb-4">🙏</div>
            <h3 class="text-xl font-semibold mb-3">Prayer</h3>
            <p class="text-gray-600">Prayer is the foundation of our relationship with God and each other.</p>
          </div>
          
          <div class="card text-center">
            <div class="text-4xl mb-4">📖</div>
            <h3 class="text-xl font-semibold mb-3">Scripture</h3>
            <p class="text-gray-600">God's Word guides all our decisions and actions.</p>
          </div>
          
          <div class="card text-center">
            <div class="text-4xl mb-4">🎵</div>
            <h3 class="text-xl font-semibold mb-3">Worship</h3>
            <p class="text-gray-600">We worship God with our whole hearts in spirit and truth.</p>
          </div>
          
          <div class="card text-center">
            <div class="text-4xl mb-4">🌱</div>
            <h3 class="text-xl font-semibold mb-3">Growth</h3>
            <p class="text-gray-600">We encourage spiritual growth and personal development.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Leadership Section -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-8">Our Leadership</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👨‍💼</span>
            </div>
            <h3 class="text-xl font-semibold mb-2">Pastor John Smith</h3>
            <p class="text-primary-600 font-medium mb-3">Senior Pastor</p>
            <p class="text-gray-600">
              With over 15 years of pastoral experience, Pastor John leads our congregation with wisdom and compassion.
            </p>
          </div>
          
          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👩‍💼</span>
            </div>
            <h3 class="text-xl font-semibold mb-2">Sarah Johnson</h3>
            <p class="text-primary-600 font-medium mb-3">Associate Pastor</p>
            <p class="text-gray-600">
              Sarah brings energy and innovation to our youth and community outreach programs.
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Join Our Family</h2>
        <p class="text-xl mb-8">
          We'd love to welcome you into our community of faith and fellowship.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/donate" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Support Our Mission
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>
