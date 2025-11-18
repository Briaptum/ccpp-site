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
          50: '#1a3a38',
          100: '#1a3533',
          200: '#18302e',
          300: '#162b29',
          400: '#142624',
          500: '#122422',
          600: '#0f1f1d',
          700: '#0d1a18',
          800: '#0b1513',
          900: '#09100e',
        },
        secondary: {
          50: '#fef5e8',
          100: '#fdebd1',
          200: '#fbd7a3',
          300: '#f9c375',
          400: '#f7af47',
          500: '#E49D40',
          600: '#d68d2a',
          700: '#b87523',
          800: '#9a5d1c',
          900: '#7c4515',
        },
        main: {
          50: '#f0f7f5',
          100: '#e1efeb',
          200: '#c3dfd7',
          300: '#a5cfc3',
          400: '#87bfaf',
          500: '#9AC2B4',
          600: '#7ba89a',
          700: '#5c8e80',
          800: '#3d7466',
          900: '#1e5a4c',
        },
        background: {
          DEFAULT: '#F9EEE4',
          50: '#fefcfb',
          100: '#fdf9f7',
          200: '#fbf3ef',
          300: '#f9ede7',
          400: '#f7e7df',
          500: '#F9EEE4',
          600: '#f5e6dc',
          700: '#f1ded4',
          800: '#edd6cc',
          900: '#e9cec4',
        }
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
