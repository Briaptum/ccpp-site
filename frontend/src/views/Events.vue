<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center px-4">
        <h1 class="text-4xl md:text-5xl font-bold mb-4">Church Events</h1>
        <p class="text-xl mb-8">Join us for worship, fellowship, and community</p>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="max-w-4xl mx-auto py-12">
      <div class="text-center">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mb-4"></div>
        <p class="text-gray-600">Loading events...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="max-w-4xl mx-auto py-12">
      <div class="card bg-red-50 border-red-200">
        <p class="text-red-600 text-center">{{ error }}</p>
      </div>
    </div>

    <!-- Events List -->
    <div v-else class="max-w-4xl mx-auto">
      <!-- Events List -->
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
              <div class="w-20 h-20 bg-primary-600 text-white rounded-lg flex flex-col items-center justify-center">
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
                  <span class="px-2 py-1 bg-primary-100 text-primary-700 rounded-full">{{ event.time }}</span>
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

