import * as firebase from 'firebase/app'
import 'firebase/auth'

const config = {
  apiKey: process.env.firebaseApiKey,
  authDomain: `${process.env.firebaseProjectId}.firebaseapp.com`,
  databaseURL: `https://${process.env.firebaseProjectId}.firebaseio.com`,
  projectId: process.env.firebaseProjectId,
  storageBucket: `${process.env.firebaseProjectId}.appspot.com`,
  messagingSenderId: process.env.firebaseMessagingSenderId,
}

if (!firebase.apps.length) {
  firebase.initializeApp(config)
}

export default ({ _ }, inject) => {
  inject('firebase', firebase)
}
