import React from 'react';
import { useSelector } from 'react-redux';
import EmailEdit from '~/screens/EmailEdit';
import { Auth } from '~/store/models';
import { useReduxDispatch } from '~/store/modules';
import { authSelector } from '~/store/selectors';
import { editEmailAsync, signOutAsync } from '~/store/usecases';

export default function ConnectedEmailEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      emailEdit(email: string): Promise<void> {
        return dispatch(editEmailAsync(email));
      },
      signOut(): Promise<void> {
        return dispatch(signOutAsync());
      },
    }),
    [dispatch],
  );

  return <EmailEdit email={auth.email} actions={actions} />;
}
