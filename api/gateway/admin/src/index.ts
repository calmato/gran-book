import express from 'express'
import bodyParser from 'body-parser'
import cors from 'cors'
import { corsOptions } from '~/config/cors'
import { authentication } from '~/lib/authenticated'
import { errorHandler } from '~/lib/error-handler'
import { accessLogHandler } from '~/lib/log-handler'
import { health, v1Auth } from '~/routes'

const app = express()

const host: string = process.env.HOST || '0.0.0.0'
const port: string = process.env.PORT || '3000'

app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.json())
app.use(cors(corsOptions))
app.use(accessLogHandler)
app.use(authentication)

app.use('/', health)
app.use('/v1/auth', v1Auth)

app.use(errorHandler)

app.listen(port, (): void => {
  console.log(`listening at http://${host}:${port}`)
})

export default app
