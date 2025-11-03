<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Our Ministries</h1>
        <p class="text-xl md:text-2xl">Serving Our Community Through Faith and Love</p>
      </div>
    </div>

    <!-- Ministries Overview -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl font-bold text-gray-900 mb-6">Ministry Overview</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            At CCPP, we believe that ministry is about serving others and sharing the love of Christ. 
            Our various ministries are designed to meet people where they are and help them grow in 
            their relationship with God.
          </p>
          <p class="text-gray-700 mb-6">
            Whether you're looking to get involved, receive support, or grow in your faith, we have 
            ministries that can help you on your journey. Each ministry is committed to serving our 
            community with excellence and love.
          </p>
        </div>
      </div>
    </div>

    <!-- Ministries Grid -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-12">Our Ministries</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="card">
            <div class="text-4xl mb-4">👥</div>
            <h3 class="text-xl font-semibold mb-3">Children's Ministry</h3>
            <p class="text-gray-600 mb-4">
              Nurturing young hearts and minds through age-appropriate Bible teaching, worship, 
              and fun activities that help children grow in their faith.
            </p>
            <p class="text-sm text-primary-600 font-medium">Ages 4-12</p>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">🎵</div>
            <h3 class="text-xl font-semibold mb-3">Worship Ministry</h3>
            <p class="text-gray-600 mb-4">
              Leading our congregation in heartfelt worship through music, prayer, and praise, 
              creating an atmosphere where people can encounter God.
            </p>
            <p class="text-sm text-primary-600 font-medium">All Ages Welcome</p>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">📖</div>
            <h3 class="text-xl font-semibold mb-3">Bible Study</h3>
            <p class="text-gray-600 mb-4">
              Deep dive into God's Word through weekly Bible study groups that provide fellowship, 
              discussion, and practical application of Scripture.
            </p>
            <p class="text-sm text-primary-600 font-medium">Wednesday Evenings</p>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">🙏</div>
            <h3 class="text-xl font-semibold mb-3">Prayer Ministry</h3>
            <p class="text-gray-600 mb-4">
              Committed to praying for the needs of our church family and community, believing 
              that prayer moves the heart of God.
            </p>
            <p class="text-sm text-primary-600 font-medium">Ongoing Support</p>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">🌍</div>
            <h3 class="text-xl font-semibold mb-3">Outreach Ministry</h3>
            <p class="text-gray-600 mb-4">
              Extending God's love beyond our walls through community service, evangelism, and 
              partnerships with local organizations.
            </p>
            <p class="text-sm text-primary-600 font-medium">Monthly Events</p>
          </div>
          
          <div class="card">
            <div class="text-4xl mb-4">👨‍👩‍👧‍👦</div>
            <h3 class="text-xl font-semibold mb-3">Family Ministry</h3>
            <p class="text-gray-600 mb-4">
              Supporting and strengthening families through resources, events, and programs designed 
              to help families grow together in faith.
            </p>
            <p class="text-sm text-primary-600 font-medium">Family Events Monthly</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Get Involved Section -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Get Involved</h2>
        <p class="text-xl mb-8">
          We'd love to have you join us in serving our community through ministry.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/events" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            View Upcoming Events
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Ministries'
}
</script>

