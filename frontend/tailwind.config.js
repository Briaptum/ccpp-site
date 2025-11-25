/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        main: '#122422',
        secondary: '#111827',
        primary: '##ffffff',
        'custom-orange': '#E49D40',
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
