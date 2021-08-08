import { Request } from 'express'
import fs from 'fs'
import { chatClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { ICreateChatMessageInput, ICreateChatRoomInput, IListChatRoomInput, IUploadChatImageInput } from '~/types/input'
import { IChatMessageOutput, IChatRoomListOutput, IChatRoomListOutputRoom, IChatRoomOutput } from '~/types/output'
import { ClientWritableStream } from '@grpc/grpc-js'
import {
  ChatMessageResponse,
  ChatRoomListResponse,
  ChatRoomResponse,
  CreateChatRoomRequest,
  CreateChatMessageRequest,
  ListChatRoomRequest,
  UploadChatImageRequest,
} from '~/proto/chat_service_pb'

export function listChatRoom(req: Request<any>, input: IListChatRoomInput): Promise<IChatRoomListOutput> {
  const request = new ListChatRoomRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  return new Promise((resolve: (res: IChatRoomListOutput) => void, reject: (reason: Error) => void) => {
    chatClient.listRoom(request, metadata, (err: any, res: ChatRoomListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setChatRoomListOutput(res))
    })
  })
}

export function createChatRoom(req: Request<any>, input: ICreateChatRoomInput): Promise<IChatRoomOutput> {
  const request = new CreateChatRoomRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserIdsList(input.userIds)

  return new Promise((resolve: (res: IChatRoomOutput) => void, reject: (reason: Error) => void) => {
    chatClient.createRoom(request, metadata, (err: any, res: ChatRoomResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setChatRoomOutput(res))
    })
  })
}

export function createChatMessage(req: Request<any>, input: ICreateChatMessageInput): Promise<IChatMessageOutput> {
  const request = new CreateChatMessageRequest()
  const metadata = getGrpcMetadata(req)

  request.setRoomId(input.roomId)
  request.setUserId(input.userId)
  request.setText(input.text)

  return new Promise((resolve: (res: IChatMessageOutput) => void, reject: (reason: Error) => void) => {
    chatClient.createMessage(request, metadata, (err: any, res: ChatMessageResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setChatMessageOuput(res))
    })
  })
}

export function uploadChatImage(req: Request<any>, input: IUploadChatImageInput): Promise<IChatMessageOutput> {
  const metadata = getGrpcMetadata(req)

  return new Promise((resolve: (res: IChatMessageOutput) => void, reject: (reason: Error) => void) => {
    const call: ClientWritableStream<UploadChatImageRequest> = chatClient.uploadImage(
      metadata,
      (err: any, res: ChatMessageResponse) => {
        if (err) {
          return reject(getGrpcError(err))
        }

        return resolve(setChatMessageOuput(res))
      }
    )

    const stream: fs.ReadStream = fs.createReadStream(input.path, { highWaterMark: 102400 })
    let count = 0 // 読み込み回数

    stream.on('data', (chunk: Buffer) => {
      const request = new UploadChatImageRequest()
      request.setRoomId(input.roomId)
      request.setUserId(input.userId)
      request.setImage(chunk)
      request.setPosition(count)

      call.write(request)
      count += 1
    })

    stream.on('end', () => {
      call.end() // TODO: try-catchとかのエラー処理必要かも
    })
  })
}

function setChatRoomOutput(res: ChatRoomResponse): IChatRoomOutput {
  const room: ChatRoomResponse.AsObject = res.toObject()

  const output: IChatRoomOutput = {
    ...room,
    userIds: room.userIdsList,
  }

  return output
}

function setChatRoomListOutput(res: ChatRoomListResponse): IChatRoomListOutput {
  const rooms: Array<IChatRoomListOutputRoom> = res.getRoomsList().map(
    (value: ChatRoomListResponse.Room): IChatRoomListOutputRoom => {
      const room: ChatRoomListResponse.Room.AsObject = value.toObject()

      const output: IChatRoomListOutputRoom = {
        ...room,
        userIds: room.userIdsList,
      }

      return output
    }
  )

  const output: IChatRoomListOutput = {
    rooms,
  }

  return output
}

function setChatMessageOuput(res: ChatMessageResponse): IChatMessageOutput {
  const message: ChatMessageResponse.AsObject = res.toObject()

  const output: IChatMessageOutput = { ...message }
  return output
}
