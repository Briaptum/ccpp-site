<template>
  <AdminLayout>
    <div class="mb-6 flex justify-end">
      <button
        @click="showCreateModal = true"
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
          <div class="flex justify-between items-center">
            <h2 class="text-xl font-semibold text-gray-900">Upcoming Events</h2>
            <div class="flex items-center space-x-2">
              <input
                type="text"
                placeholder="Search events..."
                class="px-4 py-2 border border-gray-300 rounded-lg text-sm"
              />
              <select class="px-4 py-2 border border-gray-300 rounded-lg text-sm">
                <option value="all">All Events</option>
                <option value="upcoming">Upcoming</option>
                <option value="past">Past</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div class="text-center py-12">
          <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
          <p class="text-gray-600 mb-2">No events created yet</p>
          <p class="text-sm text-gray-500 mb-4">Create your first event to get started</p>
          <button
            @click="showCreateModal = true"
            class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
          >
            Create Event
          </button>
        </div>

        <!-- Events Table (will show when events are created) -->
        <!--
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Title</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Location</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-6 py-4">Event Title</td>
                <td class="px-6 py-4">Date</td>
                <td class="px-6 py-4">Location</td>
                <td class="px-6 py-4">
                  <button class="text-blue-600 hover:text-blue-800">Edit</button>
                  <button class="text-red-600 hover:text-red-800 ml-4">Delete</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        -->
      </div>

    <!-- Create Event Modal -->
    <div
      v-if="showCreateModal"
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      @click="showCreateModal = false"
    >
      <div
        class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <div class="p-6 border-b border-gray-200">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">Create New Event</h3>
            <button
              @click="showCreateModal = false"
              class="text-gray-400 hover:text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <form class="p-6 space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Event Title *
            </label>
            <input
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              placeholder="Sunday Worship Service"
            />
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Date *
              </label>
              <input
                type="date"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Time *
              </label>
              <input
                type="time"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              />
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Location
            </label>
            <input
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              placeholder="Church Main Hall"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Description
            </label>
            <textarea
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
              placeholder="Event description..."
            ></textarea>
          </div>
          
          <div class="flex justify-end space-x-4 pt-4 border-t border-gray-200">
            <button
              type="button"
              @click="showCreateModal = false"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
            >
              Create Event
            </button>
          </div>
        </form>
      </div>
    </div>
    </AdminLayout>
  </template>

<script>
import { ref } from 'vue'
import AdminLayout from '@/components/admin/AdminLayout.vue'

export default {
  name: 'EventManagement',
  components: {
    AdminLayout
  },
  setup() {
    const showCreateModal = ref(false)
    
    return {
      showCreateModal
    }
  }
}
</script>

