import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import Support from '~/screens/Support';
import { Auth } from '~/store/models';

export default function ConnectedSupport(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <Support auth={auth} />;
}
