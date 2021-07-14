import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en'

TimeAgo.addDefaultLocale(en)
const timeAgo = new TimeAgo('en-US')

export default {
  methods: {
    ethShortAddress(address) {
      if (address !== undefined && address.length > 0) {
        return (
          address.substring(0, 6) +
          "..." +
          address.substring(address.length - 4)
        );
      }
      return "unknown address";
    },
    endDateAgo(endDate) {
      const now = Date.now()
      const text = timeAgo.format(endDate)
      if (endDate > now) {
        return "end " + text
      }
      return "ended " + text
    }
  }
}
