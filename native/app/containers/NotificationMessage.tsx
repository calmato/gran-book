import React from 'react';
import { useSelector } from 'react-redux';
import NotificationMessage from '~/screens/NotificationMessage';
import { Auth } from '~/store/models';
import { authSelector } from '~/store/selectors';

export default function ConnectedNotificationMessage(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  return <NotificationMessage auth={auth} />;
}
