import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import TwitterButton from '~/components/molecules/TwitterButton';
import { ViewStyle } from 'react-native';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';
import { Button } from 'react-native-elements';

configure({ adapter: new Adapter() });

describe('<TwitterButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<TwitterButton />);

    const buttonStyle: ViewStyle = {
      backgroundColor: COLOR.TWITTER,
      ...SOCIAL_BUTTON,
    };
    const iconStyle: ViewStyle = {
      marginRight: 10,
    };

    const button = wrapper.find(Button).get(0);

    expect(button.props.title).toEqual('Twitterでサインイン');
    expect(button.props.buttonStyle).toEqual(buttonStyle);

    const icon = button.props.icon;

    expect(icon.props.name).toEqual('logo-twitter');
    expect(icon.props.size).toEqual(24);
    expect(icon.props.color).toEqual('white');
    expect(icon.props.style).toEqual(iconStyle);
  });
});
