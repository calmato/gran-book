import React from 'react';
import SignUp from '~/screens/SignUp';
import { useReduxDispatch } from '~/store/modules';
import { signUpWithEmailAsync } from '~/store/usecases';

export default function ConnectedSignUp(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      signUpWithEmail(
        email: string,
        password: string,
        passwordConfirmation: string,
        username: string,
      ): Promise<void> {
        return dispatch(signUpWithEmailAsync(email, password, passwordConfirmation, username));
      },
    }),
    [dispatch],
  );

  return <SignUp actions={actions} />;
}
