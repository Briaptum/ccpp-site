<template>
  <div>


    <!-- Hero Section with Parallax Background -->
    <div
      ref="heroSection"
      class="relative min-h-screen flex items-center justify-center overflow-hidden"
    >
      <div
        class="absolute inset-0 bg-cover bg-center bg-no-repeat"
        :style="parallaxStyle"
      >
        <div class="absolute inset-0" style="background-image: radial-gradient(circle at center, rgba(34, 112, 130, 0.30), rgba(34, 112, 130, 0.18), transparent);"></div>
        <div class="absolute inset-0 bg-gradient-to-br from-brand-blue/20 via-brand-blue/8 to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-tl from-brand-blue/20 via-brand-blue/8 to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-t from-brand-blue/15 via-transparent to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-b from-transparent via-black/15 to-black/50"></div>
        <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/45 to-black/15"></div>
        <div class="absolute inset-0 bg-gradient-to-b from-black/30 via-transparent to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-br from-brand-orange/12 via-brand-orange/4 to-transparent"></div>
        <div class="absolute inset-0 bg-gradient-to-r from-black/25 via-transparent to-black/25"></div>
        <div class="absolute inset-0 bg-gradient-to-t from-transparent via-transparent to-black/20"></div>
      </div>
      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12 items-center min-h-[70vh] py-12">
          <div class="text-center lg:text-left text-white">
            <div class="mb-6">
              <p class="text-sm uppercase tracking-[0.4em] text-white/70 mb-4">Leadership</p>
              <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold mb-4 leading-tight drop-shadow-2xl">
                <span class="text-brand-orange uppercase tracking-[0.3em]">CALVARY CHAPEL</span><br />
                <span class="text-white text-2xl md:text-3xl lg:text-4xl uppercase font-light tracking-[0.4em]">LEADERSHIP</span>
              </h1>
              <div class="w-24 h-1 bg-brand-orange mx-auto lg:mx-0 mb-6 shadow-lg"></div>
            </div>
            <p class="text-xl md:text-2xl text-white font-light drop-shadow-xl mb-4 leading-relaxed">
              Shepherding Phnom Penh with prayer, Scripture, and presence.
            </p>
            <p class="text-base md:text-lg text-white/90 drop-shadow-md leading-relaxed font-light max-w-2xl mx-auto lg:mx-0">
              Our elders and ministry leads teach verse by verse, disciple across generations, and keep the church tethered to Jesus and His mission.
            </p>
          </div>
          <div class="flex justify-center lg:justify-end">
            <div class="p-0 md:p-2 text-white max-w-md w-full">
              <h2 class="text-2xl font-semibold text-white mb-6 text-center">How We Lead</h2>
              <div class="space-y-8 mb-8">
                <div>
                  <p class="text-sm uppercase tracking-[0.4em] text-white/80 mb-2">
                    Teaching · <span class="font-semibold tracking-normal">Verse by verse</span>
                  </p>
                  <p class="text-base text-white/85 leading-relaxed">
                    Expository preaching and pastoral care rooted in Scripture and grace.
                  </p>
                </div>
                <div>
                  <p class="text-sm uppercase tracking-[0.4em] text-white/80 mb-2">
                    Shepherding · <span class="font-semibold tracking-normal">Prayer & Counsel</span>
                  </p>
                  <p class="text-base text-white/85 leading-relaxed">
                    Walking with families through discipleship, prayer, and practical support.
                  </p>
                </div>
                <div>
                  <p class="text-sm uppercase tracking-[0.4em] text-white/80 mb-2">
                    Mission · <span class="font-semibold tracking-normal">City Presence</span>
                  </p>
                  <p class="text-base text-white/85 leading-relaxed">
                    Leading teams to serve Phnom Penh with compassion, truth, and hope.
                  </p>
                </div>
              </div>
              <router-link
                to="/contact"
                class="inline-flex items-center justify-center w-full text-sm font-semibold tracking-wide text-brand-orange hover:text-brand-blue transition-colors"
              >
                Meet the team
                <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Leaders Section -->
    <section class="py-24 bg-white">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid gap-10 sm:grid-cols-2 lg:grid-cols-3">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="rounded-3xl border border-gray-100 shadow-sm hover:shadow-xl transition-shadow bg-white p-8 text-center flex flex-col"
          >
            <div class="mx-auto h-32 w-32 overflow-hidden rounded-full border-4 border-white shadow-lg mb-6">
              <img
                :src="leader.photo"
                :alt="`${leader.name} portrait`"
                class="h-full w-full object-cover"
              />
            </div>
            <h3 class="text-2xl font-semibold text-gray-900">{{ leader.name }}</h3>
            <p class="mt-2 text-brand-blue font-medium tracking-wide uppercase text-xs">{{ leader.role }}</p>
            <p class="mt-4 text-sm text-gray-600 leading-relaxed flex-1">
              {{ leader.bio }}
            </p>
          </article>
        </div>

        <div class="text-center mt-16">
          <p class="text-gray-600 mb-3 text-sm uppercase tracking-[0.3em]">Next steps</p>
          <h4 class="text-2xl font-semibold text-gray-900 mb-6">Connect with our leadership team</h4>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors shadow-md"
          >
            Contact Us
            <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </router-link>
        </div>
      </div>
    </section>

  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import leadershipHeroImage from '@/assets/background/bg-hero4.jpg'
