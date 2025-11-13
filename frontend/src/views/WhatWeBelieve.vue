<template>
  <FullPageShell>
    <div class="w-full">
      <div class="text-center">
        <img
          src="/src/assets/images/logo3.png"
          alt="Calvary Chapel Phnom Penh Logo"
          class="mx-auto h-24 w-auto drop-shadow-lg sm:h-28"
        />
        <p class="mt-6 text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Statement of Faith</p>
        <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">What We Believe</h1>
        <p class="mt-6 text-base text-slate-100/80">
          These convictions anchor every ministry expression of Calvary Chapel Phnom Penh.
        </p>
      </div>

      <div class="mt-12 grid gap-6 lg:grid-cols-2">
        <article
          v-for="(statement, index) in statements"
          :key="index"
          class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60"
        >
          <h2 class="text-lg font-semibold uppercase tracking-[0.35em] text-primary-200">We Believe</h2>
          <p class="mt-4 text-base text-slate-100/85">{{ statement.body }}</p>
          <p v-if="statement.reference" class="mt-4 text-xs font-semibold uppercase tracking-[0.25em] text-primary-300">
            {{ statement.reference }}
          </p>
        </article>
      </div>

      <div class="mt-14 rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
        <h2 class="text-2xl font-semibold text-white">Join Our Family</h2>
        <p class="mt-4 text-base text-slate-100/80">
          We'd love to welcome you into our community of faith and fellowship.
        </p>
        <div class="mt-6 flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center rounded-full border border-primary-200 px-6 py-3 text-sm font-semibold text-primary-100 transition hover:-translate-y-0.5 hover:bg-primary-500 hover:text-white"
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
    </div>
  </FullPageShell>
</template>

<script>
import FullPageShell from '@/components/FullPageShell.vue'

export default {
  name: 'WhatWeBelieve',
  components: {
    FullPageShell,
  },
  data() {
    return {
      statements: [
        {
          body:
            'That there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.',
        },
        {
          body:
            'That the Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.',
        },
        {
          body:
            'In God the Father, an infinite, personal Spirit, perfect in holiness, wisdom, power and love; that He concerns Himself mercifully in the affairs of men, hears and answers prayer, and saves from sin and death all who come to Him through Jesus Christ.',
        },
        {
          body:
            "In Jesus Christ, God's only begotten Son, conceived by the Holy Spirit. We affirm His virgin birth, sinless life, miracles and teachings, His substitutionary atoning death, bodily resurrection, ascension into heaven, perpetual intercession for His people, and personal, visible return to earth.",
        },
        {
          body:
            'In the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.',
        },
        {
          body:
            'In the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved and the everlasting conscious punishment of the lost.',
        },
        {
          body:
            'In the Holy Spirit, who came forth from the Father and Son to convict the world of sin, righteousness, and judgment and to regenerate, sanctify and empower for ministry all who believe in Christ. We believe the Holy Spirit indwells every believer in Jesus Christ and that He is an abiding Helper, Teacher and Guide, active today in all His gifts.',
        },
        {
          body:
            'That all men are sinners by nature and choice and therefore under condemnation, that God regenerates by the Holy Spirit those who repent of their sins and confess Jesus Christ as Lord; that Jesus Christ baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.',
        },
        {
          body:
            'In the universal Church, the living spiritual body of which Christ is the Head and all regenerated persons are members.',
        },
        {
          body:
            "That the Lord Jesus Christ committed two ordinances to the Church: baptism and the Lord's Supper. We practice baptism by immersion and celebrate communion open to all believers.",
        },
        {
          body:
            'In the personal, visible return of Christ to earth and the establishment of His Kingdom; in the resurrection of the body, the final judgment and eternal blessing of the righteous and endless separation of the wicked.',
        },
        {
          body:
            'That the Bible is the inspired Word of God, complete and entire, through which the Holy Spirit gives authority to test all things and discern what is of God.',
          reference: '1 Corinthians 2:13',
        },
      ],
    };
  },
};
</script>

