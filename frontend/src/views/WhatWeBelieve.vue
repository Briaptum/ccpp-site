<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[65vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-16">
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="{ backgroundImage: `linear-gradient(120deg, rgba(17,39,84,0.35), rgba(17,39,84,0.55)), url(${beliefHeroImage})` }"
        ></div>
        <div class="absolute inset-0 bg-gradient-to-br from-black/70 via-black/55 to-black/65"></div>
        <div class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-blue/35 via-brand-blue/15 to-transparent blur-[120px] opacity-60"></div>
        <div class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-orange/30 via-brand-orange/12 to-transparent blur-[180px] opacity-65"></div>
        <div class="absolute -top-24 left-1/2 w-[420px] h-[420px] -translate-x-1/2 bg-white/15 blur-[180px] opacity-60"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/80 via-black/60 to-transparent"></div>
        <div
          class="absolute inset-0 opacity-15 mix-blend-screen"
          style="background-image: radial-gradient(circle at 18% 30%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 10%, rgba(255,255,255,0.22), transparent 30%), radial-gradient(circle at 62% 78%, rgba(255,255,255,0.18), transparent 32%);"
        ></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.35em] text-white/70">Statement of faith</p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight">
          What We Believe
        </h1>
        <p class="text-lg md:text-xl text-white/85 font-light">
          Our doctrine is historic, Christ-centered, and fuels the way we teach, serve, and love Phnom Penh.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="point in heroPoints"
            :key="point"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.25em] backdrop-blur-sm"
          >
            {{ point }}
          </span>
        </div>
      </div>
    </section>

    <!-- Detailed Statements -->
    <section class="relative py-24 bg-brand-blue text-white overflow-hidden">
      <div class="absolute inset-0 opacity-80" aria-hidden="true">
        <div class="absolute inset-0 bg-gradient-to-br from-brand-blue via-[#0a1f3e] to-[#153866]"></div>
        <div class="absolute inset-y-[-25%] left-[-20%] w-2/3 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[160px] opacity-60"></div>
        <div class="absolute inset-y-[-30%] right-[-15%] w-1/2 bg-gradient-to-l from-white/15 via-transparent to-transparent blur-[140px] opacity-70"></div>
        <div class="absolute inset-x-0 bottom-0 h-44 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="background-image: radial-gradient(circle at 25% 20%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 70% 10%, rgba(255,255,255,0.22), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"
        ></div>
      </div>
      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <div class="text-center">
          <p class="text-sm uppercase tracking-[0.35em] text-white/70 mb-3">Full statement</p>
          <h2 class="text-4xl md:text-5xl font-bold mb-4">We believe…</h2>
          <p class="text-white/80 max-w-3xl mx-auto">
            The following articles summarize the doctrines we joyfully teach. Each one shapes the way we gather, disciple, and serve Phnom Penh.
          </p>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <article
            v-for="(statement, index) in detailedStatements"
            :key="statement.id"
            class="rounded-3xl border border-white/15 bg-white/5 backdrop-blur-lg p-8 space-y-4"
          >
            <div class="flex items-center gap-4">
              <span class="inline-flex h-12 w-12 items-center justify-center rounded-2xl bg-white/10 text-lg font-semibold">
                {{ (index + 1).toString().padStart(2, '0') }}
              </span>
              <h3 class="text-xl font-semibold leading-snug">{{ statement.title }}</h3>
            </div>
            <p class="text-base leading-relaxed text-white/85">{{ statement.body }}</p>
            <p v-if="statement.reference" class="text-sm uppercase tracking-[0.3em] text-white/60">
              {{ statement.reference }}
            </p>
          </article>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="py-20 bg-white">
      <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Connect</p>
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900">Want to talk through these beliefs?</h2>
        <p class="text-lg text-gray-600">
          Our pastors and team would love to open Scripture with you, answer questions, and help you take the next step with Jesus.
        </p>
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            Contact Us
          </router-link>
          <router-link
            to="/about"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-blue/10 text-brand-blue font-semibold rounded-full border border-brand-blue/20 hover:bg-brand-blue/20 transition-colors"
          >
            Discover Calvary Chapel
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import beliefHeroImage from '@/assets/background/bg-hero3.jpg'

export default {
  name: 'WhatWeBelieve',
  setup() {
    const heroPoints = ['Scripture', 'Spirit-led', 'Christ-centered', 'Hope-filled']

    const detailedStatements = [
      {
        id: 'triune-god',
        title: 'The Triune God',
          body:
          'That there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.'
        },
        {
        id: 'scripture',
        title: 'Authority of Scripture',
          body:
          'That the Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.'
        },
        {
        id: 'father',
        title: 'God the Father',
          body:
          'In God the Father, an infinite, personal Spirit, perfect in holiness, wisdom, power and love; that He concerns Himself mercifully in the affairs of men, hears and answers prayer, and saves from sin and death all who come to Him through Jesus Christ.'
        },
        {
        id: 'jesus',
        title: 'Jesus Christ the Son',
          body:
          "In Jesus Christ, God's only begotten Son, conceived by the Holy Spirit. We affirm His virgin birth, sinless life, miracles and teachings, His substitutionary atoning death, bodily resurrection, ascension into heaven, perpetual intercession for His people, and personal, visible return to earth."
        },
        {
        id: 'blessed-hope',
        title: 'The Blessed Hope',
        body: 'In the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.'
        },
        {
        id: 'resurrection',
        title: 'Resurrection & Eternity',
          body:
          'In the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved and the everlasting conscious punishment of the lost.'
        },
        {
        id: 'holy-spirit',
        title: 'The Holy Spirit',
          body:
          'In the Holy Spirit, who came forth from the Father and Son to convict the world of sin, righteousness, and judgment and to regenerate, sanctify and empower for ministry all who believe in Christ. We believe the Holy Spirit indwells every believer in Jesus Christ and that He is an abiding Helper, Teacher and Guide, active today in all His gifts.'
        },
        {
        id: 'salvation',
        title: 'Salvation & New Life',
          body:
          'That all men are sinners by nature and choice and therefore under condemnation, that God regenerates by the Holy Spirit those who repent of their sins and confess Jesus Christ as Lord; that Jesus Christ baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.'
        },
        {
        id: 'church',
        title: 'The Church',
        body: 'In the universal Church, the living spiritual body of which Christ is the Head and all regenerated persons are members.'
        },
        {
        id: 'ordinances',
        title: 'Baptism & Communion',
          body:
          "That the Lord Jesus Christ committed two ordinances to the Church: baptism and the Lord's Supper. We practice baptism by immersion and celebrate communion open to all believers."
        },
        {
        id: 'kingdom',
        title: 'Christ’s Kingdom',
          body:
          'In the personal, visible return of Christ to earth and the establishment of His Kingdom; in the resurrection of the body, the final judgment and eternal blessing of the righteous and endless separation of the wicked.'
        },
        {
        id: 'discernment',
        title: 'Discernment by the Word',
          body:
            'That the Bible is the inspired Word of God, complete and entire, through which the Holy Spirit gives authority to test all things and discern what is of God.',
        reference: '1 Corinthians 2:13'
      }
    ]

    return {
      beliefHeroImage,
      heroPoints,
      detailedStatements
    }
  }
}
</script>
