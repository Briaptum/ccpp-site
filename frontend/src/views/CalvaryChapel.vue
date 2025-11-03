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
        <h2 class="text-3xl text-light text-gray-900 mb-6">Calvary Chapel</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            Calvary Chapel is a fellowship of churches that began in 1965 with Pastor Chuck Smith in Costa Mesa, California. What started as a small church of 25 people has grown into a movement of thousands of churches worldwide.
          </p>
          <p class="text-gray-700 mb-6">
            The distinctive characteristics of Calvary Chapel include:
          </p>
          <ul class="text-gray-700 mb-6 list-disc list-inside space-y-2">
            <li><strong>Verse-by-verse teaching:</strong> We go through the Bible systematically, teaching verse by verse, chapter by chapter, book by book</li>
            <li><strong>Balance:</strong> We emphasize both the Word and the Spirit, avoiding extremes</li>
            <li><strong>Contemporary worship:</strong> We worship with contemporary music while maintaining reverence for God</li>
            <li><strong>Autonomous churches:</strong> Each Calvary Chapel is independently governed while maintaining fellowship with others in the movement</li>
            <li><strong>Grace-centered:</strong> We emphasize God's grace and love, not legalism</li>
          </ul>
          <p class="text-gray-700 mb-6">
            Calvary Chapel Phnom Penh is part of this global movement, bringing the same commitment to faithful Bible teaching and authentic worship to the people of Cambodia.
          </p>
          <p class="text-gray-700 mb-6">
            Our approach to ministry is characterized by simplicity, sincerity, and a focus on what matters most: knowing Jesus Christ and making Him known. We don't get caught up in complicated programs or structures—we simply want to love God, love people, and teach the Bible.
          </p>
        </div>
      </div>
    </div>

    <!-- Calvary Chapel Values -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-4xl mx-auto">
        <h2 class="text-3xl text-light text-center text-gray-900 mb-8">Calvary Chapel Distinctives</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">📖</div>
            <h3 class="text-xl text-light mb-3">Bible Teaching</h3>
            <p class="text-gray-600">Systematic, verse-by-verse exposition of Scripture</p>
          </div>
          
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">💝</div>
            <h3 class="text-xl text-light mb-3">Grace</h3>
            <p class="text-gray-600">Emphasizing God's unconditional love and grace</p>
          </div>
          
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">⚖️</div>
            <h3 class="text-xl text-light mb-3">Balance</h3>
            <p class="text-gray-600">Maintaining balance between Word and Spirit</p>
          </div>
          
          <div class="card text-center bg-white">
            <div class="text-4xl mb-4">🌍</div>
            <h3 class="text-xl text-light mb-3">Fellowship</h3>
            <p class="text-gray-600">Part of a global movement of like-minded churches</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl text-light mb-4">Experience Calvary Chapel</h2>
        <p class="text-xl mb-8">
          Join us for worship and Bible teaching in the Calvary Chapel tradition.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/events" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Upcoming Events
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CalvaryChapel'
}
</script>

