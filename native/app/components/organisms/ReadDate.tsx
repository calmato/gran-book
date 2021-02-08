import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, Text, Switch, Platform, Button } from 'react-native';
import { Divider } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';
import DefaultDataPicker from '~/components/molecules/DefaultDataPicker';
import dayjs from 'dayjs';

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
  },
});

interface Props{
  date: Date,
  handleSetDate: (date: Date) => void,
}

const ReadDate = function ReadDate(props: Props): ReactElement {
  const [isEnabled, setIsEnabled] = useState(false);
  const toggleSwitch = () => setIsEnabled(previousState => !previousState);
  
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
          !isEnabled &&        
           <Button onPress={showDatepicker}
             title={`${dayjs(props.date).format('YYYY/MM/DD')}`}
           />
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
          value={isEnabled}
        />
      </View>
      {!isEnabled && show && (
        <DefaultDataPicker
          date={props.date}
          onChange={onChange}
        />
      )}
    </View>
  );
};

export default ReadDate;
