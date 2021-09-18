import React, { useContext } from 'react';
import { AuthContext } from '~/context/auth';
import MyPage from '~/screens/MyPage';

export default function ConnectedMyPage(): JSX.Element {
  const { authState } = useContext(AuthContext);

  return <MyPage auth={authState} />;
}
