import axios from '~/lib/axios';
import { ISignUpResponse } from '~/types/response';

export function signUpWithEmailAsync(email: string, password: string, passwordConfirmation: string, username: string) {
  return async (): Promise<void> => {
    return await axios
      .post('/v1/auth', {
        email,
        password,
        passwordConfirmation,
        username,
      })
      .then((res: ISignUpResponse) => {
        // TODO: レスポンス処理
        console.log('debug', res);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}
