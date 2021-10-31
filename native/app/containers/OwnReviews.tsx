import React from 'react';
import { useSelector } from 'react-redux';
import { OwnReviews } from '~/screens/OwnReviews';
import { Auth } from '~/store/models';
import { authSelector } from '~/store/selectors';

export default function ConnentedOwnProfile(): JSX.Element {
  const auth: Auth.Model = useSelector(authSelector);
  return <OwnReviews auth={auth} />;
}
