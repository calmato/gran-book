import { Auth } from '~/store/models';
import { AppState } from '~/store/modules';

export const authSelector = (state: AppState): Auth.Model => {
  return state.auth;
};
