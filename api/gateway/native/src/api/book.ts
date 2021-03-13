import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import { BookResponse, CreateBookRequest } from '~/proto/book_apiv1_pb'

export function createBook(req: Request<any>): Promise<void> {
  const request = new CreateBookRequest()
  const metadata = getGrpcMetadata(req)

  const author = new CreateBookRequest.Author()
  author.setName('岸本斉史')

  const category = new CreateBookRequest.Category()
  category.setName('Comics & Graphic Novels')

  request.setTitle('NARUTO―ナルト― モノクロ版 29')
  request.setDescription('囚われの身となった我愛羅を救出するため、ナルトたちは“暁”のアジトを目指す！ しかし、その前にイタチが立ちふさがる!! 一方、別動隊のガイ一行の前に現れたのは鬼鮫!! 果たして我愛羅の運命や如何に…!?')
  request.setIsbn('08873849872840315501')
  request.setThumbnailUrl('http://books.google.com/books/content?id=7Z3LCwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api')
  request.setVersion('1.2.2.0.preview.3')
  request.setPublisher('集英社')
  request.setPublishedOn('2005-08-04')
  request.setAuthorsList([author])
  request.setCategoriesList([category])

  return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
    bookClient.createBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      console.log('debug', res)
      resolve()
    })
  })
}
