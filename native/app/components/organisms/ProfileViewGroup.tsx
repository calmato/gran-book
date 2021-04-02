import React, { ReactElement } from 'react';
import { View } from 'react-native';
import ProfileBasicInfoGroup from '../molecules/ProfileBasicInfoGroup';
import ProfileFollowFollwer from '~/components/molecules/ProfileFollowFollower';

interface Props {
  name: string;
  avatarUrl: string | undefined;
  rating: number;
  reviewCount: number;
  saleCount: number;
  followerCount: number;
  followCount: number;
  buttonTitle: string;
  handleClick: () => void;
}

const ProfileViewGroup = function ProfileViewGroup(props: Props): ReactElement {
  return (
    <View>
      <ProfileBasicInfoGroup
        name={props.name}
        avatarUrl={props.avatarUrl}
        rating={props.rating}
        reviewCount={props.reviewCount}
        buttonTitle={props.buttonTitle}
        handleClick={props.handleClick}
      />
      <ProfileFollowFollwer
        saleCount={props.saleCount}
        followerCount={props.followerCount}
        followCount={props.followCount}
      />
    </View>
  );
};

export default ProfileViewGroup;
