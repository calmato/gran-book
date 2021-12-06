import React, { createContext, useReducer } from 'react';
import { initialState, Model, UserValues } from '~/store/models/user';

interface UserContextProps {
  userState: Model;
  dispatch: React.Dispatch<UserStateAction>;
}

const UserContext = createContext<UserContextProps>({
  userState: initialState,
  dispatch: () => {
    return;
  },
});

type ActionType = 'SET_USER';
type Payload = UserValues;

interface UserStateAction {
  type: ActionType;
  payload: Payload;
}

const reducer: React.Reducer<Model, UserStateAction> = function reducer(
  state: Model,
  action: UserStateAction,
): Model {
  switch (action.type) {
    case 'SET_USER':
      return {
        ...state,
        ...action.payload,
      };
  }
};

interface Props {
  children?: React.ReactNode;
}

const UserProvider = function UserProvider({ children }: Props) {
  const [userState, dispatch] = useReducer(reducer, initialState);

  return (
    <UserContext.Provider
      value={{
        userState,
        dispatch,
      }}>
      {children}
    </UserContext.Provider>
  );
};

export { UserContext, UserProvider };
