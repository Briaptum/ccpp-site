<template>
  <div class="px-4 sm:px-0">
    <!-- Statement of Faith -->
    <div class="max-w-4xl mx-auto mb-12">
      <div class="card">
        <img 
          src="/src/assets/images/logo3.png" 
          alt="Calvary Chapel Phnom Penh Logo" 
          class="h-20 md:h-20 w-auto object-contain mb-6"
        />
        <h2 class="text-3xl text-light text-gray-900 mb-6">What We Believe</h2>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> that there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> that the Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in God the Father, an infinite, personal Spirit, perfect in holiness, wisdom, power and love; that He concerns Himself mercifully in the affairs of men, that he hears and answers prayer; and that He saves from sin and death all who come to Him through Jesus Christ.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in Jesus Christ, God's only begotten Son, conceived by the Holy Spirit. We believe in His virgin birth, sinless life, miracles and teachings, His substitutionary atoning death, bodily resurrection, ascension into heaven, perpetual intercession for His people and personal, visible return to earth.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved and the everlasting conscious punishment of the lost.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in the Holy Spirit, who came forth from the Father and Son to convict the world of sin, righteousness, and judgment and to regenerate, sanctify and empower for ministry all who believe in Christ; we believe the Holy Spirit indwells every believer in Jesus Christ and that He is an abiding Helper, Teacher and Guide. We believe in the present ministry of the Holy Spirit and in the exercise of all the Bible gifts of the Spirit.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> that all men are sinners by nature and choice and, therefore, under condemnation, that God regenerates by the Holy Spirit, those who repent of their sins and confess Jesus Christ as Lord; that Jesus Christ baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in the universal Church, that living spiritual body, of which Christ is the Head and all regenerated persons are members.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> that the Lord Jesus Christ committed two ordinances to the Church; 1) baptism and 2) the Lord's Supper. We believe in baptism by immersion and communion open to all believers.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> in the personal, visible return of Christ to earth and the establishment of His Kingdom; in the resurrection of the body, the final judgment and eternal blessing of the righteous and endless suffering of the wicked.
          </p>
          <p class="text-gray-700 mb-6">
            <strong class="text-[#39B4FB]">WE BELIEVE</strong> that the Bible is the inspired Word of God, complete and entire, by which authority is given through the Holy Spirit to test all things to see whether they are of God or not. 
            (I Corinthians 2:13)
          </p>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl text-light mb-4">Join Our Family</h2>
        <p class="text-xl mb-8">
          We'd love to welcome you into our community of faith and fellowship.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/about" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Learn More
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'WhatWeBelieve'
}
</script>

