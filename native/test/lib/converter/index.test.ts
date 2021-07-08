import 'jest';
import { convertToIBook } from '~/lib/converter';
import { IBook } from '~/types/response';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';

describe('convert to IBook', () => {
  it('Rakuten Books API response data convert to IBook object', () => {
    const sampleResponse: ISearchResultItem = {
      limitedFlag: 0,
      authorKana: 'スエツグ ユキ',
      author: '末次 由紀',
      subTitle: '',
      seriesNameKana: 'BELOVEKCビーラブコミックス',
      title: 'ちはやふる（46）',
      subTitleKana: '',
      itemCaption: '',
      publisherName: '講談社',
      listPrice: 0,
      isbn: '9784065225608',
      largeImageUrl:
        'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/5608/9784065225608.jpg?_ex=200x200',
      mediumImageUrl:
        'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/5608/9784065225608.jpg?_ex=120x120',
      titleKana: 'チハヤフル46',
      availability: '1',
      postageFlag: 2,
      salesDate: '2021年03月12日',
      contents: '',
      smallImageUrl:
        'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/5608/9784065225608.jpg?_ex=64x64',
      discountPrice: 0,
      itemPrice: 495,
      size: 'コミック',
      booksGenreId: '001001002020',
      affiliateUrl: '',
      seriesName: 'BE　LOVE　KC',
      reviewCount: 22,
      reviewAverage: '4.26',
      discountRate: 0,
      chirayomiUrl: '',
      itemUrl: 'https://books.rakuten.co.jp/rb/16609483/',
    };

    const actual: IBook = convertToIBook(sampleResponse);

    expect(actual.isbn).toBe(sampleResponse.isbn);
    expect(actual.author).toBe(sampleResponse.author);
    expect(actual.authorKana).toBe(sampleResponse.authorKana);
    expect(actual.title).toBe(sampleResponse.title);
    expect(actual.titleKana).toBe(sampleResponse.titleKana);
    expect(actual.description).toBe(sampleResponse.itemCaption);
    expect(actual.rakutenUrl).toBe(sampleResponse.itemUrl);
    expect(actual.thumbnailUrl).toBe(sampleResponse.largeImageUrl);
  });
});
