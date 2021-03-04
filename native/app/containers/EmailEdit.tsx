import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import { Auth } from '~/store/models';
import EmailEdit from '~/screens/EmailEdit';

export default function ConnectedEmailEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <EmailEdit email={auth.email} />;
}