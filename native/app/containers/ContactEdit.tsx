import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import { Auth } from '~/store/models';
import ContactEdit from '~/screens/ContactEdit';

export default function ConnectedContactEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <ContactEdit email={auth.email} />;
}
