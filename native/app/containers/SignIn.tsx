import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { signInWithEmailAsync } from '~/store/usecases/auth';
import SignIn from '~/screens/SignIn';

export default function ConnectedSignIn(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signInWithEmail(email: string, password: string): Promise<void> {
        console.log('debug', 'email', email);
        console.log('debug', 'password', password);
        return dispatch(signInWithEmailAsync(email, password));
      },
    }), [dispatch],
  );

  return <SignIn actions={actions} />;
}
