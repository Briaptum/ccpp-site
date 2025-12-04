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
            to="/about"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors"
          >
            About Us
          </router-link>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-white/10 text-white font-semibold rounded-md border border-white/20 hover:bg-white/20 transition-colors backdrop-blur-sm"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

    <!-- Regular Study Times -->
    <section class="py-20 bg-gradient-to-b from-white to-gray-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400 mb-3">Weekly Rhythm</p>
          <h2 class="text-4xl font-bold text-gray-900 mb-4">Regular Study & Worship</h2>
          <p class="text-lg text-gray-600 max-w-3xl mx-auto">
            Two anchor gatherings each week keep us rooted in Scripture and connected as a family.
          </p>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="relative overflow-hidden rounded-3xl border border-gray-100 bg-white shadow-2xl">
            <div class="absolute inset-0 opacity-60 pointer-events-none" style="background: radial-gradient(circle at top left, rgba(255,164,91,0.18), transparent 50%), radial-gradient(circle at bottom right, rgba(34,112,130,0.15), transparent 45%);"></div>
            <div class="relative flex flex-col md:flex-row items-center md:items-center gap-8 p-8">
              <div class="flex flex-col items-center justify-center text-center bg-brand-blue/10 text-brand-blue rounded-2xl px-6 py-7 w-full md:w-40 shrink-0">
                <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/70">Sunday</span>
                <span class="text-4xl md:text-5xl font-bold text-gray-900 mt-2 leading-none">9:00</span>
                <span class="text-sm tracking-[0.3em] text-gray-500 mt-1">a.m.</span>
              </div>
              <div class="flex-1 text-gray-700">
                <div class="inline-flex items-center bg-brand-blue/10 text-brand-blue px-3 py-1 rounded-full uppercase tracking-[0.35em] text-xs font-semibold mb-4">
                  Sunday Worship
                </div>
                <h3 class="text-3xl font-semibold text-gray-900 mb-3">Sunday Morning Worship</h3>
                <p class="text-base leading-relaxed text-gray-600 mb-5">
                  Verse-by-verse teaching, Spirit-led worship, and community for every generation gathered in the sanctuary each Sunday morning at 9:00 a.m.
                </p>
                <router-link
                  to="/contact"
                  class="inline-flex items-center text-brand-blue font-semibold hover:text-brand-orange transition-colors"
                >
                  Contact Us
                  <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </router-link>
              </div>
            </div>
          </div>

          <div class="relative overflow-hidden rounded-3xl border border-gray-100 bg-white shadow-2xl">
            <div class="absolute inset-0 opacity-60 pointer-events-none" style="background: radial-gradient(circle at top right, rgba(34,112,130,0.2), transparent 50%), radial-gradient(circle at bottom left, rgba(255,164,91,0.15), transparent 40%);"></div>
            <div class="relative flex flex-col md:flex-row items-center md:items-center gap-8 p-8">
              <div class="flex flex-col items-center justify-center text-center bg-brand-orange/10 text-brand-orange rounded-2xl px-6 py-7 w-full md:w-40 shrink-0">
                <span class="text-xs uppercase tracking-[0.4em] text-brand-orange/80">Wednesday</span>
                <span class="text-4xl md:text-5xl font-bold text-gray-900 mt-2 leading-none">7:00</span>
                <span class="text-sm tracking-[0.3em] text-gray-500 mt-1">p.m.</span>
              </div>
              <div class="flex-1 text-gray-700">
                <div class="inline-flex items-center bg-brand-orange/10 text-brand-orange px-3 py-1 rounded-full uppercase tracking-[0.35em] text-xs font-semibold mb-4">
                  Wednesday Church
                </div>
                <h3 class="text-3xl font-semibold text-gray-900 mb-3">Wednesday Night Church</h3>
                <p class="text-base leading-relaxed text-gray-600 mb-5">
                  Wednesday night church family gathering featuring Scripture, discussion, and prayer beginning promptly at 7:00 p.m.
                </p>
                <router-link
                  to="/events/calendar"
                  class="inline-flex items-center text-brand-orange font-semibold hover:text-brand-blue transition-colors"
                >
                  Contact Us
                  <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Upcoming Section -->
    <div class="relative py-24 bg-brand-blue overflow-hidden">
      <div class="absolute inset-0 opacity-80" aria-hidden="true">
        <div class="absolute inset-y-0 left-0 w-1/2 bg-gradient-to-r from-black/40 via-transparent to-transparent"></div>
        <div class="absolute inset-y-0 right-0 w-2/3 bg-gradient-to-l from-brand-orange/20 via-transparent to-transparent"></div>
        <div class="absolute inset-x-0 top-0 h-1/2 bg-gradient-to-b from-white/15 via-transparent to-transparent"></div>
        <div class="absolute inset-x-0 bottom-0 h-56 bg-gradient-to-t from-black/30 via-brand-blue/20 to-transparent"></div>
        <div class="absolute inset-0 opacity-35 mix-blend-screen" style="background-image: radial-gradient(circle at 25% 20%, rgba(255,255,255,0.25), transparent 40%), radial-gradient(circle at 80% 0%, rgba(255,255,255,0.15), transparent 35%), radial-gradient(circle at 60% 85%, rgba(255,255,255,0.2), transparent 40%);"></div>
      </div>
      <div class="relative">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-white">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Upcoming</p>
          <h2 class="text-4xl md:text-5xl font-bold mb-4">What’s ahead this season</h2>
          <p class="text-white/80 max-w-2xl mx-auto">
            A glimpse at the moments, gatherings, and opportunities we’re preparing for together.
          </p>
        </div>

        <div class="max-w-2xl mx-auto space-y-6">
          <div class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]">
            <div class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8">
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2">Mar</span>
              <span class="text-5xl font-bold leading-none text-brand-blue">02</span>
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3">Sat</span>
            </div>
            <div class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8">
              <div class="flex-1">
                <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">City Prayer Night</h3>
                <p class="text-white/80 text-base leading-relaxed">
                  An evening of prayer and worship for Phnom Penh, hosted in the main sanctuary at 6:30 PM.
                </p>
              </div>
              <router-link
                to="/events"
                class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
              >
                More Details
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]">
            <div class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8">
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2">Mar</span>
              <span class="text-5xl font-bold leading-none text-brand-blue">10</span>
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3">Sun</span>
            </div>
            <div class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8">
              <div class="flex-1">
                <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">Foundations Class</h3>
                <p class="text-white/80 text-base leading-relaxed">
                  Four weeks exploring what we believe, how to grow, and where to serve. Sundays at 11:15 AM.
                </p>
              </div>
              <router-link
                to="/ministries"
                class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
              >
                More Details
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>

          <div class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]">
            <div class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8">
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2">Mar</span>
              <span class="text-5xl font-bold leading-none text-brand-blue">16</span>
              <span class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3">Sat</span>
            </div>
            <div class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8">
              <div class="flex-1">
                <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">Serve Phnom Penh</h3>
                <p class="text-white/80 text-base leading-relaxed">
                  Neighborhood outreach day providing meals, prayer, and practical care. Teams launch at 8:00 AM.
                </p>
              </div>
              <router-link
                to="/ministries/outreaches"
                class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
              >
                More Details
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div class="relative py-24 bg-gradient-to-br from-[#0B111E] via-[#112437] to-[#1C4053] text-white overflow-hidden">
      <div class="absolute inset-0 opacity-70" aria-hidden="true">
        <div class="absolute inset-0" style="background-image: radial-gradient(circle at 20% 15%, rgba(228,157,64,0.25), transparent 45%), radial-gradient(circle at 80% 10%, rgba(255,255,255,0.12), transparent 40%), radial-gradient(circle at 60% 85%, rgba(34,112,130,0.3), transparent 45%);"></div>
        <div class="absolute inset-0 mix-blend-soft-light opacity-40" style="background-image: linear-gradient(120deg, rgba(255,255,255,0.15) 0%, rgba(255,255,255,0) 45%), linear-gradient(300deg, rgba(34,112,130,0.25) 0%, rgba(255,255,255,0) 55%);"></div>
        <div class="absolute inset-0 opacity-20" style="background-image: linear-gradient(to right, rgba(255,255,255,0.06) 1px, transparent 1px), linear-gradient(to bottom, rgba(255,255,255,0.06) 1px, transparent 1px); background-size: 80px 80px;"></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">Latest Teachings</p>
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">Recent messages from our pulpit</h2>
          <p class="text-lg text-white/80 max-w-2xl mx-auto font-light">
            Teachings captured each week to keep you connected to the verse-by-verse journey through Scripture.
          </p>
        </div>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="teaching in latestTeachings"
            :key="teaching.id"
            class="bg-white/10 border border-white/15 rounded-2xl backdrop-blur-xl shadow-2xl hover:bg-white/15 transition-all duration-300 overflow-hidden flex flex-col"
          >
            <div class="relative h-48 overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover transition-transform duration-500 hover:scale-110"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/55 via-transparent to-transparent"></div>
              <div class="absolute bottom-4 left-4 inline-flex items-center px-3 py-1 bg-white text-gray-900 rounded-full text-xs font-semibold tracking-wide shadow">
                Teaching Series
              </div>
            </div>
            <div class="p-6 flex flex-col flex-1">
              <h3 class="text-lg font-semibold text-white mb-2 leading-snug line-clamp-2">{{ teaching.title }}</h3>
              <p class="text-sm text-white/70 mb-5 leading-relaxed line-clamp-3 flex-1">
                {{ teaching.description }}
              </p>
              <div class="flex items-center justify-between border-t border-white/20 pt-4">
                <span class="text-xs uppercase tracking-[0.3em] text-white/50">YouTube</span>
                <a
                  :href="teaching.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center text-sm font-semibold text-brand-orange hover:text-white transition-colors"
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
