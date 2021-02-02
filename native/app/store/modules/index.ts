import { AppState } from 'react-native';
import { useDispatch } from 'react-redux';
import { combineReducers, Action } from 'redux';
import { ThunkDispatch } from 'redux-thunk';
import * as Auth from './auth';

export function createInitialState() {
  return {
    auth: Auth.createInitialState(),
  };
}

export type AppState = Readonly<ReturnType<typeof createInitialState>>;

export default combineReducers<AppState>({
  auth: Auth.default,
});

export type ReduxDispatch = ThunkDispatch<AppState, any, Action>;

export function useReduxDispatch(): ReduxDispatch {
  return useDispatch<ReduxDispatch>();
}
