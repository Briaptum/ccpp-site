<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"></div>
      <div class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"></div>
    </div>

    <section class="relative flex min-h-screen items-center px-4 py-20 sm:px-8 lg:px-12">
      <div class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"></div>
        <div class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"></div>
      </div>
      <div class="w-full text-slate-100">
        <div class="w-full">
          <div class="text-center">
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Our Ministries</p>
            <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
              Serving Phnom Penh with faith and love.
            </h1>
            <p class="mt-6 text-base text-slate-100/80">
              Every ministry is designed to meet people where they are and help them grow in their relationship with Jesus.
            </p>
          </div>

          <div class="mt-12 grid gap-10 lg:grid-cols-[1.1fr,0.9fr] lg:items-start">
            <div class="space-y-6 text-left text-base text-slate-100/85">
              <p>
                At CCPP, ministry is about serving others and sharing the love of Christ. From the youngest to the oldest, we are committed to creating spaces where people encounter Jesus and are formed into disciples.
              </p>
              <p>
                Whether you're looking to get involved, receive support, or grow in your faith, we have ministries that can help you on your journey. Each ministry is committed to serving our community with excellence, humility, and love.
              </p>
              <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur">
                <h2 class="text-lg font-semibold text-white">Ministry Overview</h2>
                <p class="mt-4 text-sm text-slate-100/80">
                  We gather, equip, and send disciples who reflect Christ in their homes, workplaces, and neighborhoods.
                </p>
              </div>
            </div>
            <div class="rounded-3xl border border-primary-300/40 bg-primary-500/10 p-10 shadow-xl backdrop-blur">
              <h2 class="text-lg font-semibold uppercase tracking-[0.3em] text-primary-200">Core Rhythms</h2>
              <ul class="mt-6 space-y-4 text-sm text-slate-100/80">
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Gathering weekly to worship Jesus and study His Word.</span>
                </li>
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Growing together in smaller communities of prayer and encouragement.</span>
                </li>
                <li class="flex items-start gap-3">
                  <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                  <span>Going out to serve our city with the compassion of Christ.</span>
                </li>
              </ul>
            </div>
          </div>

          <div class="mt-16">
            <h2 class="text-center text-2xl font-semibold text-white">Ministries We Offer</h2>
            <div class="mt-10 grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">👥</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Children's Ministry</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Nurturing young hearts with age-appropriate Bible teaching, worship, and joyful community.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">Ages 4 – 12</p>
              </article>

              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">🎵</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Worship Ministry</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Leading our church in heartfelt praise through music, prayer, and a posture of surrender.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">All Ages Welcome</p>
              </article>

              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">📖</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Bible Study</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Weekly gatherings to explore Scripture deeply and apply it in community.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">Wednesday Evenings</p>
              </article>

              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">🙏</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Prayer Ministry</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Standing in the gap for our church family, our city, and the nations through intercession.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">Ongoing Support</p>
              </article>

              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">🌍</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Outreach Ministry</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Extending God’s love beyond our walls through service, evangelism, and partnerships.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">Monthly Events</p>
              </article>

              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">👨‍👩‍👧‍👦</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Family Ministry</h3>
                <p class="mt-3 text-sm text-slate-100/80">
                  Resources and gatherings that strengthen families to walk with Jesus together.
                </p>
                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.3em] text-primary-200">Monthly Gatherings</p>
              </article>
            </div>
          </div>

          <div class="mt-16 rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
            <h2 class="text-2xl font-semibold text-white">Get Involved</h2>
            <p class="mt-4 text-base text-slate-100/80">
              We’d love to have you join us in serving our community through ministry.
            </p>
            <div class="mt-6 flex flex-wrap justify-center gap-4">
              <router-link
                to="/contact"
                class="inline-flex items-center justify-center rounded-full border border-primary-200 px-6 py-3 text-sm font-semibold text-primary-100 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
              >
                Contact Us
              </router-link>
              <router-link
                to="/events"
                class="inline-flex items-center justify-center rounded-full border border-white/20 px-6 py-3 text-sm font-semibold text-slate-100 transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
              >
                View Upcoming Events
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'Ministries'
}
</script>

