<template>
  <AdminLayout>
      <!-- Filters -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="flex flex-wrap items-center gap-4">
          <select class="px-4 py-2 border border-gray-300 rounded-lg text-sm">
            <option value="all">All Messages</option>
            <option value="unread">Unread</option>
            <option value="read">Read</option>
            <option value="replied">Replied</option>
          </select>
          <input
            type="text"
            placeholder="Search messages..."
            class="px-4 py-2 border border-gray-300 rounded-lg text-sm flex-1 min-w-64"
          />
          <button class="px-4 py-2 bg-main text-white rounded-lg hover:opacity-90 transition-colors text-sm">
            Search
          </button>
        </div>
      </div>

      <!-- Messages List -->
      <div class="bg-white rounded-lg shadow">
        <!-- Empty State -->
        <div class="text-center py-12">
          <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
          </svg>
          <p class="text-gray-600 mb-2">No messages yet</p>
          <p class="text-sm text-gray-500">Contact form submissions will appear here</p>
        </div>

        <!-- Messages Table (will show when messages exist) -->
        <!--
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Email</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Subject</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-6 py-4">John Doe</td>
                <td class="px-6 py-4">john@example.com</td>
                <td class="px-6 py-4">Question about services</td>
                <td class="px-6 py-4">2024-01-15</td>
                <td class="px-6 py-4"><span class="px-2 py-1 bg-yellow-100 text-yellow-800 rounded text-xs">Unread</span></td>
                <td class="px-6 py-4">
                  <button class="text-blue-600 hover:text-blue-800">View</button>
                  <button class="text-red-600 hover:text-red-800 ml-4">Delete</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        -->
      </div>
    </AdminLayout>
  </template>

<script>
import AdminLayout from '@/components/admin/AdminLayout.vue'

export default {
  name: 'ContactMessages',
  components: {
    AdminLayout
  }
}
</script>

