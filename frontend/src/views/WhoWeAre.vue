<template>
  <FullPageShell>
    <div class="grid w-full gap-12 lg:grid-cols-[1.2fr,0.8fr] lg:items-start">
      <div>
        <img
          src="/src/assets/images/logo3.png"
          alt="Calvary Chapel Phnom Penh Logo"
          class="mx-auto h-24 w-auto drop-shadow-lg sm:mx-0 sm:h-28"
        />
        <p class="mt-6 text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Who We Are</p>
        <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
          A family following Jesus together in Phnom Penh.
        </h1>
        <div class="mt-8 space-y-5 text-base text-slate-100/85">
          <p>
            Calvary Chapel Phnom Penh is a community of believers dedicated to serving God and sharing His love with the people of Cambodia. We are a part of the Calvary Chapel movement, which emphasizes verse-by-verse, chapter-by-chapter teaching through the Bible.
          </p>
          <p>
            Our church family is diverse, welcoming people from all walks of life. We believe in creating a warm, loving environment where everyone can grow in their relationship with Jesus Christ.
          </p>
          <p>We are committed to:</p>
          <ul class="space-y-2 text-sm leading-relaxed text-slate-100/80">
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Teaching the Word of God faithfully and accurately</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Fostering genuine fellowship and community</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Serving our local community in Phnom Penh and throughout Cambodia</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Making disciples who make disciples</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Sharing the Gospel with those who don't know Jesus</span>
            </li>
          </ul>
          <p>
            Whether you're new to faith or have been walking with Christ for years, you're welcome here. We'd love to get to know you and have you join us as we seek to follow Jesus together.
          </p>
        </div>
      </div>
      <div class="space-y-6">
        <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur">
          <h2 class="text-lg font-semibold text-white">Join Our Family</h2>
          <p class="mt-3 text-sm text-slate-100/80">
            We'd love to welcome you into our community of faith and fellowship.
          </p>
          <div class="mt-6 flex flex-wrap gap-4">
            <router-link
              to="/contact"
              class="inline-flex items-center justify-center rounded-full border border-primary-400 px-6 py-3 text-sm font-semibold text-primary-200 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
            >
              Contact Us
            </router-link>
            <router-link
              to="/about"
              class="inline-flex items-center justify-center rounded-full border border-white/20 px-6 py-3 text-sm font-semibold text-slate-100 transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
            >
              Learn More
            </router-link>
          </div>
        </div>
        <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur">
          <h3 class="text-base font-semibold uppercase tracking-[0.3em] text-primary-200">Our Heartbeat</h3>
          <p class="mt-4 text-sm text-slate-100/80">
            We grow together through worship, the Word, and walking out the gospel every day.
          </p>
        </div>
      </div>
    </div>
  </FullPageShell>
</template>

<script>
import FullPageShell from '@/components/FullPageShell.vue'

export default {
  name: 'WhoWeAre',
  components: {
    FullPageShell,
  },
}
</script>

