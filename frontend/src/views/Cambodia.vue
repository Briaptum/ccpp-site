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
            <img
              src="/src/assets/images/logo3.png"
              alt="Calvary Chapel Phnom Penh Logo"
              class="mx-auto h-24 w-auto drop-shadow-lg sm:h-28"
            />
            <p class="mt-6 text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Cambodia</p>
            <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
              Loving Cambodia with the hope of the gospel.
            </h1>
            <p class="mt-6 text-base text-slate-100/80">
              We are honored to serve in Phnom Penh—standing with a resilient nation rich in culture, history, and promise.
            </p>
          </div>

          <div class="mt-12 space-y-6 text-left text-base text-slate-100/85">
            <p>
              Cambodia, a beautiful Southeast Asian nation, is home to over 16 million people. The country has a rich cultural heritage and a resilient spirit, having overcome tremendous challenges throughout its history.
            </p>
            <p>
              Located in Phnom Penh, the capital city, Calvary Chapel Phnom Penh serves a diverse community of Cambodians and expatriates. We are committed to being a light in this nation and sharing the hope of the Gospel with all people.
            </p>
            <h2 class="mt-8 text-2xl font-semibold text-white">Our Heart for Cambodia</h2>
            <p>We love Cambodia and its people. Our ministry is focused on:</p>
            <ul class="space-y-3 text-sm text-slate-100/80">
              <li class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Sharing the Gospel with Cambodians in their heart language</span>
              </li>
              <li class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Supporting local communities through outreaches and practical care</span>
              </li>
              <li class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Training and equipping Cambodian believers for ministry</span>
              </li>
              <li class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Partnering with local ministries and organizations</span>
              </li>
              <li class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Praying for spiritual awakening across the nation</span>
              </li>
            </ul>
            <h2 class="mt-8 text-2xl font-semibold text-white">Why Cambodia?</h2>
            <p>
              Cambodia is a nation open to the Gospel. The people are friendly, welcoming, and hungry for truth. As we serve here, we see God moving in powerful ways, transforming lives and communities.
            </p>
            <p>
              Whether you're a Cambodian national or an expatriate living here, we invite you to join us as we share the love of Jesus Christ with all who call Cambodia home.
            </p>
          </div>

          <div class="mt-16">
            <h2 class="text-center text-2xl font-semibold text-white">About Cambodia</h2>
            <div class="mt-10 grid gap-6 sm:grid-cols-3">
              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 text-center shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">🇰🇭</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Kingdom of Cambodia</h3>
                <p class="mt-3 text-sm text-slate-100/80">Rich in history, culture, and hospitality.</p>
              </article>
              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 text-center shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">👥</div>
                <h3 class="mt-4 text-lg font-semibold text-white">16+ Million People</h3>
                <p class="mt-3 text-sm text-slate-100/80">A diverse and welcoming population.</p>
              </article>
              <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 text-center shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-4xl drop-shadow-sm">🏛️</div>
                <h3 class="mt-4 text-lg font-semibold text-white">Phnom Penh</h3>
                <p class="mt-3 text-sm text-slate-100/80">The capital city where we serve and gather.</p>
              </article>
            </div>
          </div>

          <div class="mt-16 rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
            <h2 class="text-2xl font-semibold text-white">Serving Cambodia Together</h2>
            <p class="mt-4 text-base text-slate-100/80">
              Join us as we serve and share God's love in this beautiful nation.
            </p>
            <div class="mt-6 flex flex-wrap justify-center gap-4">
              <router-link
                to="/contact"
                class="inline-flex items-center justify-center rounded-full border border-primary-200 px-6 py-3 text-sm font-semibold text-primary-100 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
              >
                Get Involved
              </router-link>
              <router-link
                to="/how-to-help"
                class="inline-flex items-center justify-center rounded-full border border-white/20 px-6 py-3 text-sm font-semibold text-slate-100 transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
              >
                How to Help
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
  name: 'Cambodia'
}
</script>

