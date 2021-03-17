import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { ListItem, Button, Rating, Text, Avatar } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    padding: 10,
  },
  reviewContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingTop:10,
    paddingBottom:10,
  },
  reviewRate: {
    paddingStart:10,
    paddingEnd:10,
    color: COLOR.TEXT_DEFAULT,
  },
  icon:{
    paddingStart:30,
  },
  button: {
    height:35,
    alignSelf:'flex-end',
  },
  buttonTitle: {
    color: COLOR.TEXT_TITLE,
  },
});

interface Props {
  name: string
  avatarUrl: string | undefined
  rating: number
  reviewCount: number
  buttonTitle: string
  handleClick: () => void
}

const ProfileBasicInfoGroup = function ProfileBasicInfoGroup(props: Props): ReactElement {
  return (
    <ListItem style={styles.container}>
      <Avatar source={{uri: props.avatarUrl}} rounded size='medium'/>
      <ListItem.Content>
        <ListItem.Title>{props.name}</ListItem.Title>
        <View style={styles.reviewContainer}>
          <Rating 
            fractions={1}
            readonly={true}
            startingValue={props.rating}
            imageSize={20}
          />
          <Text style={styles.reviewRate}>{props.rating}</Text>
          <Text>({props.reviewCount}件)</Text>
          <MaterialIcons name="keyboard-arrow-right" size={24} color="black" style={styles.icon}/>
        </View>
        <Button title={props.buttonTitle} buttonStyle={styles.button} titleStyle={styles.buttonTitle} onPress={props.handleClick}/>
      </ListItem.Content>
    </ListItem>
  );
};

export default ProfileBasicInfoGroup;
