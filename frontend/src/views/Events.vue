<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"></div>
      <div class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"></div>
    </div>

    <section class="relative flex min-h-screen flex-col px-4 py-20 sm:px-8 lg:px-12">
      <div class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"></div>
        <div class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"></div>
      </div>

      <div class="w-full text-slate-100">
        <div class="w-full">
          <div class="text-center">
            <p class="text-xs font-semibold uppercase tracking-[0.35em] text-primary-200/80">Gather Together</p>
            <h1 class="mt-6 text-4xl font-semibold tracking-tight text-white sm:text-5xl">Church Events</h1>
            <p class="mt-6 text-base text-slate-100/80">
              Join us for worship, fellowship, and community throughout the week.
            </p>
          </div>

          <div v-if="loading" class="mt-16 flex flex-col items-center justify-center">
            <div class="mb-4 inline-block h-12 w-12 animate-spin rounded-full border-b-2 border-primary-300"></div>
            <p class="text-sm text-slate-100/70">Loading events...</p>
          </div>

          <div v-else-if="error" class="mt-16 rounded-3xl border border-red-300/40 bg-red-900/30 p-8 text-center shadow-xl backdrop-blur">
            <p class="text-sm font-semibold text-red-200">{{ error }}</p>
          </div>

          <div v-else class="mt-16">
            <div v-if="events.length > 0" class="space-y-8">
              <article
                v-for="event in events"
                :key="event.id"
                @click="openModal(event)"
                class="group cursor-pointer rounded-3xl border border-white/10 bg-slate-900/70 p-6 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60"
              >
                <div class="flex flex-col gap-6 md:flex-row">
                  <div class="flex-shrink-0">
                    <div class="flex h-24 w-24 flex-col items-center justify-center rounded-3xl bg-primary-500/30 text-white shadow-lg shadow-primary-500/40">
                      <span class="text-3xl font-bold">{{ formatDay(event.date) }}</span>
                      <span class="text-xs font-semibold uppercase tracking-[0.4em] text-primary-100">{{ formatMonth(event.date) }}</span>
                    </div>
                  </div>
                  <div class="flex-1 text-left">
                    <h2 class="text-2xl font-semibold text-white">{{ event.title }}</h2>
                    <p class="mt-3 text-sm text-slate-100/80">{{ event.summary }}</p>
                    <div class="mt-4 flex flex-wrap items-center gap-3 text-xs font-semibold uppercase tracking-[0.25em] text-slate-100/60">
                      <span>{{ formatFullDate(event.date) }}</span>
                      <span v-if="event.time" class="rounded-full border border-primary-200/50 px-3 py-1 text-primary-100">
                        {{ event.time }}
                      </span>
                      <span v-if="event.location" class="rounded-full border border-white/20 px-3 py-1 text-slate-100">
                        {{ event.location }}
                      </span>
                    </div>
                  </div>
                </div>
              </article>
            </div>

            <div
              v-else
              class="rounded-3xl border border-white/10 bg-slate-900/70 p-12 text-center shadow-xl backdrop-blur"
            >
              <div class="text-6xl drop-shadow-lg">📅</div>
              <h3 class="mt-6 text-2xl font-semibold text-white">No Upcoming Events</h3>
              <p class="mt-4 text-sm text-slate-100/70">Check back soon for upcoming church events and activities.</p>
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="selectedEvent"
        @click="closeModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4"
      >
        <div
          @click.stop
          class="w-full max-w-3xl overflow-hidden rounded-3xl border border-white/10 bg-slate-950/90 shadow-2xl backdrop-blur"
        >
          <div class="flex items-center justify-between border-b border-white/10 px-6 py-4">
            <h2 class="text-2xl font-semibold text-white">{{ selectedEvent.title }}</h2>
            <button
              @click="closeModal"
              class="text-slate-200 transition hover:text-white"
            >
              <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          <div class="px-6 py-6">
            <div class="flex flex-wrap items-center gap-3 text-xs font-semibold uppercase tracking-[0.25em] text-slate-100/60">
              <span>{{ formatFullDate(selectedEvent.date) }}</span>
              <span v-if="selectedEvent.time" class="rounded-full border border-primary-200/40 px-3 py-1 text-primary-100">
                {{ selectedEvent.time }}
              </span>
              <span v-if="selectedEvent.location" class="rounded-full border border-white/20 px-3 py-1 text-slate-100">
                {{ selectedEvent.location }}
              </span>
            </div>
            <div class="mt-6 text-sm leading-relaxed text-slate-100/80">
              <p class="whitespace-pre-line">{{ selectedEvent.content }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'Events',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const events = ref([])
    const selectedEvent = ref(null)

    // Sample events data - In production, this would come from an API
    const sampleEvents = [
      {
        id: 1,
        title: 'Sunday Worship Service',
        summary: 'Join us every Sunday for worship, prayer, and fellowship.',
        content: `We gather every Sunday morning for worship and teaching from God's Word.\n\nService Schedule:\n- Praise and Worship: 9:00 AM\n- Prayer: 9:30 AM\n- Message: 10:00 AM\n- Fellowship: 11:30 AM\n\nAll are welcome! We look forward to worshipping with you.`,
        date: new Date('2024-02-04'),
        time: '9:00 AM',
        location: 'Main Sanctuary'
      },
      {
        id: 2,
        title: 'Midweek Bible Study',
        summary: 'Deep dive into God\'s Word every Wednesday evening.',
        content: `Join us for our weekly Bible study where we explore Scripture together and grow in our understanding of God's Word.\n\nCurrent Study: The Book of Romans\n\nTime: Every Wednesday at 7:00 PM\nLocation: Main Sanctuary\n\nThis is a great opportunity for fellowship, discussion, and spiritual growth.`,
        date: new Date('2024-02-07'),
        time: '7:00 PM',
        location: 'Main Sanctuary'
      },
      {
        id: 3,
        title: 'Community Outreach Day',
        summary: 'Serving our community together with love and compassion.',
        content: `We're organizing a community outreach day to serve our neighbors and share the love of Christ.\n\nActivities:\n- Food distribution\n- Clothing drive\n- Prayer and counseling\n- Children's activities\n\nDate: Saturday, February 10th\nTime: 10:00 AM - 2:00 PM\nLocation: Church Grounds\n\nVolunteers are welcome! Please contact us if you'd like to help.`,
        date: new Date('2024-02-10'),
        time: '10:00 AM - 2:00 PM',
        location: 'Church Grounds'
      },
      {
        id: 4,
        title: 'Youth Group Meeting',
        summary: 'Fellowship and fun for teens and young adults.',
        content: `Our youth group meets regularly for games, discussion, and growing together in faith.\n\nAges: 13-25\nTime: Friday, February 9th at 6:00 PM\nLocation: Youth Center\n\nActivities include:\n- Group games\n- Bible discussion\n- Prayer time\n- Snacks and fellowship\n\nCome join us for an evening of fun and spiritual growth!`,
        date: new Date('2024-02-09'),
        time: '6:00 PM',
        location: 'Youth Center'
      },
      {
        id: 5,
        title: 'Prayer Meeting',
        summary: 'Gathering together for prayer and intercession.',
        content: `Join us for our monthly prayer meeting where we lift up the needs of our church, community, and world.\n\nWe will pray for:\n- Our church family\n- Community needs\n- Mission work\n- Personal requests\n\nTime: Friday, February 9th at 7:00 PM\nLocation: Prayer Room\n\nAll are welcome to join us in prayer.`,
        date: new Date('2024-02-09'),
        time: '7:00 PM',
        location: 'Prayer Room'
      }
    ]

    const formatDay = (date) => {
      return date.getDate()
    }

    const formatMonth = (date) => {
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
      return months[date.getMonth()]
    }

    const formatFullDate = (date) => {
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
    }

    const fetchEvents = async () => {
      loading.value = true
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 500))
        
        // In production, you would fetch from an API here
        events.value = sampleEvents.sort((a, b) => a.date - b.date)
        
        loading.value = false
      } catch (err) {
        console.error('Error loading events:', err)
        error.value = 'Failed to load events'
        loading.value = false
      }
    }

    const openModal = (event) => {
      selectedEvent.value = event
    }

    const closeModal = () => {
      selectedEvent.value = null
    }

    onMounted(() => {
      fetchEvents()
    })

    return {
      loading,
      error,
      events,
      selectedEvent,
      formatDay,
      formatMonth,
      formatFullDate,
      openModal,
      closeModal
    }
  }
}
</script>

