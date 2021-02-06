import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';
import PasswordInput from '~/components/molecules/PasswordInput';

configure({ adapter: new Adapter() });

describe('<PasswordInput />', () => {
  it('has default props', () => {
    const placeholder = 'test';
    const value = 'test';
    const wrapper = shallow(<PasswordInput
      placeholder={placeholder}
      value={value}
      onChangeText={() => console.log('test')}
    />);
  });
});
