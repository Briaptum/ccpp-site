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
        <div class="w-full text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Youth Ministry</p>
          <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
            Empowering the next generation to follow Jesus.
          </h1>
          <p class="mt-6 text-base text-slate-100/80">
            A vibrant community where teens encounter God, build lasting friendships, and discover their identity in Christ.
          </p>

          <div class="mt-12 grid gap-10 lg:grid-cols-[1.1fr,0.9fr] lg:items-start">
            <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-10 text-left shadow-xl backdrop-blur">
              <h2 class="text-2xl font-semibold text-white">About Our Youth Ministry</h2>
              <div class="mt-6 space-y-5 text-base text-slate-100/85">
                <p>
                  Our Youth Ministry helps teenagers and young adults discover who they are in Christ and develop a deep, personal relationship with God. We cultivate an environment where young people can ask questions, explore faith, and build authentic community.
                </p>
                <p>
                  Weekly gatherings combine engaging teaching, heartfelt worship, fun activities, and meaningful service opportunities so students can grow as disciples of Jesus.
                </p>
              </div>
            </div>
            <div class="space-y-6">
              <div class="rounded-3xl border border-primary-200/30 bg-primary-500/10 p-10 shadow-xl backdrop-blur">
                <h3 class="text-lg font-semibold text-white">What to Expect</h3>
                <ul class="mt-4 space-y-3 text-sm text-slate-100/80 text-left">
                  <li class="flex items-start gap-3">
                    <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                    <span>Interactive Bible teaching that speaks to everyday life.</span>
                  </li>
                  <li class="flex items-start gap-3">
                    <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                    <span>Authentic worship led by and for students.</span>
                  </li>
                  <li class="flex items-start gap-3">
                    <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200"></span>
                    <span>Service projects that put faith into action.</span>
                  </li>
                </ul>
              </div>
              <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-10 shadow-xl backdrop-blur">
                <h3 class="text-lg font-semibold text-white">Join Our Youth</h3>
                <p class="mt-4 text-sm text-slate-100/80">
                  We’d love to welcome you to our youth community. Reach out to get connected or learn more.
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
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'YouthMinistry'
}
</script>

