var path = require('path')

module.exports = {
  target: 'node',
  entry: [
    'babel-polyfill',
    './src/app.js'
  ],
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'main.bundle.js'
  },
  devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.js$/,
        include: path.join(__dirname, 'src'),
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          query: {
            presets: ['es2015-node4']
          }
        }
      }
    ]
  },
  stats: {
    colors: true
  }
}