import pastorRithPhoto from '@/assets/leadership/rith.jpg'
import pastorLongPhoto from '@/assets/leadership/Paster Long.jpg'
import randyPhoto from '@/assets/leadership/randy.jpg'
import vuthaPhoto from '@/assets/leadership/longdee.jpg'
import chengPhoto from '@/assets/leadership/Paster Borin.jpg'

export default {
  name: 'Leadership',
  setup() {
    const heroSection = ref(null)
    const scrollY = ref(0)
    let animationFrameId = null

    const parallaxStyle = computed(() => {
      const yPos = scrollY.value * 0.5
      return {
        backgroundImage: `url(${leadershipHeroImage})`,
        transform: `translateY(${yPos}px)`,
        willChange: 'transform'
      }
    })

    const handleScroll = () => {
      if (animationFrameId !== null) {
        return
      }

      animationFrameId = requestAnimationFrame(() => {
        if (heroSection.value) {
          scrollY.value = window.scrollY
        }
        animationFrameId = null
      })
    }

    onMounted(() => {
      window.addEventListener('scroll', handleScroll, { passive: true })
      handleScroll()
    })

    onBeforeUnmount(() => {
      window.removeEventListener('scroll', handleScroll)
      if (animationFrameId !== null) {
        cancelAnimationFrame(animationFrameId)
      }
    })

    const leaders = [
      {
        name: 'Pastor Rith',
        role: 'Senior Pastor',
        photo: pastorRithPhoto,
        bio: "With over 15 years of pastoral experience, Pastor Rith leads our congregation with wisdom and compassion. He faithfully teaches God's Word and shepherds our church family.",
      },
      {
        name: 'Pastor Long',
        role: 'Associate Pastor',
        photo: pastorLongPhoto,
        bio: 'Pastor Long brings energy to youth and community outreach. He is passionate about discipleship and seeing lives transformed by the Gospel.',
      },
      {
        name: 'Randy',
        role: 'Leadership Team',
        photo: randyPhoto,
        bio: 'Randy supports the church family through faithful service and a commitment to caring for every ministry area.',
      },
      {
        name: 'Vutha',
        role: 'Leadership Team',
        photo: vuthaPhoto,
        bio: 'Vutha encourages meaningful worship and helps guide teams that cultivate fellowship and spiritual growth.',
      },
      {
        name: 'Cheng',
        role: 'Leadership Team',
        photo: chengPhoto,
        bio: 'Cheng invests in discipleship and helps coordinate opportunities for members to grow together in faith.',
      },
    ]

    return {
      heroSection,
      parallaxStyle,
      leaders
    }
  }
}
</script>

