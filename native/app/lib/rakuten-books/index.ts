import axios, { AxiosResponse } from 'axios';
import { IErrorResponse, ISearchResponse } from '~/types/response/external/rakuten-books';

const baseUrl = 'https://app.rakuten.co.jp/services/api/BooksBook/Search';
const version = '20170404';
const format = 'json';
const formatVersion = 2;
const hits = 30;

const applicationId = process.env.RAKUTEN_BOOKS_APPLICATION_ID;

export async function searchBookByTitle(title: string, page = 1) {
  const url = `${baseUrl}/${version}?format=${format}&title=${encodeURI(title)}&formatVersion=${formatVersion}&applicationId=${applicationId}&page=${page}&hits=${hits}`;

  return axios.get(url)
    .then((res: AxiosResponse<ISearchResponse>) => {
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}
