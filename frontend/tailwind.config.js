/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'brand-orange': '#E49D40',
        'brand-blue': '#227082',
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
