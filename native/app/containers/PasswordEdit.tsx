import React from 'react';
import PasswordEdit from '~/screens/PasswordEdit';
import { useReduxDispatch } from '~/store/modules';
import { editPasswordAsync } from '~/store/usecases';

export default function ConnectedPasswordEdit(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      editPassword(password: string, passwordConfirmation: string): Promise<void> {
        return dispatch(editPasswordAsync(password, passwordConfirmation));
      },
    }),
    [dispatch],
  );

  return <PasswordEdit actions={actions} />;
}
