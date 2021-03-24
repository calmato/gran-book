import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Text, ViewStyle } from 'react-native';
import { colors } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  icon: {},
  text: {},
});

interface Props {
  styles?: ViewStyle;
  title: string;
  checked: boolean;
  onPress: () => void | undefined;
}

const CheckBox = function CheckBox(props: Props): ReactElement {
  return (
    <View style={{ ...styles.container, ...props.styles }}>
      <TouchableOpacity onPress={props.onPress}>
        <Ionicons
          name="md-checkbox"
          size={24}
          style={styles.icon}
          color={props.checked ? colors.primary : colors.grey0}
        />
      </TouchableOpacity>
      <Text style={styles.text}>{props.title}</Text>
    </View>
  );
};

export default CheckBox;
