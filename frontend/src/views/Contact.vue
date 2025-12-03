<template>
  <div class="px-4 sm:px-0 contact-page">
    <div class="max-w-6xl mx-auto py-12 contact-shell">
      <div class="text-center mb-12 contact-hero">
        <h1 class="text-4xl md:text-6xl font-bold mb-4 text-gray-900">Contact Us</h1>
        <p class="text-xl text-gray-700">We'd Love to Hear from You</p>
      </div>
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 mb-12">
        <!-- Contact Information -->
        <div>
          <div class="card contact-card info-card">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">Get in Touch</h2>
            
            <div class="space-y-6">
              <div class="flex items-start space-x-4 contact-info-row">
                <div class="contact-icon">📍</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Address</h3>
                  <p class="text-gray-600">
                    Street 26BT, House 428<br>
                    Boeung Tompun<br>
                    Phnom Penh, Cambodia
                  </p>
                </div>
              </div>
              
              <div class="flex items-start space-x-4 contact-info-row">
                <div class="contact-icon">📞</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Phone</h3>
                  <p class="text-gray-600">+855 15 81 44 40</p>
                </div>
              </div>
              
              <div class="flex items-start space-x-4 contact-info-row">
                <div class="contact-icon">📧</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Email</h3>
                  <p class="text-gray-600">calvarychapelphnompenh@gmail.com</p>
                </div>
              </div>

            </div>
          </div>

          <!-- Service Times -->
          <div class="card contact-card mt-8 contact-hours">
            <h3 class="text-xl font-bold text-gray-900 mb-4">Service Times</h3>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="font-medium">Sunday Worship</span>
                <span class="text-gray-600">10:00 AM</span>
              </div>

              <div class="flex justify-between">
                <span class="font-medium">Wednesday Bible Study</span>
                <span class="text-gray-600">7:00 PM</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Contact Form -->
        <div>
          <div class="card">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">Send us a Message</h2>
            
            <form @submit.prevent="submitForm" class="space-y-6 contact-form">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    First Name *
                  </label>
                  <input
                    v-model="form.firstName"
                    type="text"
                    required
                    class="input-field contact-input"
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
                    class="input-field contact-input"
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
                  class="input-field contact-input"
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
                  class="input-field contact-input"
                  placeholder="+855 15 81 44 40"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Subject *
                </label>
                <select
                  v-model="form.subject"
                  required
                  class="input-field contact-input"
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
                  class="input-field contact-input"
                  placeholder="How can we help you?"
                ></textarea>
              </div>
              
              <button
                type="submit"
                :disabled="submitting"
                class="btn-primary w-full contact-submit"
                :class="{ 'opacity-50 cursor-not-allowed': submitting }"
              >
                {{ submitting ? 'Sending...' : 'Send Message' }}
              </button>
            </form>
            
            <!-- Success/Error Messages -->
            <div
              v-if="message"
              class="mt-4 p-4 rounded-lg contact-feedback"
              :class="messageType === 'success' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'"
            >
              {{ message }}
            </div>
          </div>
        </div>
      </div>

      <!-- Map Section -->
      <div class="card contact-card map-card mb-12">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Find Us</h2>
        <div class="relative w-full h-96 rounded-lg overflow-hidden shadow-xl map-frame">
          <iframe
            src="https://www.google.com/maps?q=Street+26BT,+House+428,+Phnom+Penh,+Cambodia&output=embed"
            width="100%"
            height="100%"
            style="border:0;"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
            class="rounded-lg map-iframe"
          ></iframe>
        </div>
        <div class="mt-4 text-center">
          <a 
            href="https://maps.app.goo.gl/a4YGCcvWRLy6BvRz9" 
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center hover:opacity-80 font-medium map-link"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            Open in Google Maps
          </a>
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
      alert('Visit scheduling feature coming soon! Please call us at +855 15 81 44 40 to schedule your visit.')
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

