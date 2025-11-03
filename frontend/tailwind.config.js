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
          50: '#fbf7ed',
          100: '#f6efdb',
          200: '#ecdfb7',
          300: '#e1b65d',
          400: '#dba94b',
          500: '#E1B65D',
          600: '#c9a34d',
          700: '#a8873f',
          800: '#866b31',
          900: '#644f23',
        },
        secondary: {
          50: '#f7f2ef',
          100: '#efe5df',
          200: '#dfcbbe',
          300: '#cfb19d',
          400: '#bf977c',
          500: '#BC6D4F',
          600: '#a86246',
          700: '#8a503a',
          800: '#6c3e2e',
          900: '#4e2c21',
        },
        main: {
          50: '#e8edef',
          100: '#d1dbdf',
          200: '#a3b7bf',
          300: '#75939f',
          400: '#476f7f',
          500: '#265267',
          600: '#1f4858',
          700: '#193e49',
          800: '#13343a',
          900: '#0d2a2b',
        }
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
