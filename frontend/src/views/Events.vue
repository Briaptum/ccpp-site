<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[50vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#101c33] via-[#172b4a] to-[#1f3d61] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 20% 30%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 12%, rgba(255,255,255,0.18), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.2), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">{{ hero.title }}</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          {{ hero.subtitle }}
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="chip in hero.chips"
            :key="chip"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm"
          >
            {{ chip }}
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <!-- Navigation Links -->
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/events/upcoming-events"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            {{ nav.upcoming }}
          </router-link>
          <router-link
            to="/events/calendar"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-blue/10 text-brand-blue font-semibold rounded-full border border-brand-blue/20 hover:bg-brand-blue/20 transition-colors"
          >
            {{ nav.calendar }}
          </router-link>
        </div>

        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">{{ intro.eyebrow }}</p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">{{ intro.title }}</h2>
          <div class="space-y-4 text-left text-lg text-gray-700 leading-relaxed max-w-3xl mx-auto">
            <p v-for="(paragraph, idx) in intro.body" :key="idx" v-html="paragraph"></p>
          </div>
        </section>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <article
            v-for="card in cards"
            :key="card.id"
            class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-3"
          >
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">{{ card.tag }}</p>
            <h3 class="text-2xl font-semibold text-gray-900">{{ card.title }}</h3>
            <p class="text-gray-600">
              {{ card.body }}
            </p>
          </article>
        </div>

        <div class="text-center text-gray-700 text-base">
          {{ footerNote }}
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            {{ cta.primary }}
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useLanguageStore } from '@/stores/language'

