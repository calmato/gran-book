import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    paddingBottom: 10,
  },
  number: {
    paddingBottom: 10,
    color: COLOR.TEXT_DEFAULT,
    fontSize: FONT_SIZE.TEXT_INFO,
  },
});

interface Props {
  name: string;
  numberOfThings: number;
}

const ProfileNumberOfThings = function ProfileNumberOfThings(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <Text style={styles.number}>{props.numberOfThings}</Text>
      <Text>{props.name}</Text>
    </View>
  );
};

export default ProfileNumberOfThings;
