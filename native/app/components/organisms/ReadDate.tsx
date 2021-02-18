import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, Text, Switch, Platform } from 'react-native';
import { Divider } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';
import DefaultDataPicker from '~/components/molecules/DefaultDatePicker';
import dayjs from 'dayjs';
import { MaterialIcons } from '@expo/vector-icons';
import { TouchableOpacity } from 'react-native-gesture-handler';

const styles = StyleSheet.create({
  containerStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    marginTop: 20,
  },
  childStyle: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignSelf: 'stretch',
    alignItems: 'center',
    height:60,
    paddingStart: 20,
    paddingEnd: 20,
    paddingTop: 10,
    paddingBottom: 10,
  },
  textStyle: {
    fontSize: 16,
    color: COLOR.GREY,
  },
});

interface Props{
  date: Date,
  handleSetDate: (date: Date) => void,
  isDateUnknown: boolean,
  handleIsDateUnknown: (isDateUnknown: boolean) => void,
}

const ReadDate = function ReadDate(props: Props): ReactElement {
  const toggleSwitch = () => props.handleIsDateUnknown(!props.isDateUnknown);
  
  const [show, setShow] = useState(false);
  const onChange = (_event, selectedDate) => {
    const currentDate = selectedDate;
    setShow(Platform.OS === 'ios');
    props.handleSetDate(currentDate);
  };
  const showDatepicker = () => {
    setShow(true);
  };

  return(
    <View style={styles.containerStyle}>
      <View style={styles.childStyle}>
        <Text style={styles.textStyle}>読んだ日</Text>
        {
          !props.isDateUnknown &&      
          <TouchableOpacity style={{flexDirection:'row', alignItems:'center'}} onPress={showDatepicker}>
            <Text style={[styles.textStyle]}>{`${dayjs(props.date).format('YYYY/MM/DD')}`}</Text> 
            <MaterialIcons name="keyboard-arrow-right" size={24} color={COLOR.GREY} />
          </TouchableOpacity>
        }
      </View>
      <Divider/>
      <View style={styles.childStyle}>
        <Text style={styles.textStyle}>不明</Text>
        <Switch
          trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
          thumbColor={COLOR.BACKGROUND_WHITE}
          ios_backgroundColor= {COLOR.TEXT_GRAY}
          onValueChange={toggleSwitch}
          value={props.isDateUnknown}
        />
      </View>
      {!props.isDateUnknown && show && (
        <DefaultDataPicker
          date={props.date}
          onChange={onChange}
        />
      )}
    </View>
  );
};

export default ReadDate;
