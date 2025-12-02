<template>
  <div>
    <!-- Hero Section -->
    <div class="relative min-h-screen flex items-center overflow-hidden pt-24 md:pt-28">
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="{ backgroundImage: `linear-gradient(120deg, rgba(17, 39, 84, 0.25), rgba(17, 39, 84, 0.45)), url(${aboutHeroImage})` }"
        ></div>
        <div class="absolute inset-0 bg-gradient-to-br from-black/70 via-black/55 to-black/65"></div>
        <div class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-blue/35 via-brand-blue/15 to-transparent blur-[120px] opacity-60"></div>
        <div class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-orange/30 via-brand-orange/12 to-transparent blur-[180px] opacity-65"></div>
        <div class="absolute -top-24 left-1/2 w-[420px] h-[420px] -translate-x-1/2 bg-white/15 blur-[180px] opacity-60"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/80 via-black/60 to-transparent"></div>
        <div
          class="absolute inset-0 opacity-15 mix-blend-screen"
          style="background-image: radial-gradient(circle at 20% 28%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 10%, rgba(255,255,255,0.22), transparent 30%), radial-gradient(circle at 62% 78%, rgba(255,255,255,0.18), transparent 32%);"
        ></div>
      </div>
      <div class="relative z-10 w-full max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <p class="text-sm uppercase tracking-[0.4em] text-white/70 mb-4">Calvary Chapel Phnom Penh</p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight mb-6">
          A worshiping community rooted in Scripture and present for Phnom Penh.
        </h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light mb-10">
          We gather to teach verse by verse through the Bible, cultivate authentic relationships,
          and serve the city with the hope of Jesus.
        </p>
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/what-we-believe"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors"
          >
            What We Believe
          </router-link>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-white/10 text-white font-semibold rounded-md border border-white/20 hover:bg-white/20 transition-colors backdrop-blur-sm"
          >
            Plan a Visit
          </router-link>
        </div>
      </div>
    </div>

    <!-- Upcoming Section -->
    <div class="py-24 bg-brand-blue">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-white">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Upcoming</p>
          <h2 class="text-4xl md:text-5xl font-bold mb-4">What’s ahead this season</h2>
          <p class="text-white/80 max-w-2xl mx-auto">
            A glimpse at the moments, gatherings, and opportunities we’re preparing for together.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 2</p>
            <h3 class="text-2xl font-semibold mb-3">City Prayer Night</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              An evening of prayer and worship for Phnom Penh, hosted in the main sanctuary at 6:30 PM.
            </p>
            <router-link
              to="/events"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Details
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 10</p>
            <h3 class="text-2xl font-semibold mb-3">Foundations Class</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              Four weeks exploring what we believe, how to grow, and where to serve. Sundays at 11:15 AM.
            </p>
            <router-link
              to="/ministries"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Save a seat
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 p-6 backdrop-blur-md shadow-lg">
            <p class="text-xs uppercase tracking-[0.4em] text-white/60 mb-2">Mar 16</p>
            <h3 class="text-2xl font-semibold mb-3">Serve Phnom Penh</h3>
            <p class="text-white/80 text-sm leading-relaxed mb-5">
              Neighborhood outreach day providing meals, prayer, and practical care. Teams launch at 8:00 AM.
            </p>
            <router-link
              to="/ministries/outreaches"
              class="inline-flex items-center text-brand-orange font-semibold tracking-wide text-sm hover:text-white transition-colors"
            >
              Volunteer
              <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="py-24 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-gray-500 mb-3">Latest Teachings</p>
          <h2 class="text-4xl md:text-5xl font-bold text-gray-700 mb-4">Recent messages from our pulpit</h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto font-light">
            Teachings captured each week to keep you connected to the verse-by-verse journey through Scripture.
          </p>
        </div>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="teaching in latestTeachings"
            :key="teaching.id"
            class="bg-white rounded-2xl border border-gray-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden flex flex-col"
          >
            <div class="relative h-48 overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover transition-transform duration-500 hover:scale-105"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/40 via-transparent to-transparent"></div>
              <div class="absolute bottom-4 left-4 inline-flex items-center px-3 py-1 bg-white/90 rounded-full text-xs font-semibold tracking-wide text-gray-900 shadow">
                Teaching Series
              </div>
            </div>
            <div class="p-6 flex flex-col flex-1">
              <h3 class="text-lg font-semibold text-gray-800 mb-2 leading-snug line-clamp-2">{{ teaching.title }}</h3>
              <p class="text-sm text-gray-600 mb-5 leading-relaxed line-clamp-3 flex-1">
                {{ teaching.description }}
              </p>
              <div class="flex items-center justify-between border-t border-gray-100 pt-4">
                <span class="text-xs uppercase tracking-[0.3em] text-gray-500">YouTube</span>
                <a
                  :href="teaching.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center text-sm font-semibold text-brand-blue hover:text-gray-900 transition-colors"
                >
                  Watch now
                  <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import aboutHeroImage from '@/assets/background/about-bg.jpg'

export default {
  name: 'Home',
  setup() {
    const latestTeachings = [
      {
        id: 'IrTzMEBtbaI',
        title: 'Latest Teaching 1',
        description: 'Watch our latest message and be encouraged.',
        url: 'https://www.youtube.com/watch?v=IrTzMEBtbaI'
      },
      {
        id: 'AdDxynFrn8E',
        title: 'Latest Teaching 2',
        description: 'Join us for this week\'s teaching.',
        url: 'https://www.youtube.com/watch?v=AdDxynFrn8E'
      },
      {
        id: 'FC0Tm4bNPkk',
        title: 'Latest Teaching 3',
        description: 'Stay connected with our community through God\'s Word.',
        url: 'https://www.youtube.com/watch?v=FC0Tm4bNPkk'
      }
    ]

    return {
      aboutHeroImage,
      latestTeachings
    }
  }
}
</script>
