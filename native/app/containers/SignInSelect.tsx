import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { getAuthAsync, registerForPushNotificationsAsync  } from '~/store/usecases';
import SignInSelect from '~/screens/SignInSelect';

export default function ConnectedSignInSelect(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      getAuth(): Promise<void> {
        return dispatch(getAuthAsync());
      },
      registerForPushNotifications(): Promise<void> {
        return dispatch(registerForPushNotificationsAsync());
      },
    }),
    [dispatch],
  );

  return <SignInSelect actions={actions} />;
}
