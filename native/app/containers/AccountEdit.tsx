import React from 'react';
import { useReduxDispatch } from '~/store/modules';
import { searchAddress } from '~/store/usecases';
import AccountEdit from '~/screens/AccountEdit';

export default function SearchAddress(): JSX.Element {
  const dispatch = useReduxDispatch();
  const actions = React.useMemo(
    () => ({
      searchAddress(postalCode: string): Promise<string> {
        return dispatch(searchAddress(postalCode));
      },
    }), [dispatch],
  );

  return <AccountEdit actions={actions} />;
}
