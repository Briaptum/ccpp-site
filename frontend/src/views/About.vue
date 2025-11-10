<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"></div>
      <div class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"></div>
    </div>

    <section class="relative px-4 py-20 sm:px-8 lg:px-12">
      <div class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"></div>
        <div class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"></div>
      </div>
      <div class="mx-auto grid max-w-6xl gap-12 text-slate-100 lg:grid-cols-[1.1fr,0.9fr] lg:items-center">
        <div class="relative">
          <div class="absolute -left-6 -top-6 hidden h-32 w-32 rounded-3xl bg-primary-500/30 blur-2xl sm:block"></div>
          <img src="/src/assets/images/logo3.png" alt="Calvary Chapel Phnom Penh Logo" class="relative mx-auto mb-10 h-24 w-auto drop-shadow-lg sm:mx-0 sm:h-28" />
          <p class="relative text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Calvary Chapel Phnom Penh</p>
          <h1 class="relative mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">Walking in Scripture, empowered by the Spirit, united in love.</h1>
          <p class="relative mt-6 text-lg text-slate-100/85">
            Our church family gathers around the living Word of God, welcomes the transforming work of the Holy Spirit, and walks side by side as disciples of Jesus in Phnom Penh.
          </p>
          <div class="relative mt-10 grid gap-5 sm:grid-cols-2">
            <article
              v-for="highlight in heroHighlights"
              :key="highlight.title"
              class="rounded-2xl border border-white/10 bg-slate-900/60 p-6 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60 hover:bg-slate-800/70"
            >
              <h2 class="text-lg font-semibold text-white">{{ highlight.title }}</h2>
              <p class="mt-2 text-sm text-slate-100/85">{{ highlight.description }}</p>
            </article>
          </div>
        </div>
        <div class="relative">
          <div class="absolute inset-0 -z-10 rounded-3xl bg-primary-500/25 blur-3xl"></div>
          <div class="relative overflow-hidden rounded-3xl border border-white/10 bg-slate-900/70 p-10 shadow-[0_30px_80px_-40px_rgba(15,23,42,0.9)] backdrop-blur-md">
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200">Scripture Focus</p>
            <blockquote class="mt-6 text-2xl font-semibold text-white">
              &ldquo;These things we also speak, not in words taught by human wisdom, but in words taught by the Spirit.&rdquo;
            </blockquote>
            <p class="mt-6 text-sm font-medium text-primary-100">1 Corinthians 2:13</p>
            <div class="mt-10 space-y-4 text-sm text-slate-100/85">
              <div class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Teaching verse by verse through Scripture to know God deeply.</span>
              </div>
              <div class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Creating space for the Holy Spirit to renew minds and hearts.</span>
              </div>
              <div class="flex items-start gap-3">
                <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
                <span>Living out the gospel together in authentic community.</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="bg-white px-4 py-20 text-slate-900 sm:px-8 lg:px-12">
      <div class="mx-auto max-w-6xl">
        <div class="mx-auto max-w-3xl text-center">
          <h2 class="text-3xl font-semibold text-slate-900 sm:text-4xl">What We Believe</h2>
          <p class="mt-4 text-lg text-slate-600">
            These statements anchor every ministry expression of Calvary Chapel Phnom Penh.
          </p>
        </div>
        <div class="mt-12 grid gap-6 lg:grid-cols-2">
          <article
            v-for="statement in beliefStatements"
            :key="statement.title"
            class="group relative overflow-hidden rounded-3xl border border-slate-200 bg-white p-8 shadow-lg transition hover:-translate-y-1 hover:border-primary-400 hover:shadow-2xl"
          >
            <div class="absolute -right-20 -top-20 h-48 w-48 rounded-full bg-primary-200/30 blur-3xl transition group-hover:scale-125"></div>
            <h3 class="relative text-xl font-semibold text-slate-900">{{ statement.title }}</h3>
            <p class="relative mt-4 text-sm leading-relaxed text-slate-600">{{ statement.body }}</p>
            <p
              v-if="statement.reference"
              class="relative mt-4 text-xs font-semibold uppercase tracking-[0.2em] text-primary-500"
            >
              {{ statement.reference }}
            </p>
          </article>
        </div>
      </div>
    </section>

    <section class="bg-slate-100 px-4 py-20 text-slate-900 sm:px-8 lg:px-12">
      <div class="mx-auto max-w-6xl">
        <div class="flex flex-col gap-8 lg:flex-row lg:items-end lg:justify-between">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-500">Our Culture</p>
            <h2 class="mt-4 text-3xl font-semibold text-slate-900 sm:text-4xl">Values we live by</h2>
            <p class="mt-4 max-w-2xl text-base text-slate-600">
              Our rhythms and relationships flow out of a devotion to Jesus and one another.
            </p>
          </div>
          <router-link
            to="/outreaches"
            class="inline-flex items-center justify-center rounded-full border border-primary-500 px-6 py-3 text-sm font-semibold text-primary-600 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
          >
            See How We Serve
          </router-link>
        </div>
        <div class="mt-12 grid gap-6 md:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="value in values"
            :key="value.title"
            class="group flex flex-col gap-4 rounded-3xl border border-white bg-white p-8 shadow-lg transition hover:-translate-y-1 hover:border-primary-300 hover:shadow-xl"
          >
            <div class="text-4xl drop-shadow-sm">{{ value.icon }}</div>
            <div>
              <h3 class="text-xl font-semibold text-slate-900">{{ value.title }}</h3>
              <p class="mt-2 text-sm text-slate-600">{{ value.description }}</p>
            </div>
            <span class="mt-auto text-xs font-semibold uppercase tracking-[0.3em] text-primary-400 transition group-hover:text-primary-600"
              >Learn &amp; Live</span
            >
          </article>
        </div>
      </div>
    </section>

  </div>
