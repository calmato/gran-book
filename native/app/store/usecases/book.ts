import { Dispatch } from 'redux';
import { AppState } from '~/store/modules';
import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { IErrorResponse, ISearchResultItem } from '~/types/response/external/rakuten-books';
import { IBookResponse } from '~/types/response';
import { setBooks } from '../modules/book';
import { BookValues } from '../models/book';


/**
 * バックエンドAPIにアクセスし書籍登録を行う関数
 * @param  {Partial<ISearchResultItem>} 登録する書籍情報(楽天BooksAPIのレスポンス形式)
 * @return {Promise<AxiosResponse<any>|AxiosResponse<IErrorResponse> } 成功時:登録した書籍情報, 失敗時:HTTPエラーオブジェクト
 */
export async function addBookAsync(book: Partial<ISearchResultItem>): Promise<AxiosResponse<any> | AxiosResponse<IErrorResponse>> {
  return internal.post('/v1/books', book)
    .then((res: AxiosResponse<any>) => {
      console.log('[success]', res.data);
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}

export async function registerReadBookAsync(userId: string, bookId: number) {
  return internal.post(`v1/users/${userId}/books/${bookId}/read`)
    .then((res) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}


export async function registerReadingBookAsync(userId: string, bookId: number) {
  return internal.post(`v1/users/${userId}/books/${bookId}/reading`)
    .then((res) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}


export async function registerStackBookAsync(userId: string, bookId: number) {
  return internal.post(`v1/users/${userId}/books/${bookId}/stack`)
    .then((res) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}


export async function registerWantBookAsync(userId: string, bookId: number) {
  return internal.post(`v1/users/${userId}/books/${bookId}/want`)
    .then((res) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

export async function registerReleaseBookAsync(userId: string, bookId: number) {
  return internal.post(`v1/users/${userId}/books/${bookId}/release`)
    .then((res) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

/**
 * バックエンドAPIにアクセスし書籍を全件取得する関数
 * @param  {Partial<ISearchResultItem>} 登録する書籍情報(楽天BooksAPIのレスポンス形式)
 * @return {Promise<AxiosResponse<any>|AxiosResponse<IErrorResponse> } 成功時:登録した書籍情報, 失敗時:HTTPエラーオブジェクト
 */
export async function getAllBookByUserId(userId: string) {
  return internal.get(`/v1/users/${userId}/books`)
    .then((res: AxiosResponse<IBookResponse>) => {
      return res;
    }).catch((err: AxiosResponse<IErrorResponse>) => {
      console.log('Called');
      return Promise.reject(err);
    });
}

// TODO: 例外処理
/**
 * 書籍を全件取得しreduxのstateに保存する関数
 * @returns Promise<void>
 */
export function getAllBookAsync() {
  return async (dispatch: Dispatch, getState: () => AppState) => {
    const user = getState().auth;
    const res = await getAllBookByUserId(user.id);
    const books = res.data.books;
    dispatch(setBooks({books}));
  };
}
