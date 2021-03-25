import React from 'react';
import { useSelector } from 'react-redux';
import { authSelector, userSelector } from '~/store/selectors';
import OwnProfile from '~/screens/OwnProfile';
import { useReduxDispatch } from '~/store/modules';
import { Auth, User } from '~/store/models';
import { getOwnProfileAsync } from '~/store/usecases/user';

export default function ConnectedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  const user: User.Model = useSelector(userSelector);
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      getOwnProfile(): Promise<void> {
        return dispatch(getOwnProfileAsync(auth.id));
      },
    }),
    [dispatch],
  );

  return (
    <OwnProfile
      username={auth.username}
      selfIntroduction={auth.selfIntroduction}
      thumbnailUrl={user.thumbnailUrl}
      gender={auth.gender}
      rating={user.rating}
      reviewCount={user.reviewCount}
      saleCount={user.products.length}
      followCount={user.followCount}
      followerCount={user.followerCount}
      actions={actions}
    />
  );
}
