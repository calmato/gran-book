import 'jest';
import {searchBookByTitle} from '~/lib/rakuten-books';

describe('search book', () => {
  it('[Normal](real-api) return 200 response when send valid request', async () => {
    const title = 'ちはやふる';

    await searchBookByTitle(title)
      .then((got) => {
        expect(got).toBeTruthy;
        expect(got.status).toEqual(200);
      });
  });

  it('[Normal](real-api) send valid request with page', async () => {
    const title = 'ちはやふる';

    const got = await searchBookByTitle(title);
    expect(got).toBeTruthy;
    expect(got.data.page).toBe(1);
    expect(got.data.Items.length).toBe(30);

    const got2 = await searchBookByTitle(title, got.data.page + 1);
    expect(got2).toBeTruthy;
    expect(got2.data.page).toBe(2);
    expect(got2.data.Items.length).toBe(30);
  });
});
