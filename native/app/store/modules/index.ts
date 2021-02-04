import { useDispatch } from 'react-redux';
import { combineReducers, Action } from 'redux';
import { ThunkDispatch } from 'redux-thunk';
import * as AuthState from './auth';
import { Auth } from '~/store/models';

interface InitialState {
  auth: Auth.Model
}

export function createInitialState(): InitialState {
  return {
    auth: AuthState.createInitialState(),
  };
}

export type AppState = Readonly<ReturnType<typeof createInitialState>>;

export default combineReducers<AppState>({
  auth: AuthState.default,
});

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type ReduxDispatch = ThunkDispatch<AppState, InitialState, Action>;

export function useReduxDispatch(): ReduxDispatch {
  return useDispatch<ReduxDispatch>();
}
