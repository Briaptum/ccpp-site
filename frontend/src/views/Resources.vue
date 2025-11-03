<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Resources</h1>
        <p class="text-xl md:text-2xl">Tools and Materials for Your Spiritual Journey</p>
      </div>
    </div>

    <!-- Resources Overview -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">Welcome to Our Resource Center</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            We've compiled helpful resources to support you on your spiritual journey. Whether you're 
            looking for Bible study materials, sermon recordings, prayer guides, or information about 
            our programs, you'll find it here.
          </p>
          <p class="text-gray-700 mb-6">
            Our resources are designed to help you grow in your relationship with God, understand 
            Scripture better, and connect with our church community. Explore the sections below to 
            find what you're looking for.
          </p>
        </div>
      </div>
    </div>

    <!-- Resources Sections -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-12">Available Resources</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="card">
            <div class="text-4xl mb-4">📖</div>
            <h3 class="text-xl font-semibold mb-3">Bible Study Materials</h3>
            <p class="text-gray-600 mb-4">
              Study guides, commentaries, and resources to help you dive deeper into God's Word.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Weekly study guides</li>
              <li>• Bible reading plans</li>
              <li>• Devotional materials</li>
            </ul>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">🎵</div>
            <h3 class="text-xl font-semibold mb-3">Sermon Recordings</h3>
            <p class="text-gray-600 mb-4">
              Access recordings of our weekly messages and teaching series.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Audio recordings</li>
              <li>• Video messages</li>
              <li>• Message transcripts</li>
            </ul>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">🙏</div>
            <h3 class="text-xl font-semibold mb-3">Prayer Resources</h3>
            <p class="text-gray-600 mb-4">
              Prayer guides, intercession lists, and materials to strengthen your prayer life.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Prayer guides</li>
              <li>• Intercession lists</li>
              <li>• Prayer request forms</li>
            </ul>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">📚</div>
            <h3 class="text-xl font-semibold mb-3">Library Resources</h3>
            <p class="text-gray-600 mb-4">
              Books, reference materials, and educational resources available in our church library.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Christian books</li>
              <li>• Reference materials</li>
              <li>• Educational resources</li>
            </ul>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">👨‍👩‍👧‍👦</div>
            <h3 class="text-xl font-semibold mb-3">Family Resources</h3>
            <p class="text-gray-600 mb-4">
              Materials to help families grow together in faith and navigate life's challenges.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Family devotionals</li>
              <li>• Parenting resources</li>
              <li>• Children's materials</li>
            </ul>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">💡</div>
            <h3 class="text-xl font-semibold mb-3">New Believer Resources</h3>
            <p class="text-gray-600 mb-4">
              Getting started guides and foundational materials for those new to the faith.
            </p>
            <ul class="text-sm text-gray-600 space-y-2">
              <li>• Getting started guide</li>
              <li>• Foundation classes</li>
              <li>• Discipleship materials</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Links Section -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-8">Quick Links</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <router-link to="/contact" class="flex items-center p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            <div class="text-3xl mr-4">📧</div>
            <div>
              <h3 class="font-semibold text-gray-900 mb-1">Request Resources</h3>
              <p class="text-sm text-gray-600">Contact us for specific materials or information</p>
            </div>
          </router-link>
          
          <router-link to="/events" class="flex items-center p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            <div class="text-3xl mr-4">📅</div>
            <div>
              <h3 class="font-semibold text-gray-900 mb-1">Upcoming Events</h3>
              <p class="text-sm text-gray-600">See our calendar of classes and workshops</p>
            </div>
          </router-link>
          
          <router-link to="/ministries" class="flex items-center p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            <div class="text-3xl mr-4">👥</div>
            <div>
              <h3 class="font-semibold text-gray-900 mb-1">Join a Ministry</h3>
              <p class="text-sm text-gray-600">Get connected and find resources through our ministries</p>
            </div>
          </router-link>
          
          <router-link to="/about" class="flex items-center p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
            <div class="text-3xl mr-4">ℹ️</div>
            <div>
              <h3 class="font-semibold text-gray-900 mb-1">Learn More</h3>
              <p class="text-sm text-gray-600">Discover more about our church and beliefs</p>
            </div>
          </router-link>
        </div>
      </div>
    </div>

    <!-- Contact Section -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Need Help Finding Resources?</h2>
        <p class="text-xl mb-8">
          Our team is here to help you find what you need for your spiritual growth.
        </p>
        <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
          Contact Us
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Resources'
}
</script>

