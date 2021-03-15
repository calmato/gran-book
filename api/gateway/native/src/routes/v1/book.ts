import express, { NextFunction, Request, Response } from 'express'
import { createBook } from '~/api'
import { ICreateBookRequest } from '~/types/request'
import { ICreateBookInput } from '~/types/input'
import { GrpcError } from '~/types/exception'
import { IBookOutput, IBookOutputAuthor, IBookOutputCategory } from '~/types/output'
import { IBookResponse, IBookResponseAuthor, IBookResponseCategory } from '~/types/response'

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

    if (industryIdentifiers.length > 0) {
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
  const authors: IBookResponseAuthor[] = output.authors.map((author: IBookOutputAuthor) => {
    return { id: author.id, name: author.name }
  })

  const categories: IBookResponseCategory[] = output.categories.map((category: IBookOutputCategory) => {
    return { id: category.id, name: category.name }
  })

  const response: IBookResponse = {
    id: output.id,
    publisherId: output.publisherId,
    title: output.title,
    description: output.description,
    isbn: output.isbn,
    thumbnailUrl: output.thumbnailUrl,
    version: output.version,
    publishedOn: output.publishedOn,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
    authors,
    categories,
  }

  return response
}

export default router
