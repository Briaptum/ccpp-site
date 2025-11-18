<template>
  <div class="px-4 sm:px-0">
    <div class="max-w-4xl mx-auto py-12">
      <div class="text-center mb-12">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-gray-900">Church Calendar</h1>
        <p class="text-xl text-gray-700">Monthly view of all our church activities</p>
      </div>

    <!-- Loading State -->
    <div v-if="loading" class="py-12">
      <div class="text-center">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mb-4"></div>
        <p class="text-gray-600">Loading calendar...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="py-12">
      <div class="card bg-red-50 border-red-200">
        <p class="text-red-600 text-center">{{ error }}</p>
      </div>
    </div>

    <!-- Calendar View -->
    <div v-else>
      <div class="card">
        <div class="text-center mb-6">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">{{ currentMonthName }} {{ currentYear }}</h2>
          <div class="flex justify-between items-center mb-4">
            <button 
              @click="previousMonth"
              class="text-primary-600 hover:text-primary-800 transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
              </svg>
            </button>
            <button 
              @click="resetToCurrentMonth"
              class="btn-primary"
            >
              Today
            </button>
            <button 
              @click="nextMonth"
              class="text-primary-600 hover:text-primary-800 transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- Calendar Grid -->
        <div class="grid grid-cols-7 gap-1 mb-4">
          <div v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" :key="day" 
               class="text-center font-semibold text-gray-700 py-2">
            {{ day }}
          </div>
        </div>

        <div class="grid grid-cols-7 gap-1">
          <div 
            v-for="day in calendarDays" 
            :key="day.date"
            :class="[
              'min-h-20 p-2 border rounded cursor-pointer transition-colors',
              day.isCurrentMonth ? 'bg-white hover:bg-gray-50' : 'bg-gray-100 text-gray-400',
              day.isToday ? 'ring-2 ring-primary-600' : '',
              day.hasEvents ? 'border-primary-300' : 'border-gray-200'
            ]"
            @click="selectDate(day)"
          >
            <div class="text-sm font-medium mb-1">{{ day.day }}</div>
            <div v-if="day.events && day.events.length > 0" class="space-y-1">
              <div 
                v-for="event in day.events.slice(0, 2)" 
                :key="event.id"
                class="text-xs bg-primary-100 text-primary-700 rounded px-1 py-0.5 truncate"
              >
                {{ event.title }}
              </div>
              <div v-if="day.events.length > 2" class="text-xs text-gray-500">
                +{{ day.events.length - 2 }} more
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Selected Date Events -->
      <div v-if="selectedDateEvents.length > 0" class="card mt-8">
        <h3 class="text-2xl font-bold text-gray-900 mb-4">
          Events on {{ formatFullDate(selectedDay) }}
        </h3>
        <div class="space-y-4">
          <div 
            v-for="event in selectedDateEvents" 
            :key="event.id"
            class="border-l-4 border-primary-600 pl-4 py-2 hover:bg-gray-50 transition-colors"
            @click="openModal(event)"
          >
            <h4 class="font-semibold text-gray-900">{{ event.title }}</h4>
            <p class="text-sm text-gray-600">{{ event.summary }}</p>
            <div class="flex items-center text-xs text-gray-500 mt-2">
              <span v-if="event.time">{{ event.time }}</span>
              <span v-if="event.location" class="ml-3">{{ event.location }}</span>
            </div>
          </div>
        </div>
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

