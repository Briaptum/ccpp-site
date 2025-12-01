<template>
  <div>


    <!-- Hero Section -->
    <section class="py-24 bg-gradient-to-b from-brand-blue via-brand-blue/90 to-brand-blue/85 text-white border-b border-white/10">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.4em] text-white/70">Leadership</p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
          Shepherding Phnom Penh with prayer, Scripture, and presence.
        </h1>
        <p class="text-lg md:text-xl text-white/85 font-light max-w-3xl mx-auto">
          Our elders and ministry leads teach verse by verse, disciple across generations, and keep the
          church tethered to Jesus and His mission.
        </p>
      </div>
    </section>

    <!-- Leaders Section -->
    <section class="py-24 bg-gray-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid gap-10 sm:grid-cols-2 lg:grid-cols-3">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="rounded-3xl bg-white/80 border border-white/60 shadow-md backdrop-blur hover:shadow-2xl transition-shadow p-8 text-center flex flex-col"
          >
            <div class="mx-auto h-32 w-32 overflow-hidden rounded-full border-4 border-white shadow-lg mb-6">
              <img
                :src="leader.photo"
                :alt="`${leader.name} portrait`"
                class="h-full w-full object-cover"
              />
            </div>
            <h3 class="text-2xl font-semibold text-gray-900">{{ leader.name }}</h3>
            <p class="mt-2 text-brand-blue font-medium tracking-wide uppercase text-xs">{{ leader.role }}</p>
            <p class="mt-4 text-sm text-gray-600 leading-relaxed flex-1">
              {{ leader.bio }}
            </p>
          </article>
        </div>

        <div class="text-center mt-16">
          <p class="text-gray-600 mb-3 text-sm uppercase tracking-[0.3em]">Next steps</p>
          <h4 class="text-2xl font-semibold text-gray-900 mb-6">Connect with our leadership team</h4>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors shadow-md"
          >
            Contact Us
            <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </router-link>
        </div>
      </div>
    </section>

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
  setup() {
    const leaders = [
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
    ]

    return {
      leaders
    }
  }
}
</script>

