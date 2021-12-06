import React from 'react';
import AccountSetting from '~/screens/AccoutSetting';
import { useReduxDispatch } from '~/store/modules';
import { signOut } from '~/store/usecases/v2/auth';

export default function ConnectAccountSetting(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signOut(): Promise<void> {
        return signOut();
      },
    }),
    [dispatch],
  );

  return <AccountSetting actions={actions} />;
}
