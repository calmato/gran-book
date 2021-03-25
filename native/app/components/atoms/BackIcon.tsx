import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';

interface Props {
  size?: number;
  color?: string;
}

const BackIcon = function BackIcon(props: Props): ReactElement {
  return <Ionicons name="ios-arrow-back" size={props.size} color={props.color} />;
};

export default BackIcon;
