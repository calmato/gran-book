import 'jest';
import { fullWidth2halfWidth } from '~/lib/util';

describe('fullWidth2halfWidth', () => {
  it('return react when ｒｅａｃｔ', () => {
    const str = 'ｒｅａｃｔ';
    const got = fullWidth2halfWidth(str);
    expect(got).toBe('react');
  });

  it('return 1 when １', () => {
    const str = '１';
    const got = fullWidth2halfWidth(str);
    expect(got).toBe('1');
  });

  it('return ,. when ，．', () => {
    const str = '，．';
    const got = fullWidth2halfWidth(str);
    expect(got).toBe(',.');
  });
});
