<template>
  <FullPageShell>
    <div class="flex w-full flex-col gap-12">
      <div class="text-center lg:text-left">
        <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Leadership</p>
        <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">
          Shepherds who serve with humility and courage.
        </h1>
        <p class="mt-6 text-base text-slate-100/80">
          Our leadership team is committed to faithful teaching, compassionate care, and equipping the body of Christ to thrive in Phnom Penh and beyond.
        </p>
      </div>

      <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
        <article
          v-for="leader in leaders"
          :key="leader.name"
          class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 text-center shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60"
        >
          <div class="mx-auto h-28 w-28 overflow-hidden rounded-full border border-white/20 shadow-lg">
            <img
              :src="leader.photo"
              :alt="`${leader.name} portrait`"
              class="h-full w-full object-cover"
            />
          </div>
          <h3 class="mt-6 text-xl font-semibold text-white">{{ leader.name }}</h3>
          <p class="mt-2 text-primary-200/90">{{ leader.role }}</p>
          <p class="mt-4 text-sm text-slate-100/80">
            {{ leader.bio }}
          </p>
        </article>
      </div>

      <div class="grid gap-8 lg:grid-cols-[1.1fr,0.9fr]">
        <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-10 shadow-xl backdrop-blur">
          <h2 class="text-2xl font-semibold text-white">Our Leadership Philosophy</h2>
          <p class="mt-4 text-base text-slate-100/80">
            Our leadership team is committed to servant leadership, following the example of Jesus Christ who came not to be served, but to serve.
          </p>
          <ul class="mt-6 space-y-3 text-sm text-slate-100/80">
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Humbly serving the body of Christ</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Faithfully teaching and applying God's Word</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Leading by example in love, integrity, and devotion</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Equipping others for ministry and service</span>
            </li>
            <li class="flex items-start gap-3">
              <span class="mt-1 inline-flex h-2.5 w-2.5 rounded-full bg-primary-200/90"></span>
              <span>Praying for and caring for the congregation</span>
            </li>
          </ul>
        </div>
        <div class="rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
          <h2 class="text-2xl font-semibold text-white">Get In Touch</h2>
          <p class="mt-4 text-base text-slate-100/80">
            Have questions or want to connect with our leadership team?
          </p>
          <router-link
            to="/contact"
            class="mt-8 inline-flex items-center justify-center rounded-full border border-white/25 px-8 py-3 text-sm font-semibold text-white transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>
  </FullPageShell>
</template>

<script>
import FullPageShell from '@/components/FullPageShell.vue'

import pastorRithPhoto from '@/assets/gallery/547212625_1236675858494072_4425998334324483137_n.jpg'
import pastorLongPhoto from '@/assets/gallery/547266955_1236675715160753_108576014740250024_n.jpg'
import randyPhoto from '@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg'
import vuthaPhoto from '@/assets/gallery/547541557_1236675618494096_195308602630672666_n.jpg'
import chengPhoto from '@/assets/gallery/547588178_1236675705160754_8393862384233332823_n.jpg'

export default {
  name: 'Leadership',
  components: {
    FullPageShell,
  },
  data() {
    return {
      leaders: [
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
      ],
    };
  },
};
</script>

