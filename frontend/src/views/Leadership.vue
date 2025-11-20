<template>
  <div>


    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Leadership</h1>
        <p class="text-base md:text-lg text-gray-700 mt-4 max-w-2xl mx-auto">
          Our leadership team is committed to serving God and shepherding our church family with wisdom, compassion, and faithful teaching of God's Word. Each leader brings unique gifts and a heart for ministry, working together to guide Calvary Chapel Phnom Penh in fulfilling our mission to share the Gospel and make disciples in Cambodia.
        </p>
      </div>
    </div>

    <!-- Leaders Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3 mb-12 justify-items-center">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="bg-primary rounded-lg p-6 text-center shadow-md"
          >
            <div class="mx-auto h-28 w-28 overflow-hidden rounded-full border-4 border-secondary shadow-lg mb-4">
              <img
                :src="leader.photo"
                :alt="`${leader.name} portrait`"
                class="h-full w-full object-cover"
              />
            </div>
            <h3 class="text-xl font-semibold text-gray-900">{{ leader.name }}</h3>
            <p class="mt-2 text-main font-medium">{{ leader.role }}</p>
            <p class="mt-4 text-sm text-gray-600">
              {{ leader.bio }}
            </p>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-main text-white font-semibold rounded-lg hover:opacity-90 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import pastorRithPhoto from '@/assets/leadership/rith.jpg'
import pastorLongPhoto from '@/assets/leadership/Paster Long.jpg'
import randyPhoto from '@/assets/leadership/randy.jpg'
import vuthaPhoto from '@/assets/leadership/longdee.jpg'
import chengPhoto from '@/assets/leadership/Paster Borin.jpg'

export default {
  name: 'Leadership',
  data() {
    return {
      leaders: [
        {
          name: 'Pastor Rith',
          role: 'Senior Pastor',
          photo: pastorRithPhoto,
          bio: "With over 15 years of pastoral experience, Pastor Rith leads our congregation with wisdom and compassion. He faithfully teaches God's Word and shepherds our church family.",
        },
        {
          name: 'Pastor Long',
          role: 'Associate Pastor',
          photo: pastorLongPhoto,
          bio: 'Pastor Long brings energy to youth and community outreach. He is passionate about discipleship and seeing lives transformed by the Gospel.',
        },
        {
          name: 'Randy',
          role: 'Leadership Team',
          photo: randyPhoto,
          bio: 'Randy supports the church family through faithful service and a commitment to caring for every ministry area.',
        },
        {
          name: 'Vutha',
          role: 'Leadership Team',
          photo: vuthaPhoto,
          bio: 'Vutha encourages meaningful worship and helps guide teams that cultivate fellowship and spiritual growth.',
        },
        {
          name: 'Cheng',
          role: 'Leadership Team',
          photo: chengPhoto,
          bio: 'Cheng invests in discipleship and helps coordinate opportunities for members to grow together in faith.',
        },
      ],
    }
  },
}
</script>

