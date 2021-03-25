import { User } from "../models";
import { AppState } from '~/store/modules';

export const userSelector = (state: AppState): User.Model => {
  return state.user;
};
