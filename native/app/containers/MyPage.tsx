import React from 'react';
import { useSelector } from 'react-redux';
import MyPage from '~/screens/MyPage';
import { Auth } from '~/store/models';
import { authSelector } from '~/store/selectors';

export default function ConnectedMyPage(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <MyPage auth={auth} />;
}
