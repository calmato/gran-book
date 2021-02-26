import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Input } from 'react-native-elements';
import ChangeIconGroup from '~/components/organisms/ChangeIconGroup';
import ChangeNickname from '~/components/organisms/ChangeNickname';
import GenderRadioGroup from '~/components/organisms/GenderRadioGroup';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { ProfileEditForm } from '~/types/forms';

const styles = StyleSheet.create({
  container: {

  },
  bio: {
    minHeight: 100,
    textAlignVertical: 'top',
    paddingRight: 20,
  },
});

const ProfileEdit = function ProfileEdit(props: Props): ReactElement {
  const [userInfo, setValue] = useState<ProfileEditForm>({
    name: 'hamachans',
    avatarUrl: 'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
    bio: 'よろしくお願いします。',
    gender: 3,
  })

  const nameError: boolean = useMemo((): boolean => {
    return (userInfo.name === '');
  }, [userInfo.name])

  return (
  <View>
    <HeaderWithBackButton 
      title='プロフィール編集'
      onPress={()=>undefined}
    />
    <ChangeIconGroup
      avatarUrl={userInfo.avatarUrl}
      handleOnPressed={()=>undefined}
    />
    <ChangeNickname
    defaultValue={userInfo.name}
    handelOnChangeText={(text)=>setValue({...userInfo, name: text})}
    />
    <Input
    style={styles.bio}
    placeholder={'自己紹介を入力してください'}
    multiline={true}
    maxLength={256}
    onChangeText={(text)=>setValue({...userInfo, bio: text})}
    value={userInfo.bio}
    />
    <GenderRadioGroup
    gender={userInfo.gender}
    />
  </View>
  );
}

export default ProfileEdit
