import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import firebase from '~/lib/firebase';
import { Auth } from '~/store/models';
import { AuthV1Response } from '~/types/api/auth_apiv1_response_pb';
import { SignInForm, SingUpForm, PasswordResetForm } from '~/types/forms';

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
    return Promise.reject(e);
  }
}

/**
 * firebase authenticationを使ってユーザーをログアウトします。
 * @returns
 */
export async function signOut() {
  await firebase.auth().signOut();
}

/**
 * バックエンドAPIにリクエストを送り、現在ログインしているユーザーのプロフィール情報を取得します。
 * @returns Promise<Auth.ProfileValues>
 */
export async function getProfile() {
  try {
    const { data }: AxiosResponse<AuthV1Response.AsObject> = await internal.get(
      `${API_VERSION}/auth`,
    );

    const {
      username,
      gender,
      phoneNumber,
      thumbnailUrl,
      selfIntroduction,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
      postalCode,
      prefecture,
      city,
      addressLine1,
      addressLine2,
      createdAt,
      updatedAt,
    } = data;

    // const values: Auth.ProfileValues
    return {
      username,
      gender,
      phoneNumber,
      role: 0, // TODO: remove
      thumbnailUrl,
      selfIntroduction,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
      postalCode,
      prefecture,
      city,
      addressLine1,
      addressLine2,
      createdAt,
      updatedAt,
    } as Auth.ProfileValues;
  } catch (e) {
    return Promise.reject(e);
  }
}

/**
 * firebase authenticationを使って入力されたメールアドレスにパスワードリセットのリンクを送信します。
 * @param payload
 * @returns
 */
export async function sendPasswordResetEmail(payload: PasswordResetForm) {
  try {
    const { email } = payload;
    await firebase.auth().sendPasswordResetEmail(email);
  } catch (e) {
    return Promise.reject(e);
  }
}
