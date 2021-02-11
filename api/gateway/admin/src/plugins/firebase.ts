import * as admin from 'firebase-admin'

const firebaseProjectId: string = process.env.FIREBASE_PROJECT_ID || ''
const firebaseClientEmail: string = process.env.FIREBASE_CLIENT_EMAIL || ''
const firebasePrivateKey: string = process.env.FIREBASE_PRIVATE_KEY || ''

const cert: admin.ServiceAccount = {
  projectId: firebaseProjectId,
  clientEmail: firebaseClientEmail,
  privateKey: firebasePrivateKey.replace(/\\n/g, '\n'),
}

const app = admin.initializeApp({
  credential: admin.credential.cert(cert),
})

const auth = app.auth()

export { auth }
