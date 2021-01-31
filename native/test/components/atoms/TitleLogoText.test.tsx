import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import { View, ViewStyle } from 'react-native';
import { Text } from 'react-native-elements';

configure({ adapter: new Adapter() });

describe('<TitleLogoText />', () => {

  it('has default props', () => {
    const text = 'test';
    const wrapper = shallow(<TitleLogoText text={text} />);

    expect(wrapper.find(View).get(0).props.style).toEqual(undefined);
    expect(wrapper.find(Text).get(0).props.h1).toBeTruthy();
    expect(wrapper.find(Text).get(0).props.children).toEqual(text);
  });

  it('has setting props', () => {
    const text = 'test';

    const viewStyle: ViewStyle = {
      padding: 10,
      margin: 10,
    };

    const wrapper = shallow(<TitleLogoText text={text} style={viewStyle} />);

    expect(wrapper.find(View).get(0).props.style).toEqual(viewStyle);
    expect(wrapper.find(Text).get(0).props.h1).toBeTruthy();
    expect(wrapper.find(Text).get(0).props.children).toEqual(text);
  });
});
