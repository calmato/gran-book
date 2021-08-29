import { BookshelfV1Response } from '~/types/api/bookshelf_apiv1_response_pb';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';

/**
 * 楽天Books APIの書籍のレスポンスをバックエンドAPIから取得するレスポンス形式に変換する関数
 * @param i 楽天Books APIの書籍のレスポンス
 * @returns IBook バックエンドAPIから取得するレスポンス形式
 */
export function convertToIBook(i: ISearchResultItem): BookshelfV1Response.AsObject {
  return {
    id: 0,
    isbn: i.isbn,
    author: i.author,
    authorKana: i.authorKana,
    title: i.title,
    titleKana: i.titleKana,
    description: i.itemCaption,
    thumbnailUrl: i.largeImageUrl,
    rakutenUrl: i.itemUrl,
    publishedOn: i.salesDate,
    publisher: i.publisherName,
    size: i.size,
    createdAt: '',
    updatedAt: '',
  };
}
