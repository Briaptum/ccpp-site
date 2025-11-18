<template>
  <div>


    <!-- Leaders Section -->
    <div class="bg-background py-32 md:py-40">
      <h1 class="text-4xl md:text-6xl font-bold mb-4 drop-shadow-lg text-center">Leadership</h1>
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3 mb-12">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="bg-background-50 rounded-lg p-6 text-center shadow-md"
          >
            <div class="mx-auto h-28 w-28 overflow-hidden rounded-full border-4 border-primary-200 shadow-lg mb-4">
              <img
                :src="leader.photo"
                :alt="`${leader.name} portrait`"
                class="h-full w-full object-cover"
              />
            </div>
            <h3 class="text-xl font-semibold text-gray-900">{{ leader.name }}</h3>
            <p class="mt-2 text-primary-600 font-medium">{{ leader.role }}</p>
            <p class="mt-4 text-sm text-gray-600">
              {{ leader.bio }}
            </p>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-8 py-3 bg-primary-600 text-white font-semibold rounded-lg hover:bg-primary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import background2Image from '@/assets/images/background2.jpg'
import pastorRithPhoto from '@/assets/gallery/547212625_1236675858494072_4425998334324483137_n.jpg'
import pastorLongPhoto from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'
import randyPhoto from '@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg'
import vuthaPhoto from '@/assets/gallery/547541557_1236675618494096_195308602630672666_n.jpg'
import chengPhoto from '@/assets/gallery/547588178_1236675705160754_8393862384233332823_n.jpg'

export default {
  name: 'Leadership',
  setup() {
    const heroSection = ref(null)
    const parallaxOffset = ref(0)
    let animationFrameId = null

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

    const updateParallax = () => {
      if (!heroSection.value) {
        return
      }

      const rect = heroSection.value.getBoundingClientRect()
      parallaxOffset.value = rect.top * -0.3
    }

    const handleScroll = () => {
      if (animationFrameId !== null) {
        return
      }

      animationFrameId = requestAnimationFrame(() => {
        updateParallax()
        animationFrameId = null
      })
    }

    const parallaxBackgroundStyle = computed(() => ({
      backgroundImage: `url(${background2Image})`,
      transform: `translateY(${parallaxOffset.value}px) scale(1.1)`
    }))

    onMounted(() => {
      updateParallax()
      window.addEventListener('scroll', handleScroll, { passive: true })
    })

    onBeforeUnmount(() => {
      window.removeEventListener('scroll', handleScroll)
      if (animationFrameId !== null) {
        cancelAnimationFrame(animationFrameId)
      }
    })

    return {
      background2Image,
      leaders,
      parallaxBackgroundStyle,
      heroSection
    }
  }
}
</script>

