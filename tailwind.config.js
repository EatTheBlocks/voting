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
        primary: 'var(--main-primary)',
        bg: 'var(--main-bg)',
        text: 'var(--main-text)',
        link: 'var(--main-link)',
        heading: 'var(--main-heading)',
        border: 'var(--main-border)',
        block: 'var(--main-block)',
        blockText: 'var(--main-blockText)',
        progress: 'var(--main-progress)',
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
