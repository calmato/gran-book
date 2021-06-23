import { MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { ListItem, Button, Rating, Text, Avatar } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  reviewContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginTop: 6,
  },
  reviewRate: {
    marginStart: 10,
    color: COLOR.TEXT_DEFAULT,
  },
  button: {
    height: 35,
    alignSelf: 'center',
    marginVertical: 8,
  },
  buttonTitle: {
    color: COLOR.TEXT_TITLE,
  },
  name: {
    color: COLOR.TEXT_DEFAULT,
    fontSize: 18,
  },
});

interface Props {
  name: string;
  avatarUrl: string | undefined;
  rating: number;
  reviewCount: number;
  buttonTitle: string;
  handleClick: () => void;
}

const ProfileBasicInfoGroup = function ProfileBasicInfoGroup(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <ListItem containerStyle={styles.container}>
        {props.avatarUrl !== '' ? (
          <Avatar source={{ uri: props.avatarUrl }} rounded size="medium" />
        ) : (
          <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
        )}
        <ListItem.Content>
          <ListItem.Title style={styles.name}>{props.name}</ListItem.Title>
          <View style={styles.reviewContainer}>
            <Rating fractions={1} readonly={true} startingValue={props.rating} imageSize={16} />
            <Text style={styles.reviewRate}>{props.rating}</Text>
            <Text>（{props.reviewCount}件）</Text>
          </View>
        </ListItem.Content>
        <MaterialIcons name="keyboard-arrow-right" size={24} color={COLOR.GREY} />
      </ListItem>
      <Button
        title={props.buttonTitle}
        buttonStyle={styles.button}
        titleStyle={styles.buttonTitle}
        onPress={props.handleClick}
      />
    </View>
  );
};

export default ProfileBasicInfoGroup;
