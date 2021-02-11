import * as admin from 'firebase-admin'

const cert: admin.ServiceAccount = {
  projectId: process.env.FIREBASE_PROJECT_ID!,
  clientEmail: process.env.FIREBASE_CLIENT_EMAIL!,
  privateKey: process.env.FIREBASE_PRIVATE_KEY!.replace(/\\n/g, '\n'),
}

const app = admin.initializeApp({
  credential: admin.credential.cert(cert),
})

const auth = app.auth()

export { auth }
