<template>
  <div>


    <!-- Hero Section -->
    <section class="py-24 bg-gradient-to-br from-brand-blue via-brand-blue/80 to-brand-orange/90 text-white border-b border-white/10">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.4em] text-white/70">Leadership</p>
        <h1 class="text-2xl md:text-3xl lg:text-4xl font-bold leading-tight">
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
        <div class="grid gap-10 md:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="group relative overflow-hidden rounded-[2.3rem] border border-white/15 bg-white/10 shadow-[0_20px_50px_rgba(15,23,42,0.18)] backdrop-blur-xl p-8 text-center flex flex-col items-center"
          >
            <div class="pointer-events-none absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-500">
              <div class="absolute inset-0 bg-gradient-to-br from-brand-blue/30 via-transparent to-brand-orange/40"></div>
            </div>
            <div class="relative flex flex-col items-center flex-1">
              <div class="relative mb-6">
                <div class="absolute inset-0 rounded-full bg-brand-orange/40 blur-3xl opacity-0 group-hover:opacity-60 transition-opacity duration-500"></div>
                <div class="mx-auto h-28 w-28 overflow-hidden rounded-full border-4 border-white/80 shadow-[0_10px_35px_rgba(15,23,42,0.3)]">
                  <img
                    :src="leader.photo"
                    :alt="`${leader.name} portrait`"
                    class="h-full w-full object-cover"
                  />
                </div>
              </div>
              <h3 class="text-xl font-semibold text-gray-900">{{ leader.name }}</h3>
              <p class="mt-2 text-brand-blue font-medium tracking-[0.3em] uppercase text-[0.65rem]">{{ leader.role }}</p>
              <p class="mt-3 text-sm text-gray-700 leading-relaxed max-h-32 overflow-hidden">
                {{ leader.bio }}
              </p>
            </div>
            <button
              type="button"
              class="relative mt-6 inline-flex items-center gap-2 rounded-full border border-brand-blue/40 px-5 py-2 text-[0.65rem] font-semibold tracking-[0.3em] uppercase text-brand-blue transition-all duration-300 hover:bg-brand-blue hover:text-white"
            >
              Know More
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </button>
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

