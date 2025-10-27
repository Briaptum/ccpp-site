/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#e8f7fe',
          100: '#d1effd',
          200: '#a3dffc',
          300: '#75cffa',
          400: '#57c4fc',
          500: '#39B4FB',
          600: '#1ca5fa',
          700: '#0891e8',
          800: '#0673b8',
          900: '#045588',
        }
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
