const nodeExternals = require('webpack-node-externals');
const path = require('path');

module.exports = {
  target: 'node',
  entry: path.resolve(__dirname, '..', 'src', 'index.ts'),
  externals: [
    nodeExternals(),
  ],
  output: {
    path: path.resolve(__dirname, '..', 'dist'),
    filename: 'server.js',
  },
  resolve: {
    extensions: ['.ts', '.js', '.json'],
    alias: {
      '~': path.resolve(__dirname, '..', 'src'),
      '@': path.resolve(__dirname, '..', 'src'),
      '~~': path.resolve(__dirname, '..'),
      '@@': path.resolve(__dirname, '..'),
    }
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        exclude: /node_modules/,
        loader: 'ts-loader',
        options: {
          configFile: path.resolve(__dirname, '..', 'tsconfig.json'),
        },
      },
    ],
  },
};
