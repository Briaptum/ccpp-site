<template>
  <AdminLayout>
    <div class="mb-6 flex justify-end">
      <button
        @click="openCreateModal"
        class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors shadow-lg font-medium flex items-center space-x-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        <span>New Event</span>
      </button>
    </div>
    
    <!-- Events List -->
    <div class="bg-white rounded-lg shadow">
      <div class="p-6 border-b border-gray-200">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <h2 class="text-xl font-semibold text-gray-900">Events</h2>
          <div class="flex items-center space-x-2 w-full sm:w-auto">
            <input
              v-model="searchQuery"
              @input="filterEvents"
              type="text"
              placeholder="Search events..."
              class="flex-1 sm:flex-none px-4 py-2 border border-gray-300 rounded-lg text-sm min-w-48"
            />
            <select 
              v-model="filterType"
              @change="loadEvents"
              class="px-4 py-2 border border-gray-300 rounded-lg text-sm"
            >
              <option value="all">All Events</option>
              <option value="upcoming">Upcoming</option>
              <option value="past">Past</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"></div>
        <p class="text-gray-600">Loading events...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-12">
        <p class="text-red-600 mb-4">{{ error }}</p>
        <button 
          @click="loadEvents"
          class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
        >
          Retry
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredEvents.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
        </svg>
        <p class="text-gray-600 mb-2">{{ searchQuery ? 'No events found' : 'No events created yet' }}</p>
        <p class="text-sm text-gray-500 mb-4">{{ searchQuery ? 'Try a different search term' : 'Create your first event to get started' }}</p>
        <button
          v-if="!searchQuery"
          @click="openCreateModal"
          class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors"
        >
          Create Event
        </button>
      </div>

      <!-- Events Table -->
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Time</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Location</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr 
              v-for="event in filteredEvents" 
              :key="event.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ event.title }}</div>
                <div v-if="event.summary" class="text-sm text-gray-500 truncate max-w-xs">{{ event.summary }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ formatDate(event.date) }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ event.time || 'N/A' }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ event.location || 'N/A' }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button 
                  @click="openEditModal(event)"
                  class="text-blue-600 hover:text-blue-800 mr-4"
                >
                  Edit
                </button>
                <button 
                  @click="deleteEvent(event.id)"
                  class="text-red-600 hover:text-red-800"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Create/Edit Event Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      @click="closeModal"
    >
      <div
        class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-6 border-b border-gray-200">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">
              {{ editingEvent ? 'Edit Event' : 'Create New Event' }}
            </h3>
            <button
              @click="closeModal"
              class="text-gray-400 hover:text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <form @submit.prevent="saveEvent" class="p-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Event Title *
            </label>
            <input
              v-model="eventForm.title"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Sunday Worship Service"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Summary (Short Description)
            </label>
            <input
              v-model="eventForm.summary"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Brief summary of the event"
            />
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Date *
              </label>
              <input
                v-model="eventForm.date"
                type="date"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Time
              </label>
              <input
                v-model="eventForm.time"
                type="time"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              />
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Location
            </label>
            <input
              v-model="eventForm.location"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Church Main Hall"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Description
            </label>
            <textarea
              v-model="eventForm.description"
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main"
              placeholder="Event description..."
            ></textarea>
          </div>

          <!-- Error Message -->
          <div v-if="formError" class="p-4 bg-red-50 border border-red-200 rounded-lg">
            <p class="text-red-600 text-sm">{{ formError }}</p>
          </div>
          
          <div class="flex justify-end space-x-4 pt-4 border-t border-gray-200">
            <button
              type="button"
              @click="closeModal"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="saving"
              class="px-6 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {{ saving ? 'Saving...' : (editingEvent ? 'Update Event' : 'Create Event') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </AdminLayout>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import AdminLayout from '@/components/admin/AdminLayout.vue'
import { eventService } from '@/services/eventService'

export default {
  name: 'EventManagement',
  components: {
    AdminLayout
  },
  setup() {
    const showModal = ref(false)
    const editingEvent = ref(null)
    const events = ref([])
    const loading = ref(false)
    const error = ref('')
    const saving = ref(false)
    const formError = ref('')
    const filterType = ref('all')
    const searchQuery = ref('')

    const eventForm = ref({
      title: '',
      summary: '',
      description: '',
      date: '',
      time: '',
      location: ''
    })

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      })
    }

    const loadEvents = async () => {
      loading.value = true
      error.value = ''
      try {
        const upcoming = filterType.value === 'upcoming'
        const response = await eventService.getEvents(upcoming)
        events.value = response.events || []
      } catch (err) {
        console.error('Error loading events:', err)
        error.value = err.message || 'Failed to load events'
      } finally {
        loading.value = false
      }
    }

    const filterEvents = () => {
      // Filtering is handled by computed property
    }

    const filteredEvents = computed(() => {
      let filtered = [...events.value]

      // Apply search filter
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase()
        filtered = filtered.filter(event => 
          event.title.toLowerCase().includes(query) ||
          (event.summary && event.summary.toLowerCase().includes(query)) ||
          (event.description && event.description.toLowerCase().includes(query)) ||
          (event.location && event.location.toLowerCase().includes(query))
        )
      }

      // Apply past filter
      if (filterType.value === 'past') {
        const now = new Date()
        filtered = filtered.filter(event => new Date(event.date) < now)
      }

      return filtered.sort((a, b) => new Date(a.date) - new Date(b.date))
    })

    const openCreateModal = () => {
      editingEvent.value = null
      eventForm.value = {
        title: '',
        summary: '',
        description: '',
        date: '',
        time: '',
        location: ''
      }
      formError.value = ''
      showModal.value = true
    }

    const openEditModal = (event) => {
      editingEvent.value = event
      const eventDate = new Date(event.date)
      eventForm.value = {
        title: event.title,
        summary: event.summary || '',
        description: event.description || '',
        date: eventDate.toISOString().split('T')[0],
        time: event.time || '',
        location: event.location || ''
      }
      formError.value = ''
      showModal.value = true
    }

    const closeModal = () => {
      showModal.value = false
      editingEvent.value = null
      formError.value = ''
    }

    const saveEvent = async () => {
      saving.value = true
      formError.value = ''

      try {
        const eventData = {
          ...eventForm.value,
          date: new Date(eventForm.value.date).toISOString()
        }

        if (editingEvent.value) {
          await eventService.updateEvent(editingEvent.value.id, eventData)
        } else {
          await eventService.createEvent(eventData)
        }

        await loadEvents()
        closeModal()
      } catch (err) {
        console.error('Error saving event:', err)
        formError.value = err.message || 'Failed to save event'
      } finally {
        saving.value = false
      }
    }

    const deleteEvent = async (id) => {
      if (!confirm('Are you sure you want to delete this event?')) {
        return
      }

      try {
        await eventService.deleteEvent(id)
        await loadEvents()
      } catch (err) {
        console.error('Error deleting event:', err)
        alert(`Failed to delete event: ${err.message}`)
      }
    }

    onMounted(() => {
      loadEvents()
    })

    return {
      showModal,
      editingEvent,
      events,
      loading,
      error,
      saving,
      formError,
      filterType,
      searchQuery,
      eventForm,
      filteredEvents,
      formatDate,
      loadEvents,
      filterEvents,
      openCreateModal,
      openEditModal,
      closeModal,
      saveEvent,
      deleteEvent
    }
  }
}
</script>
