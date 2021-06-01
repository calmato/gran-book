import React from 'react';
import SignIn from '~/screens/SignIn';
import { useReduxDispatch } from '~/store/modules';
import {
  getAuthAsync,
  registerForPushNotificationsAsync,
  signInWithEmailAsync,
} from '~/store/usecases';

export default function ConnectedSignIn(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signInWithEmail(email: string, password: string): Promise<void> {
        return dispatch(signInWithEmailAsync(email, password));
      },
      getAuth(): Promise<void> {
        return dispatch(getAuthAsync());
      },
      registerForPushNotifications(): Promise<void> {
        return dispatch(registerForPushNotificationsAsync());
      },
    }),
    [dispatch],
  );

  return <SignIn actions={actions} />;
}
