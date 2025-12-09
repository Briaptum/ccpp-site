<template>
  <div class="bg-white">
    <!-- Hero -->
    <section class="relative min-h-[55vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-16 bg-gradient-to-br from-[#0f172a] via-[#11263f] to-[#1d3a54] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-30 mix-blend-screen"
          style="background-image: radial-gradient(circle at 20% 30%, rgba(255,255,255,0.25), transparent 45%), radial-gradient(circle at 70% 15%, rgba(255,255,255,0.18), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.2), transparent 30%);"
        ></div>
        <div class="absolute inset-y-[-35%] left-[-15%] w-1/2 bg-gradient-to-r from-brand-orange/30 via-transparent to-transparent blur-[150px] opacity-70"></div>
        <div class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-70"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/70 via-black/40 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.35em] text-white/70">We'd love to hear from you</p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Contact Calvary Chapel Phnom Penh</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          Questions, prayer requests, or planning a visit? Reach out and our team will respond quickly.
        </p>
        <div class="flex justify-center gap-4 pt-4 text-white/80">
          <a
            href="https://www.facebook.com/calvarychapelpp"
            target="_blank"
            rel="noopener noreferrer"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="Facebook"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
            </svg>
          </a>
          <a
            href="mailto:calvarychapelphnompenh@gmail.com"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="Email"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </a>
          <a
            href="https://www.youtube.com/@CalvaryChapelPhnomPenh"
            target="_blank"
            rel="noopener noreferrer"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="YouTube"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>
            </svg>
          </a>
        </div>
      </div>
    </section>

    <!-- Info + Form -->
    <section class="py-16 md:py-20 bg-gradient-to-b from-white via-gray-50 to-gray-100">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid gap-10 lg:grid-cols-2">
        <div class="space-y-8">
          <div class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-6">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Visit & connect</p>
            <h2 class="text-3xl font-semibold text-gray-900">Get in touch</h2>
            <p class="text-gray-600">
              Our team is available throughout the week to answer questions, pray with you, and help you plug in.
            </p>
            <div class="space-y-6">
              <div class="flex gap-4">
                <div class="rounded-2xl bg-brand-blue/10 text-brand-blue flex items-center justify-center h-14 w-14 text-2xl">📍</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Address</h3>
                  <p class="text-gray-600 leading-relaxed">
                    Street 26BT, House 428<br />
                    Boeung Tompun<br />
                    Phnom Penh, Cambodia
                  </p>
                </div>
              </div>
              <div class="flex gap-4">
                <div class="rounded-2xl bg-brand-blue/10 text-brand-blue flex items-center justify-center h-14 w-14 text-2xl">📞</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Phone</h3>
                  <p class="text-gray-600">+855 15 81 44 40</p>
                </div>
              </div>
              <div class="flex gap-4">
                <div class="rounded-2xl bg-brand-blue/10 text-brand-blue flex items-center justify-center h-14 w-14 text-2xl">📧</div>
                <div>
                  <h3 class="font-semibold text-gray-900">Email</h3>
                  <p class="text-gray-600">calvarychapelphnompenh@gmail.com</p>
                </div>
              </div>
            </div>
          </div>

          <div class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-orange/80">Gatherings</p>
            <h3 class="text-2xl font-semibold text-gray-900">Regular Service</h3>
            <ul class="space-y-4 text-gray-600">
              <li class="flex items-center justify-between border-b border-gray-100 pb-3">
                <div>
                  <p class="font-semibold text-gray-900">Sunday Bible Study</p>
                  <p class="text-sm text-gray-500">Main sanctuary</p>
              </div>
                <span class="font-medium text-gray-700">9:00 AM</span>
              </li>
              <li class="flex items-center justify-between">
                <div>
                  <p class="font-semibold text-gray-900">Wednesday Bible Study</p>
                  <p class="text-sm text-gray-500">Mid-week gathering</p>
              </div>
                <span class="font-medium text-gray-700">7:00 PM</span>
              </li>
            </ul>
          </div>
        </div>

        <div class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8">
          <div class="space-y-2 mb-6 text-center lg:text-left">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">Reach out</p>
            <h2 class="text-3xl font-semibold text-gray-900">Send us a message</h2>
            <p class="text-gray-600">We typically respond within one business day.</p>
          </div>
            <form @submit.prevent="submitForm" class="space-y-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Full Name *</label>
                  <input
                    v-model="form.name"
                    type="text"
                    required
                    class="input-field"
                    placeholder="Your full name"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Email *</label>
                  <input
                    v-model="form.email"
                    type="email"
                    required
                    class="input-field"
                    placeholder="your.email@example.com"
                  />
                </div>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Phone</label>
                  <input
                    v-model="form.phone"
                    type="tel"
                    class="input-field"
                    placeholder="+855 15 81 44 40"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Why are you visiting? *</label>
                  <select
                    v-model="form.reason"
                    required
                    class="input-field"
                  >
                    <option value="">Select one</option>
                    <option value="Plan a visit">Plan a visit</option>
                    <option value="Prayer request">Prayer request</option>
                    <option value="Counseling / care">Counseling / care</option>
                    <option value="Serving / volunteer">Serving / volunteer</option>
                    <option value="Membership / next steps">Membership / next steps</option>
                    <option value="General question">General question</option>
                    <option value="Other">Other</option>
                  </select>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Message *</label>
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
          <div
            v-if="message"
            class="mt-6 rounded-2xl p-4 border"
            :class="messageType === 'success' ? 'bg-green-50 text-green-700 border-green-100' : 'bg-red-50 text-red-700 border-red-100'"
          >
              {{ message }}
          </div>
        </div>
      </div>
    </section>

    <!-- Map -->
    <section class="py-16 bg-white">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid gap-10 lg:grid-cols-2 items-center">
        <div class="space-y-4">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Visit us</p>
          <h2 class="text-3xl font-semibold text-gray-900">Sunday gatherings in Phnom Penh</h2>
          <p class="text-gray-600">
            We meet just south of central Phnom Penh near Boeung Tompun. Use the map for directions or open the route in Google Maps.
          </p>
          <a
            href="https://maps.app.goo.gl/a4YGCcvWRLy6BvRz9"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center text-brand-blue font-semibold hover:text-brand-orange transition-colors"
          >
            Open directions
            <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </a>
        </div>
        <div class="rounded-[32px] overflow-hidden border border-gray-100 shadow-2xl">
          <iframe
            src="https://www.google.com/maps?q=Street+26BT,+House+428,+Phnom+Penh,+Cambodia&output=embed"
            width="100%"
            height="420"
            style="border:0;"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
            class="w-full"
          ></iframe>
        </div>
      </div>
    </section>

  </div>
</template>

<script>
import { ref } from 'vue'
import { contactRequestService } from '@/services/contactRequestService'

export default {
  name: 'Contact',
  setup() {
    const form = ref({
      name: '',
      email: '',
      phone: '',
      reason: '',
      message: ''
    })
    
    const submitting = ref(false)
    const message = ref('')
    const messageType = ref('')

    const submitForm = async () => {
      submitting.value = true
      message.value = ''
      
      try {
        await contactRequestService.createContactRequest({
          name: form.value.name,
          email: form.value.email,
          phone: form.value.phone || undefined,
          reason: form.value.reason,
          message: form.value.message,
        })
        
        message.value = 'Thank you for your message! We\'ll get back to you soon.'
        messageType.value = 'success'
        
        form.value = {
          name: '',
          email: '',
          phone: '',
          reason: '',
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
