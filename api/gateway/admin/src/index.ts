import express from 'express'
import { urlencoded, json } from 'body-parser'
import cors from 'cors'
import { corsOptions } from '~/config/cors'
import { authentication, authorization } from '~/lib/authenticated'
import { notFoundErrorHandler, otherErrorHandler } from '~/lib/error-handler'
import { accessLogHandler } from '~/lib/log-handler'
import { common, v1Admin, v1Auth, v1Book, v1User } from '~/routes'

const app = express()

const host: string = process.env.HOST || '0.0.0.0'
const port: string = process.env.PORT || '3000'

app.use(urlencoded({ limit: '4mb', extended: true }))
app.use(json({ limit: '4mb' }))
app.use(cors(corsOptions))
app.use(accessLogHandler)
app.use(authentication)
app.use(authorization)

app.use(common)
app.use(v1Admin) // /v1/admin
app.use(v1Auth) // /v1/auth
app.use(v1Book) // /v1/books
app.use(v1User) // /v1/users

app.use(notFoundErrorHandler)
app.use(otherErrorHandler)

app.listen(port, (): void => {
  console.log(`listening at http://${host}:${port}`)
})

export default app
