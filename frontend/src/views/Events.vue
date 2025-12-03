<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[50vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#101c33] via-[#172b4a] to-[#1f3d61] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 20% 30%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 12%, rgba(255,255,255,0.18), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.2), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Events</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          Weekly rhythms and special moments that keep Calvary Chapel Phnom Penh connected to Jesus and each other.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Worship & Word
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Prayer
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Youth & Families
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            City Serve
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <!-- Navigation Links -->
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/events/upcoming-events"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            Upcoming Events
          </router-link>
          <router-link
            to="/events/calendar"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-blue/10 text-brand-blue font-semibold rounded-full border border-brand-blue/20 hover:bg-brand-blue/20 transition-colors"
          >
            View Calendar
          </router-link>
        </div>

        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">What happens each week?</h2>
          <div class="space-y-4 text-left text-lg text-gray-700 leading-relaxed max-w-3xl mx-auto">
            <p>
              Our weekly rhythm keeps us rooted in <span class="font-semibold text-brand-orange">Scripture, prayer, and community</span>. Gatherings happen in the sanctuary, homes, and neighborhoods across Phnom Penh.
            </p>
            <p>
              Sundays and Wednesdays anchor the church, while smaller circles—youth, prayer, and house churches—carry grace throughout the week.
            </p>
          </div>
        </section>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Sunday</p>
            <h3 class="text-2xl font-semibold text-gray-900">9:00 AM Worship</h3>
            <p class="text-gray-600">
              Verse-by-verse teaching, bilingual worship, kids ministry, and hospitality every Sunday morning.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Wednesday</p>
            <h3 class="text-2xl font-semibold text-gray-900">Bible study & prayer</h3>
            <p class="text-gray-600">
              7:00 PM midweek gathering focused on teaching, table discussion, and intercession for Phnom Penh.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Prayer</p>
            <h3 class="text-2xl font-semibold text-gray-900">City prayer nights</h3>
            <p class="text-gray-600">
              Monthly prayer for revival, leaders, and neighborhoods rotating through homes and the sanctuary.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Outreach</p>
            <h3 class="text-2xl font-semibold text-gray-900">Serve Phnom Penh</h3>
            <p class="text-gray-600">
              Monthly outreaches bring meals, medical care, and the Gospel into schools, villages, and communities.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Youth</p>
            <h3 class="text-2xl font-semibold text-gray-900">Fridays 6:30 PM</h3>
            <p class="text-gray-600">
              Students gather for worship, teaching, and serving together—open to ages 12–18.
            </p>
          </article>
          <article class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">House churches</p>
            <h3 class="text-2xl font-semibold text-gray-900">Grace communities</h3>
            <p class="text-gray-600">
              Mid-week meals, communion, and discipleship circles throughout the city.
            </p>
          </article>
        </div>

        <div class="text-center text-gray-700 text-base">
          For detailed schedules and new events, check the calendar or contact us—new opportunities are added regularly.
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'Events',
};
</script>

