import React, { ReactElement, useState } from 'react';
import { StyleSheet, Text } from 'react-native';
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


const ChangeIconGroup = function ChangeIconGroup(props: Props): ReactElement {
  const [image, setImage] = useState(null);
  const pickImage = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [4, 3],
      quality: 1,
    });

    console.log(result);

    if (!result.cancelled) {
      //
    }
  };

  return (
    <ListItem style={{alignItems:'flex-start'}} Component={TouchableOpacity} onPress={pickImage}>
      <Avatar source={{uri: props.avatarUrl}} rounded size='medium'/>
      <ListItem.Content>
        <Text style={styles.text}>アイコン変更</Text>
      </ListItem.Content>
      <MaterialIcons name="keyboard-arrow-right" size={24} color="black"/>
    </ListItem>
  );
};

export default ChangeIconGroup;
