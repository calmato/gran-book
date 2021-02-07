import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import FacebookButton from '~/components/molecules/FacebookButton';
import { Button } from 'react-native-elements';
import { ViewStyle } from 'react-native';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<FacebookButton />', () => {
  it('has default props', () => {

    const buttonStyle: ViewStyle = {
      backgroundColor: COLOR.FACEBOOK,
      ...SOCIAL_BUTTON,
    };

    const iconStyle: ViewStyle = {
      marginRight: 10,
    };

    const wrapper = shallow(<FacebookButton />);

    const button = wrapper.find(Button).get(0);
    expect(button.props.title).toEqual('Facebookでサインイン');
    expect(button.props.buttonStyle).toEqual(buttonStyle);
    expect(button.props.containerStyle).toEqual(undefined);

    
    const icon = button.props.icon;
    expect(icon.props.size).toEqual(24);
    expect(icon.props.name).toEqual('logo-facebook');
    expect(icon.props.color).toEqual('white');
    expect(icon.props.style).toEqual(iconStyle);
  });
  
  it('has setting props', () => {
    const containerProps: ViewStyle = {
      margin: 10,
      padding: 10,
    };
    const wrapper = shallow(<FacebookButton style={containerProps} />);

    expect(wrapper.find(Button).get(0).props.containerStyle).toEqual(containerProps);
  });
});
