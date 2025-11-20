<template>
  <div class="px-4 sm:px-0">
    <div class="max-w-4xl mx-auto py-12">
      <div class="text-center mb-12">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-gray-900">Upcoming Events</h1>
        <p class="text-xl text-gray-700">Join us for worship, fellowship, and community</p>
      </div>

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
      <div v-if="events.length > 0" class="space-y-8">
        <div 
          v-for="event in events" 
          :key="event.id"
          class="card hover:shadow-xl transition-shadow cursor-pointer"
          @click="openModal(event)"
        >
          <div class="flex flex-col md:flex-row gap-6">
            <!-- Date Badge -->
            <div class="flex-shrink-0">
              <div class="w-20 h-20 bg-main text-white rounded-lg flex flex-col items-center justify-center">
                <span class="text-2xl font-bold">{{ formatDay(event.date) }}</span>
                <span class="text-xs uppercase">{{ formatMonth(event.date) }}</span>
              </div>
            </div>
            
            <!-- Content -->
            <div class="flex-1">
              <h2 class="text-2xl font-bold text-gray-900 mb-2">{{ event.title }}</h2>
              <p class="text-gray-600 mb-3">{{ event.summary }}</p>
              <div class="flex items-center text-sm text-gray-500">
                <span>{{ formatFullDate(event.date) }}</span>
                <span v-if="event.time" class="ml-4">
                  <span class="px-2 py-1 bg-secondary text-main rounded-full">{{ event.time }}</span>
                </span>
                <span v-if="event.location" class="ml-4">
                  <span class="px-2 py-1 bg-gray-100 text-gray-700 rounded-full">{{ event.location }}</span>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="card text-center py-12">
        <div class="text-6xl mb-4">📅</div>
        <h3 class="text-2xl font-bold text-gray-900 mb-2">No Upcoming Events</h3>
        <p class="text-gray-600">Check back soon for upcoming church events and activities.</p>
      </div>

      <!-- Event Modal -->
      <div
        v-if="selectedEvent"
        @click="closeModal"
        class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      >
        <div 
          @click.stop
          class="bg-white rounded-lg max-w-3xl w-full max-h-[90vh] overflow-y-auto"
        >
          <div class="sticky top-0 bg-white border-b px-6 py-4 flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">{{ selectedEvent.title }}</h2>
            <button
              @click="closeModal"
              class="text-gray-500 hover:text-gray-700 transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          
          <div class="px-6 py-6">
            <div class="flex items-center text-sm text-gray-500 mb-4 flex-wrap gap-2">
              <span>{{ formatFullDate(selectedEvent.date) }}</span>
              <span v-if="selectedEvent.time" class="px-2 py-1 bg-primary-100 text-primary-700 rounded-full">
                {{ selectedEvent.time }}
              </span>
              <span v-if="selectedEvent.location" class="px-2 py-1 bg-gray-100 text-gray-700 rounded-full">
                {{ selectedEvent.location }}
              </span>
            </div>
            
            <div class="prose prose-lg max-w-none">
              <p class="text-gray-700 whitespace-pre-line">{{ selectedEvent.content }}</p>
            </div>
          </div>
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

