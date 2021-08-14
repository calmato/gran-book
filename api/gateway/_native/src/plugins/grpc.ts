import { credentials } from '@grpc/grpc-js'
import { AuthServiceClient, IAuthServiceClient } from '~/proto/auth_service_grpc_pb'
import { BookServiceClient, IBookServiceClient } from '~/proto/book_apiv1_grpc_pb'
import { ChatServiceClient, IChatServiceClient } from '~/proto/chat_service_grpc_pb'
import { IUserServiceClient, UserServiceClient } from '~/proto/user_service_grpc_pb'

const userAPIURL: string = process.env.USER_API_URL || 'user_api:8080'
const bookAPIURL: string = process.env.BOOK_API_URL || 'book_api:8080'

const authClient: IAuthServiceClient = new AuthServiceClient(userAPIURL, credentials.createInsecure())
const bookClient: IBookServiceClient = new BookServiceClient(bookAPIURL, credentials.createInsecure())
const chatClient: IChatServiceClient = new ChatServiceClient(userAPIURL, credentials.createInsecure())
const userClient: IUserServiceClient = new UserServiceClient(userAPIURL, credentials.createInsecure())

export { authClient, bookClient, chatClient, userClient }
