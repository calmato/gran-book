import axios, { AxiosResponse } from 'axios';
import { ISearchResponse } from '~/types/response/search';

const url = 'https://www.googleapis.com/books/v1/volumes';

export async function searchBook(param: string) {
  return axios.get(`${url}?q=${param}`)
    .then((res: AxiosResponse<ISearchResponse>) => {
      return res.data;
    })
    .catch((err) => {
      console.log('[debug: search]',err.response);
      return null;
    });
}
