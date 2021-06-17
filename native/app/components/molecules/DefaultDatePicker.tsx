import DateTimePicker from '@react-native-community/datetimepicker';
import React, { ReactElement } from 'react';

interface Props {
  date: Date;
  onChange: (event, selectedDate) => void;
}

const DefaultDatePicker = function DefaultDatePicker(props: Props): ReactElement {
  return (
    <DateTimePicker
      testID="dateTimePicker"
      value={props.date}
      mode={'date'}
      display="spinner"
      onChange={(event, selectedDate) => props.onChange(event, selectedDate)}
    />
  );
};

export default DefaultDatePicker;
