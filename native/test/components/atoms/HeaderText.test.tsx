import 'jest';
import HeaderText from '@/components/atoms/HeaderText';
import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';


import { View, Text, TextStyle, ViewStyle } from 'react-native';
import { COLOR } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<HeaderText />', () => {
  it('has default props', () => {
    const title = 'test';
    const textStyle: TextStyle = {
      fontSize: 16,
      fontWeight: 'bold',
      color: COLOR.TEXT_TITLE,
    };
    const viewStyle: ViewStyle = {
      justifyContent: 'center',
    };

    const wrapper = shallow(<HeaderText title={title} />);

    expect(wrapper.find(View).get(0).props.style).toEqual(viewStyle);
    expect(wrapper.find(Text).get(0).props.style).toEqual(textStyle);
    expect(wrapper.find(Text).get(0).props.children).toEqual(title);
  });
});
