import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import { BookListResponse, BookResponse, CreateAndUpdateBooksRequest, CreateBookRequest } from '~/proto/book_apiv1_pb'
import { IBookItemInput, ICreateAndUpdateBooksInput } from '~/types/input'
import { IBookListOutput, IBookOutput } from '~/types/output'

export function createBook(req: Request<any>, input: IBookItemInput): Promise<IBookOutput> {
  const request = new CreateBookRequest()
  const metadata = getGrpcMetadata(req)

  request.setTitle(input.title)
  request.setDescription(input.description)
  request.setIsbn(input.isbn)
  request.setThumbnailUrl(input.thumbnailURL)
  request.setVersion(input.version)
  request.setPublisher(input.publisher)
  request.setPublishedOn(input.publishedOn)
  request.setAuthorsList(input.authors)
  request.setCategoriesList(input.categories)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.createBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setBookOutput(res))
    })
  })
}

export function createAndUpdateBooks(req: Request<any>, input: ICreateAndUpdateBooksInput): Promise<IBookListOutput> {
  const request = new CreateAndUpdateBooksRequest()
  const metadata = getGrpcMetadata(req)

  const books: CreateAndUpdateBooksRequest.Book[] = input.items.map((item: IBookItemInput) => {
    const params = new CreateAndUpdateBooksRequest.Book()

    params.setTitle(item.title)
    params.setDescription(item.description)
    params.setIsbn(item.isbn)
    params.setThumbnailUrl(item.thumbnailURL)
    params.setVersion(item.version)
    params.setPublisher(item.publisher)
    params.setPublishedOn(item.publishedOn)
    params.setAuthorsList(item.authors)
    params.setCategoriesList(item.categories)

    return params
  })

  request.setBooksList(books)

  return new Promise((resolve: (output: IBookListOutput) => void, reject: (reason: Error) => void) => {
    bookClient.createAndUpdateBooks(request, metadata, (err: any, res: BookListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setBookListOutput(res))
    })
  })
}

function setBookOutput(res: BookResponse): IBookOutput {
  const output: IBookOutput = {
    id: res.getId(),
    title: res.getTitle(),
    description: res.getDescription(),
    isbn: res.getIsbn(),
    thumbnailUrl: res.getThumbnailUrl(),
    version: res.getVersion(),
    publisher: res.getPublisher(),
    publishedOn: res.getPublishedOn(),
    authors: res.getAuthorsList(),
    categories: res.getCategoriesList(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}

function setBookListOutput(res: BookListResponse): IBookListOutput {
  const books = res.getBooksList().map((item: BookListResponse.Book) => {
    const book: IBookOutput = {
      id: item.getId(),
      title: item.getTitle(),
      description: item.getDescription(),
      isbn: item.getIsbn(),
      thumbnailUrl: item.getThumbnailUrl(),
      version: item.getVersion(),
      publisher: item.getPublisher(),
      publishedOn: item.getPublishedOn(),
      authors: item.getAuthorsList(),
      categories: item.getCategoriesList(),
      createdAt: item.getCreatedAt(),
      updatedAt: item.getUpdatedAt(),
    }

    return book
  })

  const output: IBookListOutput = {
    books,
  }

  return output
}
