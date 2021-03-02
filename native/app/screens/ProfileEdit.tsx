import { RouteProp, useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Input } from 'react-native-elements';
import ChangeIconGroup from '~/components/organisms/ChangeIconGroup';
import ChangeNickname from '~/components/organisms/ChangeNickname';
import GenderRadioGroup from '~/components/organisms/GenderRadioGroup';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { ProfileEditForm } from '~/types/forms';
import { UserInfoStackParamList } from '~/types/navigation';

const styles = StyleSheet.create({
  bio: {
    minHeight: 100,
    textAlignVertical: 'top',
    paddingRight: 20,
  },
  button: {
    marginTop: 20,
    alignSelf: 'center',
  },
});

interface Props {
  username: string, 
  selfIntroduction: string | undefined, 
  thumbnailUrl: string | undefined, 
  gender: number
}

const ProfileEdit = function ProfileEdit(props: Props): ReactElement {
  const [userInfo, setValue] = useState<ProfileEditForm>({
    name: props.username,
    avatar: props.thumbnailUrl,
    bio: props.selfIntroduction,
    gender: props.gender,
  });
  const navigation = useNavigation();

  const nameError: boolean = useMemo((): boolean => {
    return (userInfo.name === '');
  }, [userInfo.name]);

  const handleGenderChange = (value: string) => {
    switch (value) {
    case '男性':
      setValue({...userInfo, gender: 1});
      break;
    case '女性':
      setValue({...userInfo, gender: 2});
      break;
    default:
      setValue({...userInfo, gender: 0});
      break;
    }
  };
  return (
    <View>
      <HeaderWithBackButton 
        title='プロフィール編集'
        onPress={()=>navigation.goBack()}
      />
      <ChangeIconGroup
        avatarUrl={userInfo.avatar}
        handleOnClicked={()=>undefined}
      />
      <ChangeNickname
        value={userInfo.name}
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
        handleOnChange={(value)=>handleGenderChange(value)}
        data={[{label:'男性'}, {label:'女性'}, {label:'未選択'}]}
        title={'性別'}
        initial={(userInfo.gender === 0) ? 3 : userInfo.gender}
      />
      <Button
        title={'保存する'}
        onPress={()=>undefined}
        containerStyle={styles.button} 
      />
    </View>
  );
};

export default ProfileEdit;
