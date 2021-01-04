import React, { ReactElement } from 'react';
import { TouchableOpacity } from 'react-native';
import BackIcon from '~/components/atoms/BackIcon';

interface Props {
  size?: number,
  color?: string,
  onPress?: () => void | undefined,
}

const BackButton = function BackButton(props: Props): ReactElement {
  return (
    <TouchableOpacity onPress={props.onPress}>
      <BackIcon
        size={props.size}
        color={props.color}
      />
    </TouchableOpacity>
  
  );
};

export default BackButton;
