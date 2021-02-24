import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import ChangeIconGroup from '~/components/organisms/ChangeIconGroup';

const styles = StyleSheet.create({});

interface Props {}

const ProfileEdit = function ProfileEdit(props: Props): ReactElement {
  return (
  <View>
    <ChangeIconGroup/>
  </View>
  );
}

export default ProfileEdit
