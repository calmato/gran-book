import React, { ReactElement } from 'react';
import { Ionicons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';

interface Props {
  size?: number;
  color?: string;
}

const SearchIcon = function SearchIcon(props: Props): ReactElement {
  return <Ionicons name="md-search" size={props.size} color={props.color} />;
};

SearchIcon.defaultProps = {
  size: 18,
  color: COLOR.GREY,
};

export default SearchIcon;
