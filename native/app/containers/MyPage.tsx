import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import MyPage from '~/screens/MyPage';
import { Auth } from '~/store/models';

export default function ConnectedMyPage(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <MyPage auth={auth} />;
}
