import { AxiosResponse } from "axios";
import { internal } from "~/lib/axios";
import { getAuthHeader } from "~/lib/axios/util";
import {
  BookshelfListV1Response,
  BookshelfV1Response,
} from "~/types/api/bookshelf_apiv1_response_pb";
import { BookReviewListV1Response } from "~/types/api/review_apiv1_response_pb";
import { ImpressionForm } from "~/types/forms";
import { ISearchResultItem } from "~/types/response/external/rakuten-books";

/**
 * バックエンドAPIにリクエストを送りユーザーが登録している書籍を全件取得する非同期関数
 * @param userId
 * @param token
 * @returns
 */
export async function getAllBookByUserId(userId: string, token: string) {
  try {
    const { data }: AxiosResponse<BookshelfListV1Response.AsObject> =
      await internal.get(`/v1/users/${userId}/books`, getAuthHeader(token));
    return data;
  } catch (e) {
    return Promise.reject(e);
  }
}

/**
 * バックエンドAPIにアクセスして書籍をユーザーの本棚に登録する非同期関数
 * @param payload
 * @param token 認証トークン
 * @returns
 */
export async function registerOwnBook(
  payload: {
    userId: string;
    bookId: number;
    status: "reading" | "read" | "stack" | "release" | "want";
    impressionForm?: ImpressionForm;
  },
  token: string
) {
  const { userId, bookId, status, impressionForm } = payload;
  try {
    const { data }: AxiosResponse<BookshelfV1Response.AsObject> =
      await internal.post(
        `v1/users/${userId}/books/${bookId}/${status}`,
        impressionForm,
        getAuthHeader(token)
      );

    return data;
  } catch (e) {
    console.log("[error]", e);
  }
}

/**
 * バックエンドAPIにアクセスし書籍をindexに登録する非同期関数
 * @param payload
 * @param token 認証トークン
 * @returns
 */
export async function addBook(
  payload: { book: Partial<ISearchResultItem> },
  token: string
) {
  const { book } = payload;

  const res: AxiosResponse<BookshelfV1Response.AsObject> = await internal.post(
    "/v1/books",
    book,
    getAuthHeader(token)
  );
  return res.data;
}

/**
 * バックエンドAPIにアクセスしISBNコードから登録済みの書籍を取得する非同期関数
 * @param payload
 * @param token
 * @returns
 */
export async function getBookByISBN(payload: { isbn: string }, token: string) {
  const { isbn } = payload;
  const res: AxiosResponse<BookshelfV1Response.AsObject> = await internal.get(
    `/v1/books/${isbn}?key=isbn`,
    getAuthHeader(token)
  );
  return res.data;
}

/**
 * バックエンドAPIにアクセスし書籍IDから関連する感想を全て取得する非同期関数
 * @param payload
 * @param token
 * @returns
 */
export async function getAllImpressionByBookId(
  payload: { bookId: number },
  token: string
) {
  const { bookId } = payload;
  const res: AxiosResponse<BookReviewListV1Response.AsObject> =
    await internal.get(`/v1/books/${bookId}/reviews`, getAuthHeader(token));
  return res.data;
}
