const NodePolyfillPlugin = require('node-polyfill-webpack-plugin')

module.exports = {
  transpileDependencies: ['vuex-persist'],
  configureWebpack: {
    plugins: [
      new NodePolyfillPlugin()
    ]
  }
}
