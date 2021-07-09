import React from 'react';
import { useSelector } from 'react-redux';
import ContactEdit from '~/screens/ContactEdit';
import { Auth } from '~/store/models';
import { authSelector } from '~/store/selectors';

export default function ConnectedContactEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <ContactEdit email={auth.email} />;
}
