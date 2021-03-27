import React, { ReactElement, useEffect, useState } from 'react';
import { Platform, StyleSheet, Text } from 'react-native';
import { Avatar, ListItem } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';
import { MaterialIcons } from '@expo/vector-icons';
import * as ImagePicker from 'expo-image-picker';

const styles = StyleSheet.create({
  text: {
    fontSize: 16,
  },
});

interface Props {
  avatarUrl: string | undefined,
}

const state = {
  iconState: false,
  iconUri: '',
};

const ChangeIconGroup = function ChangeIconGroup(props: Props): ReactElement {
  const [image, setImage] = useState(String);

  useEffect(() => {
    (async () => {
      if (Platform.OS !== 'web') {
        const { status } = await ImagePicker.requestCameraRollPermissionsAsync();
        if (status !== 'granted') {
          alert('Sorry, we need camera roll permissions to make this work!');
        }
      }
    })();
  }, []);

  const pickImage = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      base64: true,
      aspect: [4, 3],
      quality: 1,
    });

    if (!result.cancelled) {
      const imageEncode64 = result ? `data:image/jpg;base64,${result.base64}` : null;
      state.iconState = true;
      state.iconUri = result.uri;
      setImage(result.uri);
    }
  };

  return (
    <ListItem style={{alignItems:'flex-start'}} Component={TouchableOpacity} onPress={pickImage}>
      <Avatar source={{uri: state.iconState ? state.iconUri : props.avatarUrl}} rounded size='medium'/>
      <ListItem.Content>
        <Text style={styles.text}>アイコン変更</Text>
      </ListItem.Content>
      <MaterialIcons name="keyboard-arrow-right" size={24} color="black"/>
    </ListItem>
  );
};

export default ChangeIconGroup;
