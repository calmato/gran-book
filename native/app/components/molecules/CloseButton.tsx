import React, { ReactElement } from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';
import CloseIcon from '~/components/atoms/CloseIcon';

interface Props {
  size?: number;
  color?: string;
  onPress?: () => void | undefined;
}

const CloseButton = function CloseButton(props: Props): ReactElement {
  return (
    <TouchableOpacity onPress={props.onPress}>
      <CloseIcon size={props.size} color={props.color} />
    </TouchableOpacity>
  );
};

export default CloseButton;
