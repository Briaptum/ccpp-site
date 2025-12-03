<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#0f1f33] via-[#182f4a] to-[#1f4061] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-25 mix-blend-screen" style="background-image: radial-gradient(circle at 20% 30%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 80% 10%, rgba(255,255,255,0.18), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.2), transparent 30%);"></div>
        <div class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/35 via-transparent to-transparent blur-[200px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Church Calendar</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          One place to see every worship gathering, prayer night, outreach, and family moment we have planned.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Monthly Rhythm
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Worship & Prayer
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Youth & Families
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-10">
        <section class="text-center space-y-4">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">How do I use the calendar?</h2>
          <p class="text-lg text-gray-700 max-w-3xl mx-auto leading-relaxed">
            Browse by month, tap any highlighted day to see its events, then open an event for full details—dates, times, and locations.
          </p>
        </section>
        <!-- Loading State -->
        <div v-if="loading" class="py-12">
          <div class="text-center">
            <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-main mb-4"></div>
            <p class="text-gray-600">Loading calendar...</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="py-12 text-center">
          <p class="text-red-600">{{ error }}</p>
        </div>

        <!-- Calendar View -->
        <div v-else>
          <div class="bg-white/95 backdrop-blur-sm rounded-[32px] shadow-2xl p-6 md:p-8 border border-gray-100">
            <!-- Calendar Header -->
            <div class="text-center mb-8">
              <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">{{ currentMonthName }} {{ currentYear }}</h2>
              <div class="flex justify-between items-center max-w-md mx-auto mb-6">
                <button 
                  @click="previousMonth"
                  class="p-2 text-main hover:bg-gray-100 rounded-lg transition-colors"
                  aria-label="Previous month"
                >
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                  </svg>
                </button>
                <button 
                  @click="resetToCurrentMonth"
                  class="px-6 py-2 bg-main text-white font-semibold rounded-full hover:bg-main/90 transition-colors shadow-md shadow-main/30"
                >
                  Today
                </button>
                <button 
                  @click="nextMonth"
                  class="p-2 text-main hover:bg-gray-100 rounded-lg transition-colors"
                  aria-label="Next month"
                >
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Calendar Grid -->
            <div class="grid grid-cols-7 gap-2 mb-4">
              <div 
                v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" 
                :key="day" 
                class="text-center font-bold text-gray-700 py-3 text-sm md:text-base"
              >
                {{ day }}
              </div>
            </div>

            <div class="grid grid-cols-7 gap-2">
              <div 
                v-for="day in calendarDays" 
                :key="day.date"
                :class="[
                  'min-h-24 md:min-h-32 p-3 border rounded-2xl cursor-pointer transition-all',
                  day.isToday
                    ? 'bg-main text-white border-main shadow-lg shadow-main/20'
                    : day.isCurrentMonth 
                      ? 'bg-white hover:bg-gray-50 border-gray-200 hover:border-main/30' 
                      : 'bg-gray-50 text-gray-400 border-gray-100 opacity-70',
                  day.hasEvents && !day.isToday ? 'border-brand-orange/50' : ''
                ]"
                @click="selectDate(day)"
              >
                <div 
                  :class="[
                    'text-sm md:text-base font-semibold mb-2',
                    day.isToday ? 'text-white' : day.isCurrentMonth ? 'text-gray-900' : 'text-gray-400'
                  ]"
                >
                  {{ day.day }}
                </div>
                <div v-if="day.events && day.events.length > 0" class="space-y-1">
                  <div 
                    v-for="event in day.events.slice(0, 2)" 
                    :key="event.id"
                    class="text-xs bg-brand-orange/15 text-brand-orange rounded-full px-2 py-1 truncate font-medium"
                  >
                    {{ event.title }}
                  </div>
                  <div v-if="day.events.length > 2" class="text-xs text-gray-500 font-medium">
                    +{{ day.events.length - 2 }} more
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Selected Date Events -->
          <div v-if="selectedDateEvents.length > 0" class="bg-white/95 rounded-[32px] shadow-2xl p-6 md:p-8 mt-8 border border-gray-100">
            <h3 class="text-2xl md:text-3xl font-bold text-gray-900 mb-6">
              Events on {{ formatFullDate(selectedDay) }}
            </h3>
            <div class="space-y-4">
              <div 
                v-for="event in selectedDateEvents" 
                :key="event.id"
                class="border-l-4 border-main pl-6 py-4 hover:bg-gray-50 transition-all rounded-r-lg cursor-pointer group bg-white"
                @click="openModal(event)"
              >
                <h4 class="text-xl font-bold text-gray-900 mb-2 group-hover:text-main transition-colors">{{ event.title }}</h4>
                <p class="text-gray-600 mb-3">{{ event.summary }}</p>
                <div class="flex flex-wrap items-center gap-4 text-sm">
                  <div v-if="event.time" class="flex items-center text-gray-600">
                    <svg class="w-4 h-4 mr-2 text-brand-orange" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <span class="font-medium">{{ event.time }}</span>
                  </div>
                  <div v-if="event.location" class="flex items-center text-gray-600">
                    <svg class="w-4 h-4 mr-2 text-main" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    <span class="font-medium">{{ event.location }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Event Modal -->
    <div
      v-if="selectedEvent"
      @click="closeModal"
      class="fixed inset-0 bg-black bg-opacity-60 z-50 flex items-center justify-center p-4 backdrop-blur-sm"
    >
      <div 
        @click.stop
        class="bg-white rounded-2xl max-w-4xl w-full max-h-[90vh] overflow-hidden shadow-2xl transform transition-all"
      >
        <!-- Modal Header -->
        <div class="sticky top-0 bg-gradient-to-r from-main to-secondary px-8 py-6 flex justify-between items-start z-10">
          <div class="flex-1">
            <h2 class="text-3xl font-bold text-white mb-4 pr-8">{{ selectedEvent.title }}</h2>
            <div class="flex flex-wrap items-center gap-4 text-sm">
              <div class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span class="font-medium">{{ formatFullDate(selectedEvent.date) }}</span>
              </div>
              <div v-if="selectedEvent.time" class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="font-medium">{{ selectedEvent.time }}</span>
              </div>
              <div v-if="selectedEvent.location" class="flex items-center text-white/90">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <span class="font-medium">{{ selectedEvent.location }}</span>
              </div>
            </div>
          </div>
          <button
            @click="closeModal"
            class="text-white hover:text-gray-200 transition-colors p-2 hover:bg-white/10 rounded-lg"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <!-- Modal Content -->
        <div class="px-8 py-8 overflow-y-auto max-h-[calc(90vh-180px)]">
          <div class="prose prose-lg max-w-none">
            <p class="text-gray-700 whitespace-pre-line leading-relaxed text-lg">{{ selectedEvent.content }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'

export default {
  name: 'Calendar',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const events = ref([])
    const selectedEvent = ref(null)
    const currentMonth = ref(new Date().getMonth())
    const currentYear = ref(new Date().getFullYear())
    const selectedDay = ref(null)

    // Helper function to get events for current/next month
    const generateEvents = () => {
      const today = new Date()
      const currentYear = today.getFullYear()
      const currentMonth = today.getMonth()
      const nextMonth = currentMonth === 11 ? 0 : currentMonth + 1
      const nextYear = currentMonth === 11 ? currentYear + 1 : currentYear
      
      const events = []
      let id = 1

      // Generate events for current month and next month
      for (let month = 0; month <= 1; month++) {
        const year = month === 0 ? currentYear : nextYear
        const monthIndex = month === 0 ? currentMonth : nextMonth
        
        // Get all days in the month
        const daysInMonth = new Date(year, monthIndex + 1, 0).getDate()
        
        for (let day = 1; day <= daysInMonth; day++) {
          const date = new Date(year, monthIndex, day)
          const dayOfWeek = date.getDay()
          
          // Sunday Worship Service (every Sunday)
          if (dayOfWeek === 0) {
            events.push({
              id: id++,
              title: 'Sunday Worship Service',
              summary: 'Join us every Sunday for worship, prayer, and fellowship.',
              content: `We gather every Sunday morning for worship and teaching from God's Word.\n\nService Schedule:\n- Praise and Worship: 9:00 AM\n- Prayer: 9:30 AM\n- Message: 10:00 AM\n- Fellowship: 11:30 AM\n\nAll are welcome! We look forward to worshipping with you.`,
              date: date,
              time: '9:00 AM',
              location: 'Main Sanctuary'
            })
            
            // Sunday School (every Sunday)
            events.push({
              id: id++,
              title: 'Sunday School',
              summary: 'Bible study classes for all ages.',
              content: `Join us for Sunday School classes for children, youth, and adults.\n\nAll classes run from 10:30 AM to 11:30 AM\nLocation: Various Classrooms\n\nCome grow in God's Word with fellow believers!`,
              date: date,
              time: '10:30 AM',
              location: 'Various Classrooms'
            })
          }
          
          // Youth Group (every Saturday)
          if (dayOfWeek === 6) {
            events.push({
              id: id++,
              title: 'Youth Group Meeting',
              summary: 'Fellowship and fun for teens and young adults.',
              content: `Our youth group meets every Saturday for games, discussion, and growing together in faith.\n\nAges: 13-25\nActivities include:\n- Group games\n- Bible discussion\n- Prayer time\n- Snacks and fellowship\n\nCome join us for an evening of fun and spiritual growth!`,
              date: date,
              time: '6:00 PM',
              location: 'Youth Center'
            })
          }
          
          // Prayer Night (first Saturday of month)
          if (dayOfWeek === 6 && day <= 7) {
            events.push({
              id: id++,
              title: 'Prayer Night',
              summary: 'Corporate prayer and intercession.',
              content: `Join us for our monthly prayer night where we lift up the needs of our church, community, and world.\n\nWe will pray for:\n- Our church family\n- Community needs\n- Mission work\n- Personal requests\n\nAll are welcome to join us in prayer.`,
              date: date,
              time: '7:00 PM',
              location: 'Main Sanctuary'
            })
          }
          
          // Baptism Service (third Sunday of month)
          if (dayOfWeek === 0 && day >= 15 && day <= 21) {
            events.push({
              id: id++,
              title: 'Baptism Service',
              summary: 'Public baptism and celebration of new life in Christ.',
              content: `Join us as we celebrate baptisms of new believers publicly professing their faith in Christ.\n\nThis is a special time of worship and celebration as we witness these powerful testimonies of God's grace.\n\nAfter the service, we'll have a fellowship celebration with refreshments.`,
              date: date,
              time: 'After Service',
              location: 'Main Sanctuary'
            })
          }
        }
      }

      return events
    }

    const sampleEvents = generateEvents()

    const formatFullDate = (date) => {
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`
    }

    const fetchEvents = async () => {
      loading.value = true
      
      try {
        await new Promise(resolve => setTimeout(resolve, 500))
        events.value = sampleEvents
        loading.value = false
      } catch (err) {
        console.error('Error loading events:', err)
        error.value = 'Failed to load calendar'
        loading.value = false
      }
    }

    const calendarDays = computed(() => {
      const year = currentYear.value
      const month = currentMonth.value
      
      const firstDayOfMonth = new Date(year, month, 1)
      const lastDayOfMonth = new Date(year, month + 1, 0)
      
      const startDate = new Date(firstDayOfMonth)
      startDate.setDate(startDate.getDate() - firstDayOfMonth.getDay())
      
      const endDate = new Date(lastDayOfMonth)
      endDate.setDate(endDate.getDate() + (6 - lastDayOfMonth.getDay()))
      
      const days = []
      const currentDate = new Date(startDate)
      
      while (currentDate <= endDate) {
        const dateKey = currentDate.toISOString().split('T')[0]
        const dayEvents = events.value.filter(event => {
          const eventDate = event.date.toISOString().split('T')[0]
          return eventDate === dateKey
        })
        
        days.push({
          date: dateKey,
          day: currentDate.getDate(),
          isCurrentMonth: currentDate.getMonth() === month,
          isToday: currentDate.toDateString() === new Date().toDateString(),
          hasEvents: dayEvents.length > 0,
          events: dayEvents
        })
        
        currentDate.setDate(currentDate.getDate() + 1)
      }
      
      return days
    })

    const selectedDateEvents = computed(() => {
      if (!selectedDay.value) return []
      
      const dateKey = selectedDay.value.toISOString().split('T')[0]
      return events.value.filter(event => {
        const eventDate = event.date.toISOString().split('T')[0]
        return eventDate === dateKey
      })
    })

    const currentMonthName = computed(() => {
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      return months[currentMonth.value]
    })

    const selectDate = (day) => {
      if (!day.isCurrentMonth) return
      const [year, month, date] = day.date.split('-').map(Number)
      selectedDay.value = new Date(year, month - 1, date)
    }

    const previousMonth = () => {
      if (currentMonth.value === 0) {
        currentMonth.value = 11
        currentYear.value--
      } else {
        currentMonth.value--
      }
      selectedDay.value = null
    }

    const nextMonth = () => {
      if (currentMonth.value === 11) {
        currentMonth.value = 0
        currentYear.value++
      } else {
        currentMonth.value++
      }
      selectedDay.value = null
    }

    const resetToCurrentMonth = () => {
      const now = new Date()
      currentMonth.value = now.getMonth()
      currentYear.value = now.getFullYear()
      selectedDay.value = null
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
      selectedEvent,
      currentMonth,
      currentYear,
      selectedDay,
      selectedDateEvents,
      currentMonthName,
      calendarDays,
      formatFullDate,
      selectDate,
      previousMonth,
      nextMonth,
      resetToCurrentMonth,
      openModal,
      closeModal
    }
  }
}
</script>

