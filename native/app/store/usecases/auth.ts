import { Dispatch } from 'redux';
import { AxiosResponse } from 'axios';
import Firebase from 'firebase';
import axios from '~/lib/axios';
import firebase from '~/lib/firebase';
import * as LocalStorage from '~/lib/local-storage';
import { Auth } from '~/store/models';
import { AppState } from '~/store/modules';
import { setAuth, setProfile } from '~/store/modules/auth';
import { IAuthResponse } from '~/types/response';

interface IAuth {
  user: Firebase.User
  token: string
}

export function signInWithEmailAsync(email: string, password: string) {
  return async (dispatch: Dispatch): Promise<void> => {
    return await firebase
      .auth()
      .signInWithEmailAndPassword(email, password)
      .then(async () => {
        return await onAuthStateChanged();
      })
      .then(async (res: IAuth) => {
        const { user, token } = res;
        const values: Auth.AuthValues = {
          id: user.uid,
          email: user.email || undefined,
          emailVerified: user.emailVerified,
          token,
        };

        const model: Auth.Model = {
          ...Auth.initialState,
          id: values.id,
          token: values.token,
          email: values.email || '',
          emailVerified: values.emailVerified || false,
        };

        dispatch(setAuth(values));
        await LocalStorage.AuthStorage.save(model);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}

export function signOutAsync() {
  return async(dispatch: Dispatch): Promise<void> => {
    return await firebase
      .auth()
      .signOut()
      .then(async () => {
        dispatch(setAuth(Auth.initialState));
      })
      .catch((err: Error) => {
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
      .then(async (res: AxiosResponse<IAuthResponse>) => {
        console.log('debug', res);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}

export function editPasswordAsync(password: string, passwordConfirmation: string) {
  return async (): Promise<void> => {
    return await axios
      .patch('/v1/auth/password', {
        password,
        passwordConfirmation,
      })
      .then(async (res: AxiosResponse<IAuthResponse>) => {
        console.log('debug', res);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}

export function getAuthAsync() {
  return async (dispatch: Dispatch, getState: () => AppState): Promise<void> => {
    return await axios
      .get('/v1/auth')
      .then(async (res: AxiosResponse<IAuthResponse>) => {
        const {
          username,
          gender,
          phoneNumber,
          role,
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
        } = res.data;

        const values: Auth.ProfileValues = {
          username,
          gender,
          phoneNumber,
          role,
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
        };

        dispatch(setProfile(values));

        const auth: Auth.Model = getState().auth;
        await LocalStorage.AuthStorage.save(auth);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}

function onAuthStateChanged(): Promise<IAuth> {
  return new Promise((resolve: (auth: IAuth) => void, reject: (reason: Error) => void) => {
    firebase
      .auth()
      .onAuthStateChanged(async (user: Firebase.User | null) => {
        if (!user) {
          reject(new Error('Unauthorized'));
          return;
        }

        if (!user.emailVerified) {
          sendEmailVerification();
          reject(new Error('Email address is unapproved'));
          return;
        }

        await getIdToken()
          .then((token: string) => {
            resolve({ user, token });
          })
          .catch((err: Error) => {
            reject(err);
          });
      });
  });
}

function getIdToken(): Promise<string> {
  return new Promise((resolve: (token: string) => void, reject: (reason: Error) => void) => {
    firebase
      .auth()
      .currentUser?.getIdToken(true)
      .then((token: string) => {
        resolve(token);
      })
      .catch((err: Error) => {
        reject(err);
      });
  });
}

function sendEmailVerification(): Promise<void> {
  return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
    firebase
      .auth()
      .currentUser?.sendEmailVerification()
      .then(() => {
        resolve();
      })
      .catch((err: Error) => {
        reject(err);
      });
  });
}

export function editEmailAsync(email: string) {
  return async (): Promise<void> => {
    return await axios
      .patch('/v1/auth/email', {
        email
      })
      .then(async (res: AxiosResponse<IAuthResponse>) => {
        console.log('debug', res);
      })
      .catch((err: Error) => {
        throw err;
      });
  };
}
