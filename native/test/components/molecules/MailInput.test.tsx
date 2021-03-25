import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import MailInput from '~/components/molecules/MailInput';
import { Input } from 'react-native-elements';

configure({ adapter: new Adapter() });

describe('<MailInput />', () => {
  it('has default props', () => {
    const value = 'test';
    const errorMessage = undefined;
    const wrapper = shallow(
      <MailInput
        hasError={false}
        value={value}
        onChangeText={() => {
          console.log('test');
        }}
      />,
    );

    const input = wrapper.find(Input).get(0);

    expect(input.props.hasError).toBeFalsy();
    expect(input.props.value).toEqual(value);
    expect(input.props.errorMessage).toEqual(errorMessage);
  });

  it('errorMessage is メールアドレスを入力してください．when hasError is true', () => {
    const value = 'test';
    const errorMessage = 'メールアドレスを入力してください．';
    const wrapper = shallow(
      <MailInput
        hasError={true}
        value={value}
        onChangeText={() => {
          console.log('test');
        }}
      />,
    );

    const input = wrapper.find(Input).get(0);

    expect(input.props.hasError).toBeFalsy();
    expect(input.props.value).toEqual(value);
    expect(input.props.errorMessage).toEqual(errorMessage);
  });
});
