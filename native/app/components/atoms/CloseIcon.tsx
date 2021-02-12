import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';

interface Props {
  size?: number,
  color?: string,
}

const CloseIcon = function CloseIcon(props: Props): ReactElement {
  return (
    <Ionicons
      name="md-close"
      size={props.size}
      color={props.color}
    />
  );
};

export default CloseIcon;
