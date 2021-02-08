import React, { ReactElement } from 'react';
import DateTimePicker from '@react-native-community/datetimepicker';

interface Props{
  date: Date,
  onChange: (event,selectedDate) => void,
}

const DefaultDataPicker = function DefaultDataPicker(props: Props): ReactElement {
  return(
    <DateTimePicker
    testID="dateTimePicker"
    value={props.date}
    mode={'date'}
    display="spinner"
    onChange={(event, selectedDate) => props.onChange(event,selectedDate)}
    />
  );
};

export default DefaultDataPicker;
