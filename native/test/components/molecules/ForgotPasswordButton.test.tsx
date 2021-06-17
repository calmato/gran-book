import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import ForgotPasswordButton from '~/components/molecules/ForgotPasswordButton';

configure({ adapter: new Adapter() });

describe('<ForgotPasswordButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<ForgotPasswordButton onPress={() => console.log('test')} />);
  });
});
