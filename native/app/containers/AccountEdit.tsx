import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { editAccountAsync } from '~/store/usecases';
import AccountEdit from '~/screens/AccountEdit';
import { AccountEditForm } from '~/types/forms';

export default function AccountSave(): JSX.Element {
  const dispatch = useReduxDispatch();
  const actions = React.useMemo(
    () => ({
      accountEdit(formData: AccountEditForm): Promise<void> {
        return dispatch(editAccountAsync(formData));
      },
    }),
    [dispatch],
  );

  return <AccountEdit actions={actions} />;
}
