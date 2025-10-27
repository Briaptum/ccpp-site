<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-primary-600 to-primary-800 text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">How to Help</h1>
        <p class="text-xl md:text-2xl">Join us in serving God's Kingdom</p>
      </div>
    </div>

    <!-- Prayer Section -->
    <div id="prayer" class="max-w-4xl mx-auto mb-12 scroll-mt-20">
      <div class="card">
        <div class="flex items-center mb-6">
          <div class="text-4xl mr-4">🙏</div>
          <h2 class="text-3xl font-bold text-gray-900">Prayer</h2>
        </div>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-4">
            Prayer is the foundation of everything we do. We believe in the power of prayer to transform lives, 
            heal communities, and bring people closer to God.
          </p>
          <h3 class="text-xl font-semibold text-gray-900 mb-3">How You Can Help Through Prayer:</h3>
          <ul class="list-disc list-inside text-gray-700 space-y-2 mb-4">
            <li>Join our weekly prayer meetings every Wednesday at 7:00 PM</li>
            <li>Sign up for our prayer chain to receive urgent prayer requests</li>
            <li>Commit to praying for our missionaries and outreach programs</li>
            <li>Participate in our monthly prayer and fasting days</li>
            <li>Pray for our church leadership and ministry teams</li>
          </ul>
          <p class="text-gray-700 mb-4">
            <strong class="text-primary-600">Submit a Prayer Request:</strong> If you have a prayer need, 
            please reach out to us through our contact form or speak with one of our prayer team members.
          </p>
          <div class="mt-6">
            <router-link to="/contact" class="btn-primary">
              Submit Prayer Request
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Donation Section -->
    <div id="donation" class="bg-gray-100 py-12 mb-12 scroll-mt-20">
      <div class="max-w-4xl mx-auto">
        <div class="card">
          <div class="flex items-center mb-6">
            <div class="text-4xl mr-4">💝</div>
            <h2 class="text-3xl font-bold text-gray-900">Donation</h2>
          </div>
          <div class="prose prose-lg max-w-none">
            <p class="text-gray-700 mb-4">
              Your generous donations help us continue our mission of spreading the Gospel, serving our community, 
              and supporting those in need. Every contribution, no matter the size, makes a significant difference.
            </p>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">Ways to Give:</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
              <div class="bg-white p-4 rounded-lg border border-gray-200">
                <h4 class="font-semibold text-gray-900 mb-2">General Fund</h4>
                <p class="text-gray-600 text-sm">Supports daily operations, ministry programs, and facility maintenance.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border border-gray-200">
                <h4 class="font-semibold text-gray-900 mb-2">Mission Fund</h4>
                <p class="text-gray-600 text-sm">Directly supports our missionaries and evangelism efforts.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border border-gray-200">
                <h4 class="font-semibold text-gray-900 mb-2">Building Fund</h4>
                <p class="text-gray-600 text-sm">Helps us expand and improve our church facilities.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border border-gray-200">
                <h4 class="font-semibold text-gray-900 mb-2">Benevolence Fund</h4>
                <p class="text-gray-600 text-sm">Provides assistance to those facing financial hardship.</p>
              </div>
            </div>
            <p class="text-gray-700 mb-4 italic">
              "Each of you should give what you have decided in your heart to give, not reluctantly or under compulsion, 
              for God loves a cheerful giver." - 2 Corinthians 9:7
            </p>
            <div class="mt-6">
              <router-link to="/donate" class="btn-primary">
                Make a Donation
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Needs List Section -->
    <div id="needs-list" class="max-w-4xl mx-auto mb-12 scroll-mt-20">
      <div class="card">
        <div class="flex items-center mb-6">
          <div class="text-4xl mr-4">📋</div>
          <h2 class="text-3xl font-bold text-gray-900">Needs List</h2>
        </div>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-4">
            We are grateful for the support of our community. Below are specific items and resources that would 
            greatly benefit our ministry and the people we serve.
          </p>
          
          <h3 class="text-xl font-semibold text-gray-900 mb-3">Current Ministry Needs:</h3>
          <div class="bg-primary-50 p-6 rounded-lg mb-6">
            <h4 class="font-semibold text-gray-900 mb-3">For Our Church:</h4>
            <ul class="list-disc list-inside text-gray-700 space-y-2">
              <li>Audio/Visual equipment for worship services</li>
              <li>Musical instruments (keyboard, guitars, drums)</li>
              <li>Chairs and tables for fellowship hall</li>
              <li>Bibles and study materials (English and Khmer)</li>
              <li>Office supplies and computers</li>
            </ul>
          </div>

          <div class="bg-primary-50 p-6 rounded-lg mb-6">
            <h4 class="font-semibold text-gray-900 mb-3">For Our Outreach Programs:</h4>
            <ul class="list-disc list-inside text-gray-700 space-y-2">
              <li>Food and hygiene items for community assistance</li>
              <li>Children's ministry supplies and toys</li>
              <li>School supplies for local students</li>
              <li>Transportation support for elderly members</li>
              <li>Volunteers for various ministry activities</li>
            </ul>
          </div>

          <div class="bg-primary-50 p-6 rounded-lg mb-6">
            <h4 class="font-semibold text-gray-900 mb-3">For Our Missionaries:</h4>
            <ul class="list-disc list-inside text-gray-700 space-y-2">
              <li>Monthly financial support</li>
              <li>Care packages and encouragement</li>
              <li>Technology for remote communication</li>
              <li>Translation services and materials</li>
            </ul>
          </div>

          <p class="text-gray-700 mb-4">
            If you would like to donate any of these items or have questions about other ways to help, 
            please contact our office.
          </p>
          <div class="mt-6">
            <router-link to="/contact" class="btn-primary">
              Contact Us About Donations
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Testimonial Section -->
    <div id="testimonial" class="bg-gray-100 py-12 mb-12 scroll-mt-20">
      <div class="max-w-4xl mx-auto">
        <div class="card">
          <div class="flex items-center mb-6">
            <div class="text-4xl mr-4">💬</div>
            <h2 class="text-3xl font-bold text-gray-900">Testimonial</h2>
          </div>
          <div class="prose prose-lg max-w-none">
            <p class="text-gray-700 mb-6">
              Your story matters! Sharing how God has worked in your life encourages others and glorifies Him. 
              We would love to hear about your journey of faith.
            </p>

            <h3 class="text-xl font-semibold text-gray-900 mb-4">Why Share Your Testimony?</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
              <div class="bg-white p-4 rounded-lg border-l-4 border-primary-500">
                <h4 class="font-semibold text-gray-900 mb-2">Encourage Others</h4>
                <p class="text-gray-600 text-sm">Your testimony can inspire and strengthen the faith of others who may be going through similar challenges.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border-l-4 border-primary-500">
                <h4 class="font-semibold text-gray-900 mb-2">Glorify God</h4>
                <p class="text-gray-600 text-sm">Sharing what God has done in your life gives Him glory and points others to His goodness.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border-l-4 border-primary-500">
                <h4 class="font-semibold text-gray-900 mb-2">Build Community</h4>
                <p class="text-gray-600 text-sm">Testimonies create connections and foster a sense of community among believers.</p>
              </div>
              <div class="bg-white p-4 rounded-lg border-l-4 border-primary-500">
                <h4 class="font-semibold text-gray-900 mb-2">Evangelize</h4>
                <p class="text-gray-600 text-sm">Your personal story is one of the most powerful tools for reaching the lost.</p>
              </div>
            </div>

            <h3 class="text-xl font-semibold text-gray-900 mb-3">Ways to Share Your Testimony:</h3>
            <ul class="list-disc list-inside text-gray-700 space-y-2 mb-6">
              <li>Submit a written testimony for our newsletter or website</li>
              <li>Share during Sunday service (speak with a pastor first)</li>
              <li>Record a video testimony for our social media</li>
              <li>Participate in small group discussions</li>
              <li>Mentor someone who is new to the faith</li>
            </ul>

            <p class="text-gray-700 mb-4 italic">
              "But in your hearts revere Christ as Lord. Always be prepared to give an answer to everyone 
              who asks you to give the reason for the hope that you have." - 1 Peter 3:15
            </p>

            <div class="mt-6">
              <router-link to="/contact" class="btn-primary">
                Share Your Testimony
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Service Section -->
    <div id="service" class="max-w-4xl mx-auto mb-12 scroll-mt-20">
      <div class="card">
        <div class="flex items-center mb-6">
          <div class="text-4xl mr-4">🤝</div>
          <h2 class="text-3xl font-bold text-gray-900">Service</h2>
        </div>
        <div class="prose prose-lg max-w-none">
          <p class="text-gray-700 mb-4">
            God has given each of us unique gifts and talents to serve His church and community. 
            We invite you to use your abilities to make a difference in the lives of others.
          </p>

          <h3 class="text-xl font-semibold text-gray-900 mb-4">Ministry Opportunities:</h3>
          
          <div class="space-y-6 mb-6">
            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🎵</span>
                Worship Team
              </h4>
              <p class="text-gray-700 mb-2">
                Join our worship team as a vocalist, musician, or audio/visual technician.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: Sunday mornings + weekly practice</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">👶</span>
                Children's Ministry
              </h4>
              <p class="text-gray-700 mb-2">
                Help teach, care for, and mentor children during Sunday School and special events.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: 1-2 Sundays per month</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🎓</span>
                Youth Ministry
              </h4>
              <p class="text-gray-700 mb-2">
                Mentor teenagers, lead small groups, or help organize youth activities and events.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: Weekly youth meetings + monthly events</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🤗</span>
                Hospitality Team
              </h4>
              <p class="text-gray-700 mb-2">
                Welcome visitors, serve refreshments, and help create a warm, inviting atmosphere.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: 1-2 Sundays per month</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🌍</span>
                Outreach & Evangelism
              </h4>
              <p class="text-gray-700 mb-2">
                Participate in community outreach, street evangelism, and missionary support activities.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: Monthly outreach events</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🔧</span>
                Facilities & Maintenance
              </h4>
              <p class="text-gray-700 mb-2">
                Help maintain and improve our church facilities through repairs, cleaning, and landscaping.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: As needed / project-based</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">💻</span>
                Communications & Media
              </h4>
              <p class="text-gray-700 mb-2">
                Assist with website updates, social media, graphic design, photography, or videography.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: Flexible / remote options available</p>
            </div>

            <div class="bg-primary-50 p-6 rounded-lg">
              <h4 class="font-semibold text-gray-900 mb-2 flex items-center">
                <span class="text-2xl mr-2">🙋</span>
                Administrative Support
              </h4>
              <p class="text-gray-700 mb-2">
                Help with office tasks, data entry, event coordination, or other administrative needs.
              </p>
              <p class="text-sm text-gray-600">Time Commitment: Flexible</p>
            </div>
          </div>

          <p class="text-gray-700 mb-4 italic">
            "Each of you should use whatever gift you have received to serve others, as faithful stewards 
            of God's grace in its various forms." - 1 Peter 4:10
          </p>

          <div class="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-6">
            <p class="text-yellow-800">
              <strong>Note:</strong> All volunteers working with children and youth must complete a background check 
              and attend our ministry training session.
            </p>
          </div>

          <div class="mt-6">
            <router-link to="/contact" class="btn-primary">
              Sign Up to Serve
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Call to Action -->
    <div class="bg-primary-600 text-white py-12 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Ready to Get Involved?</h2>
        <p class="text-xl mb-8">
          Whether through prayer, giving, or service, your involvement makes a difference. 
          Let's work together to serve God and bless our community.
        </p>
        <div class="space-x-4">
          <router-link to="/contact" class="btn-primary bg-white text-primary-600 hover:bg-gray-100">
            Contact Us
          </router-link>
          <router-link to="/donate" class="btn-primary border-2 border-white text-white hover:bg-white hover:text-primary-600">
            Donate Now
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HowToHelp',
  mounted() {
    // Smooth scroll to hash if present in URL
    if (this.$route.hash) {
      this.$nextTick(() => {
        const element = document.querySelector(this.$route.hash);
        if (element) {
          element.scrollIntoView({ behavior: 'smooth' });
        }
      });
    }
  },
  watch: {
    '$route.hash'(newHash) {
      if (newHash) {
        this.$nextTick(() => {
          const element = document.querySelector(newHash);
          if (element) {
            element.scrollIntoView({ behavior: 'smooth' });
          }
        });
      }
    }
  }
}
</script>

