<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[45vh] flex items-center overflow-hidden pt-16 md:pt-20 pb-12 bg-gradient-to-br from-[#111d33] via-[#1d2f48] to-[#213d5c] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 15% 25%, rgba(255,255,255,0.2), transparent 45%), radial-gradient(circle at 70% 15%, rgba(255,255,255,0.15), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Ministries</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          Serving Phnom Penh with Scripture, worship, discipleship, and practical love.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Discipleship
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Worship
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Families
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Outreaches
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">What are our ministries?</h2>
          <div class="space-y-4 text-left text-base md:text-lg text-gray-900 leading-relaxed max-w-3xl mx-auto">
            <p>
              Each ministry exists to keep <span class="font-semibold text-brand-orange">Jesus at the center</span>, meeting people in Phnom Penh with Scripture, Spirit-led care, and authentic community.
            </p>
            <p>
              We disciple kids, students, adults, and families through <span class="font-semibold text-brand-orange">Bible study</span>, <span class="font-semibold text-brand-orange">worship</span>, and <span class="font-semibold text-brand-orange">prayer</span>, while practical outreaches serve neighborhoods with compassion.
            </p>
            <p>
              Whether you’re new to faith or ready to lead, there is a place for you to grow, belong, and serve alongside Calvary Chapel Phnom Penh.
            </p>
          </div>
        </section>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="youthMinistryImage"
                alt="Youth Ministry"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
              <span class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900">
                Youth
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">Youth Ministry</h3>
              <p class="text-gray-600 flex-1">
                Teens gather weekly for worship, teaching, and discipleship that equips them to follow Jesus boldly and love Phnom Penh well.
              </p>
              <router-link
                to="/ministries/youth-ministry"
                class="inline-flex items-center justify-center px-4 py-2 text-base font-semibold text-white bg-brand-orange rounded-full shadow-lg shadow-brand-orange/30 transition-all duration-300 group-hover:-translate-y-0.5 group-hover:bg-brand-orange/90"
              >
                Learn More
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </article>

          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="graceChurchImage"
                alt="Grace Church"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
              <span class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900">
                Grace
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">Grace Church</h3>
              <p class="text-gray-600 flex-1">
                A network of house gatherings and Sunday services proclaiming the grace of God and forming communities that practice forgiveness and hospitality.
              </p>
              <router-link
                to="/ministries/grace-church"
                class="inline-flex items-center justify-center px-4 py-2 text-base font-semibold text-white bg-brand-orange rounded-full shadow-lg shadow-brand-orange/30 transition-all duration-300 group-hover:-translate-y-0.5 group-hover:bg-brand-orange/90"
              >
                Learn More
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </article>

          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="outreachesImage"
                alt="Outreaches"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
              <span class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900">
                Outreach
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">Outreaches</h3>
              <p class="text-gray-600 flex-1">
                Teams bring meals, prayer, medical care, and encouragement into neighborhoods, pointing friends and families to the hope of Jesus.
              </p>
              <router-link
                to="/ministries/outreaches"
                class="inline-flex items-center justify-center px-4 py-2 text-base font-semibold text-white bg-brand-orange rounded-full shadow-lg shadow-brand-orange/30 transition-all duration-300 group-hover:-translate-y-0.5 group-hover:bg-brand-orange/90"
              >
                Learn More
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-lg hover:bg-brand-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import youthMinistryImage from '@/assets/gallery/571314951_1275851217909869_307313164703676157_n.jpg'
import graceChurchImage from '@/assets/gallery/548188348_1236675751827416_5183731636046995004_n.jpg'
import outreachesImage from '@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg'

export default {
  name: 'Ministries',
  data() {
    return {
      youthMinistryImage,
      graceChurchImage,
      outreachesImage,
    };
  },
};
</script>

