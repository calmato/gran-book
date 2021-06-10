import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import { TextStyle, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import GoogleButton from '~/components/molecules/GoogleButton';
import google from '~~/assets/g-logo.png';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<GoogleButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<GoogleButton />);

    const buttonStyle: ViewStyle = {
      backgroundColor: COLOR.GOOGLE,
      ...SOCIAL_BUTTON,
    };
    const fontStyle: TextStyle = {
      color: 'rgba(0,0,0,0.54)',
    };
    const iconStyle: ViewStyle = {
      marginRight: 10,
      width: 24,
      height: 24,
    };

    const button = wrapper.find(Button).get(0);

    expect(button.props.title).toEqual('Googleでサインイン');
    expect(button.props.buttonStyle).toEqual(buttonStyle);
    expect(button.props.titleStyle).toEqual(fontStyle);

    const icon = button.props.icon;

    expect(icon.props.source).toEqual(google);
    expect(icon.props.style).toEqual(iconStyle);
  });
});
