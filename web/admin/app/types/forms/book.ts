import { ISelect, ITextField } from './utils'

// ---------------------------
// interface - form
// ---------------------------
export interface IBookSearchForm {
  params: IBookSearchParams
  options: IBookSearchOptions
}

// ---------------------------
// interface - params
// ---------------------------
export interface IBookSearchParams {
  title: string
  author: string
  publisher: string
  isbn: string
  size: number
}

// ---------------------------
// interface - options
// ---------------------------
export interface IBookSearchOptions {
  title: ITextField
  author: ITextField
  publisher: ITextField
  isbn: ITextField
  size: ISelect
}

// ---------------------------
// const - Options
// ---------------------------
export const BookSearchOptions: IBookSearchOptions = {
  title: {
    label: '書籍名',
    rules: {},
  } as ITextField,
  author: {
    label: '著者名',
    rules: {},
  } as ITextField,
  publisher: {
    label: '出版社名',
    rules: {},
  } as ITextField,
  isbn: {
    label: 'ISBN',
    rules: {},
  } as ITextField,
  size: {
    label: '書籍サイズ',
    rules: {},
  } as ITextField,
}
