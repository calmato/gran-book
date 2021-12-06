import React, { useState } from 'react';

export enum Status {
  LOADING = 'loading', // アプリ起動時のロード処理中
  FIRST_OPEN = 'firstOpen', // 初回起動時
  UN_AUTHORIZED = 'unAuthorized', // 初回起動後で認証が済んでいない状態
  AUTHORIZED = 'authorized', // 初回起動後で認証が済んでいる状態
}

export function createApplicationInitialState(): Status {
  return Status.LOADING;
}

const UiContext = React.createContext({
  applicationState: createApplicationInitialState(),
  setApplicationState: (_: Status) => {
    return;
  },
});

interface Props {
  children?: React.ReactNode;
}

const UiProvider = function UiProvider({ children }: Props) {
  const [applicationState, setApplicationState] = useState(createApplicationInitialState());

  return (
    <UiContext.Provider value={{ applicationState, setApplicationState }}>
      {children}
    </UiContext.Provider>
  );
};

export { UiContext, UiProvider };
