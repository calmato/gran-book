import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View, Alert, ScrollView, Text, TouchableOpacity } from 'react-native';
import { Avatar, Button, Input, ListItem } from 'react-native-elements';
import ChangeNickname from '~/components/organisms/ChangeNickname';
import GenderRadioGroup from '~/components/organisms/GenderRadioGroup';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { ProfileEditForm } from '~/types/forms';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import * as ImagePicker from 'expo-image-picker';
import { MaterialIcons } from '@expo/vector-icons';

const styles = StyleSheet.create({
  selfIntroduction: {
    minHeight: 100,
    textAlignVertical: 'top',
    paddingRight: 20,
  },
  button: {
    marginTop: 20,
    marginBottom: 20,
    alignSelf: 'center',
  },
  text: {
    fontSize: 16,
  },
});

interface Props {
  username: string;
  selfIntroduction: string | '';
  thumbnailUrl: string | undefined;
  gender: number;
  actions: {
    profileEdit: (
      username: string,
      gender: number,
      thumbnail: string | undefined,
      selfIntroduction: string,
    ) => Promise<void>;
  };
}
let imageEncode64: string | undefined = '';

const ProfileEdit = function ProfileEdit(props: Props): ReactElement {
  const [userInfo, setValue] = useState<ProfileEditForm>({
    name: props.username,
    avatar: props.thumbnailUrl,
    selfIntroduction: props.selfIntroduction,
    gender: props.gender,
  });
  const navigation = useNavigation();
  const { profileEdit } = props.actions;

  const nameError: boolean = useMemo((): boolean => {
    return userInfo.name === '';
  }, [userInfo.name]);

  const handleGenderChange = (value: string) => {
    switch (value) {
      case '男性':
        setValue({ ...userInfo, gender: 1 });
        break;
      case '女性':
        setValue({ ...userInfo, gender: 2 });
        break;
      default:
        setValue({ ...userInfo, gender: 0 });
        break;
    }
  };

  const pickImage = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      base64: true,
      aspect: [4, 3],
      quality: 1,
    });

    if(result !== undefined && result !== null) {
      if (!result.cancelled) {
        imageEncode64 = result ? result.base64 : '';
        setValue({ ...userInfo, avatar: result.uri });
      }
    }
  };

  const createAlertNotifyProfileEditError = (code: number) =>
    Alert.alert('ユーザー登録に失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSubmit = React.useCallback(async () => {
    await profileEdit(userInfo.name, userInfo.gender, imageEncode64, userInfo.selfIntroduction)
      .then(() => {
        navigation.navigate('OwnProfile');
      })
      .catch((err) => {
        createAlertNotifyProfileEditError(err.code);
      });
  }, [userInfo, profileEdit, navigation]);

  return (
    <View>
      <ScrollView stickyHeaderIndices={[0]}>
        <HeaderWithBackButton title="プロフィール編集" onPress={() => navigation.goBack()} />
        <ListItem
          style={{ alignItems: 'flex-start' }}
          Component={TouchableOpacity}
          onPress={pickImage}>
          <Avatar source={{ uri: userInfo.avatar }} rounded size="medium" />
          <ListItem.Content>
            <Text style={styles.text}>アイコン変更</Text>
          </ListItem.Content>
          <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
        </ListItem>
        <ChangeNickname
          value={userInfo.name}
          handelOnChangeText={(text) => setValue({ ...userInfo, name: text })}
        />
        <Input
          style={styles.selfIntroduction}
          placeholder={'自己紹介を入力してください'}
          multiline={true}
          maxLength={256}
          onChangeText={(text) => setValue({ ...userInfo, selfIntroduction: text })}
          value={userInfo.selfIntroduction}
        />
        <GenderRadioGroup
          handleOnChange={(value) => handleGenderChange(value)}
          data={[{ label: '男性' }, { label: '女性' }, { label: '未選択' }]}
          title={'性別'}
          initial={userInfo.gender === 0 ? 3 : userInfo.gender}
        />
        <Button title={'保存する'} onPress={handleSubmit} containerStyle={styles.button} />
      </ScrollView>
    </View>
  );
};

export default ProfileEdit;
