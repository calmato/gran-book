import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { getAuthAsync, signInWithFacebookAsync  } from '~/store/usecases';
import SignInSelect from '~/screens/SignInSelect';

export default function ConnectedSignInSelect(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      getAuth(): Promise<void> {
        return dispatch(getAuthAsync());
      },
      signInWithFacebook(): Promise<void> {
        return dispatch(signInWithFacebookAsync())
      }
    }),
    [dispatch],
  );

  return <SignInSelect actions={actions} />;
}