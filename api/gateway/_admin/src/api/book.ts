import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import { DeleteBookRequest, EmptyBook } from '~/proto/book_apiv1_pb'
import { IDeleteBookInput } from '~/types/input'

export function deleteBook(req: Request<any>, input: IDeleteBookInput): Promise<void> {
  const request = new DeleteBookRequest()
  const metadata = getGrpcMetadata(req)

  request.setBookId(input.bookId)

  return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
    bookClient.deleteBook(request, metadata, (err: any, _: EmptyBook) => {
      if (err) {
        reject(getGrpcError(err))
      }

      resolve()
    })
  })
}
