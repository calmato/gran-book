import { Ionicons, MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { colors, Input } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';

interface Props{
  placeholder: string,
}

const color = colors.grey0;

const PasswordInput = function PasswordInput(props: Props): ReactElement {

  const [hidden, setValue] = React.useState(true);

  return (
    <Input
      leftIcon={
        <MaterialIcons name="lock" size={24} color={color} />
      }
      secureTextEntry={hidden}
      placeholder={props.placeholder}
      rightIcon={
        <TouchableOpacity onPress={() => setValue(!hidden)}>
          <Ionicons name={ hidden ? 'md-eye' : 'md-eye-off'} size={24} color={color} />
        </TouchableOpacity>
      }
    />
  );
};

export default PasswordInput;
