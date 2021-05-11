import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { IErrorResponse, ISearchResultItem } from '~/types/response/external/rakuten-books';


/**
 * バックエンドAPIにアクセスし書籍登録を行う関数
 * @param  {Partial<ISearchResultItem>} 登録する書籍情報(楽天BooksAPIのレスポンス形式)
 * @return {Promise<AxiosResponse<any>|AxiosResponse<IErrorResponse> } 成功時:登録した書籍情報, 失敗時:HTTPエラーオブジェクト
 */
export async function addBookAsync(book: Partial<ISearchResultItem>) {
  return internal.post('/v1/books', book)
    .then((res: AxiosResponse<any>) => {
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}
