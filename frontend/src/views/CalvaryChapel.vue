<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Calvary Chapel</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-background py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          An international association of charismatic evangelical churches committed to faithful Bible teaching and Spirit-led ministry.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            Calvary Chapel is an international association of charismatic evangelical churches, with origins in Pentecostalism. It maintains a number of radio stations around the world and operates many local Calvary Chapel Bible College programs.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Beginning in 1965 in Southern California, this fellowship of churches grew out of Chuck Smith's Calvary Chapel Costa Mesa. It became a hub of the Jesus Movement in the late 1960s through connections with Lonnie Frisbee and John Higgins, attracting thousands of young converts and fostering contemporary Christian music through Maranatha! Music. Known for its verse-by-verse Bible teaching, casual style of worship, and emphasis on expository preaching, the movement expanded into a worldwide fellowship of independent churches.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Calvary Chapel identifies as neither a denomination nor strictly Pentecostal. It holds to evangelical doctrine with charismatic practices like tongues and prophecy while maintaining a strong pretribulationist, premillennialist eschatology.
          </p>
          <p class="text-lg text-gray-700 mb-6">
            Calvary Chapel remains influential through its Bible college, radio stations, and Harvest Crusades. Many well-known pastors and musicians, such as Greg Laurie, Skip Heitzig, Switchfoot, and P.O.D., have roots in Calvary Chapel.
          </p>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
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
  name: 'CalvaryChapel',
};
</script>

