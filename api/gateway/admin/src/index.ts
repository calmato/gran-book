import express from 'express'
import { health, v1Auth } from '~/routes'

const app = express()

const port: string = process.env.PORT || '3000'

app.use('/', health)
app.use('/v1/auth', v1Auth)

app.listen(port, (): void => {
  console.log(`listening at http://localhost:${port}`)
})

export default app
