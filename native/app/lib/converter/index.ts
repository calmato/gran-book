import { IBook } from '~/types/response';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';

/**
 * 楽天Books APIの書籍のレスポンスをバックエンドAPIから取得するレスポンス形式に変換する関数
 * @param i 楽天Books APIの書籍のレスポンス
 * @returns IBook バックエンドAPIから取得するレスポンス形式
 */
export function convertToIBook(i: ISearchResultItem): IBook {
  const book: IBook = {
    id: 0,
    readOn: '',
    status: 0,
    createdAt: '',
    updatedAt: '',
    detail: {
      id: 0,
      isbn: i.isbn,
      author: i.author,
      authorKana: i.authorKana,
      title: i.title,
      description: i.contents,
      titleKana: i.titleKana,
      thumbnailUrl: i.largeImageUrl,
      rakutenGenreId: i.booksGenreId,
      rakutenUrl: i.affiliateUrl,
      createdAt: '',
      updatedAt: '',
      publishedOn: i.salesDate,
      publisher: i.publisherName,
    }
  };
  return book;
}
