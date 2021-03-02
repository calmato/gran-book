import React, { ReactElement } from 'react';
import { StyleSheet, Text } from 'react-native';
import { Avatar, ListItem } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';
import { MaterialIcons } from '@expo/vector-icons';

const styles = StyleSheet.create({
  text: {
    fontSize: 16,
  },
});

interface Props {
  avatarUrl: string | undefined,
  handleOnClicked: () => void,
}

const ChangeIconGroup = function ChangeIconGroup(props: Props): ReactElement {
  return (
    <ListItem style={{alignItems:'flex-start'}} Component={TouchableOpacity} onPress={()=>props.handleOnClicked}>
      <Avatar source={{uri: props.avatarUrl}} rounded size='medium'/>
      <ListItem.Content>
        <Text style={styles.text}>アイコン変更</Text>
      </ListItem.Content>
      <MaterialIcons name="keyboard-arrow-right" size={24} color="black"/>
    </ListItem>
  );
};

export default ChangeIconGroup;
