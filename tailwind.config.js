/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './cmd/ui/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}

