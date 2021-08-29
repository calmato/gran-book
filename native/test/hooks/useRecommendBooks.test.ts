import 'jest';

import { act, renderHook } from '@testing-library/react-hooks/native/pure';
import MockAdapter from 'axios-mock-adapter';

import responseSample from './rakuten-books-response.sample.json';
import { useRecommendBooks } from '~/hooks/useRecommendBooks';

import externalInstance from '~/lib/axios/external';

const mockAxios = new MockAdapter(externalInstance);

const baseUrl = 'https://app.rakuten.co.jp/services/api/BooksBook/Search';
const version = '20170404';
const format = 'json';
const formatVersion = 2;
const hits = 30;
const applicationId = process.env.RAKUTEN_BOOKS_APPLICATION_ID;

const url = `${baseUrl}/${version}?format=${format}&formatVersion=${formatVersion}&applicationId=${applicationId}&sort=sales&hits=${hits}`;

describe('useRecommendBooks', () => {
  beforeEach(() => {
    mockAxios.onGet(url).reply(200, responseSample);
  });

  test('can fetch and set recommendBooks', async () => {
    const { result } = renderHook(() => useRecommendBooks());

    await act(async () => {
      await result.current.fetch();
      expect(result.current.recommendBooks.length).toBe(30);
    });
  });

  test('return error is true when fetch is failed', async () => {
    const { result } = renderHook(() => useRecommendBooks());

    mockAxios.onGet(url).reply(400);

    await act(async () => {
      await result.current.fetch();
      expect(result.current.error).toBeTruthy();
    });
  });
});
