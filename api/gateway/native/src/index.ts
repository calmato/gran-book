import express from 'express'
import bodyParser from 'body-parser'
import { health, v1Auth } from '~/routes'

const app = express()

app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.json())

const host: string = process.env.HOST || '0.0.0.0'
const port: string = process.env.PORT || '3000'

app.use('/', health)
app.use('/v1/auth', v1Auth)

app.listen(port, (): void => {
  console.log(`listening at http://${host}:${port}`)
})

export default app