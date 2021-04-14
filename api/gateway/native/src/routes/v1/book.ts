import express, { NextFunction, Request, Response } from 'express'
import { createBook, showBook, updateBook } from '~/api'
import { GrpcError } from '~/types/exception'
import { IBookInputAuthor, ICreateBookInput, IShowBookInput, IUpdateBookInput } from '~/types/input'
import { IBookOutput, IBookOutputAuthor, IBookOutputReview } from '~/types/output'
import { ICreateBookRequest, IUpdateBookRequest } from '~/types/request'
import { IBookResponse, IBookResponseBookshelf, IBookResponseReview, IBookResponseUser } from '~/types/response'

const router = express.Router()

router.get(
  '/:isbn',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const { isbn } = req.params

    const input: IShowBookInput = {
      isbn,
    }

    await showBook(req, input)
      .then((output: IBookOutput) => {
        // TODO: User情報の取得
        const response: IBookResponse = setBookResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const {
      title,
      titleKana,
      itemCaption,
      isbn,
      publisherName,
      salesDate,
      largeImageUrl,
      mediumImageUrl,
      smallImageUrl,
      itemUrl,
      booksGenreId,
      author,
      authorKana,
    } = req.body as ICreateBookRequest

    const authorNames: string[] = author.split('/')
    const authorNameKanas: string[] = authorKana.split('/')

    // 出版日のフォーマット: 2020年01月01日頃 -> 2020-01-01
    const publishedOn: string = salesDate?.replace(/[年月]/g, '-').replace(/[日頃]/g, '') || ''

    const authors: IBookInputAuthor[] = authorNames.map((val: string, i: number) => {
      const item: IBookInputAuthor = {
        name: val,
        nameKana: authorNameKanas[i],
      }

      return item
    })

    const input: ICreateBookInput = {
      title,
      titleKana,
      description: itemCaption,
      isbn,
      publisher: publisherName,
      publishedOn,
      thumbnailUrl: mediumImageUrl || smallImageUrl || largeImageUrl,
      rakutenUrl: itemUrl,
      rakutenGenreId: booksGenreId,
      authors,
    }

    await createBook(req, input)
      .then((output: IBookOutput) => {
        // TODO: User情報の取得
        const response: IBookResponse = setBookResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const {
      title,
      titleKana,
      itemCaption,
      isbn,
      publisherName,
      salesDate,
      largeImageUrl,
      mediumImageUrl,
      smallImageUrl,
      itemUrl,
      booksGenreId,
      author,
      authorKana,
    } = req.body as IUpdateBookRequest

    const authorNames: string[] = author.split('/')
    const authorNameKanas: string[] = authorKana.split('/')

    // 出版日のフォーマット: 2020年01月01日頃 -> 2020-01-01
    const publishedOn: string = salesDate?.replace(/[年月]/g, '-').replace(/[日頃]/g, '') || ''

    const authors: IBookInputAuthor[] = authorNames.map((val: string, i: number) => {
      const item: IBookInputAuthor = {
        name: val,
        nameKana: authorNameKanas[i],
      }

      return item
    })

    const input: IUpdateBookInput = {
      title,
      titleKana,
      description: itemCaption,
      isbn,
      publisher: publisherName,
      publishedOn,
      thumbnailUrl: mediumImageUrl || smallImageUrl || largeImageUrl,
      rakutenUrl: itemUrl,
      rakutenGenreId: booksGenreId,
      authors,
    }

    await updateBook(req, input)
      .then((output: IBookOutput) => {
        // TODO: User情報の取得
        const response: IBookResponse = setBookResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setBookResponse(bookOutput: IBookOutput): IBookResponse {
  const bookshelf: IBookResponseBookshelf = {
    id: bookOutput.bookshelf?.id,
    status: bookOutput.bookshelf?.status,
    readOn: bookOutput.bookshelf?.readOn,
    createdAt: bookOutput.bookshelf?.createdAt,
    updatedAt: bookOutput.bookshelf?.updatedAt,
  }

  const reviews: IBookResponseReview[] = bookOutput.reviews.map((item: IBookOutputReview) => {
    const user: IBookResponseUser = {
      id: '',
      username: '',
      thumbnailUrl: '',
    }

    const review: IBookResponseReview = {
      id: item.id,
      score: item.score,
      impression: item.impression,
      createdAt: item.createdAt,
      updatedAt: item.updatedAt,
      user,
    }

    return review
  })

  const authorNames: string[] = bookOutput.authors.map((item: IBookOutputAuthor) => {
    return item.name
  })

  const authorNameKanas: string[] = bookOutput.authors.map((item: IBookOutputAuthor) => {
    return item.nameKana
  })

  const response: IBookResponse = {
    id: bookOutput.id,
    title: bookOutput.title,
    titleKana: bookOutput.titleKana,
    description: bookOutput.description,
    isbn: bookOutput.isbn,
    publisher: bookOutput.publisher,
    publishedOn: bookOutput.publishedOn,
    thumbnailUrl: bookOutput.thumbnailUrl,
    rakutenUrl: bookOutput.rakutenUrl,
    rakutenGenreId: bookOutput.rakutenGenreId,
    author: authorNames.join('/'),
    authorKana: authorNameKanas.join('/'),
    createdAt: bookOutput.createdAt,
    updatedAt: bookOutput.updatedAt,
    bookshelf,
    reviews,
  }

  return response
}

export default router