export default {
  name: 'Events',
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore())

    const translations = {
      en: {
        hero: {
          title: 'Events',
          subtitle: 'Weekly rhythms and special moments that keep Calvary Chapel Phnom Penh connected to Jesus and each other.',
          chips: ['Worship & Word', 'Prayer', 'Youth & Families', 'City Serve']
        },
        nav: {
          upcoming: 'Upcoming Events',
          calendar: 'View Calendar'
        },
        intro: {
          eyebrow: 'Question',
          title: 'What happens each week?',
          body: [
            'Our weekly rhythm keeps us rooted in <span class="font-semibold text-brand-orange">Scripture, prayer, and community</span>. Gatherings happen in the sanctuary, homes, and neighborhoods across Phnom Penh.',
            'Sundays and Wednesdays anchor the church, while smaller circles—youth, prayer, and house churches—carry grace throughout the week.'
          ]
        },
        cards: [
          { id: 'sunday', tag: 'Sunday', title: '9:00 AM Worship', body: 'Verse-by-verse teaching, bilingual worship, kids ministry, and hospitality every Sunday morning.' },
          { id: 'wednesday', tag: 'Wednesday', title: 'Bible study & prayer', body: '7:00 PM midweek gathering focused on teaching, table discussion, and intercession for Phnom Penh.' },
          { id: 'prayer', tag: 'Prayer', title: 'City prayer nights', body: 'Monthly prayer for revival, leaders, and neighborhoods rotating through homes and the sanctuary.' },
          { id: 'outreach', tag: 'Outreach', title: 'Serve Phnom Penh', body: 'Monthly outreaches bring meals, medical care, and the Gospel into schools, villages, and communities.' },
          { id: 'youth', tag: 'Youth', title: 'Fridays 6:30 PM', body: 'Students gather for worship, teaching, and serving together—open to ages 12–18.' },
          { id: 'house', tag: 'House churches', title: 'Grace communities', body: 'Mid-week meals, communion, and discipleship circles throughout the city.' }
        ],
        footer: 'For detailed schedules and new events, check the calendar or contact us—new opportunities are added regularly.',
        cta: { primary: 'Contact Us' }
      },
      kh: {
        hero: {
          title: 'ព្រឹត្តិការណ៍',
          subtitle: 'កាលវិភាគប្រចាំសប្ដាហ៍ និងពេលពិសេស ដែលរក្សាព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ តភ្ជាប់នឹងព្រះយេស៊ូវ និងគ្នាទៅវិញទៅមក។',
          chips: ['ការថ្វាយបង្គំ និងព្រះបន្ទូល', 'ការអធិស្ឋាន', 'យុវជន និងគ្រួសារ', 'បម្រើទីក្រុង']
        },
        nav: {
          upcoming: 'ព្រឹត្តិការណ៍ជិតមកដល់',
          calendar: 'មើលប្រតិទិន'
        },
        intro: {
          eyebrow: 'សំណួរ',
          title: 'អ្វីកើតឡើងរៀងរាល់សប្ដាហ៍?',
          body: [
            'កាលវិភាគប្រចាំសប្ដាហ៍រក្សាយើងជាប់នឹង <span class="font-semibold text-brand-orange">ព្រះវចនៈ ការអធិស្ឋាន និងសហគមន៍</span>។ ការជួបជុំកើតឡើងនៅសាលធំ គេហដ្ឋាន និងសង្កាត់ទូទាំងទីក្រុងភ្នំពេញ។',
            'ថ្ងៃអាទិត្យ និងថ្ងៃពុធជាគ្រឹះសាសនាចក្រ ខណៈពេលព្រឺត្រីតូចៗ—យុវជន ការអធិស្ឋាន និងក្រុមផ្ទះ—នាំព្រះគុណឆ្លងកាត់មួយសប្ដាហ៍។'
          ]
        },
        cards: [
          { id: 'sunday', tag: 'ថ្ងៃអាទិត្យ', title: 'ការថ្វាយបង្គំ ម៉ោង 9 ព្រឹក', body: 'ការបង្រៀនតាមជួរ ការថ្វាយបង្គំពីរភាសា សេវាកម្មកុមារ និងការស្វាគមន៍រៀងរាល់ព្រឹកថ្ងៃអាទិត្យ។' },
          { id: 'wednesday', tag: 'ថ្ងៃពុធ', title: 'សិក្សាព្រះគម្ពីរ និងអធិស្ឋាន', body: 'ការជួបកណ្ដាលសប្ដាហ៍ ម៉ោង 7 ល្ងាច ផ្ដោតលើការបង្រៀន ការពិភាក្សាតាមតុ និងការអធិស្ឋានសម្រាប់ទីក្រុងភ្នំពេញ។' },
          { id: 'prayer', tag: 'ការអធិស្ឋាន', title: 'យប់អធិស្ឋានទីក្រុង', body: 'ការអធិស្ឋានប្រចាំខែសម្រាប់ការរំភើប អ្នកដឹកនាំ និងសង្កាត់ ដោយប្តូរទីតាំងគេហដ្ឋាន និងសាលធំ។' },
          { id: 'outreach', tag: 'ការចេញផ្សាយ', title: 'បម្រើទីក្រុងភ្នំពេញ', body: 'ការចេញផ្សាយប្រចាំខែនាំអាហារ ការថែទាំសុខភាព និងសារ​ព្រះ ទៅសាលារៀន ភូមិ និងសហគមន៍។' },
          { id: 'youth', tag: 'យុវជន', title: 'ថ្ងៃសុក្រ ម៉ោង 6:30 ល្ងាច', body: 'យុវជនជួបជុំសម្រាប់ការថ្វាយបង្គំ ការបង្រៀន និងបម្រើរួមគ្នា—បើកទូលាយសម្រាប់អាយុ 12–18។' },
          { id: 'house', tag: 'ក្រុមផ្ទះ', title: 'សហគមន៍ព្រះគុណ', body: 'អាហារកណ្ដាលសប្ដាហ៍ ពិធីអាហារព្រះអម្ចាស់ និងរង្វង់បណ្តុះបណ្តាលសិស្ស ឆ្លងកាត់ទីក្រុង។' }
        ],
        footer: 'សម្រាប់កាលវិភាគលម្អិត និងព្រឹត្តិការណ៍ថ្មី សូមពិនិត្យប្រតិទិន ឬទាក់ទងមកកាន់ពួកយើង—ឱកាសថ្មីៗត្រូវបានបន្ថែមជាប្រចាំ។',
        cta: { primary: 'ទាក់ទងមកកាន់ពួកយើង' }
      }
    }

    const content = computed(() => translations[selectedLanguage.value] || translations.en)

    return {
      hero: computed(() => content.value.hero),
      nav: computed(() => content.value.nav),
      intro: computed(() => content.value.intro),
      cards: computed(() => content.value.cards),
      footerNote: computed(() => content.value.footer),
      cta: computed(() => content.value.cta)
    }
  }
}
</script>

