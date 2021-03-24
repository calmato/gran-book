import React, { ReactElement } from 'react';
import { View } from 'react-native';
import ProfileBasicInfoGroup from '../molecules/ProfileBasicInfoGroup';
import ProfileFollowFollwer from '~/components/molecules/ProfileFollowFollower';

interface Props {
  name: string;
  avatarUrl: string | undefined;
  rating: number;
  reviewNum: number;
  saleNum: number;
  followerNum: number;
  followNum: number;
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
        reviewNum={props.reviewNum}
        buttonTitle={props.buttonTitle}
        handleClick={props.handleClick}
      />
      <ProfileFollowFollwer
        saleNum={props.saleNum}
        followerNum={props.followerNum}
        followNum={props.followNum}
      />
    </View>
  );
};

export default ProfileViewGroup;
