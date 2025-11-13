<template>
  <FullPageShell>
    <div class="w-full">
      <div class="text-center">
        <h1 class="text-4xl font-semibold tracking-tight text-white sm:text-5xl">Resources</h1>
        <p class="mt-6 text-base text-slate-100/80">Tools and materials for your spiritual journey.</p>
      </div>

      <div class="mt-12 grid gap-12 lg:grid-cols-[1.1fr,0.9fr] lg:items-start">
        <div class="rounded-3xl border border-white/10 bg-slate-900/70 p-10 shadow-xl backdrop-blur">
          <h2 class="text-2xl font-semibold text-white">Welcome to Our Resource Center</h2>
          <div class="mt-6 space-y-5 text-base text-slate-100/85">
            <p>
              We’ve compiled helpful resources to support you on your spiritual journey. Whether you’re looking for Bible study materials, sermon recordings, prayer guides, or information about our programs, you’ll find it here.
            </p>
            <p>
              Our resources are designed to help you grow in your relationship with God, understand Scripture better, and connect with our church community. Explore the sections below to find what you're looking for.
            </p>
          </div>
        </div>
        <div class="space-y-6">
          <div class="rounded-3xl border border-primary-300/40 bg-primary-500/10 p-10 shadow-xl backdrop-blur">
            <h3 class="text-lg font-semibold uppercase tracking-[0.35em] text-primary-200">Quick Links</h3>
            <div class="mt-6 grid gap-4">
              <router-link to="/contact" class="flex items-center gap-4 rounded-2xl border border-white/10 bg-slate-900/60 p-4 text-left transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-3xl">📧</div>
                <div>
                  <h4 class="font-semibold text-white">Request Resources</h4>
                  <p class="text-sm text-slate-100/70">Contact us for specific materials or information.</p>
                </div>
              </router-link>
              <router-link to="/events" class="flex items-center gap-4 rounded-2xl border border-white/10 bg-slate-900/60 p-4 text-left transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-3xl">📅</div>
                <div>
                  <h4 class="font-semibold text-white">Upcoming Events</h4>
                  <p class="text-sm text-slate-100/70">See our calendar of classes and workshops.</p>
                </div>
              </router-link>
              <router-link to="/ministries" class="flex items-center gap-4 rounded-2xl border border-white/10 bg-slate-900/60 p-4 text-left transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-3xl">👥</div>
                <div>
                  <h4 class="font-semibold text-white">Join a Ministry</h4>
                  <p class="text-sm text-slate-100/70">Get connected through ministry communities.</p>
                </div>
              </router-link>
              <router-link to="/about" class="flex items-center gap-4 rounded-2xl border border-white/10 bg-slate-900/60 p-4 text-left transition hover:-translate-y-1 hover:border-primary-200/60">
                <div class="text-3xl">ℹ️</div>
                <div>
                  <h4 class="font-semibold text-white">Learn More</h4>
                  <p class="text-sm text-slate-100/70">Discover more about our church and beliefs.</p>
                </div>
              </router-link>
            </div>
          </div>
          <div class="rounded-3xl border border-white/10 bg-primary-500/20 p-10 text-center shadow-xl backdrop-blur">
            <h3 class="text-2xl font-semibold text-white">Need Help Finding Resources?</h3>
            <p class="mt-4 text-sm text-slate-100/80">Our team is here to help you find exactly what you need for your spiritual growth.</p>
            <router-link
              to="/contact"
              class="mt-6 inline-flex items-center justify-center rounded-full border border-white/25 px-6 py-3 text-sm font-semibold text-white transition hover:-translate-y-0.5 hover:bg-white hover:text-primary-600"
            >
              Contact Us
            </router-link>
          </div>
        </div>
      </div>

      <div class="mt-16">
        <h2 class="text-center text-2xl font-semibold text-white">Available Resources</h2>
        <div class="mt-10 grid gap-8 md:grid-cols-2 lg:grid-cols-3">
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">📖</div>
            <h3 class="mt-4 text-lg font-semibold text-white">Bible Study Materials</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Study guides, commentaries, and resources to help you dive deeper into God’s Word.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Weekly study guides</li>
              <li>• Bible reading plans</li>
              <li>• Devotional materials</li>
            </ul>
          </article>
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">🎵</div>
            <h3 class="mt-4 text-lg font-semibold text-white">Sermon Recordings</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Access recordings of our weekly messages and teaching series.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Audio recordings</li>
              <li>• Video messages</li>
              <li>• Message transcripts</li>
            </ul>
          </article>
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">🙏</div>
            <h3 class="mt-4 text-lg font-semibold text-white">Prayer Resources</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Prayer guides, intercession lists, and materials to strengthen your prayer life.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Prayer guides</li>
              <li>• Intercession lists</li>
              <li>• Prayer request forms</li>
            </ul>
          </article>
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">📚</div>
            <h3 class="mt-4 text-lg font-semibold text-white">Library Resources</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Books, reference materials, and educational resources available in our church library.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Christian books</li>
              <li>• Reference materials</li>
              <li>• Educational resources</li>
            </ul>
          </article>
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">👨‍👩‍👧‍👦</div>
            <h3 class="mt-4 text-lg font-semibold text-white">Family Resources</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Materials to help families grow together in faith and navigate life’s challenges.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Family devotionals</li>
              <li>• Parenting resources</li>
              <li>• Children's materials</li>
            </ul>
          </article>
          <article class="rounded-3xl border border-white/10 bg-slate-900/70 p-8 shadow-xl backdrop-blur transition hover:-translate-y-1 hover:border-primary-200/60">
            <div class="text-4xl drop-shadow-sm">💡</div>
            <h3 class="mt-4 text-lg font-semibold text-white">New Believer Resources</h3>
            <p class="mt-3 text-sm text-slate-100/80">
              Getting started guides and foundational materials for those new to the faith.
            </p>
            <ul class="mt-4 space-y-2 text-sm text-slate-100/70">
              <li>• Getting started guide</li>
              <li>• Foundation classes</li>
              <li>• Discipleship materials</li>
            </ul>
          </article>
        </div>
      </div>
    </div>
  </FullPageShell>
</template>

<script>
import FullPageShell from '@/components/FullPageShell.vue'

export default {
  name: 'Resources',
  components: {
    FullPageShell,
  },
}
</script>

