import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { signOutAsync } from '~/store/usecases';
import AccountSetting from '~/screens/AccoutSetting';

export default function ConnectAccountSetting(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signOut(): Promise<void> {
        return dispatch(signOutAsync());
      },
    }),
    [dispatch],
  );

  return <AccountSetting actions={actions} />;
}
