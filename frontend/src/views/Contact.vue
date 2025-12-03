<template>
  <div class="px-4 sm:px-0">
    <div class="max-w-6xl mx-auto py-12">
      <div class="text-center mb-12">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-gray-900">Contact Us</h1>
        <p class="text-xl text-gray-700">We'd Love to Hear from You</p>
      </div>
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 mb-12">
        <!-- Contact Information -->
        <div>
          <div class="card">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">Get in Touch</h2>
            
            <div class="space-y-6">
              <div class="flex items-start space-x-4">
                <div class="text-2xl">📍</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Address</h3>
                  <p class="text-gray-600">
                    Street 26BT, House 428<br>
                    Sangkat Boeung Tompun, Khan Meanchey<br>
                    Phnom Penh, Cambodia
                  </p>
                </div>
              </div>
              
              <div class="flex items-start space-x-4">
                <div class="text-2xl">📞</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Phone</h3>
                  <p class="text-gray-600">855 77 777 470</p>
                </div>
              </div>
              
              <div class="flex items-start space-x-4">
                <div class="text-2xl">📧</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Email</h3>
                  <p class="text-gray-600">info@ccpp.org</p>
                </div>
              </div>
              
              <div class="flex items-start space-x-4">
                <div class="text-2xl">🕒</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Office Hours</h3>
                  <p class="text-gray-600">
                    Monday - Friday: 9:00 AM - 5:00 PM<br>
                    Saturday: 9:00 AM - 2:00 PM<br>
                    Sunday: Closed
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- Service Times -->
          <div class="card mt-8">
            <h3 class="text-xl font-bold text-gray-900 mb-4">Service Times</h3>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="font-medium">Sunday Worship</span>
                <span class="text-gray-600">10:00 AM</span>
              </div>
              <div class="flex justify-between">
                <span class="font-medium">Sunday School</span>
                <span class="text-gray-600">9:00 AM</span>
              </div>
              <div class="flex justify-between">
                <span class="font-medium">Wednesday Bible Study</span>
                <span class="text-gray-600">7:00 PM</span>
              </div>
              <div class="flex justify-between">
                <span class="font-medium">Friday Youth Group</span>
                <span class="text-gray-600">7:00 PM</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Contact Form -->
        <div>
          <div class="card">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">Send us a Message</h2>
            
            <form @submit.prevent="submitForm" class="space-y-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    First Name *
                  </label>
                  <input
                    v-model="form.firstName"
                    type="text"
                    required
                    class="input-field"
                    placeholder="Your first name"
                  />
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    Last Name *
                  </label>
                  <input
                    v-model="form.lastName"
                    type="text"
                    required
                    class="input-field"
                    placeholder="Your last name"
                  />
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Email *
                </label>
                <input
                  v-model="form.email"
                  type="email"
                  required
                  class="input-field"
                  placeholder="your.email@example.com"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Phone
                </label>
                <input
                  v-model="form.phone"
                  type="tel"
                  class="input-field"
                  placeholder="855 77 777 470"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Subject *
                </label>
                <select
                  v-model="form.subject"
                  required
                  class="input-field"
                >
                  <option value="">Select a subject</option>
                  <option value="general">General Inquiry</option>
                  <option value="visit">Planning a Visit</option>
                  <option value="membership">Membership Information</option>
                  <option value="ministry">Ministry Opportunities</option>
                  <option value="prayer">Prayer Request</option>
                  <option value="support">Pastoral Support</option>
                  <option value="other">Other</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Message *
                </label>
                <textarea
                  v-model="form.message"
                  required
                  rows="5"
                  class="input-field"
                  placeholder="How can we help you?"
                ></textarea>
              </div>
              
              <button
                type="submit"
                :disabled="submitting"
                class="btn-primary w-full"
                :class="{ 'opacity-50 cursor-not-allowed': submitting }"
              >
                {{ submitting ? 'Sending...' : 'Send Message' }}
              </button>
            </form>
            
            <!-- Success/Error Messages -->
            <div v-if="message" class="mt-4 p-4 rounded-lg" :class="messageType === 'success' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'">
              {{ message }}
            </div>
          </div>
        </div>
      </div>

      <!-- Map Section -->
      <div class="card mb-12">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Find Us</h2>
        <div class="relative w-full h-96 rounded-lg overflow-hidden shadow-lg">
          <iframe
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3908.7347!2d104.8975!3d11.5564!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0x0!2zMTHCsDMzJzIzLjAiTiAxMDTCsDUzJzUxLjAiRQ!5e0!3m2!1sen!2sus!4v1234567890"
            width="100%"
            height="100%"
            style="border:0;"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
            class="rounded-lg"
          ></iframe>
        </div>
        <div class="mt-4 text-center">
          <a 
            href="https://maps.app.goo.gl/99i0uWMZ5NJg16w2V" 
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center text-main hover:opacity-80 font-medium"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            Open in Google Maps
          </a>
        </div>
      </div>

      <!-- Additional Contact Options -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
        <div class="card text-center">
          <div class="text-4xl mb-4">📧</div>
          <h3 class="text-xl font-semibold mb-3">Email Us</h3>
          <p class="text-gray-600 mb-4">
            Send us an email and we'll get back to you within 24 hours.
          </p>
          <a href="mailto:info@ccpp.org" class="btn-primary">
            Send Email
          </a>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">📞</div>
          <h3 class="text-xl font-semibold mb-3">Call Us</h3>
          <p class="text-gray-600 mb-4">
            Speak directly with our staff during office hours.
          </p>
          <a href="tel:+85577777470" class="btn-primary">
            Call Now
          </a>
        </div>
        
        <div class="card text-center">
          <div class="text-4xl mb-4">📅</div>
          <h3 class="text-xl font-semibold mb-3">Schedule a Visit</h3>
          <p class="text-gray-600 mb-4">
            Plan your first visit and we'll make you feel welcome.
          </p>
          <button @click="scheduleVisit" class="btn-primary">
            Schedule Visit
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { contactService } from '@/services/contactService'

export default {
  name: 'Contact',
  setup() {
    const form = ref({
      firstName: '',
      lastName: '',
      email: '',
      phone: '',
      subject: '',
      message: ''
    })
    
    const submitting = ref(false)
    const message = ref('')
    const messageType = ref('')

    const submitForm = async () => {
      submitting.value = true
      message.value = ''
      
      try {
        // Submit form to backend
        await contactService.createContact(form.value)
        
        message.value = 'Thank you for your message! We\'ll get back to you soon.'
        messageType.value = 'success'
        
        // Reset form
        form.value = {
          firstName: '',
          lastName: '',
          email: '',
          phone: '',
          subject: '',
          message: ''
        }
      } catch (error) {
        message.value = error.message || 'Sorry, there was an error sending your message. Please try again.'
        messageType.value = 'error'
      } finally {
        submitting.value = false
      }
    }

    const scheduleVisit = () => {
      // This would typically open a scheduling modal or redirect to a booking system
      alert('Visit scheduling feature coming soon! Please call us at 855 77 777 470 to schedule your visit.')
    }

    return {
      form,
      submitting,
      message,
      messageType,
      submitForm,
      scheduleVisit
    }
  }
}
</script>
