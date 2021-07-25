import { Request } from 'express'
import fs from 'fs'
import { chatClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { ICreateChatMessageInput, ICreateChatRoomInput, IListChatRoomInput, IUploadChatImageInput } from '~/types/input'
import {
  IChatMessageOutput,
  IChatRoomListOutput,
  IChatRoomListOutputMessage,
  IChatRoomListOutputRoom,
  IChatRoomOutput,
} from '~/types/output'
import {
  ChatMessageResponse,
  ChatRoomListResponse,
  ChatRoomResponse,
  CreateChatMessageRequest,
  CreateChatRoomRequest,
  ListChatRoomRequest,
  UploadChatImageRequest,
} from '~/proto/user_apiv1_pb'
import { ClientWritableStream } from '@grpc/grpc-js'

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
  const output: IChatRoomOutput = {
    id: res.getId(),
    userIds: res.getUserIdsList(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}

function setChatRoomListOutput(res: ChatRoomListResponse): IChatRoomListOutput {
  const rooms = res.getRoomsList().map(
    (cr: ChatRoomListResponse.Room): IChatRoomListOutputRoom => {
      const room: IChatRoomListOutputRoom = {
        id: cr.getId(),
        userIds: cr.getUserIdsList(),
        createdAt: '', // TODO: 追加
        updatedAt: '', // TODO: 追加
      }

      const message: ChatRoomListResponse.Message | undefined = cr.getLatestmessage()
      if (message) {
        const m: IChatRoomListOutputMessage = {
          userId: message.getUserId(),
          text: message.getText(),
          image: message.getImage(),
          createdAt: message.getCreatedAt(),
        }

        room.latestMessage = m
      }

      return room
    }
  )

  const output: IChatRoomListOutput = {
    rooms,
  }

  return output
}

function setChatMessageOuput(res: ChatMessageResponse): IChatMessageOutput {
  const output: IChatMessageOutput = {
    id: res.getId(),
    userId: res.getUserId(),
    text: res.getText(),
    image: res.getImage(),
    createdAt: res.getCreatedAt(),
  }

  return output
}