<style scoped>
.contact-page {
  position: relative;
  min-height: 100vh;
  padding-bottom: 4rem;
  background: linear-gradient(180deg, #f6f8fb 0%, #ffffff 60%, #ffffff 100%);
}

.contact-page::before,
.contact-page::after {
  content: "";
  position: absolute;
  border-radius: 9999px;
  filter: blur(80px);
  opacity: 0.35;
}

.contact-page::before {
  width: 24rem;
  height: 24rem;
  top: -6rem;
  right: -4rem;
  background: radial-gradient(circle, rgba(34, 112, 130, 0.25), transparent 70%);
}

.contact-page::after {
  width: 18rem;
  height: 18rem;
  bottom: 4rem;
  left: -3rem;
  background: radial-gradient(circle, rgba(228, 157, 64, 0.25), transparent 70%);
}

.contact-shell {
  position: relative;
  z-index: 1;
}

.contact-hero {
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid rgba(15, 23, 42, 0.05);
  border-radius: 1.75rem;
  padding: 3rem 1.5rem;
  box-shadow: 0 45px 90px -50px rgba(15, 23, 42, 0.7);
  backdrop-filter: blur(8px);
}

.contact-hero h1 {
  letter-spacing: -0.02em;
}

.contact-card {
  position: relative;
  border-radius: 1.5rem;
  border: 1px solid rgba(15, 23, 42, 0.05);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, #ffffff 100%);
  box-shadow: 0 35px 80px -45px rgba(15, 23, 42, 0.65);
}

.contact-card::after {
  content: "";
  position: absolute;
  inset: 1px;
  border-radius: 1.4rem;
  border: 1px solid rgba(255, 255, 255, 0.7);
  pointer-events: none;
}

.info-card {
  overflow: hidden;
}

.contact-info-row {
  padding-bottom: 1rem;
  border-bottom: 1px solid rgba(15, 23, 42, 0.05);
}

.contact-info-row:last-of-type {
  border-bottom: none;
  padding-bottom: 0;
}

.contact-icon {
  width: 3rem;
  height: 3rem;
  border-radius: 9999px;
  display: grid;
  place-items: center;
  font-size: 1.5rem;
  background: rgba(34, 112, 130, 0.08);
  color: #227082;
}

.contact-hours {
  background: linear-gradient(135deg, rgba(34, 112, 130, 0.08), rgba(255, 255, 255, 0.9));
}

.contact-form {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 1.25rem;
}

.contact-input {
  border: 1px solid rgba(15, 23, 42, 0.1);
  background-color: rgba(255, 255, 255, 0.95);
  transition: border-color 0.2s ease, box-shadow 0.2s ease, transform 0.2s ease;
}

.contact-input:focus {
  border-color: #E49D40;
  box-shadow: 0 0 0 3px rgba(228, 157, 64, 0.25);
}

.contact-submit {
  border-radius: 9999px;
  padding-top: 0.85rem;
  padding-bottom: 0.85rem;
  box-shadow: 0 25px 45px -25px rgba(228, 157, 64, 0.9);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.contact-submit:not(:disabled):hover {
  transform: translateY(-2px);
  box-shadow: 0 35px 55px -25px rgba(228, 157, 64, 1);
}

.contact-feedback {
  border-radius: 1rem;
  border: 1px solid rgba(15, 23, 42, 0.04);
}

.map-card {
  border-radius: 1.75rem;
  overflow: hidden;
}

.map-frame {
  border-radius: 1.5rem;
  border: 1px solid rgba(15, 23, 42, 0.05);
}

.map-iframe {
  filter: saturate(1.05);
}

.map-link {
  color: #227082;
  text-decoration: none;
  transition: color 0.2s ease;
}

.map-link:hover {
  color: #E49D40;
}

.contact-options {
  margin-bottom: 6rem;
}

.contact-tile {
  border-radius: 1.5rem;
  border: 1px solid rgba(15, 23, 42, 0.05);
  background: linear-gradient(150deg, rgba(34, 112, 130, 0.06), rgba(255, 255, 255, 0.95));
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}

.contact-tile:hover {
  transform: translateY(-6px);
  box-shadow: 0 30px 60px -35px rgba(15, 23, 42, 0.55);
}

.contact-tile-icon {
  width: 3rem;
  height: 3rem;
  margin: 0 auto 1rem;
}

.contact-option-btn {
  width: 100%;
  border-radius: 9999px;
}

@media (min-width: 768px) {
  .contact-hero {
    padding: 3.5rem 4rem;
  }
}
</style>
