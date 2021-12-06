import React from 'react';
import SignIn from '~/screens/SignIn';

import { signInWithEmailAndPassword } from '~/store/usecases/v2/auth';

export default function ConnectedSignIn(): JSX.Element {
  const actions = React.useMemo(
    () => ({
      signInWithEmailAndPassword: signInWithEmailAndPassword,
    }),
    [],
  );

  return <SignIn actions={actions} />;
}
