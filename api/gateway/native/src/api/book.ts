import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import { BookListResponse, BookResponse, CreateAndUpdateBooksRequest, CreateBookRequest } from '~/proto/book_apiv1_pb'
import { IBookItemInput, ICreateAndUpdateBooksInput } from '~/types/input'
import { IBookListOutput, IBookOutput, IBookOutputAuthor, IBookOutputCategory, IBookOutputPublisher } from '~/types/output'

export function createBook(req: Request<any>, input: IBookItemInput): Promise<IBookOutput> {
  const request = new CreateBookRequest()
  const metadata = getGrpcMetadata(req)

  const authors: CreateBookRequest.Author[] = input.authors?.map((author: string) => {
    const item = new CreateBookRequest.Author()
    item.setName(author)

    return item
  })

  const categories: CreateBookRequest.Category[] = input.categories?.map((category: string) => {
    const item = new CreateBookRequest.Category()
    item.setName(category)

    return item
  })

  request.setTitle(input.title)
  request.setDescription(input.description)
  request.setIsbn(input.isbn)
  request.setThumbnailUrl(input.thumbnailURL)
  request.setVersion(input.version)
  request.setPublisher(input.publisher)
  request.setPublishedOn(input.publishedOn)
  request.setAuthorsList(authors)
  request.setCategoriesList(categories)

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

  const items: CreateAndUpdateBooksRequest.Item[] = input.items.map((item: IBookItemInput) => {
    const params = new CreateAndUpdateBooksRequest.Item()

    const authors: CreateAndUpdateBooksRequest.Author[] = item.authors?.map((author: string) => {
      const param = new CreateAndUpdateBooksRequest.Author()
      param.setName(author)

      return param
    })

    const categories: CreateAndUpdateBooksRequest.Category[] = item.categories?.map((category: string) => {
      const param = new CreateAndUpdateBooksRequest.Category()
      param.setName(category)

      return param
    })

    params.setTitle(item.title)
    params.setDescription(item.description)
    params.setIsbn(item.isbn)
    params.setThumbnailUrl(item.thumbnailURL)
    params.setVersion(item.version)
    params.setPublisher(item.publisher)
    params.setPublishedOn(item.publishedOn)
    params.setAuthorsList(authors)
    params.setCategoriesList(categories)

    return params
  })

  request.setItemsList(items)

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
  const publisher: IBookOutputPublisher = {
    id: res.getPublisher()?.getId() || 0,
    name: res.getPublisher()?.getName() || '',
  }

  const authors: IBookOutputAuthor[] = res.getAuthorsList().map((author: BookResponse.Author) => {
    return { id: author.getId(), name: author.getName() }
  })

  const categories: IBookOutputCategory[] = res.getCategoriesList().map((category: BookResponse.Category) => {
    return { id: category.getId(), name: category.getName() }
  })

  const output: IBookOutput = {
    id: res.getId(),
    title: res.getTitle(),
    description: res.getDescription(),
    isbn: res.getIsbn(),
    thumbnailUrl: res.getThumbnailUrl(),
    version: res.getVersion(),
    publishedOn: res.getPublishedOn(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
    publisher,
    authors,
    categories,
  }

  return output
}

function setBookListOutput(res: BookListResponse): IBookListOutput {
  const items = res.getItemsList().map((item: BookListResponse.Item) => {
    const publisher: IBookOutputPublisher = {
      id: item.getPublisher()?.getId() || 0,
      name: item.getPublisher()?.getName() || '',
    }
    const authors: IBookOutputAuthor[] = item.getAuthorsList().map((author: BookListResponse.Author) => {
      return { id: author.getId(), name: author.getName() }
    })

    const categories: IBookOutputCategory[] = item.getCategoriesList().map((category: BookListResponse.Category) => {
      return { id: category.getId(), name: category.getName() }
    })

    const output: IBookOutput = {
      id: item.getId(),
      title: item.getTitle(),
      description: item.getDescription(),
      isbn: item.getIsbn(),
      thumbnailUrl: item.getThumbnailUrl(),
      version: item.getVersion(),
      publishedOn: item.getPublishedOn(),
      createdAt: item.getCreatedAt(),
      updatedAt: item.getUpdatedAt(),
      publisher,
      authors,
      categories,
    }

    return output
  })

  const output: IBookListOutput = {
    items,
  }

  return output
}
