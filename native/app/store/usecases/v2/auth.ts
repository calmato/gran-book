import { internal } from '~/lib/axios';
import firebase from '~/lib/firebase';
import { SignInForm, SingUpForm } from '~/types/forms';

const API_VERSION = 'v1';

/**
 * firebase authenticationを使って email と password でログインを実行します。
 * @param payload
 */
export async function signInWithEmailAndPassword(payload: SignInForm) {
  const { email, password } = payload;

  try {
    const { user } = await firebase.auth().signInWithEmailAndPassword(email, password);

    return user;
  } catch (e) {
    return Promise.reject(e);
  }
}

/**
 * バックエンドAPIにリクエストを送りユーザーを新規登録します。
 * ユーザー登録後にfirebase authenticationから登録メールアドレスに認証リンク付メールを送信します。
 * @param payload
 */
export async function signUpWithEmail(payload: SingUpForm) {
  try {
    await internal.post(`/${API_VERSION}/auth`, payload);
    await firebase.auth().currentUser?.sendEmailVerification();
  } catch (e) {
    Promise.reject(e);
  }
}
