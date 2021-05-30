import { IBook } from '../response';

/**
 * コンポーネントで表示に使用する書籍のデータ構造
 */
export interface ViewBooks {
  reading: IBook[];
  read: IBook[];
  stack: IBook[];
  release: IBook[];
  want: IBook[];
}
