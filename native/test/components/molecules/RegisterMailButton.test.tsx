import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import RegisterMailButton from '~/components/molecules/RegisterMailButton';
import { Button, colors } from 'react-native-elements';
import { TextStyle, ViewStyle } from 'react-native';
import { SOCIAL_BUTTON } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<RegisterMailButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<RegisterMailButton
      onPress={() => console.log('test')}
    />);

    const color = colors.grey0;

    const buttonStyle: ViewStyle = {
      borderColor: color,
      backgroundColor: '#00000000',
      ...SOCIAL_BUTTON
    };
    const iconStyle: ViewStyle = {
      marginRight: 10
    };
    const fontStyle: TextStyle = {
      color: color
    };

    const button = wrapper.find(Button).get(0);

    expect(button.props.title).toEqual('メールアドレスで新規登録');
    expect(button.props.buttonStyle).toEqual(buttonStyle);
    expect(button.props.titleStyle).toEqual(fontStyle);
    expect(button.props.type).toEqual('outline');

    const icon = button.props.icon;
    
    expect(icon.props.name).toEqual('md-mail');
    expect(icon.props.size).toEqual(24);
    expect(icon.props.color).toEqual(color);
    expect(icon.props.style).toEqual(iconStyle);
  });
});
