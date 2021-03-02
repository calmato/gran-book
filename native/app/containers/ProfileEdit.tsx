import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import { Auth } from '~/store/models';
import ProfileEdit from '~/screens/ProfileEdit';
import { useReduxDispatch } from '~/store/modules';
import { profileEditAsync } from '~/store/usecases';

export default function ConnectedProfilrEdit(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      profileEdit(username: string, gender: number, thumbnail: string | undefined, selfIntroduction: string | undefined): Promise<void> {
        return dispatch(profileEditAsync(username, gender, thumbnail, selfIntroduction));
      },
    }), [dispatch],
  );

  return <ProfileEdit 
    username={auth.username} 
    selfIntroduction={auth.selfIntroduction} 
    thumbnailUrl={auth.thumbnailUrl}
    gender={auth.gender}
    actions={actions}
  />;
}
