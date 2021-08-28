import React from 'react';
import { useSelector } from 'react-redux';
import OtherProfile from '~/screens/OtherProfile';
import { Auth, User } from '~/store/models';
import { useReduxDispatch } from '~/store/modules';
import { authSelector, userSelector } from '~/store/selectors';
import { getOtherProfileAsync } from '~/store/usecases/user';

export default function ConnectedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  const user: User.Model = useSelector(userSelector);
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      getOtherProfile(): Promise<void> {
        return dispatch(getOtherProfileAsync(auth.id));
      },
    }),
    [dispatch, auth.id],
  );

  return (
    <OtherProfile
      username={user.username}
      selfIntroduction={user.selfIntroduction}
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
