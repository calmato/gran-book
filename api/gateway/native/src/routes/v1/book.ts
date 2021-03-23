import express, { NextFunction, Request, Response } from 'express'
import { createAndUpdateBooks } from '~/api'
import { IBookItem, ICreateAndUpdateBooksRequest } from '~/types/request'
import { IBookItemInput, ICreateAndUpdateBooksInput } from '~/types/input'
import { GrpcError } from '~/types/exception'
import { IBookListOutput, IBookOutput, IBookOutputAuthor, IBookOutputCategory } from '~/types/output'
import { IBookListResponse, IBookResponse } from '~/types/response'

const router = express.Router()

router.post(
  '/',
  async (req: Request, res: Response<IBookListResponse>, next: NextFunction): Promise<void> => {
    const { items } = req.body as ICreateAndUpdateBooksRequest

    const itemInputs: IBookItemInput[] = items?.map((item: IBookItem) => {
      const {
        title,
        authors,
        publisher,
        publishedDate,
        description,
        industryIdentifiers,
        categories,
        contentVersion,
        imageLinks,
      } = item?.volumeInfo

      const input: IBookItemInput = {
        title,
        description,
        authors,
        categories,
        isbn: '',
        thumbnailURL: '',
        version: contentVersion,
        publishedOn: publishedDate,
        publisher,
      }

      if (industryIdentifiers?.length > 0) {
        input.isbn = industryIdentifiers[0].identifier
      }

      if (imageLinks) {
        input.thumbnailURL = imageLinks.thumbnail || imageLinks.smallThumbnail || ''
      }

      return input
    })

    const input: ICreateAndUpdateBooksInput = {
      items: itemInputs,
    }

    await createAndUpdateBooks(req, input)
      .then((output: IBookListOutput) => {
        const response: IBookListResponse = setBookListResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setBookResponse(output: IBookOutput): IBookResponse {
  const publisher: string = output.publisher?.name || ''

  const authors: string[] = output.authors.map((author: IBookOutputAuthor) => {
    return author.name
  })

  const categories: string[] = output.categories.map((category: IBookOutputCategory) => {
    return category.name
  })

  const response: IBookResponse = {
    id: output.id,
    title: output.title,
    description: output.description,
    isbn: output.isbn,
    thumbnailUrl: output.thumbnailUrl,
    version: output.version,
    publishedOn: output.publishedOn,
    publisher,
    authors,
    categories,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  return response
}

function setBookListResponse(output: IBookListOutput): IBookListResponse {
  const books: IBookResponse[] = output.items?.map((item: IBookOutput) => {
    return setBookResponse(item)
  })

  const response: IBookListResponse = {
    books,
  }

  return response
}

export default router
