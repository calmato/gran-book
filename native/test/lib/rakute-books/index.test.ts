import 'jest';
import {searchBook} from '~/lib/rakuten-books';

describe('search book', () => {
  it('return 200 response when send valid request', async () => {
    const title = 'ちはやふる';
    const got = await searchBook(title);

    expect(got).toBeTruthy;
    expect(got.status).toEqual(200);
  });
});
