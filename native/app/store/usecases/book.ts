import { AxiosResponse } from 'axios';
import { Dispatch } from 'redux';
import { setBooks } from '../modules/book';
import { internal } from '~/lib/axios';
import { AppState } from '~/store/modules';
import {
  BookshelfListV1Response,
  BookshelfV1Response,
} from '~/types/api/bookshelf_apiv1_response_pb';
import { BookReviewListV1Response } from '~/types/api/review_apiv1_response_pb';
import { ImpressionForm } from '~/types/forms';
import {
  IErrorResponse,
  IReviewListResponse,
  ISearchResultItem,
} from '~/types/response/external/rakuten-books';

/**
 * バックエンドAPIにアクセスし書籍登録を行う関数
 * @param  {Partial<ISearchResultItem>} 登録する書籍情報(楽天BooksAPIのレスポンス形式)
 * @return {Promise<AxiosResponse<any>|AxiosResponse<IErrorResponse> } 成功時:登録した書籍情報, 失敗時:HTTPエラーオブジェクト
 */
export async function addBookAsync(
  book: Partial<ISearchResultItem>,
): Promise<AxiosResponse<any> | AxiosResponse<IErrorResponse>> {
  return internal
    .post('/v1/books', book)
    .then((res: AxiosResponse<any>) => {
      console.log('[success]', res.data);
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}

/**
 * バックエンドAPIにアクセスし指定したISBNの書籍の情報を返す関数
 * @param isbn 取得したい書籍のISBN
 * @returns
 */
export async function getBookByISBNAsync(isbn: string) {
  try {
    const res: AxiosResponse<BookshelfV1Response.AsObject> = await internal.get(
      `/v1/books/${isbn}?key=isbn`,
    );
    return res.data;
  } catch (err) {
    return Promise.reject(err);
  }
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
    const books = res.booksList;
    dispatch(setBooks({ books }));
  };
}

// TODO: 例外処理
/**
 * ユーザーが自分の本棚に登録する関数
 * @param status ユーザーが登録したい書籍の状態
 * @param bookId バックエンドAPIに登録されている書籍のID
 * @returns Promise<void>
 */
export function registerOwnBookAsync(status: string, bookId: number) {
  return async (_dispatch: Dispatch, getState: () => AppState) => {
    const user = getState().auth;
    let res: void;
    switch (status) {
      case 'read':
        res = await registerReadBookAsync(user.id, bookId);
        break;
      case 'reading':
        res = await registerReadingBookAsync(user.id, bookId);
        break;
      case 'stack':
        res = await registerStackBookAsync(user.id, bookId);
        break;
      case 'release':
        res = await registerReleaseBookAsync(user.id, bookId);
        break;
      case 'want':
        res = await registerWantBookAsync(user.id, bookId);
        break;
      default:
        break;
    }
  };
}

/**
 * バックエンドAPIにアクセスし書籍の感想を登録する関数
 * @param bookId 感想を登録する書籍のID
 * @param impressionForm 書籍の感想
 * TODO エラーハンドリング
 */
export function registerReadBookImpressionAsync(bookId: number, impressionForm: ImpressionForm) {
  return async (_dispatch: Dispatch, getState: () => AppState) => {
    const user = getState().auth;
    registerReadBookAsync(user.id, bookId, impressionForm);
  };
}

/**
 * バックエンドAPIにアクセスし指定した書籍の感想を取得する関数
 * @param bookId 感想を取得する書籍のID
 * @returns
 */
export async function getAllImpressionByBookIdAsync(bookId: number) {
  return internal
    .get(`/v1/books/${bookId}/reviews`)
    .then((res: AxiosResponse<BookReviewListV1Response.AsObject>) => {
      return res.data;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      console.log(err);
      return Promise.reject(err);
    });
}

/**
 * バックエンドAPIにアクセスし書籍を全件取得する関数
 * @param  {Partial<ISearchResultItem>} 登録する書籍情報(楽天BooksAPIのレスポンス形式)
 * @return {Promise<AxiosResponse<any>|AxiosResponse<IErrorResponse> } 成功時:登録した書籍情報, 失敗時:HTTPエラーオブジェクト
 */
async function getAllBookByUserId(userId: string): Promise<BookshelfListV1Response.AsObject> {
  return internal
    .get(`/v1/users/${userId}/books`)
    .then((res: AxiosResponse<BookshelfListV1Response.AsObject>) => {
      return res.data;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}

async function registerReadBookAsync(
  userId: string,
  bookId: number,
  impressionForm?: ImpressionForm,
) {
  return internal
    .post(`v1/users/${userId}/books/${bookId}/read`, impressionForm)
    .then((res: AxiosResponse<BookshelfV1Response.AsObject>) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

async function registerReadingBookAsync(userId: string, bookId: number) {
  return internal
    .post(`v1/users/${userId}/books/${bookId}/reading`)
    .then((res: AxiosResponse<BookshelfV1Response.AsObject>) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

async function registerStackBookAsync(userId: string, bookId: number) {
  return internal
    .post(`v1/users/${userId}/books/${bookId}/stack`)
    .then((res: AxiosResponse<BookshelfV1Response.AsObject>) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

async function registerWantBookAsync(userId: string, bookId: number) {
  return internal
    .post(`v1/users/${userId}/books/${bookId}/want`)
    .then((res: AxiosResponse<BookshelfV1Response.AsObject>) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

async function registerReleaseBookAsync(userId: string, bookId: number) {
  return internal
    .post(`v1/users/${userId}/books/${bookId}/release`)
    .then((res: AxiosResponse<BookshelfV1Response.AsObject>) => {
      console.log('[success]', res.data);
    })
    .catch((err) => {
      console.log('[error]', err);
    });
}

export async function getOwnReviews(userId: string) {
  return internal
    .get(`/v1/users/${userId}/reviews`)
    .then((res: AxiosResponse<IReviewListResponse>) => {
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}
