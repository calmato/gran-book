import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import { Auth } from '~/store/models';
import OwnProfile from '~/screens/OwnProfile';

export default function ConnectedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <OwnProfile 
  username={auth.username} 
  selfIntroduction={auth.selfIntroduction} 
  thumbnailUrl={auth.thumbnailUrl}
  gender={auth.gender}
  />
}