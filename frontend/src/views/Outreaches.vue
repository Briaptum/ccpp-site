<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"></div>
      <div class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"></div>
    </div>

    <section class="relative flex min-h-screen flex-col px-4 py-20 sm:px-8 lg:px-12">
      <div class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"></div>
        <div class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"></div>
      </div>

      <div class="w-full text-slate-100">
        <div class="w-full text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Outreach</p>
          <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
            Taking the gospel beyond our walls.
          </h1>
          <p class="mt-6 text-base text-slate-100/80">
            We carry the love and message of Jesus into the communities of Cambodia, serving with compassion and sharing hope.
          </p>

          <div class="mt-12 rounded-3xl border border-white/10 bg-slate-900/70 p-10 text-left shadow-xl backdrop-blur">
            <h2 class="text-2xl font-semibold text-white">Our Outreach Ministries</h2>
            <div class="mt-6 space-y-5 text-base text-slate-100/85">
              <p>
                Our outreach programs are designed to take the love and message of Jesus Christ beyond our church walls into the communities of Cambodia. Serving others is a vital expression of our faith and a powerful way to share the Gospel.
              </p>
              <p>
                Through these initiatives we partner with local organizations, serve those in need, and create opportunities for people to encounter God's love through practical acts of service.
              </p>
            </div>
          </div>

          <div class="mt-16 rounded-3xl border border-primary-300/30 bg-primary-500/15 p-10 shadow-xl backdrop-blur">
            <h2 class="text-2xl font-semibold text-white">Get Involved</h2>
            <p class="mt-4 text-base text-slate-100/80">
              Join us in reaching our city with God's love. There is a place for you on the team.
            </p>
            <div class="mt-6 flex flex-wrap justify-center gap-4">
              <router-link
                to="/contact"
                class="inline-flex items-center justify-center rounded-full border border-primary-200 px-6 py-3 text-sm font-semibold text-primary-100 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
              >
                Contact Us
              </router-link>
              <router-link
                to="/ministries"
                class="inline-flex items-center justify-center rounded-full border border-white/20 px-6 py-3 text-sm font-semibold text-slate-100 transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
              >
                Back to Ministries
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
  name: 'Outreaches'
}
</script>

