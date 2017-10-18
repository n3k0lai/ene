var path = require('path');
var webpack = require('webpack');

module.exports = {
  target:'node',
  entry: [
    'babel-polyfill', 
    './src/twitter.js'
  ],
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'main.bundle.js'
  },
  devtool: 'source-map',
  module: {
    loaders: [
      {
        test: /\.js$/,
        include: path.join(__dirname, "src"),
        loader: 'babel-loader',
        query: {
          presets: ['es2015-node4']
        }
      }
    ]
  },
  stats: {
    colors: true
  }
};
