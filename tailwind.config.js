const colors = require('tailwindcss/colors')

module.exports = {
  mode: 'jit',
  purge: [
    './public/**/*.html',
    './src/**/*.{js,jsx,ts,tsx,vue}',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      screens: {
        'xs': '370px',
      },
    },
    colors: {
      transparent: 'transparent',
      main: {
        primary: '#0657f9',
        bg: '#fff',
        text: '#586069',
        link: '#111',
        heading: '#111',
        border: '#d1d5da',
        block: '#fff',
        blockText: '#111',
        progress: '#0657f9',
      },
      dark: {
        primary: '#393939',
        bg: '#393939',
        text: '#fff',
        link: '#fff',
        heading: '#3DE4E4',
        border: '#656565',
        block: '#4b5050',
        blockText: '#fff',
        progress: '#3DE4E4',
      },
      ...colors,
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/custom-forms'),
  ],
}
