<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Upcoming Events</h1>
        <p class="text-lg text-gray-700 mt-2">Join us for worship, fellowship, and community</p>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8">
        <!-- Loading State -->
        <div v-if="loading" class="py-12">
          <div class="text-center">
            <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-main mb-4"></div>
            <p class="text-gray-600">Loading events...</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="py-12">
          <div class="card bg-red-50 border-red-200">
            <p class="text-red-600 text-center">{{ error }}</p>
          </div>
        </div>

        <!-- Events List -->
        <div v-else>
          <div v-if="events.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div 
              v-for="event in events" 
              :key="event.id"
              class="group bg-white rounded-xl shadow-md hover:shadow-2xl transition-all duration-300 cursor-pointer overflow-hidden border border-gray-100 hover:border-main/20 transform hover:-translate-y-1 flex flex-col"
              @click="openModal(event)"
            >
              <!-- Date Badge -->
              <div class="bg-gradient-to-br from-main to-secondary p-6 flex flex-col items-center justify-center">
                <span class="text-4xl font-bold text-white mb-1">{{ formatDay(event.date) }}</span>
                <span class="text-sm uppercase text-white/90 font-medium tracking-wider">{{ formatMonth(event.date) }}</span>
                <span class="text-xs text-white/80 mt-1">{{ event.date.getFullYear() }}</span>
              </div>
              
              <!-- Content -->
              <div class="flex-1 p-6 flex flex-col">
                <h2 class="text-xl font-bold text-gray-900 mb-3 group-hover:text-main transition-colors line-clamp-2">
                  {{ event.title }}
                </h2>
                <p class="text-gray-600 mb-4 text-sm leading-relaxed line-clamp-3 flex-grow">{{ event.summary }}</p>
                <div class="flex flex-col gap-3 text-sm mt-auto">
                  <div class="flex items-center text-gray-600">
                    <svg class="w-4 h-4 mr-2 text-main flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <span class="font-medium text-xs">{{ formatFullDate(event.date) }}</span>
                  </div>
                  <div v-if="event.time" class="flex items-center">
                    <svg class="w-4 h-4 mr-2 text-custom-orange flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <span class="px-2 py-1 bg-custom-orange/10 text-custom-orange rounded-full font-medium text-xs">{{ event.time }}</span>
                  </div>
                  <div v-if="event.location" class="flex items-center">
                    <svg class="w-4 h-4 mr-2 text-main flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded-full font-medium text-xs truncate">{{ event.location }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-else class="card text-center py-16">
            <div class="text-6xl mb-6">📅</div>
            <h3 class="text-3xl font-bold text-gray-900 mb-3">No Upcoming Events</h3>
            <p class="text-gray-600 text-lg">Check back soon for upcoming church events and activities.</p>
          </div>
        </div>
      </div>
    </div>

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
import { ref, onMounted } from 'vue'

export default {
  name: 'UpcomingEvents',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const events = ref([])
    const selectedEvent = ref(null)

    // Helper function to get events for current month
    const generateEvents = () => {
      const today = new Date()
      const currentYear = today.getFullYear()
      const currentMonth = today.getMonth()
      
      const monthEvents = []
      let id = 1

      // Get all days in the current month
      const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate()
      
      for (let day = 1; day <= daysInMonth; day++) {
        const date = new Date(currentYear, currentMonth, day)
        const dayOfWeek = date.getDay()
        const todayDate = new Date()
        
        // Only include today and future dates in current month
        if (date.toDateString() < todayDate.toDateString()) {
          continue
        }
        
        // Sunday Worship Service (every Sunday)
        if (dayOfWeek === 0) {
          monthEvents.push({
            id: id++,
            title: 'Sunday Worship Service',
            summary: 'Join us every Sunday for worship, prayer, and fellowship.',
            content: `We gather every Sunday morning for worship and teaching from God's Word.\n\nService Schedule:\n- Praise and Worship: 9:00 AM\n- Prayer: 9:30 AM\n- Message: 10:00 AM\n- Fellowship: 11:30 AM\n\nAll are welcome! We look forward to worshipping with you.`,
            date: date,
            time: '9:00 AM',
            location: 'Main Sanctuary'
          })
          
          // Sunday School (every Sunday)
          monthEvents.push({
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
          monthEvents.push({
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
          monthEvents.push({
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
          monthEvents.push({
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

      return monthEvents
    }

    const sampleEvents = generateEvents()

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

