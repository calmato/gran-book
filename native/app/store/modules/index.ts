import { useDispatch } from 'react-redux';
import { combineReducers, Action } from 'redux';
import { ThunkDispatch } from 'redux-thunk';
import * as AuthState from './auth';
import * as UserState from './user';
import { Auth, User } from '~/store/models';

interface InitialState {
  auth: Auth.Model;
  user: User.Model;
}

export function createInitialState(): InitialState {
  return {
    auth: AuthState.createInitialState(),
    user: UserState.createInitialState(),
  };
}

export type AppState = Readonly<ReturnType<typeof createInitialState>>;

export default combineReducers<AppState>({
  auth: AuthState.default,
  user: UserState.default,
});

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type ReduxDispatch = ThunkDispatch<AppState, InitialState, Action>;

export function useReduxDispatch(): ReduxDispatch {
  return useDispatch<ReduxDispatch>();
}
