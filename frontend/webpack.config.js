const precss = require('precss');
const autoprefixer = require('autoprefixer');
module.exports = {
  entry: './src/bundle.js',
  output: {
    filename: './public/js/bundle.js',
  },
  module: {
    rules: [{
      test: /\.scss$/,
      use: [
        { loader: 'style-loader' },
        { loader: 'css-loader' },
        {
          loader: 'postcss-loader',
          options: {
            plugins() {
              return [
                precss,
                autoprefixer,
              ];
            },
          },
        },
        { loader: 'sass-loader' },
      ],
    }],
  },
};