import firebase from '~/lib/firebase';

/**
 * firebase authenticationを使って email と password でログインを実行します。
 * @param payload
 * @returns { firebase.auth.UserCredential.user | null }
 */
export async function signInWithEmailAndPassword(payload: { email: string; password: string }) {
  const { email, password } = payload;

  try {
    const { user } = await firebase.auth().signInWithEmailAndPassword(email, password);

    return user;
  } catch (e) {
    Promise.reject(e);
  }
}
