import React, { createContext, useEffect, useReducer } from 'react';
import firebase from '~/lib/firebase';
import { AuthValues, initialState, Model, ProfileValues } from '~/store/models/auth';

interface AuthContextProps {
  authState: Model;
  dispatch: React.Dispatch<AuthStateAction>;
}

const AuthContext = createContext<AuthContextProps>({
  authState: initialState,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  dispatch: () => {},
});

export type AuthStateAction =
  | { type: 'SET_AUTH_VALUES'; payload: AuthValues }
  | { type: 'SET_PROFILE_VALUES'; payload: ProfileValues };

const reducer: React.Reducer<Model, AuthStateAction> = function reducer(
  state: Model,
  action: AuthStateAction,
): Model {
  switch (action.type) {
    case 'SET_AUTH_VALUES':
      return {
        ...state,
        ...action.payload,
      };
    case 'SET_PROFILE_VALUES':
      return {
        ...state,
        ...action.payload,
      };
    default:
      return state;
  }
};

interface Props {
  children?: React.ReactNode;
}

const AuthProvider = function AuthProvider({ children }: Props) {
  const [authState, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    const f = async () => {
      firebase.auth().onAuthStateChanged(async (user) => {
        if (user) {
          const token = await user.getIdToken();
          dispatch({
            type: 'SET_AUTH_VALUES',
            payload: {
              id: user.uid,
              email: user.email || undefined,
              emailVerified: user.emailVerified,
              token,
            },
          });
        }
      });
    };
    f();
  }, []);

  return <AuthContext.Provider value={{ authState, dispatch }}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
