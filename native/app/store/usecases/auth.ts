import axios from '~/lib/axios';

export function signUpWithEmailAsync(email: string, password: string, passwordConfirmation: string, username: string) {
  return async (): Promise<void> => {
    return await axios.post('/v1/auth')
      .then((res: any) => console.log('debug', res))
      .catch((err: Error) => {
        throw err;
      });
  }
}
