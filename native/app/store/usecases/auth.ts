import { AxiosResponse } from 'axios';
import Firebase from 'firebase';
import axios from '~/lib/axios';
import firebase from '~/lib/firebase';
import { ISignUpResponse } from '~/types/response';

export function signInWithEmailAsync(email: string, password: string) {
  return async (): Promise<void> => {
    return await firebase
      .auth()
      .signInWithEmailAndPassword(email, password)
      .then((res: Firebase.auth.UserCredential) => {
        console.log('debug', res);
      })
      .catch((err: Error) => {
        console.log('debug', err);
        throw err;
      });
  };
}

export function signUpWithEmailAsync(email: string, password: string, passwordConfirmation: string, username: string) {
  return async (): Promise<void> => {
    return await axios
      .post('/v1/auth', {
        email,
        password,
        passwordConfirmation,
        username,
      })
      .then((res: AxiosResponse<ISignUpResponse>) => {
        // TODO: レスポンス処理
        console.log('debug', res);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}
