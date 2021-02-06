import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';

configure({ adapter: new Adapter() });

describe('<PasswordInput />', () => {
  it('has default props', () => {
    const wrapper = shallow();
  });
});
