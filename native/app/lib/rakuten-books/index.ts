import externalInstance from '~/lib//axios/external';
import { AxiosResponse } from 'axios';
import { ISearchResponse } from '~/types/response/external/rakuten-books';

const baseUrl = 'https://app.rakuten.co.jp/services/api/BooksBook/Search';
const version = '20170404';
const format = 'json';
const formatVersion = 2;
const hits = 30;

const applicationId = process.env.RAKUTEN_BOOKS_APPLICATION_ID;

/**
 * 楽天 Book APIを利用して書籍を検索する関数
 * @param title 検索したい書籍のタイトル
 * @param page ページ番号（任意、デフォルト値は1）
 * @returns Promise<AxiosResponse<ISearchResponse>>
 */
export async function searchBookByTitle(title: string, page = 1) {
  const url = `${baseUrl}/${version}?format=${format}&title=${encodeURI(
    title,
  )}&formatVersion=${formatVersion}&applicationId=${applicationId}&page=${page}&hits=${hits}`;

  return externalInstance
    .get(url)
    .then((res: AxiosResponse<ISearchResponse>) => {
      return res;
    })
    .catch((err) => {
      return Promise.reject(err.response);
    });
}
