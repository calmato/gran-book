import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { getAuthHeader } from '~/lib/axios/util';
import {
  BookshelfListV1Response,
  BookshelfV1Response,
} from '~/types/api/bookshelf_apiv1_response_pb';
import { ImpressionForm } from '~/types/forms';

export async function getAllBookByUserId(userId: string, token: string) {
  try {
    const { data }: AxiosResponse<BookshelfListV1Response.AsObject> = await internal.get(
      `/v1/users/${userId}/books`,
      getAuthHeader(token),
    );
    return data;
  } catch (e) {
    return Promise.reject(e);
  }
}

export async function registerOwnBook(
  userId: string,
  bookId: number,
  status: 'reading' | 'read' | 'stack' | 'release' | 'want',
  token: string,
  impressionForm?: ImpressionForm,
) {
  try {
    const { data }: AxiosResponse<BookshelfV1Response.AsObject> = await internal.post(
      `v1/users/${userId}/books/${bookId}/${status}`,
      impressionForm,
      getAuthHeader(token),
    );

    return data;
  } catch (e) {
    console.log('[error]', e);
  }
}
