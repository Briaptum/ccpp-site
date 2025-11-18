<template>
  <div>
    <!-- Content Section -->
    <div class="bg-background py-32 md:py-40">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-center text-gray-900">Events</h1>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Join us for worship, fellowship, and community throughout the week.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            We gather regularly for worship, Bible study, prayer, and fellowship. Our regular events include:
          </p>
          <ul class="space-y-3 text-gray-700 mb-6">
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Sunday Worship Service:</strong> Every Sunday at 9:00 AM - Join us for worship, prayer, and teaching from God's Word</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Wednesday Bible Study:</strong> Every Wednesday at 7:00 PM - Deep dive into God's Word together</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Prayer Meetings:</strong> Regular gatherings for prayer and intercession</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Community Outreach:</strong> Monthly events to serve our community and share God's love</span>
            </li>
            <li class="flex items-start">
              <span class="text-primary-600 mr-2">•</span>
              <span><strong>Youth Group:</strong> Regular meetings for teens and young adults</span>
            </li>
          </ul>
          <p class="text-lg text-gray-700 mb-8">
            For specific dates and times of upcoming events, please contact us or check back regularly for updates.
          </p>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-8 py-3 bg-primary-600 text-white font-semibold rounded-lg hover:bg-primary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Events',
};
</script>

