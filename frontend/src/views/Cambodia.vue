<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[70vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-20 bg-gradient-to-br from-[#0b1c34] via-[#142b4a] to-[#1d3c5f] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div class="absolute inset-0 opacity-30 mix-blend-screen" style="background-image: radial-gradient(circle at 25% 20%, rgba(255,255,255,0.18), transparent 45%), radial-gradient(circle at 75% 15%, rgba(255,255,255,0.12), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"></div>
        <div class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"></div>
        <div class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[180px] opacity-65"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/50 via-black/30 to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Cambodia</h1>
        <p class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light">
          A vibrant nation of resilience, heritage, and Gospel opportunity in the heart of Southeast Asia.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Phnom Penh
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Khmer Heritage
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Tonlé Sap
          </span>
          <span class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm">
            Gospel Future
          </span>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section class="py-20 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-6 text-center">
        <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900">What is Cambodia?</h2>
        <div class="space-y-4 text-left text-lg text-gray-600 leading-relaxed">
          <p>
            <span class="font-semibold text-brand-orange">Cambodia</span>, officially the <span class="font-semibold text-brand-orange">Kingdom of Cambodia</span>, sits at the crossroads of <span class="font-semibold text-brand-orange">Southeast Asia</span> bordered by Thailand, Laos, Vietnam, and the Gulf of Thailand. The <span class="font-semibold text-brand-orange">Mekong River</span>, <span class="font-semibold text-brand-orange">Tonlé Sap Lake</span>, and fertile plains shape daily life for the 17+ million people who call the country home.
          </p>
          <p>
            <span class="font-semibold text-brand-orange">Phnom Penh</span>, the capital, pulses at the confluence of the Mekong and Tonlé Sap rivers. From here, <span class="font-semibold text-brand-orange">Calvary Chapel Phnom Penh</span> shares the hope of Jesus within a city marked by markets, cafes, universities, and local neighborhoods that represent every layer of Cambodian society.
          </p>
          <p>
            The legacy of the <span class="font-semibold text-brand-orange">Khmer Empire</span>—highlighted by <span class="font-semibold text-brand-orange">Angkor Wat</span>—still inspires national pride. Yet recent history includes deep trauma from the <span class="font-semibold text-brand-orange">Khmer Rouge</span> regime. Today’s Cambodia is rebuilding through education, entrepreneurship, and a new generation eager for truth, restoration, and purpose.
          </p>
          <p>
            <span class="font-semibold text-brand-orange">Ministry in Cambodia</span> involves listening well, honoring the culture, investing in relationships, and presenting the <span class="font-semibold text-brand-orange">Gospel</span> faithfully in Khmer. From house gatherings to city outreaches, we seek to equip believers and serve Phnom Penh with humility and hope.
          </p>
        </div>
      </div>
    </section>
    <section class="pb-16 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <a
          href="https://en.wikipedia.org/wiki/Cambodia"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
        >
          Learn More
        </a>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'Cambodia'
}
</script>

