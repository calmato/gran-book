import 'jest';

import React from 'react';
import * as ShallowRenderer from 'react-test-renderer/shallow';

import ClosIcon from '@/components/atoms/CloseIcon';
import { colors } from 'react-native-elements';
import CloseIcon from '@/components/atoms/CloseIcon';

describe('<CloseIcon />', () => {

  let render: ShallowRenderer.ShallowRenderer;

  beforeEach(() => {
    render = ShallowRenderer.createRenderer();
  });
  
  it('has default props', () => {
    render.render(<CloseIcon />);
    const got = render.getRenderOutput();

    expect(got).toBeTruthy();
    expect(got.props.name).toEqual('md-close');
    expect(got.props.size).toEqual(12);
    expect(got.props.color).toEqual(undefined);
  });

  it('has setting props', () => {
    const size = 24;
    const color = colors.primary;
    render.render(<ClosIcon size={size} color={color} />);
    const got = render.getRenderOutput();

    expect(got).toBeTruthy();
    expect(got.props.size).toEqual(size);
    expect(got.props.color).toEqual(color);
  });
});
