<template>
  <AdminLayout>
      <!-- Upload Section -->
      <div class="bg-white rounded-xl shadow-lg p-6 mb-8 border border-gray-100">
        <h2 class="text-xl font-bold text-gray-900 mb-6">Upload New Images</h2>
        
        <div class="border-2 border-dashed border-gray-300 rounded-xl p-12 text-center hover:border-main transition-all bg-gray-50 hover:bg-primary">
          <div class="mx-auto w-16 h-16 bg-secondary rounded-full flex items-center justify-center mb-4">
            <svg class="h-8 w-8 text-main" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
            </svg>
          </div>
          <p class="text-gray-700 font-medium mb-2">Drag and drop images here, or click to select</p>
          <p class="text-sm text-gray-500 mb-6">Supports JPG, PNG, GIF (Max 10MB per image)</p>
          <button
            class="px-6 py-3 bg-main text-white rounded-lg hover:opacity-90 transition-colors font-medium shadow-md"
          >
            Select Images
          </button>
        </div>

        <!-- Category Selection -->
        <div class="mt-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Category
          </label>
          <select class="w-full md:w-64 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-main focus:border-main bg-white">
            <option value="youth">Youth Ministry</option>
            <option value="grace-church">Grace Church</option>
            <option value="outreaches">Outreaches</option>
            <option value="general">General</option>
          </select>
        </div>
      </div>

      <!-- Gallery Images Grid -->
      <div class="bg-white rounded-xl shadow-lg p-6 border border-gray-100">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 gap-4">
          <h2 class="text-xl font-bold text-gray-900">Gallery Images</h2>
          <div class="flex items-center space-x-4 w-full sm:w-auto">
            <select class="flex-1 sm:flex-none px-4 py-2 border border-gray-300 rounded-lg text-sm bg-white focus:ring-2 focus:ring-main focus:border-main">
              <option value="all">All Categories</option>
              <option value="youth">Youth Ministry</option>
              <option value="grace-church">Grace Church</option>
              <option value="outreaches">Outreaches</option>
            </select>
            <button class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors text-sm font-medium shadow-md">
              Delete Selected
            </button>
          </div>
        </div>

        <!-- Empty State -->
        <div class="text-center py-12">
          <svg class="mx-auto h-12 w-12 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
          <p class="text-gray-600 mb-2">No images uploaded yet</p>
          <p class="text-sm text-gray-500">Upload your first image to get started</p>
        </div>

        <!-- Image Grid (will show when images are uploaded) -->
        <!--
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
          <div class="relative group">
            <img src="..." alt="..." class="w-full h-32 object-cover rounded-lg" />
            <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-50 transition-opacity rounded-lg flex items-center justify-center">
              <button class="text-white opacity-0 group-hover:opacity-100">Delete</button>
            </div>
          </div>
        </div>
        -->
      </div>
    </AdminLayout>
  </template>

<script>
import AdminLayout from '@/components/admin/AdminLayout.vue'

export default {
  name: 'GalleryManagement',
  components: {
    AdminLayout
  }
}
</script>

