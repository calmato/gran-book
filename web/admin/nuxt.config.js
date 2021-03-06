import colors from 'vuetify/es5/util/colors'

export default {
  ssr: false,
  srcDir: 'app',
  head: {
    titleTemplate: '%s - gran-book',
    title: 'gran-book',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Roboto' },
    ],
  },

  css: [],

  components: true,
  plugins: [
    '~/plugins/axios',
    '~/plugins/axios-accessor',
    '~/plugins/firebase',
    '~/plugins/persisted-state',
    '~/plugins/vee-validate',
  ],
  buildModules: ['@nuxt/typescript-build', '@nuxtjs/composition-api/module', '@nuxtjs/vuetify'],
  modules: ['@nuxtjs/axios', '@nuxt/content'],

  router: {
    middleware: ['authenticated'],
  },

  content: {},

  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: false,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },

  typescript: {
    typeCheck: {
      eslint: {
        files: './app/**/*.{ts,js,vue}',
      },
    },
  },

  env: {
    firebaseApiKey: process.env.FIREBASE_API_KEY,
    firebaseProjectId: process.env.FIREBASE_PROJECT_ID,
    firebaseMessagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID,
    rakutenAppId: process.env.RAKUTEN_APPLICATION_ID,
    apiURL: process.env.API_URL,
  },

  build: {
    babel: {
      plugins: [['@babel/plugin-proposal-private-methods', { loose: true }]],
    },
  },
}
