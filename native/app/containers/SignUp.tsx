import React from 'react';
import SignUp from '~/screens/SignUp';
import { signUpWithEmail } from '~/store/usecases/v2/auth';

export default function ConnectedSignUp(): JSX.Element {
  const actions = React.useMemo(
    () => ({
      signUpWithEmail: signUpWithEmail,
    }),
    [],
  );

  return <SignUp actions={actions} />;
}
