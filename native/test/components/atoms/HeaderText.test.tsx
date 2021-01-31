import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';

import HeaderText from '@/components/atoms/HeaderText';
import { colors } from 'react-native-elements';
import { View, Text, TextStyle, ViewStyle} from 'react-native';

configure({ adapter: new Adapter() });

describe('<HeaderText />', () => {

  it('has default props', () => {
    const title = 'test';
    const textStyle: TextStyle = {
      fontSize: 16,
      fontWeight: 'bold',
      color: colors.white,
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
