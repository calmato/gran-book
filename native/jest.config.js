module.exports = {
  preset: 'jest-expo',
  verbose: true,
  collectCoverage: true,
  collectCoverageFrom: [
    '**/app/**/*.{js,jsx,ts,tsx}',
    '!**/node_modules/**',
    '!**/tmp/**',
  ],
  coveragePathIgnorePatterns: [
    '<rootDir>/node_modules/',
    '<rootDir>/tmp/',
    '<rootDir>/app/types/api/',
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
