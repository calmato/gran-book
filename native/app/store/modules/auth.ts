import { Auth } from '~/store/models';

// Initial
export function createInitialState(): Auth.Model {
  return Auth.factory();
}

export type State = ReturnType<typeof createInitialState>;

// Actions
export const SET_AUTH = 'gran-book/auth/SET_AUTH';
export const SET_PROFILE = 'gran-book/auth/SET_PROFILE';
export const SET_OWN_PROFILE = 'gran-book/auth/SET_OWN_PROFILE';
export const RESET = 'gran-book/auth/RESET';

// Action Creattors
export function setAuth(auth: Auth.AuthValues) {
  return {
    type: SET_AUTH,
    payload: { auth },
  };
}

export function setProfile(auth: Auth.ProfileValues) {
  return {
    type: SET_PROFILE,
    payload: { auth },
  };
}

export function reset() {
  return {
    type: RESET,
    payload: {},
  };
}

export type Action =
  Readonly<ReturnType<typeof setAuth>>
  | Readonly<ReturnType<typeof setProfile>>
  | Readonly<ReturnType<typeof reset>>;

// Reducer
export default function reducer(state: State = createInitialState(), action: Action): Auth.Model {
  const { payload } = action;

  switch (action.type) {
    case SET_AUTH:
      return Auth.setAuth(state, payload.auth);
    case SET_PROFILE:
      return Auth.setProfile(state, payload.auth);
    case RESET:
      return Auth.factory();
    default:
      return state;
  }
}
