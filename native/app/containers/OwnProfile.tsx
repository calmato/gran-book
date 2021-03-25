import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import OwnProfile from '~/screens/OwnProfile';
import { useReduxDispatch } from '~/store/modules';
import { Auth } from '~/store/models';

export default function ConnectedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return (
    <OwnProfile
      id={auth.id}
      username={auth.username}
      selfIntroduction={auth.selfIntroduction}
      thumbnailUrl={auth.thumbnailUrl}
      gender={auth.gender}
      actions={undefined}
    />
  );
}
