import { MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, Text } from 'react-native';
import { Avatar, ListItem } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';

const styles = StyleSheet.create({
  listItem: {
    flexDirection:'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  text: {
    fontSize: 16,
  },
});

interface Props {
  avatarUrl: string,
  handleOnPressed: () => void,
}

const ChangeIconGroup = function ChangeIconGroup(props: Props): ReactElement {
return (
  <TouchableOpacity onPress={()=>props.handleOnPressed}>
    <ListItem>
      <Avatar source={{uri: props.avatarUrl}} rounded size='medium'/>
      <ListItem.Content style={styles.listItem}>
        <Text style={styles.text}>アイコン変更</Text>
        <MaterialIcons name="keyboard-arrow-right" size={24} color="black"/>
      </ListItem.Content>
    </ListItem>
  </TouchableOpacity>
);
}

export default ChangeIconGroup
