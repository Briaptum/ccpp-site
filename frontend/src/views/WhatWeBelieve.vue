<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">What We Believe</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          These convictions anchor every ministry expression of Calvary Chapel Phnom Penh.
        </p>

        <div class="grid gap-6 lg:grid-cols-2 mb-12">
          <article
            v-for="(statement, index) in statements"
            :key="index"
            class="bg-background-50 rounded-lg p-6 shadow-md"
          >
            <h2 class="text-lg font-semibold text-primary-600 mb-3">We Believe</h2>
            <p class="text-base text-gray-700">{{ statement.body }}</p>
            <p v-if="statement.reference" class="mt-4 text-sm font-semibold text-primary-500">
              {{ statement.reference }}
            </p>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-primary-600 text-white font-semibold rounded-lg hover:bg-secondary-700 transition-colors shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'WhatWeBelieve',
  data() {
    return {
      statements: [
        {
          body:
            'That there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.',
        },
        {
          body:
            'That the Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.',
        },
        {
          body:
            'In God the Father, an infinite, personal Spirit, perfect in holiness, wisdom, power and love; that He concerns Himself mercifully in the affairs of men, hears and answers prayer, and saves from sin and death all who come to Him through Jesus Christ.',
        },
        {
          body:
            "In Jesus Christ, God's only begotten Son, conceived by the Holy Spirit. We affirm His virgin birth, sinless life, miracles and teachings, His substitutionary atoning death, bodily resurrection, ascension into heaven, perpetual intercession for His people, and personal, visible return to earth.",
        },
        {
          body:
            'In the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.',
        },
        {
          body:
            'In the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved and the everlasting conscious punishment of the lost.',
        },
        {
          body:
            'In the Holy Spirit, who came forth from the Father and Son to convict the world of sin, righteousness, and judgment and to regenerate, sanctify and empower for ministry all who believe in Christ. We believe the Holy Spirit indwells every believer in Jesus Christ and that He is an abiding Helper, Teacher and Guide, active today in all His gifts.',
        },
        {
          body:
            'That all men are sinners by nature and choice and therefore under condemnation, that God regenerates by the Holy Spirit those who repent of their sins and confess Jesus Christ as Lord; that Jesus Christ baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.',
        },
        {
          body:
            'In the universal Church, the living spiritual body of which Christ is the Head and all regenerated persons are members.',
        },
        {
          body:
            "That the Lord Jesus Christ committed two ordinances to the Church: baptism and the Lord's Supper. We practice baptism by immersion and celebrate communion open to all believers.",
        },
        {
          body:
            'In the personal, visible return of Christ to earth and the establishment of His Kingdom; in the resurrection of the body, the final judgment and eternal blessing of the righteous and endless separation of the wicked.',
        },
        {
          body:
            'That the Bible is the inspired Word of God, complete and entire, through which the Holy Spirit gives authority to test all things and discern what is of God.',
          reference: '1 Corinthians 2:13',
        },
      ],
    };
  },
};
</script>


