import React, { useCallback, useContext, useEffect } from 'react';
import { AuthContext } from '~/context/auth';
import { UserContext } from '~/context/user';
import OwnProfile from '~/screens/OwnProfile';
import { getOwnProfile } from '~/store/usecases/v2/user';

export default function ConnectedOwnProfile(): JSX.Element {
  const { authState } = useContext(AuthContext);
  const { userState, dispatch } = useContext(UserContext);

  const fetchOwnProfile = useCallback(async () => {
    const userValue = await getOwnProfile(authState.id, authState.token);
    dispatch({
      type: 'SET_USER',
      payload: userValue,
    });
  }, [authState.id, authState.token, dispatch]);

  useEffect(() => {
    const f = async () => {
      await fetchOwnProfile();
    };
    f();
  }, []);

  return (
    <OwnProfile
      username={userState.username}
      selfIntroduction={userState.selfIntroduction}
      thumbnailUrl={userState.thumbnailUrl}
      gender={authState.gender}
      rating={userState.rating}
      reviewCount={userState.reviewCount}
      saleCount={userState.products.length}
      followCount={userState.followCount}
      followerCount={userState.followerCount}
      actions={{ fetchOwnProfile }}
    />
  );
}
