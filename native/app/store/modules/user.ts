import { User } from '../models';

export function createInitialState(): User.Model {
  return User.factory();
}

export type State = ReturnType<typeof createInitialState>;

export const SET_USER = 'gran-book/user/SET_USER';
export const RESET = 'gran-book/user/RESET';

export function setUser(user: User.UserValues) {
  return {
    type: SET_USER,
    payload: { user },
  };
}

export function reset() {
  return {
    type: RESET,
    payload: {},
  };
}

export type Action = Readonly<ReturnType<typeof setUser>> | Readonly<ReturnType<typeof reset>>;

export default function reducer(state: State = createInitialState(), action: Action): User.Model {
  const { payload } = action;

  switch (action.type) {
    case SET_USER:
      return User.setUser(state, payload.user);
    default:
      return state;
  }
}
