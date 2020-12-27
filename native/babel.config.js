module.exports = {
  presets: ['module:metro-react-native-babel-preset'],
  plugins: [
    [
      'inline-dotenv',
      {
        path: '.env',
      },
    ],
    [
      'module-resolver',
      {
        root: ['./'],
        extensions: ['.ios.js', '.android.js', '.js', '.ts', '.tsx', '.json'],
        alias: {
          '~': ['./app/'],
          '~~': ['./'],
          '@': ['./app/'],
          '@@': ['./']
        },
      },
    ],
  ],
};
