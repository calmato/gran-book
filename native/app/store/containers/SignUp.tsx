import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { signUpWithEmailAsync } from '~/store/usecases/auth';
import SignUp from '~/screens/SignUp';

export default function ConnectedSignUp() {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signUpWithEmail(email: string, password: string, passwordConfirmation: string, username: string): Promise<void> {
        return dispatch(signUpWithEmailAsync(email, password, passwordConfirmation, username));
      },
    }), [dispatch],
  );

  return <SignUp actions={actions} />;
}
