module.exports = function(api) {
  api.cache(true);
  return {
    presets: ['module:metro-react-native-babel-preset', 'babel-preset-expo'],
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
            '@@': ['./'],
            'tests': ['./test/']
          },
        },
        'react-native-reanimated/plugin'
      ],
    ],
  };
};
