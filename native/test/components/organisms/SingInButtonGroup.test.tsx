import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';

configure({ adapter: new Adapter() });

describe('<SingInButtonGroup />', () => {
  it('has default props', () => {
    // TODO
    const wrapper = shallow(
      <SignInButtonGroup handleRegisterWithMail={jest.fn()} handleSignInWithMail={jest.fn()} />,
    );
  });
});
