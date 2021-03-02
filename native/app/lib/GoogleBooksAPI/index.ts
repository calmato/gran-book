import axios, { AxiosResponse } from 'axios';
import { ISearchResponse } from '~/types/response/search';

const url = 'https://www.googleapis.com/books/v1/volumes';

// TODO: スクロールでのページネーション対応
export async function searchBook(param: string) {
  const index = 0;
  const maxResults = 40;
  return axios.get(`${url}?q=${param}&startIndex=${index}&maxResults=${maxResults}`)
    .then((res: AxiosResponse<ISearchResponse>) => {
      return res.data;
    })
    .catch((err) => {
      console.log('[debug: search]',err.response);
      return null;
    });
}
