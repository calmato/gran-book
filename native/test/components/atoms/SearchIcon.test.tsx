import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import SearchIcon from '~/components/atoms/SearchIcon';
import { Ionicons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<SearchIcon />', () => {
  it('has default props', () => {
    const wrapper = shallow(<SearchIcon />);
    const icon = wrapper.find(Ionicons).get(0);
    expect(icon.props.name).toEqual('md-search');
    expect(icon.props.size).toEqual(18);
    expect(icon.props.color).toEqual(COLOR.GREY);
  });

  it('has settings props', () => {
    const size = 23;
    const color = COLOR.PRIMARY;

    const wrapper = shallow(<SearchIcon size={size} color={color} />);
    const icon = wrapper.find(Ionicons).get(0);
    expect(icon.props.name).toEqual('md-search');
    expect(icon.props.size).toEqual(size);
    expect(icon.props.color).toEqual(color);
  });
});
