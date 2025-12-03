<template>
  <div>
    <!-- Hero Section -->
    <section class="relative min-h-[70vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-20 bg-gradient-to-br from-[#081229] via-[#0f233f] to-[#173452] text-white">
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-30 mix-blend-screen"
          style="background-image: radial-gradient(circle at 20% 20%, rgba(255,255,255,0.25), transparent 40%), radial-gradient(circle at 80% 15%, rgba(255,255,255,0.15), transparent 35%), radial-gradient(circle at 60% 80%, rgba(255,255,255,0.18), transparent 30%);"
        ></div>
        <div class="absolute inset-y-[-30%] left-[-15%] w-2/3 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[180px] opacity-60"></div>
        <div class="absolute inset-y-[-20%] right-[-10%] w-1/2 bg-gradient-to-l from-brand-blue/35 via-transparent to-transparent blur-[150px] opacity-70"></div>
        <div class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/40 via-transparent to-transparent"></div>
      </div>
      <div class="relative z-10 w-full max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">Calvary Chapel</h1>
        <p class="text-lg md:text-xl text-white/80 max-w-3xl mx-auto font-light">
          A global fellowship of churches rooted in expositional teaching, the work of the Holy Spirit, and a grace-filled culture that welcomes every generation.
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="point in heroPoints"
            :key="point"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm"
          >
            {{ point }}
          </span>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section class="py-20 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-6 text-center">
        <p class="text-sm uppercase tracking-[0.35em] text-gray-400">Question</p>
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900">What is Calvary Chapel?</h2>
        <div class="space-y-4 text-left text-lg text-gray-600 leading-relaxed">
          <p>
            Calvary Chapel is an international association of charismatic evangelical churches with roots in the 1960s Jesus Movement. Beginning in Costa Mesa, California, the fellowship embraced verse-by-verse Bible teaching, simple worship, and a welcoming culture that drew thousands of new believers.
          </p>
          <p>
            The churches identify as neither a denomination nor strictly Pentecostal. They hold to historic evangelical doctrine, practice the gifts of the Holy Spirit in an orderly way, and maintain a strong expectation of Christ’s return.
          </p>
          <p>
            Today the Calvary Chapel family includes radio ministries, Bible colleges, and church plants around the world—all focused on teaching Scripture, depending on the Spirit, and serving communities with the hope of Jesus.
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'CalvaryChapel',
  setup() {
    const heroPoints = ['Verse-by-verse', 'Spirit-led', 'Global family', 'Grace & truth']

    return {
      heroPoints
    }
  }
}
</script>

