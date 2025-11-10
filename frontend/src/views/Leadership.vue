<template>
  <div class="px-4 sm:px-0">
    <!-- Leadership Section -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <h2 class="text-3xl text-light text-gray-900 mb-8">Our Leadership</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👨‍💼</span>
            </div>
            <h3 class="text-xl text-light mb-2">Pastor Rith</h3>
            <p class="text-primary-600 mb-3">Senior Pastor</p>
            <p class="text-gray-600">
              With over 15 years of pastoral experience, Pastor Rith leads our congregation with wisdom and compassion. He is dedicated to teaching God's Word faithfully and shepherding our church family.
            </p>
          </div>
          
          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👩‍💼</span>
            </div>
            <h3 class="text-xl text-light mb-2">Pastor Long</h3>
            <p class="text-primary-600 mb-3">Associate Pastor</p>
            <p class="text-gray-600">
              Pastor Long brings energy and innovation to our youth and community outreach programs. He is passionate about discipleship and seeing lives transformed by the Gospel.
            </p>
          </div>

          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👨‍💼</span>
            </div>
            <h3 class="text-xl text-light mb-2">Randy</h3>
            <p class="text-primary-600 mb-3">Leadership Team</p>
            <p class="text-gray-600">
              Randy supports the church family through faithful service and a commitment to caring for every ministry area.
            </p>
          </div>

          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👨‍💼</span>
            </div>
            <h3 class="text-xl text-light mb-2">Vutha</h3>
            <p class="text-primary-600 mb-3">Leadership Team</p>
            <p class="text-gray-600">
              Vutha encourages meaningful worship and helps guide teams that cultivate fellowship and spiritual growth.
            </p>
          </div>

          <div class="text-center">
            <div class="w-32 h-32 bg-primary-200 rounded-full mx-auto mb-4 flex items-center justify-center">
              <span class="text-4xl">👩‍💼</span>
            </div>
            <h3 class="text-xl text-light mb-2">Cheng</h3>
            <p class="text-primary-600 mb-3">Leadership Team</p>
            <p class="text-gray-600">
              Cheng invests in discipleship and helps coordinate opportunities for members to grow together in faith.
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Leadership Philosophy -->
    <div class="bg-gray-100 py-12 mb-12">
      <div class="max-w-4xl mx-auto">
        <h2 class="text-3xl text-light text-center text-gray-900 mb-8">Our Leadership Philosophy</h2>
        <div class="card bg-white">
          <div class="prose prose-lg max-w-none">
            <p class="text-gray-700 mb-4">
              Our leadership team is committed to servant leadership, following the example of Jesus Christ who came not to be served, but to serve. We believe that leadership in the church is about:
            </p>
            <ul class="text-gray-700 list-disc list-inside space-y-2">
              <li>Humbly serving the body of Christ</li>
              <li>Faithfully teaching and applying God's Word</li>
              <li>Leading by example in love, integrity, and devotion</li>
              <li>Equipping others for ministry and service</li>
              <li>Praying for and caring for the congregation</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl text-light mb-4">Get in Touch</h2>
        <p class="text-xl mb-8">
          Have questions or want to connect with our leadership team?
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
  name: 'Leadership'
}
</script>

