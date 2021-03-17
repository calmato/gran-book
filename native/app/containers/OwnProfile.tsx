import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector } from '~/store/selectors';
import OwnProfile from '~/screens/OwnProfile';
import { getOwnProfileAsync } from '~/store/usecases';
import { useReduxDispatch } from '~/store/modules';
import { Auth } from '~/store/models';

export default function ConnectedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  const dispatch = useReduxDispatch();
  
  const actions = React.useMemo(
    () => ({
      getOwnProfile(id: string): Promise<void> {
        return dispatch(getOwnProfileAsync(id))
      },
    }), [dispatch],
  );
  
  return <OwnProfile 
    id={auth.id}
    username={auth.username} 
    selfIntroduction={auth.selfIntroduction} 
    thumbnailUrl={auth.thumbnailUrl}
    gender={auth.gender}
    followCount={auth.followCount}
    followerCount={auth.followerCount}
    reviewCount={auth.reviewCount}
    rating={auth.rating}
    actions={actions}
  />;
}
