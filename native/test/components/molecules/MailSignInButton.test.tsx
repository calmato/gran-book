import 'jest';
import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';

import { ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import MailSignInButton from '~/components/molecules/MailSignInButton';
import { SOCIAL_BUTTON } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<MailSignInButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<MailSignInButton onPress={() => console.log('test')} />);

    const buttonStyle: ViewStyle = {
      ...SOCIAL_BUTTON,
    };
    const iconStyle: ViewStyle = {
      marginRight: 10,
    };

    const button = wrapper.find(Button).get(0);

    expect(button.props.title).toEqual('メールアドレスでサインイン');
    expect(button.props.buttonStyle).toEqual(buttonStyle);

    const icon = button.props.icon;

    expect(icon.props.name).toEqual('md-mail');
    expect(icon.props.size).toEqual(24);
    expect(icon.props.color).toEqual('white');
    expect(icon.props.style).toEqual(iconStyle);
  });
});
