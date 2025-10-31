<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center px-4">
        <h1 class="text-4xl md:text-5xl font-bold mb-4">Church News & Updates</h1>
        <p class="text-xl mb-8">Stay connected with what's happening in our community</p>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="max-w-4xl mx-auto py-12">
      <div class="text-center">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mb-4"></div>
        <p class="text-gray-600">Loading news...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="max-w-4xl mx-auto py-12">
      <div class="card bg-red-50 border-red-200">
        <p class="text-red-600 text-center">{{ error }}</p>
      </div>
    </div>

    <!-- News Articles -->
    <div v-else class="max-w-4xl mx-auto">
      <!-- News List -->
      <div v-if="newsItems.length > 0" class="space-y-8">
        <div 
          v-for="item in newsItems" 
          :key="item.id"
          class="card hover:shadow-xl transition-shadow cursor-pointer"
          @click="openModal(item)"
        >
          <div class="flex flex-col md:flex-row gap-6">
            <!-- Date Badge -->
            <div class="flex-shrink-0">
              <div class="w-20 h-20 bg-primary-600 text-white rounded-lg flex flex-col items-center justify-center">
                <span class="text-2xl font-bold">{{ formatDay(item.date) }}</span>
                <span class="text-xs uppercase">{{ formatMonth(item.date) }}</span>
              </div>
            </div>
            
            <!-- Content -->
            <div class="flex-1">
              <h2 class="text-2xl font-bold text-gray-900 mb-2">{{ item.title }}</h2>
              <p class="text-gray-600 mb-3">{{ item.summary }}</p>
              <div class="flex items-center text-sm text-gray-500">
                <span>{{ formatFullDate(item.date) }}</span>
                <span v-if="item.category" class="ml-4">
                  <span class="px-2 py-1 bg-primary-100 text-primary-700 rounded-full">{{ item.category }}</span>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="card text-center py-12">
        <div class="text-6xl mb-4">📰</div>
        <h3 class="text-2xl font-bold text-gray-900 mb-2">No News Yet</h3>
        <p class="text-gray-600">Check back soon for the latest updates from our church community.</p>
      </div>

      <!-- News Modal -->
      <div
        v-if="selectedNews"
        @click="closeModal"
        class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
      >
        <div 
          @click.stop
          class="bg-white rounded-lg max-w-3xl w-full max-h-[90vh] overflow-y-auto"
        >
          <div class="sticky top-0 bg-white border-b px-6 py-4 flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">{{ selectedNews.title }}</h2>
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
            <div class="flex items-center text-sm text-gray-500 mb-4">
              <span>{{ formatFullDate(selectedNews.date) }}</span>
              <span v-if="selectedNews.category" class="ml-4">
                <span class="px-2 py-1 bg-primary-100 text-primary-700 rounded-full">{{ selectedNews.category }}</span>
              </span>
            </div>
            
            <div class="prose prose-lg max-w-none">
              <p class="text-gray-700 whitespace-pre-line">{{ selectedNews.content }}</p>
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
  name: 'News',
  setup() {
    const loading = ref(true)
    const error = ref('')
    const newsItems = ref([])
    const selectedNews = ref(null)

    // Sample news data - In production, this would come from an API
    const sampleNews = [
      {
        id: 1,
        title: 'Welcome to Our New Website!',
        summary: 'We are excited to launch our new website and stay connected with our church family.',
        content: `We are thrilled to announce the launch of our new church website! This platform will help us stay better connected with our congregation and share important news, events, and updates.\n\nOur website features:\n- Weekly service schedules\n- Church news and announcements\n- Missionary updates\n- Online giving\n- Contact information\n\nWe hope this new website will serve as a valuable resource for our church community and help us spread the Gospel message further. Thank you for your continued support and prayers.`,
        date: new Date('2024-01-15'),
        category: 'Announcement'
      },
      {
        id: 2,
        title: 'Missionary Work Update from Pastor Rith',
        summary: 'Pastor Rith shares exciting updates about our missionary work in Cambodia.',
        content: `Dear Church Family,\n\nI want to share with you the amazing work that God is doing through our missionary efforts here in Cambodia. Over the past month, we have seen tremendous growth in our outreach programs.\n\nOur team has been visiting remote villages, sharing the Gospel, and providing essential supplies to families in need. The response from the local communities has been overwhelming, with many coming to know Christ and joining our fellowship.\n\nWe want to thank you for your continued prayers and financial support. Your partnership is making a real difference in people's lives. Please continue to pray for:\n- The safety of our missionary team\n- Open hearts in the communities we serve\n- God's provision for our needs\n- Revival in Cambodia\n\nGod bless you all!`,
        date: new Date('2024-01-08'),
        category: 'Missionary'
      },
      {
        id: 3,
        title: 'Special Christmas Celebration Service',
        summary: 'Join us for a special Christmas celebration service on December 24th at 6:00 PM.',
        content: `Christmas is a time to celebrate the birth of our Savior, Jesus Christ! We invite you and your family to join us for a special Christmas celebration service on December 24th at 6:00 PM.\n\nThis special service will feature:\n- Traditional Christmas carols\n- A children's Christmas program\n- A message from Pastor Rith\n- Special fellowship time with refreshments\n\nWe encourage everyone to invite friends and family. This is a wonderful opportunity to share the true meaning of Christmas and celebrate together as a church family.\n\nPlease come early as we expect a larger than usual attendance. We look forward to celebrating this special time with all of you!`,
        date: new Date('2023-12-20'),
        category: 'Event'
      },
      {
        id: 4,
        title: 'Thanksgiving for God\'s Faithfulness',
        summary: 'We give thanks for God\'s continued faithfulness to our church community.',
        content: `As we reflect on this past year, our hearts are filled with gratitude for all that God has done in and through our church family.\n\nWe have witnessed:\n- Souls being saved\n- Lives being transformed\n- Families being restored\n- Our community growing in faith\n\nThe Lord has been faithful in meeting our needs, guiding our decisions, and blessing our ministry. Despite the challenges we've faced, we have seen God's hand at work in powerful ways.\n\nLet us continue to trust in His promises and look forward to what He has in store for us in the coming year. Thank you for being part of this journey of faith.`,
        date: new Date('2023-11-23'),
        category: 'Thanksgiving'
      },
      {
        id: 5,
        title: 'Bible Study Series: The Book of Romans',
        summary: 'Join us for a new Bible study series exploring the Book of Romans, starting this Wednesday.',
        content: `We are excited to begin a new Wednesday evening Bible study series on the Book of Romans! This profound letter from the Apostle Paul contains some of the most important teachings in all of Scripture.\n\nOver the next several months, we will explore:\n- The righteousness of God\n- Salvation by grace through faith\n- The role of the Law\n- Life in the Spirit\n- God's plan for Israel\n- Practical Christian living\n\nThis study will be an excellent opportunity to deepen your understanding of the Gospel and grow in your faith. We encourage everyone to attend and bring a friend!\n\nWednesday evenings at 7:00 PM in the main sanctuary.`,
        date: new Date('2023-10-01'),
        category: 'Study'
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

    const fetchNews = async () => {
      loading.value = true
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 500))
        
        // In production, you would fetch from an API here
        newsItems.value = sampleNews.sort((a, b) => b.date - a.date)
        
        loading.value = false
      } catch (err) {
        console.error('Error loading news:', err)
        error.value = 'Failed to load news articles'
        loading.value = false
      }
    }

    const openModal = (item) => {
      selectedNews.value = item
    }

    const closeModal = () => {
      selectedNews.value = null
    }

    onMounted(() => {
      fetchNews()
    })

    return {
      loading,
      error,
      newsItems,
      selectedNews,
      formatDay,
      formatMonth,
      formatFullDate,
      openModal,
      closeModal
    }
  }
}
</script>

