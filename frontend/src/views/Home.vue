<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-8">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Welcome to Calvary Chapel Phnom Penh</h1>
        <p class="text-xl md:text-2xl mb-8">Statement of Faith</p>
        <p class="text-lg mb-8 max-w-2xl mx-auto">    
WE BELIEVE that there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.
        </p>
        <div class="space-x-4">
          <router-link to="/about" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Learn More
          </router-link>
          <router-link to="/contact" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Get Involved
          </router-link>
        </div>
      </div>
    </div>

    <!-- Services Section -->
    <div class="max-w-6xl mx-auto mb-12">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Our Services</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="card text-center">
          <div class="text-4xl mb-4">⛪</div>
          <h3 class="text-xl font-semibold mb-3">Sunday Worship</h3>
          <p class="text-gray-600">Join us every Sunday at 9:00 AM for inspiring worship and fellowship.</p>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">📚</div>
          <h3 class="text-xl font-semibold mb-3">Bible Study</h3>
          <p class="text-gray-600">Deepen your faith through our weekly Bible study sessions.</p>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">🤝</div>
          <h3 class="text-xl font-semibold mb-3">Community Outreach</h3>
          <p class="text-gray-600">Making a positive impact in our local community through service.</p>
        </div>
      </div>
    </div>

    <!-- Mission Section -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center px-4">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">Our Mission</h2>
        <p class="text-lg text-gray-700 mb-8">
          To spread the Gospel of Jesus Christ, build a strong community of believers, and serve others with love and compassion.
        </p>
        <router-link to="/missionary" class="btn-primary">
          Learn About Our Mission Work
        </router-link>
      </div>
    </div>

    <!-- Quick Links -->
    <div class="max-w-4xl mx-auto">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Quick Links</h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <router-link to="/about" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">👥</div>
          <span class="font-medium">About Us</span>
        </router-link>
        
        <router-link to="/missionary" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">🌍</div>
          <span class="font-medium">Missionary</span>
        </router-link>
        
        <router-link to="/contact" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">📞</div>
          <span class="font-medium">Contact</span>
        </router-link>
        
        <router-link to="/donate" class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-2xl mb-2">💝</div>
          <span class="font-medium">Donate</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home'
}
</script>
