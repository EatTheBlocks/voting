import {createI18n} from 'vue-i18n'

const datetimeFormats = {
  'en-US': {
    short: {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: 'numeric',
      minute: 'numeric'
    }
  }
}

const i18n = createI18n({
  datetimeFormats,
})

export default i18n
