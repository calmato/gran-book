import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { editPasswordAsync } from '~/store/usecases';
import PasswordEdit from '~/screens/PasswordEdit';

export default function ConnectedPasswordEdit(): JSX.Element {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      editPassword(password: string, passwordConfirmation: string): Promise<void> {
        return dispatch(editPasswordAsync(password, passwordConfirmation));
      },
    }), [dispatch],
  );

  return <PasswordEdit actions={actions} />;
}
