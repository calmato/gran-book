import express, { NextFunction, Request, Response } from 'express'
import { createBook } from '~/api'
import { ICreateBookRequest } from '~/types/request'
import { ICreateBookInput } from '~/types/input'
import { GrpcError } from '~/types/exception'
import { IBookOutput, IBookOutputAuthor, IBookOutputCategory } from '~/types/output'
import { IBookResponse } from '~/types/response'

const router = express.Router()

router.post(
  '/',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
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
    } = req.body as ICreateBookRequest

    const input: ICreateBookInput = {
      title,
      description,
      authors,
      categories,
      isbn: '',
      thumbnailURL: '',
      version: contentVersion,
      publisher: publisher,
      publishedOn: publishedDate,
    }

    if (industryIdentifiers?.length > 0) {
      input.isbn = industryIdentifiers[0].identifier
    }

    if (imageLinks) {
      input.thumbnailURL = imageLinks.thumbnail || imageLinks.smallThumbnail || ''
    }

    await createBook(req, input)
      .then((output: IBookOutput) => {
        const response: IBookResponse = setBookResponse(output)
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

export default router
