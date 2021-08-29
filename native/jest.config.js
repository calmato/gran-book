module.exports = {
  preset: 'jest-expo',
  verbose: true,
  collectCoverage: true,
  collectCoverageFrom: [
    '<rootDir>/app/**/*.{js,jsx,ts,tsx}',
  ],
  coveragePathIgnorePatterns: [
    '<rootDir>/node_modules/',
    '<rootDir>/tmp/',
    '<rootDir>/app/containers/',
    '<rootDir>/app/types/',
  ],
  moduleFileExtensions: [
    'js',
    'jsx',
    'json',
    'ts',
    'tsx',
  ],
  moduleDirectories: [
    'node_modules',
  ],
  setupFiles: [
    'dotenv/config',
  ],
  testMatch: [
    '**/__tests__/**/*.(js|ts)?(x)',
    '**/?(*.)(spec|test).(js|ts)?(x)'
  ],
};
