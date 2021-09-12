import React, { createContext, useContext, useEffect, useReducer } from 'react';
import * as UiContext from '~/lib/context/ui';
import firebase from '~/lib/firebase';
import { AuthValues, initialState, Model, ProfileValues } from '~/store/models/auth';

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
    default:
      return state;
  }
};

interface Props {
  children?: React.ReactNode;
}

const AuthProvider = function AuthProvider({ children }: Props) {
  const [authState, dispatch] = useReducer(reducer, initialState);

  const uiContext = useContext(UiContext.Context);

  useEffect(() => {
    const unsubscribed = firebase.auth().onAuthStateChanged((user) => {
      if (user) {
        user.getIdToken().then((token) => {
          dispatch({
            type: 'SET_AUTH_VALUES',
            payload: {
              id: user.uid,
              email: user.email || undefined,
              emailVerified: user.emailVerified,
              token,
            },
          });
          uiContext.setApplicationState(UiContext.Status.AUTHORIZED);
        });
      }
      uiContext.setApplicationState(UiContext.Status.UN_AUTHORIZED);
    });
    return () => unsubscribed();
  }, []);

  return <AuthContext.Provider value={{ authState, dispatch }}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
