import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import axios, { AxiosResponse } from 'axios'
import { IBook, IBookState } from '~/types/store'
import { IErrorResponse, IRakutenBookItem, IRakutenBookResponse } from '~/types/responses'
import { ApiError } from '~/types/exception'
import { IBookSearchForm } from '~/types/forms'

const initialState: IBookState = {
  books: [],
}

@Module({
  name: 'book',
  stateFactory: true,
  namespaced: true,
})
export default class BookModule extends VuexModule {
  private books: IBook[] = initialState.books

  public get getBooks(): IBook[] {
    return this.books
  }

  @Mutation
  private setBooks(books: IBook[]): void {
    this.books = books
  }

  @Action({ rawError: true })
  public searchBookFromRakutenBooksAPI(payload: IBookSearchForm): Promise<void> {
    const url: string = getRakutenBooksURL(payload)

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      axios
        .get(url)
        .then((res: AxiosResponse<IRakutenBookResponse>) => {
          const data = res.data
          if (!data) {
            return
          }

          const books: IBook[] = data.Items.map((item: IRakutenBookItem) => {
            return {
              id: 0,
              title: item.title,
              titleKana: item.titleKana,
              description: item.itemCaption,
              isbn: item.isbn,
              publisher: item.publisherName,
              publishedOn: item.salesDate,
              thumbnailUrl: item.largeImageUrl,
              rakutenUrl: item.itemUrl,
              rakutenGenreId: item.booksGenreId,
              author: item.author,
              authorKana: item.authorKana,
              createdAt: '',
              updatedAt: '',
            }
          })

          this.setBooks(books)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }
}

function getRakutenBooksURL(form: IBookSearchForm): string {
  const applicationId: string = process.env.rakutenAppId || ''
  let baseURL: string = `https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?format=json&formatVersion=2&applicationId=${applicationId}`

  if (!form?.params) {
    return baseURL
  }

  if (form.params.title !== '') {
    baseURL = `${baseURL}&title=${form.params.title}`
  }

  if (form.params.author !== '') {
    baseURL = `${baseURL}&author=${form.params.author}`
  }

  if (form.params.publisher !== '') {
    baseURL = `${baseURL}&publisherName=${form.params.publisher}`
  }

  if (form.params.isbn !== '') {
    baseURL = `${baseURL}&isbn=${form.params.isbn}`
  }

  if (form.params.size !== 0) {
    baseURL = `${baseURL}&size=${form.params.size}`
  }

  return baseURL
}
