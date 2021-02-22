import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Rating, Text, Avatar } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  avator: {
    flex: 1,
  },
  group: {
    flex: 5,
  },
  reviewContainer: {
    flexDirection: 'row',
    alignItems: 'center',
  }
});

interface Props {
  name: string
  avatar_url: string
  rating: number
  numberOfReviews: number
}

const ProfileBasicInfoGroup = function ProfileBasicInfoGroup(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <Avatar source={{uri: props.avatar_url}} rounded avatarStyle={styles.avator}/>
      <View style={styles.group}>
        <Text>{props.name}</Text>
        <View style={styles.reviewContainer}>
          <Rating 
            fractions={1}
            readonly={true}
            startingValue={props.rating}
          />
          <Text>{props.rating}</Text>
          <Text>{`(${props.numberOfReviews}件)`}</Text>
          <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
        </View>
        <Button title={'プロフィールを編集'}/>
      </View>
    </View>
  );
};

export default ProfileBasicInfoGroup;
