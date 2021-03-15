import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import { BookResponse, CreateBookRequest } from '~/proto/book_apiv1_pb'
import { ICreateBookInput } from '~/types/input'
import { IBookOutput, IBookOutputAuthor, IBookOutputCategory } from '~/types/output'

export function createBook(req: Request<any>, input: ICreateBookInput): Promise<IBookOutput> {
  const request = new CreateBookRequest()
  const metadata = getGrpcMetadata(req)

  const authors: CreateBookRequest.Author[] = input.authors.map((author: string) => {
    const item = new CreateBookRequest.Author()
    item.setName(author)

    return item
  })

  const categories: CreateBookRequest.Category[] = input.categories.map((category: string) => {
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

function setBookOutput(res: BookResponse): IBookOutput {
  const authors: IBookOutputAuthor[] = res.getAuthorsList().map((author: BookResponse.Author) => {
    return { id: author.getId(), name: author.getName() }
  })

  const categories: IBookOutputCategory[] = res.getCategoriesList().map((category: BookResponse.Category) => {
    return { id: category.getId(), name: category.getName() }
  })

  const output: IBookOutput = {
    id: res.getId(),
    publisherId: res.getPublisherId(),
    title: res.getTitle(),
    description: res.getDescription(),
    isbn: res.getIsbn(),
    thumbnailUrl: res.getThumbnailUrl(),
    version: res.getVersion(),
    publishedOn: res.getPublishedOn(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
    authors,
    categories,
  }

  return output
}