</template>

<script>
export default {
  name: 'About',
  data() {
    return {
      heroHighlights: [
        {
          title: 'Rooted in Scripture',
          description: 'We treasure the Bible as our infallible guide for faith and life.',
        },
        {
          title: 'Empowered by the Spirit',
          description: 'We depend on the Holy Spirit to regenerate, sanctify, and equip every believer for ministry.',
        },
        {
          title: 'Centered on Jesus',
          description: 'We proclaim Jesus Christ - His life, death, resurrection, and promised return.',
        },
      ],
      beliefStatements: [
        {
          title: 'The Trinity',
          body: 'There is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory who created, upholds, and governs all.',
        },
        {
          title: 'The Scriptures',
          body: 'The Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.',
        },
        {
          title: 'God the Father',
          body: 'God the Father is an infinite, personal Spirit, perfect in holiness, wisdom, power, and love. He hears and answers prayer and saves from sin and death all who come to Him through Jesus Christ.',
        },
        {
          title: 'Jesus Christ',
          body: 'Jesus Christ is God\'s only begotten Son, conceived by the Holy Spirit. We affirm His virgin birth, sinless life, miracles, teachings, substitutionary death, bodily resurrection, ascension, intercession, and personal, visible return to earth.',
        },
        {
          title: 'The Blessed Hope',
          body: 'We look for the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.',
        },
        {
          title: 'Resurrection and Eternity',
          body: 'We believe in the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved, and the everlasting conscious punishment of the lost.',
        },
        {
          title: 'The Holy Spirit',
          body: 'The Holy Spirit convicts the world of sin, righteousness, and judgment and regenerates, sanctifies, and empowers for ministry all who believe in Christ. He indwells every believer as Helper, Teacher, and Guide.',
        },
        {
          title: 'Salvation',
          body: 'All people are sinners by nature and choice and therefore under condemnation. God regenerates by the Holy Spirit those who repent of their sins and confess Jesus Christ as Lord. Jesus baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.',
        },
        {
          title: 'The Church',
          body: 'We believe in the universal Church, the living spiritual body of which Christ is the Head and all regenerated persons are members.',
        },
        {
          title: 'Ordinances',
          body: 'The Lord Jesus Christ committed two ordinances to the Church: baptism and the Lord\'s Supper. We practice baptism by immersion and celebrate communion open to all believers.',
        },
        {
          title: 'Christ\'s Return',
          body: 'Jesus Christ will personally and visibly return to earth to establish His Kingdom. There will be a final judgment, eternal blessing for the righteous, and endless separation for the wicked.',
        },
        {
          title: 'The Word of God',
          body: 'The Bible is the inspired Word of God, complete and entire. Through the Holy Spirit we test all things to discern what is of God.',
          reference: '1 Corinthians 2:13',
        },
      ],
      values: [
        {
          icon: '❤️',
          title: 'Love',
          description: 'We believe in the power of unconditional love as demonstrated by Christ.',
        },
        {
          icon: '🤝',
          title: 'Community',
          description: 'We build strong relationships and support one another in faith.',
        },
        {
          icon: '🙏',
          title: 'Prayer',
          description: 'Prayer anchors our relationship with God and with each other.',
        },
        {
          icon: '📖',
          title: 'Scripture',
          description: 'God\'s Word guides every decision we make together.',
        },
        {
          icon: '🎵',
          title: 'Worship',
          description: 'We worship God with our whole hearts in spirit and truth.',
        },
        {
          icon: '🌱',
          title: 'Growth',
          description: 'We encourage spiritual maturity and personal development.',
        },
      ],
    };
  },
};
</script>
