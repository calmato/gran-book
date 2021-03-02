import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import { Auth } from '~/store/models';
import ProfileEdit from '~/screens/ProfileEdit';

export default function ConnectedProfilrEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);

  return <ProfileEdit 
  username={auth.username} 
  selfIntroduction={auth.selfIntroduction} 
  thumbnailUrl={auth.thumbnailUrl}
  gender={auth.gender}
  />
}