import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import BackIcon from '~/components/atoms/BackIcon';
import BackButton from '~/components/molecules/BackButton';

configure({ adapter: new Adapter() });

describe('<BackButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<BackButton />);

    const backIcon = wrapper.find(BackIcon).get(0);

    expect(backIcon.size).toEqual(undefined);
    expect(backIcon.color).toEqual(undefined);
  });
});
