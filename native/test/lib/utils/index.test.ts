import 'jest';
import { fullWidth2halfWidth } from '~/lib/util';

describe('fullWidth2halfWidth', () => {
  it('return true when valid email', () => {
    const str = 'ｒｅａｃｔ';
    const got = fullWidth2halfWidth(str);
    console.log(str);
    console.log(got);
  });
});
