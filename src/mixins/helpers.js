import {shorten} from '@/helpers/utils'
import numeral from 'numeral'
import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en'
import networks from '@snapshot-labs/snapshot.js/src/networks.json'

TimeAgo.addDefaultLocale(en)
const timeAgo = new TimeAgo('en-US')

export default {
  methods: {
    _shorten: shorten,
    _n(number, format = '(0.[00]a)') {
      if (number < 0.00001) return format.includes('%') ? '0%' : 0;
      return numeral(number).format(format);
    },
    _ipfsUrl(ipfsHash) {
      if (!ipfsHash) return null;
      return `https://${process.env.VUE_APP_IPFS_GATEWAY}/ipfs/${ipfsHash}`;
    },
    _explorer(network, str, type = 'address') {
      return `${networks[network].explorer}/${type}/${str}`;
    },
    _dateAgo(endDate) {
      const now = Date.now()
      const text = timeAgo.format(endDate)
      if (endDate > now) {
        return "end " + text
      }
      return "ended " + text
    }
  }
}
