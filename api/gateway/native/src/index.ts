import express from 'express'
import bodyParser from 'body-parser'
import cors from 'cors'
import { corsOptions } from '~/config/cors'
import { authentication } from '~/lib/authenticated'
import { notFoundErrorHandler, otherErrorHandler } from '~/lib/error-handler'
import { accessLogHandler } from '~/lib/log-handler'
import { common, v1Auth, v1Book, v1User } from '~/routes'

const app = express()

const host: string = process.env.HOST || '0.0.0.0'
const port: string = process.env.PORT || '3000'

app.use(bodyParser.urlencoded({ limit: '100mb', extended: true }))
app.use(bodyParser.json({ limit: '100mb' }))
app.use(cors(corsOptions))
app.use(accessLogHandler)
app.use(authentication)

app.use('/', common)
app.use('/v1/auth', v1Auth)
app.use('/v1/books', v1Book)
app.use('/v1/users', v1User)

app.use(notFoundErrorHandler)
app.use(otherErrorHandler)

app.listen(port, (): void => {
  console.log(`listening at http://${host}:${port}`)
})

export default app
