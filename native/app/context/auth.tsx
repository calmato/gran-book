import React, { createContext, useContext, useEffect, useReducer } from 'react';
import { Status, UiContext } from './ui';
import firebase from '~/lib/firebase';
import { AuthValues, initialState, Model, ProfileValues } from '~/store/models/auth';
import { getProfile, signOut } from '~/store/usecases/v2/auth';

interface AuthContextProps {
  authState: Model;
  dispatch: React.Dispatch<AuthStateAction>;
}

const AuthContext = createContext<AuthContextProps>({
  authState: initialState,
  dispatch: () => {
    return;
  },
});

type ActionType = 'SET_AUTH_VALUES' | 'SET_PROFILE_VALUES';
type Payload = AuthValues | ProfileValues;

interface AuthStateAction {
  type: ActionType;
  payload: Payload;
}

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
  }
};

interface Props {
  children?: React.ReactNode;
}

const AuthProvider = function AuthProvider({ children }: Props) {
  const [authState, dispatch] = useReducer(reducer, initialState);

  const { setApplicationState } = useContext(UiContext);

  useEffect(() => {
    const unsubscribed = firebase.auth().onAuthStateChanged((user) => {
      if (user && user.emailVerified) {
        user.getIdToken().then(async (token) => {
          dispatch({
            type: 'SET_AUTH_VALUES',
            payload: {
              id: user.uid,
              email: user.email || undefined,
              emailVerified: user.emailVerified,
              token,
            },
          });
          const profileValues = await getProfile(token);
          dispatch({
            type: 'SET_PROFILE_VALUES',
            payload: profileValues,
          });
          setApplicationState(Status.AUTHORIZED);
        });
      } else {
        signOut();
      }
      setApplicationState(Status.UN_AUTHORIZED);
    });

    return () => {
      unsubscribed();
    };
  }, [setApplicationState]);

  return <AuthContext.Provider value={{ authState, dispatch }}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
