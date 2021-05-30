import React, { ReactElement } from 'react';
import { StyleSheet, Dimensions, View, Text } from 'react-native';
import { MaterialCommunityIcons, FontAwesome5 } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { TouchableOpacity } from 'react-native-gesture-handler';

const { width } = Dimensions.get('window');
const iconSize = 25;

const styles = StyleSheet.create({
  parentStyle: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: COLOR.BACKGROUND_WHITE,
    marginBottom: 10,
    width: width,
  },
  childStyle: {
    width: width / 5,
    height: 60,
    alignItems: 'center',
    justifyContent: 'center',
    borderLeftWidth: 1,
    borderColor: COLOR.LIGHT_GREY,
  },
  textStyle: {
    color: COLOR.TEXT_DEFAULT,
    fontSize: 10,
  },
});

interface Props {
  onPress: (status: number) => void;
  status: number;
}

const IconComponents = {
  MaterialIcons: function MaterialIcons(
    iconName: 'book-plus' | 'book-open-page-variant' | 'bookshelf' | 'bookmark-plus',
    color: string,
  ) {
    return (
      <MaterialCommunityIcons
        name={iconName}
        size={iconSize}
        color={color}
        style={{ marginBottom: 6 }}
      />
    );
  },
  FontAwesome: function FontAwesome(iconName: string, color: string) {
    return (
      <FontAwesome5 name={iconName} size={iconSize} color={color} style={{ marginBottom: 6 }} />
    );
  },
};

const items = [
  {
    name: '読んでる本',
    value: 0,
    icon: (color: string) => IconComponents.MaterialIcons('book-plus', color),
  },
  {
    name: '読んだ本',
    value: 1,
    icon: (color: string) => IconComponents.MaterialIcons('book-open-page-variant', color),
  },
  {
    name: '積読本',
    value: 2,
    icon: (color: string) => IconComponents.MaterialIcons('bookshelf', color),
  },
  {
    name: '手放したい本',
    value: 3,
    icon: (color: string) => IconComponents.FontAwesome('people-carry', color),
  },
  {
    name: '欲しい本',
    value: 4,
    icon: (color: string) => IconComponents.MaterialIcons('bookmark-plus', color),
  },
];

const ButtonGroupBookFooter = function ButtonGroupBookFooter(props: Props): ReactElement {
  const status = props.status;

  const renderItem = items.map((item) => {
    const isActive: boolean = item.value === status;
    return (
      <TouchableOpacity
        key={item.name}
        style={styles.childStyle}
        onPress={() => props.onPress(item.value)}>
        {item.icon(isActive ? COLOR.PRIMARY : COLOR.TEXT_GRAY)}
        <Text
          style={
            isActive
              ? { ...styles.textStyle, color: COLOR.TEXT_TITLE, fontWeight: 'bold' }
              : styles.textStyle
          }>
          {item.name}
        </Text>
      </TouchableOpacity>
    );
  });

  return <View style={styles.parentStyle}>{renderItem}</View>;
};

export default ButtonGroupBookFooter;
