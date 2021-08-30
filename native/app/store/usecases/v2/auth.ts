import firebase from '~/lib/firebase';

/**
 * firebase authenticationを使って email と password でログインを実行します。
 * @param payload
 * @returns { Promise<firebase.auth.UserCredential.user | null > }
 */
export async function signInWithEmailAndPassword(payload: {
  email: string;
  password: string;
}): Promise<firebase.User | null> {
  const { email, password } = payload;

  try {
    const { user } = await firebase.auth().signInWithEmailAndPassword(email, password);

    return user;
  } catch (e) {
    return Promise.reject(e);
  }
}
