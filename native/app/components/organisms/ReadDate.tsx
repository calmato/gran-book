import { MaterialIcons } from '@expo/vector-icons';
import dayjs from 'dayjs';
import React, { ReactElement, useCallback, useState } from 'react';
import { StyleSheet, View, Text, Switch, Platform } from 'react-native';
import { ListItem } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';
import DefaultDataPicker from '~/components/molecules/DefaultDatePicker';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  containerStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    marginTop: 20,
  },
  textStyle: {
    fontSize: FONT_SIZE.LISTITEM_TITLE,
    color: COLOR.TEXT_DEFAULT,
  },
});

interface Props {
  date: Date;
  handleSetDate: (date: Date) => void;
  isDateUnknown: boolean;
  handleIsDateUnknown: (isDateUnknown: boolean) => void;
}

const ReadDate = function ReadDate(props: Props): ReactElement {
  const [show, setShow] = useState(false);

  const toggleSwitch = () => props.handleIsDateUnknown(!props.isDateUnknown);

  const handleChangeDate = (_event, selectedDate) => {
    const currentDate = selectedDate;
    setShow(Platform.OS === 'ios');
    props.handleSetDate(currentDate);
  };

  const showDatePicker = useCallback(() => {
    setShow(true);
  }, []);

  return (
    <View style={styles.containerStyle}>
      <ListItem bottomDivider>
        <ListItem.Content>
          <Text style={styles.textStyle}>読んだ日</Text>
        </ListItem.Content>
        {!props.isDateUnknown && (
          <TouchableOpacity
            style={{ flexDirection: 'row', alignItems: 'center' }}
            onPress={showDatePicker}>
            <Text style={[styles.textStyle]}>{`${dayjs(props.date).format('YYYY/MM/DD')}`}</Text>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={COLOR.GREY} />
          </TouchableOpacity>
        )}
      </ListItem>
      <ListItem>
        <ListItem.Content>
          <Text style={styles.textStyle}>不明</Text>
        </ListItem.Content>
        <Switch
          trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
          thumbColor={COLOR.BACKGROUND_WHITE}
          ios_backgroundColor={COLOR.TEXT_GRAY}
          onValueChange={toggleSwitch}
          value={props.isDateUnknown}
        />
      </ListItem>
      {!props.isDateUnknown && show && (
        <DefaultDataPicker date={props.date} onChange={handleChangeDate} />
      )}
    </View>
  );
};

export default ReadDate;
