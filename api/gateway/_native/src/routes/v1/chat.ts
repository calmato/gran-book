import express, { NextFunction, Request, Response } from 'express'
import multer from '~/plugins/multer'
import { createChatMessage, createChatRoom, getUser, listChatRoom, multiGetUser, uploadChatImage } from '~/api'
import {
  ICreateChatMessageInput,
  ICreateChatRoomInput,
  IGetUserInput,
  IListChatRoomInput,
  IMultiGetUserInput,
  IUploadChatImageInput,
} from '~/types/input'
import {
  IChatMessageOutput,
  IChatRoomListOutput,
  IChatRoomListOutputRoom,
  IChatRoomOutput,
  IUserMapOutput,
  IUserOutput,
} from '~/types/output'
import { ICreateChatMessageRequest, ICreateChatRoomRequest } from '~/types/request'
import {
  IChatMessageResponse,
  IChatMessageResponseUser,
  IChatRoomListResponse,
  IChatRoomListResponseMessage,
  IChatRoomListResponseRoom,
  IChatRoomListResponseUser,
  IChatRoomResponse,
  IChatRoomResponseUser,
} from '~/types/response'
import { badRequest } from '~/lib/http-exception'

const router = express.Router()

router.get(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomListResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const roomsInput: IListChatRoomInput = {
      userId: userId,
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listChatRoom(req, roomsInput)
      .then(async (roomsOutput: IChatRoomListOutput) => {
        let userIds: string[] = []
        roomsOutput.rooms.forEach((cr: IChatRoomListOutputRoom) => {
          userIds = userIds.concat(cr.userIds)
        })

        const userListInput: IMultiGetUserInput = {
          userIds,
        }

        const usersOutput: IUserMapOutput = await multiGetUser(req, userListInput).catch(() => {
          return {} as IUserMapOutput
        })

        const response: IChatRoomListResponse = setChatRoomListResponse(roomsOutput, usersOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.post(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { users } = req.body as ICreateChatRoomRequest

    const ids: string[] = users.concat(userId)

    const usersInput: IMultiGetUserInput = {
      userIds: Array.from(new Set(ids)).sort(),
    }

    const roomInput: ICreateChatRoomInput = {
      userIds: usersInput.userIds,
    }

    await multiGetUser(req, usersInput)
      .then(async (usersOutput: IUserMapOutput) => {
        if (usersInput.userIds.length !== Object.keys(usersOutput).length) {
          throw badRequest(['One of the user ids is incorrect'])
        }

        const roomOutput: IChatRoomOutput = await createChatRoom(req, roomInput)

        const response: IChatRoomResponse = setChatRoomResponse(roomOutput, usersOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.post(
  '/v1/users/:userId/chat/:chatId/messages/text',
  async (req: Request, res: Response<IChatMessageResponse>, next: NextFunction): Promise<void> => {
    const { userId, chatId } = req.params
    const { text } = req.body as ICreateChatMessageRequest

    const userInput: IGetUserInput = {
      userId,
    }

    const messageInput: ICreateChatMessageInput = {
      roomId: chatId,
      userId,
      text,
    }

    await getUser(req, userInput)
      .then(async (userOutput: IUserOutput) => {
        const messageOutput: IChatMessageOutput = await createChatMessage(req, messageInput)

        const response: IChatMessageResponse = setChatMessageResponse(messageOutput, userOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.post(
  '/v1/users/:userId/chat/:chatId/messages/image',
  multer.single('image'),
  async (req: Request, res: Response<IChatMessageResponse>, next: NextFunction): Promise<void> => {
    if (!req.file) {
      next(badRequest([{ message: 'image is not exists' }]))
      return
    }

    const { userId, chatId } = req.params

    const userInput: IGetUserInput = {
      userId,
    }

    const messageInput: IUploadChatImageInput = {
      roomId: chatId,
      userId,
      path: req.file.path,
    }

    await getUser(req, userInput)
      .then(async (userOutput: IUserOutput) => {
        const messageOutput: IChatMessageOutput = await uploadChatImage(req, messageInput)

        const response: IChatMessageResponse = setChatMessageResponse(messageOutput, userOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

function setChatRoomResponse(roomOutput: IChatRoomOutput, usersOutput: IUserMapOutput): IChatRoomResponse {
  const users = roomOutput.userIds.map(
    (userId: string): IChatRoomResponseUser => {
      const user: IChatRoomResponseUser = {
        id: userId,
        username: usersOutput[userId]?.username || 'unknown',
        thumbnailUrl: usersOutput[userId]?.thumbnailUrl || '',
      }

      return user
    }
  )

  const response: IChatRoomResponse = {
    id: roomOutput.id,
    users,
    createdAt: roomOutput.createdAt,
    updatedAt: roomOutput.updatedAt,
  }

  return response
}

function setChatRoomListResponse(roomsOutput: IChatRoomListOutput, usersOutput: IUserMapOutput): IChatRoomListResponse {
  const rooms: Array<IChatRoomListResponseRoom> = roomsOutput.rooms.map(
    (value: IChatRoomListOutputRoom): IChatRoomListResponseRoom => {
      const users = value.userIds.map(
        (userId: string): IChatRoomListResponseUser => {
          const user: IChatRoomListResponseUser = {
            id: userId,
            username: usersOutput[userId]?.username || 'unknown',
            thumbnailUrl: usersOutput[userId]?.thumbnailUrl || '',
          }

          return user
        }
      )

      const room: IChatRoomListResponseRoom = {
        id: value.id,
        users,
        createdAt: value.createdAt,
        updatedAt: value.updatedAt,
      }

      if (value.latestMessage) {
        const m: IChatRoomListResponseMessage = {
          userId: value.latestMessage.userId,
          text: value.latestMessage.text,
          image: value.latestMessage.image,
          createdAt: value.latestMessage.createdAt,
        }

        room.latestMessage = m
      }

      return room
    }
  )

  const response: IChatRoomListResponse = {
    rooms,
  }

  return response
}

function setChatMessageResponse(messageOutput: IChatMessageOutput, userOutput?: IUserOutput): IChatMessageResponse {
  const user: IChatMessageResponseUser = {
    id: messageOutput.userId,
    username: userOutput?.username || 'unknown',
    thumbnailUrl: userOutput?.thumbnailUrl || '',
  }

  const response: IChatMessageResponse = {
    id: messageOutput.id,
    text: messageOutput.text,
    image: messageOutput.image,
    user,
    createdAt: messageOutput.createdAt,
  }

  return response
}

export default router
